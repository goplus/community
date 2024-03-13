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
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"

	"github.com/goplus/community/translation"
	"github.com/qiniu/x/xlog"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	_ "github.com/go-sql-driver/mysql"
	"gocloud.dev/blob"
	"golang.org/x/oauth2"
	language "golang.org/x/text/language"
)

var (
	ErrNotExist   = os.ErrNotExist
	ErrPermission = os.ErrPermission
	Twitter       = "twitter"
	FaceBook      = "facebook"
	WeChat        = "wechat"
)

type Config struct {
	AppConfig     AppConfig
	DBConfig      DBConfig
	QiNiuConfig   QiNiuConfig
	CasdoorConfig CasdoorConfig
}

type AppConfig struct {
	EndPoint string
	Debug    bool
}

type DBConfig struct {
	Driver string
	DSN    string
}

type QiNiuConfig struct {
	AccessKey      string
	SecretKey      string
	BlobUS         string
	Domain         string
	TranslationKey string
}

type CasdoorConfig struct {
	EndPoint         string
	ClientId         string
	ClientSecret     string
	Certificate      string
	OrganizationName string
	ApplicationName  string
}

type PlatformCount struct {
	ArticleId string
	Platform  string
	ViewCount string
}

type ArticleEntry struct {
	ID            string
	Title         string
	UId           string
	Cover         string
	Tags          string
	User          User
	Abstract      string
	Label         string
	Ctime         time.Time
	Mtime         time.Time
	ViewCount     int
	LikeCount     int
	PlatformCount []PlatformCount
}

type ArticleLike struct {
	Id        int
	ArticleId int
	UserId    int
}

type Article struct {
	ArticleEntry
	Content string // in markdown
	Trans   bool
	Vtt_id  string
}

type Community struct {
	bucket        *blob.Bucket
	db            *sql.DB
	domain        string
	casdoorConfig *CasdoorConfig
	xLog          *xlog.Logger
	translation   *Translation
	bucketName    string
}

type Account struct {
}

type Translation struct {
	Engine         *translation.Engine
	VideoTaskCache *VideoTaskCache
}

func NewConfigFromEnv() *Config {
	return &Config{
		AppConfig: AppConfig{
			EndPoint: os.Getenv("GOP_COMMUNITY_ENDPOINT"),
			Debug:    os.Getenv("GOP_COMMUNITY_DEBUG") == "true",
		},
		DBConfig: DBConfig{
			Driver: os.Getenv("GOP_COMMUNITY_DRIVER"),
			DSN:    os.Getenv("GOP_COMMUNITY_DSN"),
		},
		QiNiuConfig: QiNiuConfig{
			AccessKey:      os.Getenv("GOP_COMMUNITY_ACCESSKEY"),
			SecretKey:      os.Getenv("GOP_COMMUNITY_SECRETKEY"),
			BlobUS:         os.Getenv("GOP_COMMUNITY_BLOBUS"),
			Domain:         os.Getenv("GOP_COMMUNITY_DOMAIN"),
			TranslationKey: os.Getenv("NIUTRANS_API_KEY"),
		},
		CasdoorConfig: CasdoorConfig{
			EndPoint:         os.Getenv("GOP_CASDOOR_ENDPOINT"),
			ClientId:         os.Getenv("GOP_CASDOOR_CLIENTID"),
			ClientSecret:     os.Getenv("GOP_CASDOOR_CLIENTSECRET"),
			Certificate:      os.Getenv("GOP_CASDOOR_CERTIFICATE"),
			OrganizationName: os.Getenv("GOP_CASDOOR_ORGANIZATIONNAME"),
			ApplicationName:  os.Getenv("GOP_CASDOOR_APPLICATONNAME"),
		},
	}
}

