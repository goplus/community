package core

import (
	"fmt"
	"os"
	"context"
	"io"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
// Mysql config.
type Mysql struct {
	Ip       string
	Port     int
	Username string
	Password string
	Database string
}
type Config struct {
	Mysql
}
type ArticleEntry struct {
	ID    string
	Title string
	Ctime time.Time
	Mtime time.Time
}
type Article struct {
	ArticleEntry
	AccountId    string
	Tags         string
	Status       int
	CoverId      int
	MarkdownText string
	HtmlText     string
}
type Community struct {
}

const contentSummary = `
Content
====

Text body
`
const (
	MarkBegin = ""
	MarkEnd   = "eof"
)

var ErrNotExist = os.ErrNotExist
var DB *sql.DB
// Article returns an article.
//
//line internal/core/community.gop:127:1
func (p *Community) Article(ctx context.Context, id string) (article *Article, err error) {
//line internal/core/community.gop:128:1
	article = &Article{}
//line internal/core/community.gop:129:1
	sqlStr := "select id,title,account_id,status,markdown_text,html_text from article where id=?"
//line internal/core/community.gop:130:1
	err = DB.QueryRow(sqlStr, id).Scan(&article.ID, &article.Title, &article.AccountId, &article.Status, &article.MarkdownText, &article.HtmlText)
//line internal/core/community.gop:131:1
	if err != nil {
//line internal/core/community.gop:132:1
		fmt.Println("not found the article")
//line internal/core/community.gop:133:1
		return article, ErrNotExist
	}
//line internal/core/community.gop:135:1
	return
}
// CanEditable
//
//line internal/core/community.gop:139:1
func (p *Community) CanEditable(ctx context.Context, uid string, id string) (editable bool, err error) {
//line internal/core/community.gop:140:1
	return
}
// PutArticle adds new article (ID == "") or edits an existing article (ID != "").
//
//line internal/core/community.gop:144:1
func (p *Community) PutArticle(ctx context.Context, uid string, article *Article) (id string, err error) {
//line internal/core/community.gop:145:1
	return
}
//line internal/core/community.gop:148:1
func (p *Community) DeleteArticle(ctx context.Context, uid string, id string) (err error) {
//line internal/core/community.gop:149:1
	return
}
// ListArticle lists articles from an position.
//
//line internal/core/community.gop:158:1
func (p *Community) ListArticle(ctx context.Context, from string, limit int) (items []*ArticleEntry, next string, err error) {
//line internal/core/community.gop:159:1
	if from == MarkBegin {
//line internal/core/community.gop:160:1
		item := &ArticleEntry{ID: "123", Title: "Title"}
//line internal/core/community.gop:164:1
		return []*ArticleEntry{item}, MarkEnd, nil
	}
//line internal/core/community.gop:166:1
	return nil, MarkEnd, io.EOF
}
// PutMedia uploads media.
//
//line internal/core/media.gop:6:1
func (p *Community) PutMedia(ctx context.Context, uid string, media []byte) (id string, err error) {
//line internal/core/media.gop:7:1
	return
}
//line internal/core/media.gop:10:1
func (p *Community) DeleteMedia(ctx context.Context, uid string, id string) (err error) {
//line internal/core/media.gop:11:1
	return
}
//line internal/core/media.gop:14:1
func (p *Community) MediaURL(id string) (url string) {
//line internal/core/media.gop:15:1
	return
}
//line internal/core/community.gop:69:1
func New(conf *Config) (*Community, error) {
//line internal/core/community.gop:70:1
	return &Community{}, nil
}
// InitMysqlDB init mysql connection.
//
//line internal/core/community.gop:74:1
func InitMysqlDB(conf *Config) {
//line internal/core/community.gop:75:1
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Ip, conf.Mysql.Port, conf.Mysql.Database)
//line internal/core/community.gop:83:1
	db, err := sql.Open("mysql", dsn)
//line internal/core/community.gop:84:1
	if err != nil {
//line internal/core/community.gop:85:1
		fmt.Println(err)
//line internal/core/community.gop:86:1
		return
	}
//line internal/core/community.gop:89:1
	err = db.Ping()
//line internal/core/community.gop:90:1
	if err != nil {
//line internal/core/community.gop:91:1
		fmt.Println(err)
//line internal/core/community.gop:92:1
		return
	}
//line internal/core/community.gop:95:1
	db.SetMaxOpenConns(10)
//line internal/core/community.gop:96:1
	db.SetMaxIdleConns(10)
//line internal/core/community.gop:98:1
	DB = db
//line internal/core/community.gop:100:1
	fmt.Println("mysql connected")
}
