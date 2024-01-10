package core

import (
	"os"
	"context"
	"io"
	"time"
)

type Config struct {
}
type ArticleEntry struct {
	ID    string
	Title string
	Ctime time.Time
	Mtime time.Time
}
type Article struct {
	ArticleEntry
	Content []byte
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
// Article returns an article.
//
//line internal/core/community.gop:60:1
func (p *Community) Article(ctx context.Context, id string) (article *Article, err error) {
//line internal/core/community.gop:61:1
	if id == "123" {
//line internal/core/community.gop:62:1
		article = &Article{ArticleEntry{ID: id, Title: "Title"}, []byte(contentSummary)}
//line internal/core/community.gop:69:1
		return
	}
//line internal/core/community.gop:71:1
	return nil, ErrNotExist
}
// CanEditable
//
//line internal/core/community.gop:75:1
func (p *Community) CanEditable(ctx context.Context, uid string, id string) (editable bool, err error) {
//line internal/core/community.gop:76:1
	return
}
// PutArticle adds new article (ID == "") or edits an existing article (ID != "").
//
//line internal/core/community.gop:80:1
func (p *Community) PutArticle(ctx context.Context, uid string, article *Article) (id string, err error) {
//line internal/core/community.gop:81:1
	return
}
//line internal/core/community.gop:84:1
func (p *Community) DeleteArticle(ctx context.Context, uid string, id string) (err error) {
//line internal/core/community.gop:85:1
	return
}
// ListArticle lists articles from an position.
//
//line internal/core/community.gop:94:1
func (p *Community) ListArticle(ctx context.Context, from string, limit int) (items []*ArticleEntry, next string, err error) {
//line internal/core/community.gop:95:1
	if from == MarkBegin {
//line internal/core/community.gop:96:1
		item := &ArticleEntry{ID: "123", Title: "Title"}
//line internal/core/community.gop:100:1
		return []*ArticleEntry{item}, MarkEnd, nil
	}
//line internal/core/community.gop:102:1
	return nil, MarkEnd, io.EOF
}
// PutMedia uploads media.
//
//line internal/core/media.gop:6:1
func (p *Community) PutMedia(ctx context.Context, uid string, media []byte) (id string, err error) {
}
//line internal/core/media.gop:9:1
func (p *Community) DeleteMedia(ctx context.Context, uid string, id string) (err error) {
//line internal/core/media.gop:10:1
	return
}
//line internal/core/media.gop:13:1
func (p *Community) MediaURL(id string) (url string) {
//line internal/core/media.gop:14:1
	return
}
//line internal/core/community.gop:48:1
func New(conf *Config) (*Community, error) {
//line internal/core/community.gop:49:1
	return &Community{}, nil
}