func New(ctx context.Context, conf *Config) (ret *Community, err error) {
	// Init log
	xLog := xlog.New("")

	if conf == nil {
		conf = new(Config)
	}

	// Init casdoor
	casdoorConfigInit(conf)
	casdoorConf := &conf.CasdoorConfig
	driver := conf.DBConfig.Driver
	dsn := conf.DBConfig.DSN
	bus := conf.QiNiuConfig.BlobUS
	domain := conf.QiNiuConfig.Domain
	// default
	if driver == "" {
		driver = "mysql"
	}

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

	// Init translation engine
	translationEngine := &Translation{
		Engine: translation.New(
			conf.QiNiuConfig.TranslationKey,
			conf.QiNiuConfig.AccessKey,
			conf.QiNiuConfig.SecretKey,
		),
		VideoTaskCache: NewVideoTaskCache(),
	}
	bucketName, err := getBucketName(bus, "@", "?")
	if err != nil {
		xLog.Error("get bucket name error:", err.Error())
		return
	}
	return &Community{bucket, db, domain, casdoorConf, xLog, translationEngine, bucketName}, nil
}

func getBucketName(url, startSymbol, endSymbol string) (string, error) {
	startIndex := strings.Index(url, startSymbol)
	if startIndex == -1 {
		return "", fmt.Errorf("start symbol '%s' not found", startSymbol)
	}

	startIndex += len(startSymbol)

	endIndex := strings.Index(url[startIndex:], endSymbol)
	if endIndex == -1 {
		return "", fmt.Errorf("end symbol '%s' not found", endSymbol)
	}

	return url[startIndex : startIndex+endIndex], nil
}

func (p *Community) TranslateMarkdownText(ctx context.Context, src string, from string, to language.Tag) (string, error) {
	return p.translation.Engine.TranslateMarkdownText(src, from, to)
}

func (p *Community) TranslateArticle(ctx context.Context, inputArticle *Article) (translatedArticle *Article, err error) {
	translatedArticle = inputArticle

	transTitle, _ := p.translation.Engine.TranslatePlainText(inputArticle.Title, "auto", language.English)
	translatedArticle.Title = transTitle

	transTags, _ := p.translation.Engine.TranslateBatchPlain(strings.Split(inputArticle.Tags, ","), "auto", language.English)
	translatedArticle.Tags = strings.Join(transTags, ";")

	transAbstract, _ := p.translation.Engine.TranslatePlainText(inputArticle.Abstract, "auto", language.English)
	translatedArticle.Abstract = transAbstract

	transContent, _ := p.translation.Engine.TranslateMarkdownText(inputArticle.Content, "auto", language.English)
	translatedArticle.Content = transContent

	return translatedArticle, nil
}

func (p *Community) GetTranslateArticle(ctx context.Context, id string) (article *Article, err error) {
	article = &Article{}
	var content interface{}
	sqlStr := "select trans_content, trans_title, trans_tags, trans_abstract from article where id=?"
	err = p.db.QueryRow(sqlStr, id).Scan(&content, &article.Title, &article.Tags, &article.Abstract)
	if content == "" || content == nil || reflect.ValueOf(content).Len() == 0 {
		sqlStr := "select content, title, tags, abstract from article where id=?"
		err = p.db.QueryRow(sqlStr, id).Scan(&article.Content, &article.Title, &article.Tags, &article.Abstract)
		if err != nil {
			return &Article{}, err
		}
		// database has no tranlation article, so we must translate it
		article, _ = p.TranslateArticle(ctx, article)
		// save transltation result
		sqlStr = "update article set trans_content=?, trans_title=?, trans_tags=?, trans_abstract=? where id=?"
		_, err = p.db.Exec(sqlStr, &article.Content, &article.Title, &article.Tags, &article.Abstract, id)
		return
	}
	contentStr, _ := content.([]byte)
	article.Content = string(contentStr)
	return
}

