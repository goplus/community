package main

import (
	"fmt"
	"os"
	"github.com/goplus/yap"
	"context"
	"time"
	"log"
	"net/http"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"golang.org/x/text/language"
)

type community struct {
	yap.App
	community *core.Community
	trans     *translation.Engine
}
//line cmd/gopcomm/community_yap.gox:19
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:19:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:21:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:22:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:23:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:24:1
		ctx.Yap__1("article", map[string]string{"ID": id, "Title": article.Title, "Content": article.Content})
	})
//line cmd/gopcomm/community_yap.gox:30:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:31:1
		ctx.Yap__1("add", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:34:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:35:1
		articles, _, _ := this.community.ListArticle(todo, core.MarkBegin, 20)
//line cmd/gopcomm/community_yap.gox:36:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:40:1
	this.Get("/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:41:1
		uid := "3"
//line cmd/gopcomm/community_yap.gox:42:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:43:1
		doc := map[string]string{"ID": id}
//line cmd/gopcomm/community_yap.gox:46:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:47:1
			if
//line cmd/gopcomm/community_yap.gox:47:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:49:1
				log.Println("no permissions")
//line cmd/gopcomm/community_yap.gox:50:1
				return
			}
//line cmd/gopcomm/community_yap.gox:52:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:53:1
			doc["Title"] = article.Title
//line cmd/gopcomm/community_yap.gox:54:1
			doc["Content"] = article.Content
		}
//line cmd/gopcomm/community_yap.gox:56:1
		ctx.Yap__1("edit", doc)
	})
//line cmd/gopcomm/community_yap.gox:58:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:59:1
		uid := "12"
//line cmd/gopcomm/community_yap.gox:60:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:61:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: "Sample Title", UId: "1", Cover: "sample_cover", Tags: "tag1", Ctime: time.Now(), Mtime: time.Now()}, Content: "Sample Markdown Content"}
//line cmd/gopcomm/community_yap.gox:76:1
		_, _ = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:77:1
		ctx.Yap__1("edit", *article)
	})
//line cmd/gopcomm/community_yap.gox:80:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:81:1
		mdData := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:82:1
		transData, _ := this.trans.TranslateMarkdownText(mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:83:1
		ctx.Json__1(map[string]string{"data": transData})
	})
//line cmd/gopcomm/community_yap.gox:88:1
	this.Get("/getMedia", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:89:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:91:1
		fileKey, _ := this.community.GetMediaUrl(todo, mediaId)
//line cmd/gopcomm/community_yap.gox:93:1
		ctx.Redirect("qiniu demain"+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:97:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:98:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:99:1
	os.Setenv("NIUTRANS_API_KEY", "6b7b501febb7b62b5424c1c761e0531c")
//line cmd/gopcomm/community_yap.gox:100:1
	this.trans = translation.New(os.Getenv("NIUTRANS_API_KEY"), "", "")
//line cmd/gopcomm/community_yap.gox:102:1
	fmt.Println("start")
//line cmd/gopcomm/community_yap.gox:103:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community))
}
