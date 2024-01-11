package core

import (
	"fmt"
	"os"
	"strconv"
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Config struct {
	DNS string
}
type ArticleEntry struct {
	ID    string
	Title string
	UId   string
	Cover string
	Tags  string
	Ctime time.Time
	Mtime time.Time
}
type Article struct {
	ArticleEntry
	Status  int
	Content string
	HtmlUrl string
}
type Community struct {
}

const (
	published = 1
	draft     = 0
)

var ErrNotExist = os.ErrNotExist
var DB *sql.DB
// Article returns an article.
//
//line internal/core/community.gop:87:1
func (p *Community) Article(ctx context.Context, id string) (article *Article, err error) {
//line internal/core/community.gop:88:1
	article = &Article{}
//line internal/core/community.gop:89:1
	sqlStr := "select id,title,user_id,status,content,html_url from article where id=?"
//line internal/core/community.gop:90:1
	err = DB.QueryRow(sqlStr, id).Scan(&article.ID, &article.Title, &article.UId, &article.Status, &article.Content, &article.HtmlUrl)
//line internal/core/community.gop:91:1
	if err != nil {
//line internal/core/community.gop:92:1
		fmt.Println("not found the article")
//line internal/core/community.gop:93:1
		return article, ErrNotExist
	}
//line internal/core/community.gop:96:1
	return
}
// CanEditable
//
//line internal/core/community.gop:100:1
func (p *Community) CanEditable(ctx context.Context, uid string, id string) (editable bool, err error) {
//line internal/core/community.gop:101:1
	sqlStr := "select id from article where id=? and user_id = ?"
//line internal/core/community.gop:102:1
	err = DB.QueryRow(sqlStr, id, uid).Scan(&id)
//line internal/core/community.gop:103:1
	if err != nil {
//line internal/core/community.gop:104:1
		return false, err
	}
//line internal/core/community.gop:106:1
	return true, nil
}
// PutArticle adds new article (ID == "") or edits an existing article (ID != "").
//
//line internal/core/community.gop:110:1
func (p *Community) PutArticle(ctx context.Context, uid string, article *Article) (id string, err error) {
//line internal/core/community.gop:112:1
	if article.ID == "" {
//line internal/core/community.gop:113:1
		sqlStr := "insert into article (title, ctime, mtime, user_id, tags, status, cover, content, html_url) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
//line internal/core/community.gop:114:1
		res, err := DB.Exec(sqlStr, &article.Title, time.Now(), time.Now(), &article.UId, &article.Tags, &article.Status, &article.Cover, &article.Content, &article.HtmlUrl)
//line internal/core/community.gop:115:1
		if err != nil {
//line internal/core/community.gop:116:1
			return "", err
		}
//line internal/core/community.gop:118:1
		idInt, err := res.LastInsertId()
//line internal/core/community.gop:119:1
		return strconv.FormatInt(idInt, 10), nil
	}
//line internal/core/community.gop:123:1
	canEdit, err := p.CanEditable(ctx, uid, article.ID)
//line internal/core/community.gop:124:1
	if !canEdit {
//line internal/core/community.gop:125:1
		fmt.Println("no permissions")
//line internal/core/community.gop:126:1
		return
	}
//line internal/core/community.gop:128:1
	sqlStr := "update article set title=?, mtime=?, tags=?, status=?, cover=?, content=?, html_url=? where id=?"
//line internal/core/community.gop:129:1
	_, err = DB.Exec(sqlStr, &article.Title, time.Now(), &article.Tags, &article.Status, &article.Cover, &article.Content, &article.HtmlUrl, &article.ID)
//line internal/core/community.gop:130:1
	return article.ID, err
}
//line internal/core/community.gop:133:1
func (p *Community) DeleteArticle(ctx context.Context, uid string, id string) (err error) {
//line internal/core/community.gop:135:1
	canEdit, err := p.CanEditable(ctx, uid, id)
//line internal/core/community.gop:136:1
	if !canEdit {
//line internal/core/community.gop:137:1
		fmt.Println("no permissions")
//line internal/core/community.gop:138:1
		return
	}
//line internal/core/community.gop:140:1
	sqlStr := "delete from article where id=?"
//line internal/core/community.gop:141:1
	_, err = DB.Exec(sqlStr, id)
//line internal/core/community.gop:143:1
	return
}
// ListArticle lists articles from an position.
//
//line internal/core/community.gop:152:1
func (p *Community) ListArticle(ctx context.Context, from int, limit int) (items []*ArticleEntry, next int, err error) {
//line internal/core/community.gop:153:1
	sqlStr := "select id, title, ctime, user_id, tags, cover from article where status = ? limit ? offset ?"
//line internal/core/community.gop:154:1
	rows, err := DB.Query(sqlStr, published, limit, from)
//line internal/core/community.gop:155:1
	if err != nil {
//line internal/core/community.gop:156:1
		fmt.Println(err)
//line internal/core/community.gop:157:1
		return []*ArticleEntry{}, from, err
	}
//line internal/core/community.gop:159:1
	defer rows.Close()
//line internal/core/community.gop:161:1
	var rowLen int
//line internal/core/community.gop:162:1
	for rows.Next() {
//line internal/core/community.gop:163:1
		article := &ArticleEntry{}
//line internal/core/community.gop:164:1
		err := rows.Scan(&article.ID, &article.Title, &article.Ctime, &article.UId, &article.Tags, &article.Cover)
//line internal/core/community.gop:165:1
		if err != nil {
//line internal/core/community.gop:166:1
			fmt.Println(err)
//line internal/core/community.gop:167:1
			return []*ArticleEntry{}, from, err
		}
//line internal/core/community.gop:170:1
		items = append(items, article)
//line internal/core/community.gop:171:1
		rowLen++
	}
//line internal/core/community.gop:173:1
	return items, from + rowLen, nil
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
//line internal/core/community.gop:60:1
func New(conf *Config) (*Community, error) {
//line internal/core/community.gop:61:1
	return &Community{}, nil
}
// InitMysqlDB init mysql connection.
//
//line internal/core/community.gop:65:1
func InitMysqlDB(conf *Config) {
//line internal/core/community.gop:66:1
	db, err := sql.Open("mysql", conf.DNS)
//line internal/core/community.gop:67:1
	if err != nil {
//line internal/core/community.gop:68:1
		fmt.Println(err)
//line internal/core/community.gop:69:1
		return
	}
//line internal/core/community.gop:72:1
	err = db.Ping()
//line internal/core/community.gop:73:1
	if err != nil {
//line internal/core/community.gop:74:1
		fmt.Println(err)
//line internal/core/community.gop:75:1
		return
	}
//line internal/core/community.gop:78:1
	db.SetMaxOpenConns(10)
//line internal/core/community.gop:79:1
	db.SetMaxIdleConns(10)
//line internal/core/community.gop:81:1
	DB = db
//line internal/core/community.gop:83:1
	fmt.Println("mysql connected")
}