// Article returns an article.
func (p *Community) Article(ctx context.Context, id string) (article *Article, err error) {
	article = &Article{}
	// var htmlId string
	sqlStr := "select id,title,user_id,cover,tags,abstract,content,ctime,mtime,label,like_count,view_count,vtt_id from article where id=?"
	err = p.db.QueryRow(sqlStr, id).Scan(&article.ID, &article.Title, &article.UId, &article.Cover, &article.Tags, &article.Abstract, &article.Content, &article.Ctime, &article.Mtime, &article.Label, &article.LikeCount, &article.ViewCount, &article.Vtt_id)
	if err == sql.ErrNoRows {
		p.xLog.Warn("not found the article")
		return article, ErrNotExist
	} else if err != nil {
		return article, err
	}
	// vtt don't finished when adding
	if article.Vtt_id != "" {
		save_vid := []string{}
		vids := strings.Split(article.Vtt_id, ";")
		var fileKey string
		var status string
		for _, vid := range vids {
			row := p.db.QueryRow(`select output, status from video_task where resource_id = ?`, vid)
			err = row.Scan(&fileKey, &status)
			if err != nil {
				return article, err
			}
			// vtt finish
			if status == "1" {
				article.Content = strings.Replace(article.Content, "("+p.domain+vid+")", "("+p.domain+fileKey+")", -1)
			} else {
				save_vid = append(save_vid, vid)
			}
		}
		if len(save_vid) != len(vids) {
			sqlStr := "update article set content=?, vtt_id=? where id=?"
			_, err := p.db.Exec(sqlStr, article.Content, strings.Join(save_vid, ";"), id)
			if err != nil {
				return article, err
			}
		}
	}
	// add author info
	user, err := p.GetUserById(article.UId)
	if err != nil {
		return
	}
	article.User = *user
	return
}

func (p *Community) ArticleLikeState(ctx context.Context, userId, articleId string) (bool, error) {
	if userId == "" {
		return false, nil
	}
	sqlStr := "select count(*) from article_like where article_id = ? and user_id = ?"
	var count int64
	if err := p.db.QueryRow(sqlStr, articleId, userId).Scan(&count); err != nil {
		return false, err
	}
	return count == 1, nil
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
	ext := mimetype.Detect([]byte(htmlStr)).String()
	htmlId, err := p.SaveMedia(ctx, uid, []byte(htmlStr), ext)
	if err != nil {
		return "", err
	}
	if id == "" {
		// save to database
		sqlStr := "insert into article (user_id, html_id, ctime, mtime, content) values (?, ?, ?)"
		res, err := p.db.Exec(sqlStr, uid, htmlId, time.Now(), time.Now(), mdData)
		if err != nil {
			return "", err
		}
		articleId, err := res.LastInsertId()
		if err != nil {
			return "", err
		}
		return strconv.FormatInt(articleId, 10), nil
	}
	sqlStr := "update article set content=?, html_id=?, mtime=? where id=?"
	_, err = p.db.Exec(sqlStr, mdData, htmlId, time.Now(), id)
	if err != nil {
		return "", err
	}
	return id, err
}

// PutArticle adds new article (ID == "") or edits an existing article (ID != "").
func (p *Community) PutArticle(ctx context.Context, uid string, article *Article) (id string, err error) {
	// new article
	if article.ID == "" {
		sqlStr := "insert into article (title, ctime, mtime, user_id, tags, abstract, cover, content, trans, label, vtt_id) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
		res, err := p.db.Exec(sqlStr, &article.Title, time.Now(), time.Now(), uid, &article.Tags, &article.Abstract, &article.Cover, &article.Content, &article.Trans, &article.Label, &article.Vtt_id)
		if err != nil {
			return "", err
		}
		idInt, err := res.LastInsertId()
		if err != nil {
			return "", err
		}
		return strconv.FormatInt(idInt, 10), nil
	}
	// edit article
	sqlStr := "update article set title=?, mtime=?, tags=?, abstract=?, cover=?, content=?, trans=?, label=?, vtt_id=? where id=?"
	_, err = p.db.Exec(sqlStr, &article.Title, time.Now(), &article.Tags, &article.Abstract, &article.Cover, &article.Content, &article.Trans, &article.Label, &article.Vtt_id, &article.ID)
	return article.ID, err
}

