package core

import (
	"os"
	"context"
	"database/sql"
	"io"
	"time"
)

type Config struct {
	Driver string
	DSN    string
	BlobUS string
}
type ArticleEntry struct {
	ID    string
	Title string
	Ctime time.Time
	Mtime time.Time
}
type Article struct {
	ArticleEntry
	Content string
}
type Community struct {
	db *sql.DB
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
// Article returns an article.
//
//line internal/core/community.gop:84:1
func (p *Community) Article(ctx context.Context, id string) (article *Article, err error) {
//line internal/core/community.gop:85:1
	if id == "123" {
//line internal/core/community.gop:86:1
		article = &Article{ArticleEntry{ID: id, Title: "Title"}, contentSummary}
//line internal/core/community.gop:93:1
		return
	}
//line internal/core/community.gop:95:1
	return nil, ErrNotExist
}
// CanEditable
//
//line internal/core/community.gop:99:1
func (p *Community) CanEditable(ctx context.Context, uid string, id string) (editable bool, err error) {
//line internal/core/community.gop:100:1
	return true, nil
}
// PutArticle adds new article (ID == "") or edits an existing article (ID != "").
//
//line internal/core/community.gop:104:1
func (p *Community) PutArticle(ctx context.Context, uid string, article *Article) (id string, err error) {
//line internal/core/community.gop:105:1
	return
}
//line internal/core/community.gop:108:1
func (p *Community) DeleteArticle(ctx context.Context, uid string, id string) (err error) {
//line internal/core/community.gop:109:1
	return
}
// ListArticle lists articles from an position.
//
//line internal/core/community.gop:118:1
func (p *Community) ListArticle(ctx context.Context, from string, limit int) (items []*ArticleEntry, next string, err error) {
//line internal/core/community.gop:119:1
	if from == MarkBegin {
//line internal/core/community.gop:120:1
		item := &ArticleEntry{ID: "123", Title: "Title"}
//line internal/core/community.gop:124:1
		return []*ArticleEntry{item}, MarkEnd, nil
	}
//line internal/core/community.gop:126:1
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
//line internal/core/community.gop:53:1
func New(conf *Config) (ret *Community, err error) {
//line internal/core/community.gop:54:1
	if conf == nil {
//line internal/core/community.gop:55:1
		conf = new(Config)
	}
//line internal/core/community.gop:57:1
	driver := conf.Driver
//line internal/core/community.gop:58:1
	dsn := conf.DSN
//line internal/core/community.gop:59:1
	bus := conf.BlobUS
//line internal/core/community.gop:60:1
	if driver == "" {
//line internal/core/community.gop:61:1
		driver = "mysql"
	}
//line internal/core/community.gop:63:1
	if dsn == "" {
//line internal/core/community.gop:64:1
		dsn = os.Getenv("GOP_COMMUNITY_DSN")
	}
//line internal/core/community.gop:66:1
	if bus == "" {
//line internal/core/community.gop:67:1
		bus = os.Getenv("GOP_COMMUNITY_BLOBUS")
	}
//line internal/core/community.gop:69:1
	db, err := sql.Open(driver, dsn)
//line internal/core/community.gop:70:1
	if err != nil {
//line internal/core/community.gop:71:1
		return
	}
//line internal/core/community.gop:73:1
	return &Community{db}, nil
}
