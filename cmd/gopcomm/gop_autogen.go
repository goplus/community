package main

import (
	"github.com/goplus/yap"
	"context"
	"github.com/goplus/community/internal/core"
)

type community struct {
	yap.App
	community *core.Community
}
//line cmd/gopcomm/community_yap.gox:11
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:11:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:13:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:14:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:15:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:16:1
		ctx.Yap__1("article", map[string]string{"ID": id, "Title": article.Title, "Content": article.Content})
	})
//line cmd/gopcomm/community_yap.gox:22:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:23:1
		articles, _, _ := this.community.ListArticle(todo, core.MarkBegin, 20)
//line cmd/gopcomm/community_yap.gox:24:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:28:1
	this.Get("/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:29:1
		uid := ""
//line cmd/gopcomm/community_yap.gox:30:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:31:1
		doc := map[string]string{"ID": id}
//line cmd/gopcomm/community_yap.gox:34:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:35:1
			if
//line cmd/gopcomm/community_yap.gox:35:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:37:1
				return
			}
//line cmd/gopcomm/community_yap.gox:39:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:40:1
			doc["Title"] = article.Title
//line cmd/gopcomm/community_yap.gox:41:1
			doc["Content"] = article.Content
		}
//line cmd/gopcomm/community_yap.gox:43:1
		ctx.Yap__1("edit", doc)
	})
//line cmd/gopcomm/community_yap.gox:45:1
	this.Post("/commit", func(ctx *yap.Context) {
	})
//line cmd/gopcomm/community_yap.gox:49:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:50:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:52:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community))
}
