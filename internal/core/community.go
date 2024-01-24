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
	"os"
	"strconv"
	"time"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"golang.org/x/oauth2"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gocloud.dev/blob"
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
	ID    string
	Title string
	UId   string
	Cover string
	Tags  string
	User  User
	Ctime time.Time
	Mtime time.Time
}

type Article struct {
	ArticleEntry
	Content string // in markdown
	// Status  int    // published or draft
	HtmlUrl  string // parsed html file url
	HtmlData string
}

type Community struct {
	bucket *blob.Bucket
	db     *sql.DB
	domain string
	casdoorConfig *CasdoorConfig
	zlog *zap.SugaredLogger
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

	zlog *zap.SugaredLogger
}
func New(ctx context.Context, conf *Config) (ret *Community, err error) {
	// Init log
	logger, _ := zap.NewProduction()
	zlog := logger.Sugar()

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
		zlog.Error(err)
		return
	}
	db, err := sql.Open(driver, dsn)
	if err != nil {
		zlog.Error(err)
		return
	}
	return &Community{bucket, db, domain,casdoorConf, zlog}, nil
}

// Article returns an article.
func (p *Community) Article(ctx context.Context, id string) (article *Article, err error) {
	article = &Article{}
	var htmlId string
	sqlStr := "select id,title,user_id,cover,tags,content,html_id,ctime,mtime from article where id=?"
	err = p.db.QueryRow(sqlStr, id).Scan(&article.ID, &article.Title, &article.UId, &article.Cover, &article.Tags, &article.Content, &htmlId, &article.Ctime, &article.Mtime)
	if err == sql.ErrNoRows {
		p.zlog.Warn("not found the article")
		return article, ErrNotExist
	} else if err != nil {
		return article, err
	}
	// add author info
	// user, err := p.GetUser(article.UId)
	// if err != nil {
	// 	return
	// }
	// article.User = *user
	// get html url
	fileKey, err := p.GetMediaUrl(ctx, htmlId)
	article.HtmlUrl = fmt.Sprintf("%s%s", p.domain, fileKey)
	if err != nil {
		p.zlog.Warn("have no html media")
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
		p.zlog.Warn("not found the translation html")
		return "", ErrNotExist
	}

	// get html url
	fileKey, err := p.GetMediaUrl(ctx, htmlId)
	htmlUrl = fmt.Sprintf("%s%s", p.domain, fileKey)
	if err != nil {
		p.zlog.Warn("have no html media")
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
		sqlStr := "insert into article (title, ctime, mtime, user_id, tags, cover, content, html_id) values (?, ?, ?, ?, ?, ?, ?, ?)"
		res, err := p.db.Exec(sqlStr, &article.Title, time.Now(), time.Now(), uid, &article.Tags, &article.Cover, &article.Content, htmlId)
		if err != nil {
			return "", err
		}
		idInt, err := res.LastInsertId()
		return strconv.FormatInt(idInt, 10), nil
	}
	if trans != "" {
		// add article except html_id, content (trans)
		sqlStr := "update article set title=?, mtime=?, ctime=?, tags=?, cover=?, trans_content=?, trans_html_id=? where id=?"
		_, err = p.db.Exec(sqlStr, &article.Title, time.Now(), time.Now(), &article.Tags, &article.Cover, &article.Content, htmlId, &article.ID)
		return article.ID, err
	}

	// edit article
	sqlStr := "update article set title=?, mtime=?, tags=?, cover=?, content=?, html_id=? where id=?"
	_, err = p.db.Exec(sqlStr, &article.Title, time.Now(), &article.Tags, &article.Cover, &article.Content, htmlId, &article.ID)
	return article.ID, err
}

