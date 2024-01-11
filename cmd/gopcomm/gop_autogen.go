package main

import (
	"fmt"
	"github.com/goplus/yap"
	"github.com/goplus/community/internal/core"
	"time"
)

type community struct {
	yap.App
	community *core.Community
}
//line cmd/gopcomm/community_yap.gox:11
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:11:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:12:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:13:1
		article, _ := this.community.Article(nil, id)
//line cmd/gopcomm/community_yap.gox:15:1
		ctx.Yap__1("article", map[string]string{"ID": id, "Title": article.Title, "Body": article.Content})
	})
//line cmd/gopcomm/community_yap.gox:21:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:22:1
		articles, _, _ := this.community.ListArticle(nil, 0, 20)
//line cmd/gopcomm/community_yap.gox:23:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:27:1
	this.Get("/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:28:1
		uid := "12"
//line cmd/gopcomm/community_yap.gox:29:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:30:1
		doc := map[string]string{"ID": id}
//line cmd/gopcomm/community_yap.gox:33:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:34:1
			if
//line cmd/gopcomm/community_yap.gox:34:1
			editable, _ := this.community.CanEditable(nil, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:36:1
				fmt.Println("no permissions")
//line cmd/gopcomm/community_yap.gox:37:1
				return
			}
//line cmd/gopcomm/community_yap.gox:39:1
			article, _ := this.community.Article(nil, id)
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
//line cmd/gopcomm/community_yap.gox:46:1
		uid := "12"
//line cmd/gopcomm/community_yap.gox:47:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:48:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: "Sample Title", UId: "1", Cover: "sample_cover", Tags: "tag1, tag2", Ctime: time.Now(), Mtime: time.Now()}, Status: 1, Content: "Sample Markdown Content", HtmlUrl: "/sample-html-url"}
//line cmd/gopcomm/community_yap.gox:63:1
		_, _ = this.community.PutArticle(nil, uid, article)
//line cmd/gopcomm/community_yap.gox:64:1
		ctx.Yap__1("edit", *article)
	})
//line cmd/gopcomm/community_yap.gox:67:1
	config := &core.Config{DNS: "root:password@tcp(127.0.0.1:3306)/community?parseTime=true"}
//line cmd/gopcomm/community_yap.gox:70:1
	this.community, _ = core.New(config)
//line cmd/gopcomm/community_yap.gox:71:1
	core.InitMysqlDB(config)
//line cmd/gopcomm/community_yap.gox:73:1
	fmt.Println("start")
//line cmd/gopcomm/community_yap.gox:74:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community))
}
