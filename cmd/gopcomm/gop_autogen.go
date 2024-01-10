package main

import (
	"github.com/goplus/yap"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/markdown"
)

type community struct {
	yap.App
	community *core.Community
}
//line cmd/gopcomm/community_yap.gox:10
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:10:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:11:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:12:1
		article, _ := this.community.Article(nil, id)
//line cmd/gopcomm/community_yap.gox:13:1
		html, _ := markdown.Render(article.Content)
//line cmd/gopcomm/community_yap.gox:14:1
		ctx.Yap__1("article", map[string]string{"ID": id, "Title": article.Title, "Body": html})
	})
//line cmd/gopcomm/community_yap.gox:20:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:21:1
		articles, _, _ := this.community.ListArticle(nil, core.MarkBegin, 20)
//line cmd/gopcomm/community_yap.gox:22:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:26:1
	this.Get("/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:27:1
		uid := ""
//line cmd/gopcomm/community_yap.gox:28:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:29:1
		doc := map[string]string{"ID": id}
//line cmd/gopcomm/community_yap.gox:32:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:33:1
			if
//line cmd/gopcomm/community_yap.gox:33:1
			editable, _ := this.community.CanEditable(nil, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:35:1
				return
			}
//line cmd/gopcomm/community_yap.gox:37:1
			article, _ := this.community.Article(nil, id)
//line cmd/gopcomm/community_yap.gox:38:1
			doc["Title"] = article.Title
//line cmd/gopcomm/community_yap.gox:39:1
			doc["Content"] = article.Content
		}
//line cmd/gopcomm/community_yap.gox:41:1
		ctx.Yap__1("edit", doc)
	})
//line cmd/gopcomm/community_yap.gox:43:1
	this.Post("/commit", func(ctx *yap.Context) {
	})
//line cmd/gopcomm/community_yap.gox:47:1
	config := &core.Config{}
//line cmd/gopcomm/community_yap.gox:48:1
	this.community, _ = core.New(config)
//line cmd/gopcomm/community_yap.gox:50:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community))
}
