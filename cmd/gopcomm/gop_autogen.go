package main

import (
	"context"
	"fmt"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/yap"
	"time"
)

type community struct {
	yap.App
	community *core.Community
}

//line cmd/gopcomm/community_yap.gox:13
func (this *community) MainEntry() {
	//line cmd/gopcomm/community_yap.gox:13:1
	todo := context.TODO()
	//line cmd/gopcomm/community_yap.gox:15:1
	this.Get("/p/:id", func(ctx *yap.Context) {
		//line cmd/gopcomm/community_yap.gox:16:1
		id := ctx.Param("id")
		//line cmd/gopcomm/community_yap.gox:17:1
		article, _ := this.community.Article(todo, id)
		//line cmd/gopcomm/community_yap.gox:18:1
		ctx.Yap__1("article", map[string]string{"ID": id, "Title": article.Title, "Content": article.Content})
	})
	//line cmd/gopcomm/community_yap.gox:24:1
	this.Get("/", func(ctx *yap.Context) {
		//line cmd/gopcomm/community_yap.gox:25:1
		articles, _, _ := this.community.ListArticle(todo, 0, 20)
		//line cmd/gopcomm/community_yap.gox:26:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
	//line cmd/gopcomm/community_yap.gox:30:1
	this.Get("/edit", func(ctx *yap.Context) {
		//line cmd/gopcomm/community_yap.gox:31:1
		uid := ""
		//line cmd/gopcomm/community_yap.gox:32:1
		id := ctx.Param("id")
		//line cmd/gopcomm/community_yap.gox:33:1
		doc := map[string]string{"ID": id}
		//line cmd/gopcomm/community_yap.gox:36:1
		if id != "" {
			//line cmd/gopcomm/community_yap.gox:37:1
			if
			//line cmd/gopcomm/community_yap.gox:37:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
				//line cmd/gopcomm/community_yap.gox:39:1
				fmt.Println("no permissions")
				//line cmd/gopcomm/community_yap.gox:40:1
				return
			}
			//line cmd/gopcomm/community_yap.gox:42:1
			article, _ := this.community.Article(todo, id)
			//line cmd/gopcomm/community_yap.gox:43:1
			doc["Title"] = article.Title
			//line cmd/gopcomm/community_yap.gox:44:1
			doc["Content"] = article.Content
		}
		//line cmd/gopcomm/community_yap.gox:46:1
		ctx.Yap__1("edit", doc)
	})
	//line cmd/gopcomm/community_yap.gox:48:1
	this.Post("/commit", func(ctx *yap.Context) {
		//line cmd/gopcomm/community_yap.gox:49:1
		uid := "12"
		//line cmd/gopcomm/community_yap.gox:50:1
		id := ctx.Param("id")
		//line cmd/gopcomm/community_yap.gox:51:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: "Sample Title", UId: "1", Cover: "sample_cover", Tags: "tag1", Ctime: time.Now(), Mtime: time.Now()}, Content: "Sample Markdown Content"}
		//line cmd/gopcomm/community_yap.gox:66:1
		_, _ = this.community.PutArticle(todo, uid, article)
		//line cmd/gopcomm/community_yap.gox:67:1
		ctx.Yap__1("edit", *article)
	})
	//line cmd/gopcomm/community_yap.gox:70:1
	conf := &core.Config{}
	//line cmd/gopcomm/community_yap.gox:71:1
	this.community, _ = core.New(todo, conf)
	//line cmd/gopcomm/community_yap.gox:73:1
	fmt.Println("start")
	//line cmd/gopcomm/community_yap.gox:74:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community))
}