func (p *Community) deleteMedias(ctx context.Context, uid, id string) (err error) {
	// get htmlIds
	var htmlIds []string
	sqlStr := "select html_id from article where id=? and user_id=?"
	rows, err := p.db.Query(sqlStr, id, uid)
	defer func() {
		err = rows.Close()
		if err != nil {
			p.xLog.Error(err)
		}
	}()
	for rows.Next() {
		var htmlId string
		if err := rows.Scan(&htmlId); err != nil {
			return err
		}
		htmlIds = append(htmlIds, htmlId)
	}
	err = p.DelMedias(ctx, uid, htmlIds)
	return
}

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
			err = tx.Rollback()
			if err != nil {
				p.xLog.Error(err)
			}
		}
	}()

	// Delete medias (calling deleteMedias within the transaction)
	err = p.deleteMedias(ctx, uid, id)
	if err != nil {
		return err
	}

	// Delete article
	sqlStr := "delete from article where id=? and user_id=?"
	res, err := tx.ExecContext(ctx, sqlStr, id, uid)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows != 1 {
		return fmt.Errorf("no need to delete")
	}

	// Commit the transaction
	err = tx.Commit()
	return
}

const (
	MarkBegin = ""
	MarkEnd   = "eof"
)

func (p *Community) getPageArticles(sqlStr string, from string, limit int, value string, label string, key string) (items []*ArticleEntry, next string, err error) {
	if from == MarkBegin {
		if key == "search" {
			from = "0"
		} else {
			from = "1"
		}
	} else if from == MarkEnd {
		return []*ArticleEntry{}, from, nil
	}
	fromInt, err := strconv.Atoi(from)
	if err != nil {
		return []*ArticleEntry{}, from, err
	}
	var rows *sql.Rows
	if key == "search" {
		rows, err = p.db.Query(sqlStr, value, value, label, limit, fromInt)
	} else {
		rows, err = p.db.Query(sqlStr, value, limit, (fromInt-1)*limit)
	}
	// rows, err := p.db.Query(sqlStr, value, value, label, limit, fromInt)
	if err != nil {
		return []*ArticleEntry{}, from, err
	}

	defer rows.Close()

	var rowLen int
	var articleIds []string
	for rows.Next() {
		article := &ArticleEntry{}
		err := rows.Scan(&article.ID, &article.Title, &article.Ctime, &article.UId, &article.Tags, &article.Abstract, &article.Cover, &article.Label, &article.LikeCount, &article.ViewCount)
		articleIds = append(articleIds, article.ID)
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
		return []*ArticleEntry{}, MarkEnd, nil
	}
	sqlStr = "SELECT article_id, platform, COUNT(*) AS view_count FROM article_view WHERE article_id IN (" + strings.Join(articleIds, ",") + ") AND platform IS NOT NULL AND platform <> '' GROUP BY article_id, platform;"
	rows, err = p.db.Query(sqlStr)
	if err != nil {
		return []*ArticleEntry{}, MarkEnd, nil
	}
	var m = make(map[string][]PlatformCount)
	for rows.Next() {
		p1 := PlatformCount{}
		err := rows.Scan(&p1.ArticleId, &p1.Platform, &p1.ViewCount)
		if err != nil {
			return []*ArticleEntry{}, from, err
		}
		if _, ok := m[p1.ArticleId]; !ok {
			m[p1.ArticleId] = make([]PlatformCount, 0)
		}
		m[p1.ArticleId] = append(m[p1.ArticleId], p1)
	}

	for _, article := range items {
		article.PlatformCount = m[article.ID]
	}

	if key == "search" {
		if rowLen < limit {
			return items, MarkEnd, nil
		}
		next = strconv.Itoa(fromInt + rowLen)
	} else {
		err = p.db.QueryRow("select count(*) from article where user_id = ?", value).Scan(&next)
		if err != nil {
			return items, next, nil
		}
	}

	return items, next, nil
}

// ListArticle lists articles from a position.
func (p *Community) ListArticle(ctx context.Context, from string, limit int, searchValue string, label string) (items []*ArticleEntry, next string, err error) {
	sqlStr := "select id, title, ctime, user_id, tags, abstract, cover, label, like_count, view_count from article where (title like ? or tags like ?) and label = ? order by ctime desc limit ? offset ?"
	return p.getPageArticles(sqlStr, from, limit, "%"+searchValue+"%", label, "search")
}

// GetArticlesByUid get articles by user id.
func (p *Community) GetArticlesByUid(ctx context.Context, uid string, page string, limit int) (items []*ArticleEntry, next string, err error) {
	sqlStr := "select id, title, ctime, user_id, tags, abstract, cover, label, like_count, view_count from article where user_id = ? order by ctime desc limit ? offset ?"
	return p.getPageArticles(sqlStr, page, limit, uid, "", "user")
}

