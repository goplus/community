package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/goplus/community/internal/core"
	"github.com/goplus/yap"
)

type community_yap struct {
	yap.App
	community *core.Community
}

//line cmd/gopcomm/community_yap.gox:12
func (this *community_yap) MainEntry() {
//line cmd/gopcomm/community_yap.gox:12:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:14:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:15:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:16:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:17:1
		ctx.Yap__1("article", map[string]string{"ID": id, "Title": article.Title, "Content": article.Content})
	})
//line cmd/gopcomm/community_yap.gox:23:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:24:1
		articles, _, _ := this.community.ListArticle(todo, core.MarkBegin, 20)
//line cmd/gopcomm/community_yap.gox:25:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:29:1
	this.Get("/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:30:1
		uid := ""
//line cmd/gopcomm/community_yap.gox:31:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:32:1
		doc := map[string]string{"ID": id}
//line cmd/gopcomm/community_yap.gox:35:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:36:1
			if
//line cmd/gopcomm/community_yap.gox:36:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:38:1
				return
			}
//line cmd/gopcomm/community_yap.gox:40:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:41:1
			doc["Title"] = article.Title
//line cmd/gopcomm/community_yap.gox:42:1
			doc["Content"] = article.Content
		}
//line cmd/gopcomm/community_yap.gox:44:1
		ctx.Yap__1("edit", doc)
	})
//line cmd/gopcomm/community_yap.gox:46:1
	this.Post("/commit", func(ctx *yap.Context) {
	})
//line cmd/gopcomm/community_yap.gox:50:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:51:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:53:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:55:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, "qiniu demain"+fileKey, http.StatusTemporaryRedirect)
		file, header, err := ctx.FormFile("file")
		filename := header.Filename

		ctx.ParseMultipartForm(10 << 20)

		if err != nil {
			log.Fatalln("upload file error:", filename)
			ctx.JSON(500, err.Error())
			return
		}

		dst, err := os.Create(filename)
		if err != nil {
			log.Fatalln("create file error:", file)
			ctx.JSON(500, err.Error())
			return
		}
		defer func() {
			file.Close()

			dst.Close()
			err = os.Remove(filename)
			if err != nil {
				log.Fatalln("delete file error:", filename)
				return
			}
		}()

		_, err = io.Copy(dst, file)
		if err != nil {
			log.Fatalln("copy file errer:", filename)
			ctx.JSON(500, err.Error())
			return
		}
		bytes, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalln("read file errer:", filename)
			ctx.JSON(500, err.Error())
			return
		}
		cookie, err := ctx.Request.Cookie("user_id")
		if err != nil {
			log.Fatalln("token不存在")
			ctx.JSON(500, err.Error())
			return
		}

		id, err := this.community.SaveMedia(context.Background(), cookie.Value, bytes)
		if err != nil {

			log.Fatalln("save file", err.Error())
			ctx.JSON(500, err.Error())
			return
		}
		// todo append current project ip and getMedia
		// sample: 127.0.0.1:8080/getMedia/ + id
		ctx.JSON(200, id)

	})

//line cmd/gopcomm/community_yap.gox:59:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:60:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:62:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community_yap))
}
