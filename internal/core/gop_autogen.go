package core

import (
	"os"
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
//line internal/core/community.gop:59:1
func (p *Community) Article(id string) (article *Article, err error) {
//line internal/core/community.gop:60:1
	if id == "123" {
//line internal/core/community.gop:61:1
		article = &Article{ArticleEntry{ID: id, Title: "Title"}, []byte(contentSummary)}
//line internal/core/community.gop:68:1
		return
	}
//line internal/core/community.gop:70:1
	return nil, ErrNotExist
}
// PutArticle adds new article (ID == "") or edits an existing article (ID != "").
//
//line internal/core/community.gop:74:1
func (p *Community) PutArticle(article *Article) (id string, err error) {
//line internal/core/community.gop:75:1
	return
}
// ListArticle lists articles from an position.
//
//line internal/core/community.gop:84:1
func (p *Community) ListArticle(from string, limit int) (items []*ArticleEntry, next string, err error) {
//line internal/core/community.gop:85:1
	if from == MarkBegin {
//line internal/core/community.gop:86:1
		item := &ArticleEntry{ID: "123", Title: "Title"}
//line internal/core/community.gop:90:1
		return []*ArticleEntry{item}, MarkEnd, nil
	}
//line internal/core/community.gop:92:1
	return nil, MarkEnd, io.EOF
}
//line internal/core/community.gop:47:1
func New(conf *Config) (*Community, error) {
//line internal/core/community.gop:48:1
	return &Community{}, nil
}
