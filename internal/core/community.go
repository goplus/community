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
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/goplus/community/translation"
	"github.com/qiniu/x/xlog"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
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

// S3Service is the interface for s3 service and easy to mock
type S3Service interface {
	NewReader(ctx context.Context, key string, opts *blob.ReaderOptions) (_ S3Reader, err error)
	NewWriter(ctx context.Context, key string, opts *blob.WriterOptions) (_ S3Writer, err error)
	Delete(ctx context.Context, key string) (err error)
}

type S3ServiceAdapter struct {
	bucket *blob.Bucket
}

func (s *S3ServiceAdapter) NewReader(ctx context.Context, key string, opts *blob.ReaderOptions) (_ S3Reader, err error) {
	if s.bucket == nil {
		return nil, errors.New("not implemented")
	}

	return s.bucket.NewReader(ctx, key, opts)
}

func (s *S3ServiceAdapter) NewWriter(ctx context.Context, key string, opts *blob.WriterOptions) (_ S3Writer, err error) {
	if s.bucket == nil {
		return nil, errors.New("not implemented")
	}

	return s.bucket.NewWriter(ctx, key, opts)
}

func (s *S3ServiceAdapter) Delete(ctx context.Context, key string) (err error) {
	if s.bucket == nil {
		return errors.New("not implemented")
	}

	return s.bucket.Delete(ctx, key)
}

type S3Reader interface {
	Close() error
	Read(p []byte) (int, error)
	ContentType() string
	Size() int64
}

type S3Writer interface {
	Close() error
	Write(p []byte) (n int, err error)
}

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
	VttId   string
}