// DeleteArticle delete the article.
func (p *Community) DeleteArticle(ctx context.Context, uid, id string) (err error) {
	// get htmlId
	var htmlId string
	sqlStr := "select html_id from article where id=? and user_id=?"
	err = p.db.QueryRow(sqlStr, id, uid).Scan(&htmlId)
	if err != nil {
		return
	}
	sqlStr = "delete from article where id=? and user_id=?"
	_, err = p.db.Exec(sqlStr, id, uid)
	if err != nil {
		return
	}
	// delete the media
	err = p.DelMedia(ctx, uid, htmlId)
	return
}

// DeleteArticles delete the articles by uid.
func (p *Community) DeleteArticles(ctx context.Context, uid string) (err error) {
	// get htmlIds
	var htmlIds []string
	sqlStr := "select html_id from article where user_id=?"
	rows, err := p.db.Query(sqlStr, uid)
	defer rows.Close()
	for rows.Next() {
		var htmlId string
		err = rows.Scan(&htmlId)
		if err != nil {
			return
		}
		htmlIds = append(htmlIds, htmlId)
	}
	sqlStr = "delete from article where user_id=?"
	_, err = p.db.Exec(sqlStr, uid)
	// delete the media
	err = p.DelMedias(ctx, uid, htmlIds)
	return
}

const (
	MarkBegin = ""
	MarkEnd   = "eof"
)

// ListArticle lists articles from a position.
func (p *Community) ListArticle(ctx context.Context, from string, limit int) (items []*ArticleEntry, next string, err error) {
	p.zlog.Info("sss")
	if from == MarkBegin {
		from = "0"
	} else if from == MarkEnd {
		return []*ArticleEntry{}, from, nil
	}
	fromInt, err := strconv.Atoi(from)
	if err != nil {
		return []*ArticleEntry{}, from, err
	}

	sqlStr := "select id, title, ctime, user_id, tags, cover from article order by ctime desc limit ? offset ?"
	rows, err := p.db.Query(sqlStr, limit, fromInt)
	if err != nil {
		return []*ArticleEntry{}, from, err
	}
	defer rows.Close()

	var rowLen int
	for rows.Next() {
		article := &ArticleEntry{}
		err := rows.Scan(&article.ID, &article.Title, &article.Ctime, &article.UId, &article.Tags, &article.Cover)
		if err != nil {
			p.zlog.Info("ArticleEntryArticleEntryArticleEntry", err)
			return []*ArticleEntry{}, from, err
		}
		// add author info
		// user, err := p.GetUser(article.UId)
		// if err != nil {
		// 	return []*ArticleEntry{}, from, err
		// }
		// article.User = *user

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

// SearchArticle search articles by title.
func (p *Community) SearchArticle(ctx context.Context, searchValue string) (items []*ArticleEntry, err error) {
	sqlStr := "select id, title, ctime, user_id, tags, cover from article where title like ?"
	rows, err := p.db.Query(sqlStr, "%"+searchValue+"%")
	if err != nil {
		return []*ArticleEntry{}, err
	}
	defer rows.Close()

	for rows.Next() {
		article := &ArticleEntry{}
		err := rows.Scan(&article.ID, &article.Title, &article.Ctime, &article.UId, &article.Tags, &article.Cover)
		if err != nil {
			return []*ArticleEntry{}, err
		}
		// add author info
		// user, err := p.GetUser(article.UId)
		// if err != nil {
		// 	return []*ArticleEntry{}, err
		// }
		// article.User = *user

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
	ResponseType := "code"
	Scope := "read"
	State := "casdoor"
	loginURL = fmt.Sprintf(
		"%s/login/oauth/authorize?client_id=%s&response_type=%s&redirect_uri=%s&scope=%s&state=%s",
		a.casdoorConfig.endPoint,
		a.casdoorConfig.clientId,
		ResponseType,
		redirect,
		Scope,
		State,
	)

	return loginURL
}

func (a *Community) GetAccessToken(code, state string) (token *oauth2.Token, err error) {
	token, err = casdoorsdk.GetOAuthToken(code, state)
	if err != nil {
		a.zlog.Error(err)

		return nil, ErrNotExist
	}

	return token, nil
}