func casdoorConfigInit(conf *Config) {
	endPoint := conf.CasdoorConfig.EndPoint
	clientID := conf.CasdoorConfig.ClientId
	clientSecret := conf.CasdoorConfig.ClientSecret
	certificate := conf.CasdoorConfig.Certificate
	organizationName := conf.CasdoorConfig.OrganizationName
	applicationName := conf.CasdoorConfig.ApplicationName

	casdoorsdk.InitConfig(endPoint, clientID, clientSecret, certificate, organizationName, applicationName)
}

func (a *Community) RedirectToCasdoor(redirect string) (loginURL string) {
	// TODO: Check whitelist from referer
	responseType := "code"
	scope := "read"
	state := "casdoor"
	redirectEncodeURL := url.QueryEscape(redirect)

	loginURL = fmt.Sprintf(
		"%s/login/oauth/authorize?client_id=%s&response_type=%s&redirect_uri=%s&scope=%s&state=%s",
		a.casdoorConfig.EndPoint,
		a.casdoorConfig.ClientId,
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

// share count
func (a *Community) Share(ip, platform, userId, articleId string) {

	if platform == Twitter || platform == FaceBook || platform == WeChat {
		go func(ip, platform, userId, articleId string) {
			a.xLog.Printf("user: %s, ip: %s, share to platform: %s, articleId: %s", userId, ip, platform, articleId)
			if userId == "" {
				userId = "0"
			}
			sqlStr := "INSERT INTO share (platform,user_Id,ip,index_key,create_at,article_id) values (?,?,?,?,?,?)"
			index := ip + platform + articleId
			_, err := a.db.Exec(sqlStr, platform, userId, ip, index, time.Now(), articleId)
			if err != nil {
				return
			}

		}(ip, platform, userId, articleId)
	}
}

// get community application information
func (a *Community) GetApplicationInfo() (*casdoorsdk.Application, error) {
	a2, err := casdoorsdk.GetApplication("application_x8aevk")
	if err != nil {
		a.xLog.Error(err)
	}
	return a2, err
}

// todo Optimization: Use short links to transform the sharing function in the future
// Delete async update,Real-time update of the number of views is user-friendly
func (a *Community) ArticleLView(ctx context.Context, articleId, ip, userId, platform string) {
	if platform == Twitter || platform == FaceBook || platform == WeChat || platform == "" {
		a.xLog.Debugf("user: %s, ip: %s, share to platform: %s, articleId: %s", userId, ip, platform, articleId)
		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			userIdInt = 0
		}
		articleIdInt, err := strconv.Atoi(articleId)
		if err != nil {
			a.xLog.Error(err.Error())
			return
		}
		sqlStr := "INSERT INTO article_view (ip,user_id,article_id,created_at,`index`,platform) values (?,?,?,?,?,?)"
		index := articleId + userId + ip + platform
		if _, err = a.db.Exec(sqlStr, ip, userIdInt, articleIdInt, time.Now(), index, platform); err == nil {
			// success article views incr
			sqlStr = "update article set view_count = view_count + 1 where id=?"
			_, err = a.db.Exec(sqlStr, articleId)
			if err != nil {
				a.xLog.Error(err.Error())
				return
			}
		}
	}
}

func (a *Community) GetClientIP(r *http.Request) string {
	ip := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

func (a *Community) ArticleLike(ctx context.Context, articleId int, userId string) (bool, error) {
	db := a.db
	tx, _ := db.Begin()
	sqlStr := "insert into article_like (article_id,user_id) values (?,?)"
	_, err := tx.Exec(sqlStr, articleId, userId)
	var f bool = true
	var num int = 1
	if err != nil {
		sqlStr = "delete from article_like where article_id = ? and user_id = ?"
		_, err = tx.Exec(sqlStr, articleId, userId)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return false, err
			}
			return false, err
		}
		f = false
		num = -1
	}
	sqlStr = "update article set like_count = like_count+ ? where id=?"
	_, err = tx.Exec(sqlStr, num, articleId)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return false, err
		}
		return false, err
	}
	err = tx.Commit()
	return f, err
}