type Community struct {
	db            *sql.DB
	domain        string
	casdoorConfig *CasdoorConfig
	xLog          *xlog.Logger
	translation   *Translation

	// Casdoor Service for mock
	CasdoorSDKService CasdoorSDKService
	S3Service         S3Service
	bucketName        string
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

	db, err := sql.Open(driver, dsn)
	if err != nil {
		xLog.Error(err)
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

	bucket, err := blob.OpenBucket(ctx, bus)
	if err != nil {
		xLog.Error(err)
	}
	bucketName, err := getBucketName(bus, "@", "?")
	if err != nil {
		xLog.Error("get bucket name error:", err.Error())
	}

	s3Service := &S3ServiceAdapter{bucket: bucket}
	casdoorSDKService := &CasdoorSDKServiceAdapter{}

	return &Community{db, domain, casdoorConf, xLog, translationEngine, casdoorSDKService, s3Service, bucketName}, nil
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
	err = p.db.QueryRow(sqlStr, id).Scan(&article.ID, &article.Title, &article.UId, &article.Cover, &article.Tags, &article.Abstract, &article.Content, &article.Ctime, &article.Mtime, &article.Label, &article.LikeCount, &article.ViewCount, &article.VttId)
	if err == sql.ErrNoRows {
		p.xLog.Warn("not found the article")
		return article, ErrNotExist
	} else if err != nil {
		p.xLog.Errorf("get article error: %v", err)
		return article, err
	}
	// vtt don't finished when adding
	if article.VttId != "" {
		save_vid := []string{}
		vids := strings.Split(article.VttId, ";")
		var fileKey string
		// var originKey string
		var status string
		for _, vid := range vids {
			row := p.db.QueryRow(`select output, status from video_task where resource_id = ?`, vid)
			err = row.Scan(&fileKey, &status)
			if err == sql.ErrNoRows {
				continue
			}
			// vtt finish
			if status == "1" {
				article.Content = strings.Replace(article.Content, "("+p.domain+vid+")", "("+p.domain+fileKey+")", -1)
				// article.Content = strings.Replace(article.Content, "("+p.domain+"origin"+vid+")", "("+p.domain+originKey+")", -1)
			} else {
				save_vid = append(save_vid, vid)
			}
		}
		if len(save_vid) != 0 && len(save_vid) != len(vids) {
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
		p.xLog.Errorf("get user error: %v", err)
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

// CanEditable determine whether the user has the permission to operate.
func (p *Community) CanEditable(ctx context.Context, uid, id string) (editable bool, err error) {
	sqlStr := "select id from article where id=? and user_id = ?"
	err = p.db.QueryRow(sqlStr, id, uid).Scan(&id)
	if err != nil {
		p.xLog.Errorf("CanEditable: %v", err)
		return false, ErrPermission
	}
	return true, nil
}

// PutArticle adds new article (ID == "") or edits an existing article (ID != "").
func (p *Community) PutArticle(ctx context.Context, uid string, article *Article) (id string, err error) {
	// new article
	if article.ID == "" {
		sqlStr := "insert into article (title, ctime, mtime, user_id, tags, abstract, cover, content, trans, label, vtt_id) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
		res, err := p.db.Exec(sqlStr, &article.Title, time.Now(), time.Now(), uid, &article.Tags, &article.Abstract, &article.Cover, &article.Content, &article.Trans, &article.Label, &article.VttId)
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
	_, err = p.db.Exec(sqlStr, &article.Title, time.Now(), &article.Tags, &article.Abstract, &article.Cover, &article.Content, &article.Trans, &article.Label, &article.VttId, &article.ID)
	return article.ID, err
}

// DeleteArticle delete the article.
func (p *Community) DeleteArticle(ctx context.Context, uid, id string) (err error) {
	// Start a transaction
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		p.xLog.Errorf("DeleteArticle: %v", err)
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

	// Delete article
	sqlStr := "delete from article where id=? and user_id=?"
	res, err := tx.ExecContext(ctx, sqlStr, id, uid)
	if err != nil {
		p.xLog.Errorf("DeleteArticle: %v", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		p.xLog.Errorf("DeleteArticle: %v", err)
		return err
	}
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

func (p *Community) RedirectToCasdoor(redirect string) (loginURL string) {
	// TODO: Check whitelist from referer
	responseType := "code"
	scope := "read"
	state := "casdoor"
	redirectEncodeURL := url.QueryEscape(redirect)

	loginURL = fmt.Sprintf(
		"%s/login/oauth/authorize?client_id=%s&response_type=%s&redirect_uri=%s&scope=%s&state=%s",
		p.casdoorConfig.EndPoint,
		p.casdoorConfig.ClientId,
		responseType,
		redirectEncodeURL,
		scope,
		state,
	)

	fmt.Println("loginURL", loginURL)

	return loginURL
}

func (p *Community) GetAccessToken(code, state string) (token *oauth2.Token, err error) {
	token, err = p.CasdoorSDKService.GetOAuthToken(code, state)
	if err != nil {
		p.xLog.Error(err)

		return nil, ErrNotExist
	}

	return token, nil
}

// share count
func (p *Community) Share(ip, platform, userId, articleId string) {
	go func(ip, platform, userId, articleId string) {
		sqlStr := "INSERT INTO share (platform,user_Id,ip,index_key,create_at) values (?,?,?,?,?)"
		index := ip + platform + articleId
		_, err := p.db.Exec(sqlStr, platform, userId, ip, index, time.Now())
		if err != nil {
			p.xLog.Fatalln(err.Error())
			return
		}
		p.xLog.Printf("user: %s, ip: %s, share to platform: %s, articleId: %s", userId, ip, platform, articleId)
	}(ip, platform, userId, articleId)
}

// get community application information
func (p *Community) GetApplicationInfo() (*casdoorsdk.Application, error) {
	// a2, err := p.CasdoorSDKService.GetApplication("application_x8aevk")
	// TODO: Check env
	a2, err := p.CasdoorSDKService.GetApplication(os.Getenv("GOP_CASDOOR_APPLICATONNAME"))
	if err != nil {
		p.xLog.Error(err)
	}
	return a2, err
}
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

func (p *Community) GetClientIP(r *http.Request) string {
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

func (p *Community) ArticleLike(ctx context.Context, articleId int, userId string) (bool, error) {
	db := p.db
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
