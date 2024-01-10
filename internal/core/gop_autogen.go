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
// ListArticle lists articles from an position.
//
//line internal/core/community.gop:90:1
func (p *Community) ListArticle(ctx context.Context, from string, limit int) (items []*ArticleEntry, next string, err error) {
//line internal/core/community.gop:91:1
	if from == MarkBegin {
//line internal/core/community.gop:92:1
		item := &ArticleEntry{ID: "123", Title: "Title"}
//line internal/core/community.gop:96:1
		return []*ArticleEntry{item}, MarkEnd, nil
	}
//line internal/core/community.gop:98:1
	return nil, MarkEnd, io.EOF
}
//line internal/core/community.gop:48:1
func New(conf *Config) (*Community, error) {
//line internal/core/community.gop:49:1
	return &Community{}, nil
}
