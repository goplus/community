/*
 * Copyright (c) 2023 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/qiniu/x/xlog"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	_ "github.com/go-sql-driver/mysql"
	"gocloud.dev/blob"
	"golang.org/x/oauth2"
)

var (
	ErrNotExist   = os.ErrNotExist
	ErrPermission = os.ErrPermission
)

type Config struct {
	Driver string // database driver. default is `mysql`.
	DSN    string // database data source name
	CAS    string // casdoor database data source name
	BlobUS string // blob URL scheme
}

type ArticleEntry struct {
	ID       string
	Title    string
	UId      string
	Cover    string
	Tags     string
	User     User
	Abstract string
	Ctime    time.Time
	Mtime    time.Time
}

type Article struct {
	ArticleEntry
	Content string // in markdown
	// Status  int    // published or draft
	HtmlUrl  string // parsed html file url
	HtmlData string
}

type Community struct {
	bucket        *blob.Bucket
	db            *sql.DB
	domain        string
	casdoorConfig *CasdoorConfig
	xLog          *xlog.Logger
}
type CasdoorConfig struct {
	endPoint         string
	clientId         string
	clientSecret     string
	certificate      string
	organizationName string
	applicationName  string
}

type Account struct {
	casdoorConfig *CasdoorConfig

	xLog *xlog.Logger
}

func New(ctx context.Context, conf *Config) (ret *Community, err error) {
	// Init log
	xLog := xlog.New("")

	if conf == nil {
		conf = new(Config)
	}
	casdoorConf := casdoorConfigInit()
	driver := conf.Driver
	dsn := conf.DSN
	bus := conf.BlobUS
	if driver == "" {
		driver = "mysql"
	}
	if dsn == "" {
		dsn = os.Getenv("GOP_COMMUNITY_DSN")
	}
	if bus == "" {
		bus = os.Getenv("GOP_COMMUNITY_BLOBUS")
	}
	domain := os.Getenv("GOP_COMMUNITY_DOMAIN")

	bucket, err := blob.OpenBucket(ctx, bus)
	if err != nil {
		xLog.Error(err)
		return
	}
	db, err := sql.Open(driver, dsn)
	if err != nil {
		xLog.Error(err)
		return
	}
	return &Community{bucket, db, domain, casdoorConf, xLog}, nil
}

func (p *Community) getTotal(ctx context.Context, searchValue string) (total int, err error) {
	if searchValue != "" {
		sqlStr := "select count(*) from article where title like ?"
		err = p.db.QueryRow(sqlStr, "%"+searchValue+"%").Scan(&total)
	} else {
		sqlStr := "select count(*) from article"
		err = p.db.QueryRow(sqlStr).Scan(&total)
	}
	return
}

// Article returns an article.
func (p *Community) Article(ctx context.Context, id string) (article *Article, err error) {
	article = &Article{}
	var htmlId string
	sqlStr := "select id,title,user_id,cover,tags,abstract,content,html_id,ctime,mtime from article where id=?"
	err = p.db.QueryRow(sqlStr, id).Scan(&article.ID, &article.Title, &article.UId, &article.Cover, &article.Tags, &article.Abstract, &article.Content, &htmlId, &article.Ctime, &article.Mtime)
	if err == sql.ErrNoRows {
		p.xLog.Warn("not found the article")
		return article, ErrNotExist
	} else if err != nil {
		return article, err
	}
	// add author info
	user, err := p.GetUserById(article.UId)
	if err != nil {
		return
	}
	article.User = *user
	// get html url
	fileKey, err := p.GetMediaUrl(ctx, htmlId)
	article.HtmlUrl = fmt.Sprintf("%s%s", p.domain, fileKey)
	if err != nil {
		p.xLog.Warn("have no html media")
		article.HtmlUrl = ""
	}
	return
}

// TransHtmlUrl get translation html url
func (p *Community) TransHtmlUrl(ctx context.Context, id string) (htmlUrl string, err error) {
	var htmlId string
	sqlStr := "select trans_html_id from article where id=?"
	err = p.db.QueryRow(sqlStr, id).Scan(&htmlId)
	if err == sql.ErrNoRows {
		p.xLog.Warn("not found the translation html")
		return "", ErrNotExist
	}

	// get html url
	fileKey, err := p.GetMediaUrl(ctx, htmlId)
	htmlUrl = fmt.Sprintf("%s%s", p.domain, fileKey)
	if err != nil {
		p.xLog.Warn("have no html media")
		htmlUrl = ""
	}
	return
}

// CanEditable determine whether the user has the permission to operate.
func (p *Community) CanEditable(ctx context.Context, uid, id string) (editable bool, err error) {
	sqlStr := "select id from article where id=? and user_id = ?"
	err = p.db.QueryRow(sqlStr, id, uid).Scan(&id)
	if err != nil {
		return false, ErrPermission
	}
	return true, nil
}

// SaveHtml upload origin html(string) to media for html id and save id to database
func (p *Community) SaveHtml(ctx context.Context, uid, htmlStr, mdData, id string) (articleId string, err error) {
	htmlId, err := p.SaveMedia(ctx, uid, []byte(htmlStr))
	if id == "" {
		// save to database
		sqlStr := "insert into article (user_id, html_id, ctime, mtime, content) values (?, ?, ?)"
		res, err := p.db.Exec(sqlStr, uid, htmlId, time.Now(), time.Now(), mdData)
		if err != nil {
			return "", err
		}
		articleId, err := res.LastInsertId()
		return strconv.FormatInt(articleId, 10), nil
	}
	sqlStr := "update article set content=?, html_id=?, mtime=? where id=?"
	_, err = p.db.Exec(sqlStr, mdData, htmlId, time.Now(), id)
	if err != nil {
		return "", err
	}
	return id, err
}

// uploadHtml upload html(string) to media for html id
func (p *Community) uploadHtml(ctx context.Context, uid, htmlStr string) (htmlId int64, err error) {
	htmlId, err = p.SaveMedia(ctx, uid, []byte(htmlStr))
	return
}

// PutArticle adds new article (ID == "") or edits an existing article (ID != "").
func (p *Community) PutArticle(ctx context.Context, uid string, trans string, article *Article) (id string, err error) {
	htmlId, err := p.uploadHtml(ctx, uid, article.HtmlData)
	if err != nil {
		htmlId = 0
	}
	// new article
	if article.ID == "" {
		sqlStr := "insert into article (title, ctime, mtime, user_id, tags, abstract, cover, content, html_id) values (?, ?, ?, ?, ?, ?, ?, ?,?)"
		res, err := p.db.Exec(sqlStr, &article.Title, time.Now(), time.Now(), uid, &article.Tags, &article.Abstract, &article.Cover, &article.Content, htmlId)
		if err != nil {
			return "", err
		}
		idInt, err := res.LastInsertId()
		return strconv.FormatInt(idInt, 10), nil
	}
	if trans != "" {
		// add article except html_id, content (trans)
		sqlStr := "update article set title=?, mtime=?, ctime=?, tags=?, abstract=?, cover=?, trans_content=?, trans_html_id=? where id=?"
		_, err = p.db.Exec(sqlStr, &article.Title, time.Now(), time.Now(), &article.Tags, &article.Abstract, &article.Cover, &article.Content, htmlId, &article.ID)
		return article.ID, err
	}

	// edit article
	sqlStr := "update article set title=?, mtime=?, tags=?, abstract=?, cover=?, content=?, html_id=? where id=?"
	_, err = p.db.Exec(sqlStr, &article.Title, time.Now(), &article.Tags, &article.Abstract, &article.Cover, &article.Content, htmlId, &article.ID)
	return article.ID, err
}

func (p *Community) deleteMedias(ctx context.Context, uid, id string) (err error) {
	// get htmlIds
	var htmlIds []string
	sqlStr := "select html_id from article where id=? and user_id=?"
	rows, err := p.db.Query(sqlStr, id, uid)
	defer rows.Close()
	for rows.Next() {
		var htmlId string
		rows.Scan(&htmlId)
		htmlIds = append(htmlIds, htmlId)
	}
	err = p.DelMedias(ctx, uid, htmlIds)
	return
}

// DeleteArticle delete the article.
// func (p *Community) DeleteArticle(ctx context.Context, uid, id string) (err error) {
// 	// get htmlId
// 	err = p.deleteMedias(ctx, uid ,id)
// 	if err != nil{
// 		return
// 	}
// 	sqlStr := "delete from article where id=? and user_id=?"
// 	_, err = p.db.Exec(sqlStr, id, uid)
// 	return
// }

// DeleteArticle delete the article.
func (p *Community) DeleteArticle(ctx context.Context, uid, id string) (err error) {
	// Start a transaction
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			// Rollback if there's an error
			tx.Rollback()
		}
	}()

	// Delete medias (calling deleteMedias within the transaction)
	err = p.deleteMedias(ctx, uid, id)
	if err != nil {
		return err
	}

	// Delete article
	sqlStr := "delete from article where id=? and user_id=?"
	_, err = tx.ExecContext(ctx, sqlStr, id, uid)
	if err != nil {
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	return
}

// // DeleteArticles delete the articles by uid.
// func (p *Community) DeleteArticles(ctx context.Context, uid string) (err error) {
// 	// get htmlIds
// 	var htmlIds []string
// 	sqlStr := "select html_id from article where user_id=?"
// 	rows, err := p.db.Query(sqlStr, uid)
// 	defer rows.Close()
// 	for rows.Next() {
// 		var htmlId string
// 		err = rows.Scan(&htmlId)
// 		if err != nil {
// 			return
// 		}
// 		htmlIds = append(htmlIds, htmlId)
// 	}
// 	sqlStr = "delete from article where user_id=?"
// 	_, err = p.db.Exec(sqlStr, uid)
// 	// delete the media
// 	err = p.DelMedias(ctx, uid, htmlIds)
// 	return
// }

const (
	MarkBegin = ""
	MarkEnd   = "eof"
)

// Articles lists articles from a position.
func (p *Community) Articles(ctx context.Context, page int, limit int, searchValue string) (items []*ArticleEntry, total int, err error) {
	total, err = p.getTotal(ctx, searchValue)
	if err != nil || total == 0 {
		return []*ArticleEntry{}, 0, err
	}

	sqlStr := "select id, title, ctime, user_id, tags, abstract, cover from article order by ctime desc limit ? offset ?"
	rows, err := p.db.Query(sqlStr, limit, (page-1)*limit)
	if searchValue != "" {
		sqlStr := "select id, title, ctime, user_id, tags, abstract, cover from article where title like ? order by ctime desc limit ? offset ?"
		rows, err = p.db.Query(sqlStr, "%"+searchValue+"%", limit, (page-1)*limit)
	}
	if err != nil {
		return []*ArticleEntry{}, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		article := &ArticleEntry{}
		err := rows.Scan(&article.ID, &article.Title, &article.Ctime, &article.UId, &article.Tags, &article.Abstract, &article.Cover)
		if err != nil {
			return []*ArticleEntry{}, 0, err
		}
		// add author info
		user, err := p.GetUserById(article.UId)
		if err != nil {
			return []*ArticleEntry{}, 0, err
		}
		article.User = *user

		items = append(items, article)
	}
	return items, total, nil
}

// ListArticle lists articles from a position.
func (p *Community) ListArticle(ctx context.Context, from string, limit int, searchValue string) (items []*ArticleEntry, next string, err error) {
	if from == MarkBegin {
		from = "0"
	} else if from == MarkEnd {
		return []*ArticleEntry{}, from, nil
	}
	fromInt, err := strconv.Atoi(from)
	if err != nil {
		return []*ArticleEntry{}, from, err
	}
	
	sqlStr := "select id, title, ctime, user_id, tags, abstract, cover from article where title like ? order by ctime desc limit ? offset ?"
	rows, err := p.db.Query(sqlStr, "%"+searchValue+"%", limit, fromInt)
	if err != nil {
		return []*ArticleEntry{}, from, err
	}

	defer rows.Close()

	var rowLen int
	for rows.Next() {
		article := &ArticleEntry{}
		err := rows.Scan(&article.ID, &article.Title, &article.Ctime, &article.UId, &article.Tags, &article.Abstract, &article.Cover)
		if err != nil {
			return []*ArticleEntry{}, from, err
		}
		// add author info
		user, err := p.GetUserById(article.UId)
		if err != nil {
			return []*ArticleEntry{}, from, err
		}
		article.User = *user

		items = append(items, article)
		rowLen++
	}
	// have no article
	if rowLen == 0 {
		return []*ArticleEntry{}, MarkEnd, io.EOF
	}
	next = strconv.Itoa(fromInt + rowLen)
	return items, next, nil
}

// GetArticlesByUid get articles by user id.
func (p *Community) GetArticlesByUid(ctx context.Context, uid string) (items []*ArticleEntry, err error) {
	sqlStr := "select id, title, ctime, user_id, tags, abstract, cover from article where user_id = ?"
	rows, err := p.db.Query(sqlStr, uid)
	if err != nil {
		return []*ArticleEntry{}, err
	}
	defer rows.Close()

	for rows.Next() {
		article := &ArticleEntry{}
		err := rows.Scan(&article.ID, &article.Title, &article.Ctime, &article.UId, &article.Tags, &article.Abstract, &article.Cover)
		if err != nil {
			return []*ArticleEntry{}, err
		}
		// add author info
		user, err := p.GetUserById(article.UId)
		if err != nil {
			return []*ArticleEntry{}, err
		}
		article.User = *user

		items = append(items, article)
	}
	return items, nil
}

func casdoorConfigInit() *CasdoorConfig {
	endPoint := os.Getenv("GOP_CASDOOR_ENDPOINT")
	clientID := os.Getenv("GOP_CASDOOR_CLIENTID")
	clientSecret := os.Getenv("GOP_CASDOOR_CLIENTSECRET")
	certificate := os.Getenv("GOP_CASDOOR_CERTIFICATE")
	organizationName := os.Getenv("GOP_CASDOOR_ORGANIZATIONNAME")
	applicationName := os.Getenv("GOP_CASDOOR_APPLICATONNAME")

	casdoorsdk.InitConfig(endPoint, clientID, clientSecret, certificate, organizationName, applicationName)

	return &CasdoorConfig{
		endPoint:         endPoint,
		clientId:         clientID,
		clientSecret:     clientSecret,
		certificate:      certificate,
		organizationName: organizationName,
		applicationName:  applicationName,
	}
}

func (a *Community) RedirectToCasdoor(redirect string) (loginURL string) {
	// TODO: Check whitelist from referer
	responseType := "code"
	scope := "read"
	state := "casdoor"
	redirectEncodeURL := url.QueryEscape(redirect)

	loginURL = fmt.Sprintf(
		"%s/login/oauth/authorize?client_id=%s&response_type=%s&redirect_uri=%s&scope=%s&state=%s",
		a.casdoorConfig.endPoint,
		a.casdoorConfig.clientId,
		responseType,
		redirectEncodeURL,
		scope,
		state,
	)

	fmt.Println("loginURL", loginURL)

	return loginURL
}

func (a *Community) GetAccessToken(code, state string) (token *oauth2.Token, err error) {
	token, err = casdoorsdk.GetOAuthToken(code, state)
	if err != nil {
		a.xLog.Error(err)

		return nil, ErrNotExist
	}

	return token, nil
}
