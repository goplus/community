package main

import (
	"os"
	"github.com/goplus/yap"
	"context"
	"time"
	"io"
	"net/http"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	_ "github.com/joho/godotenv/autoload"
)

type community struct {
	yap.App
	community *core.Community
	trans     *translation.Engine
}
//line cmd/gopcomm/community_yap.gox:21
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:21:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:22:1
	endpoint := os.Getenv("GOP_COMMUNITY_ENDPOINT")
//line cmd/gopcomm/community_yap.gox:23:1
	logger, _ := zap.NewProduction()
//line cmd/gopcomm/community_yap.gox:24:1
	defer logger.Sync()
//line cmd/gopcomm/community_yap.gox:25:1
	zlog := logger.Sugar()
//line cmd/gopcomm/community_yap.gox:27:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:28:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:29:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:30:1
		ctx.Yap__1("article", map[string]string{"ID": id, "Title": article.Title, "Content": article.Content})
	})
//line cmd/gopcomm/community_yap.gox:36:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:37:1
		ctx.Yap__1("add", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:40:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:41:1
		articles, _, _ := this.community.ListArticle(todo, core.MarkBegin, 20)
//line cmd/gopcomm/community_yap.gox:42:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:46:1
	this.Get("/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:47:1
		uid := "3"
//line cmd/gopcomm/community_yap.gox:48:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:49:1
		doc := map[string]string{"ID": id}
//line cmd/gopcomm/community_yap.gox:52:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:53:1
			if
//line cmd/gopcomm/community_yap.gox:53:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:55:1
				zlog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:56:1
				return
			}
//line cmd/gopcomm/community_yap.gox:58:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:59:1
			doc["Title"] = article.Title
//line cmd/gopcomm/community_yap.gox:60:1
			doc["Content"] = article.Content
		}
//line cmd/gopcomm/community_yap.gox:62:1
		ctx.Yap__1("edit", doc)
	})
//line cmd/gopcomm/community_yap.gox:64:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:65:1
		uid := "12"
//line cmd/gopcomm/community_yap.gox:66:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:67:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: "Sample Title", UId: "1", Cover: "sample_cover", Tags: "tag1", Ctime: time.Now(), Mtime: time.Now()}, Content: "Sample Markdown Content"}
//line cmd/gopcomm/community_yap.gox:82:1
		_, _ = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:83:1
		ctx.Yap__1("edit", *article)
	})
//line cmd/gopcomm/community_yap.gox:87:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:88:1
		mdData := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:89:1
		transData, _ := this.trans.TranslateMarkdownText(mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:90:1
		ctx.Json__1(map[string]string{"data": transData})
	})
//line cmd/gopcomm/community_yap.gox:95:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:96:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:98:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:100:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, "qiniu demain"+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:102:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:103:1
		file, header, err := ctx.FormFile("file")
//line cmd/gopcomm/community_yap.gox:104:1
		filename := header.Filename
//line cmd/gopcomm/community_yap.gox:106:1
		ctx.ParseMultipartForm(10 << 20)
//line cmd/gopcomm/community_yap.gox:108:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:109:1
			zlog.Error("upload file error:", filename)
//line cmd/gopcomm/community_yap.gox:110:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:111:1
			return
		}
//line cmd/gopcomm/community_yap.gox:115:1
		dst, err := os.Create(filename)
//line cmd/gopcomm/community_yap.gox:116:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:117:1
			zlog.Error("create file error:", file)
//line cmd/gopcomm/community_yap.gox:118:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:119:1
			return
		}
//line cmd/gopcomm/community_yap.gox:121:1
		defer func() {
//line cmd/gopcomm/community_yap.gox:122:1
			file.Close()
//line cmd/gopcomm/community_yap.gox:123:1
			dst.Close()
//line cmd/gopcomm/community_yap.gox:124:1
			err = os.Remove(filename)
//line cmd/gopcomm/community_yap.gox:125:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:126:1
				zlog.Error("delete file error:", filename)
//line cmd/gopcomm/community_yap.gox:127:1
				return
			}
		}()
//line cmd/gopcomm/community_yap.gox:132:1
		_, err = io.Copy(dst, file)
//line cmd/gopcomm/community_yap.gox:133:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:134:1
			zlog.Error("copy file errer:", filename)
//line cmd/gopcomm/community_yap.gox:135:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:136:1
			return
		}
//line cmd/gopcomm/community_yap.gox:138:1
		bytes, err := os.ReadFile(filename)
//line cmd/gopcomm/community_yap.gox:139:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:140:1
			zlog.Error("read file errer:", filename)
//line cmd/gopcomm/community_yap.gox:141:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:142:1
			return
		}
//line cmd/gopcomm/community_yap.gox:144:1
		cookie, err := ctx.Request.Cookie("user_id")
//line cmd/gopcomm/community_yap.gox:145:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:146:1
			zlog.Error("token不存在")
//line cmd/gopcomm/community_yap.gox:147:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:148:1
			return
		}
//line cmd/gopcomm/community_yap.gox:151:1
		id, err := this.community.SaveMedia(context.Background(), cookie.Value, bytes)
//line cmd/gopcomm/community_yap.gox:152:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:153:1
			zlog.Error("save file", err.Error())
//line cmd/gopcomm/community_yap.gox:154:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:155:1
			return
		}
//line cmd/gopcomm/community_yap.gox:159:1
		ctx.JSON(200, id)
	})
//line cmd/gopcomm/community_yap.gox:162:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:163:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:164:1
	this.trans = translation.New(os.Getenv("NIUTRANS_API_KEY"), "", "")
//line cmd/gopcomm/community_yap.gox:166:1
	zlog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:167:1
	this.Run(endpoint)
}
func main() {
	yap.Gopt_App_Main(new(community))
}
