package main

import (
	"github.com/goplus/yap"
	"github.com/goplus/community/internal/core"
)

type community struct {
	yap.App
	community *core.Community
}
//line cmd/gopcomm/community_yap.gox:9
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:9:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:10:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:11:1
		article, _ := this.community.Article(nil, id)
//line cmd/gopcomm/community_yap.gox:12:1
		ctx.Yap__1("article", map[string]string{"ID": id, "Title": article.Title, "Content": article.Content})
	})
//line cmd/gopcomm/community_yap.gox:18:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:19:1
		articles, _, _ := this.community.ListArticle(nil, core.MarkBegin, 20)
//line cmd/gopcomm/community_yap.gox:20:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:24:1
	this.Get("/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:25:1
		uid := ""
//line cmd/gopcomm/community_yap.gox:26:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:27:1
		doc := map[string]string{"ID": id}
//line cmd/gopcomm/community_yap.gox:30:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:31:1
			if
//line cmd/gopcomm/community_yap.gox:31:1
			editable, _ := this.community.CanEditable(nil, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:33:1
				return
			}
//line cmd/gopcomm/community_yap.gox:35:1
			article, _ := this.community.Article(nil, id)
//line cmd/gopcomm/community_yap.gox:36:1
			doc["Title"] = article.Title
//line cmd/gopcomm/community_yap.gox:37:1
			doc["Content"] = article.Content
		}
//line cmd/gopcomm/community_yap.gox:39:1
		ctx.Yap__1("edit", doc)
	})
//line cmd/gopcomm/community_yap.gox:41:1
	this.Post("/commit", func(ctx *yap.Context) {
	})
//line cmd/gopcomm/community_yap.gox:45:1
	config := &core.Config{}
//line cmd/gopcomm/community_yap.gox:46:1
	this.community, _ = core.New(config)
//line cmd/gopcomm/community_yap.gox:48:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community))
}
