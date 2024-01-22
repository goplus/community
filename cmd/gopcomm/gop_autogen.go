package main

import (
	"fmt"
	"os"
	"github.com/goplus/yap"
	"context"
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
//line cmd/gopcomm/community_yap.gox:22
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:22:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:23:1
	endpoint := os.Getenv("GOP_COMMUNITY_ENDPOINT")
//line cmd/gopcomm/community_yap.gox:24:1
	domain := os.Getenv("GOP_COMMUNITY_DOMAIN")
//line cmd/gopcomm/community_yap.gox:25:1
	logger, _ := zap.NewProduction()
//line cmd/gopcomm/community_yap.gox:26:1
	defer logger.Sync()
//line cmd/gopcomm/community_yap.gox:27:1
	zlog := logger.Sugar()
//line cmd/gopcomm/community_yap.gox:29:1
	this.Static__0("/")
//line cmd/gopcomm/community_yap.gox:31:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:32:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:33:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:34:1
		ctx.Yap__1("article", map[string]interface {
		}{"ID": id, "Title": article.Title, "Content": article.Content, "Tags": article.Tags, "Cover": article.Cover, "Ctime": article.Ctime, "User": article.User})
	})
//line cmd/gopcomm/community_yap.gox:45:1
	this.Get("/getArticle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:46:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:47:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:48:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:54:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:55:1
		ctx.Yap__1("edit", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:58:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:59:1
		articles, _, _ := this.community.ListArticle(todo, core.MarkBegin, 20)
//line cmd/gopcomm/community_yap.gox:60:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:65:1
	this.Post("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:66:1
		searchValue := ctx.Param("search")
//line cmd/gopcomm/community_yap.gox:67:1
		if searchValue == "" {
//line cmd/gopcomm/community_yap.gox:68:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "err": "value can not be ''."})
		}
//line cmd/gopcomm/community_yap.gox:73:1
		articles, _ := this.community.SearchArticle(todo, searchValue)
//line cmd/gopcomm/community_yap.gox:74:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:79:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:80:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:81:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:82:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:87:1
		uid, err := this.community.GetUserId(token.Value)
//line cmd/gopcomm/community_yap.gox:88:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:89:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:95:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:99:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:100:1
			if
//line cmd/gopcomm/community_yap.gox:100:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:102:1
				zlog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:103:1
				ctx.Json__1(map[string]interface {
				}{"code": 403, "err": "no permissions"})
			}
//line cmd/gopcomm/community_yap.gox:108:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:115:1
			ctx.Yap__1("edit", article)
		}
	})
//line cmd/gopcomm/community_yap.gox:119:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:120:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:121:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:122:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:123:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:128:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:135:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:137:1
		trans := ctx.Param("trans")
//line cmd/gopcomm/community_yap.gox:138:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:139:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:140:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:142:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:143:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:144:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:149:1
		uid, err := this.community.GetUserId(token.Value)
//line cmd/gopcomm/community_yap.gox:150:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:151:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:158:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags")}, Content: mdData, HtmlData: htmlData}
//line cmd/gopcomm/community_yap.gox:169:1
		id, _ = this.community.PutArticle(todo, uid, trans, article)
//line cmd/gopcomm/community_yap.gox:170:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:178:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:180:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:181:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:182:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:187:1
		uid, err := this.community.GetUserId(token.Value)
//line cmd/gopcomm/community_yap.gox:188:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:189:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:195:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:196:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:197:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:198:1
		zlog.Info(mdData)
//line cmd/gopcomm/community_yap.gox:200:1
		transData, err := this.trans.TranslateMarkdownText(mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:201:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:202:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:207:1
		id, _ = this.community.SaveHtml(todo, uid, htmlData, mdData, id)
//line cmd/gopcomm/community_yap.gox:208:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "id": id, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:215:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:216:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:218:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:220:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:223:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:224:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:225:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:226:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:227:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:228:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:233:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:239:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:240:1
		file, header, err := ctx.FormFile("file")
//line cmd/gopcomm/community_yap.gox:241:1
		filename := header.Filename
//line cmd/gopcomm/community_yap.gox:243:1
		ctx.ParseMultipartForm(10 << 20)
//line cmd/gopcomm/community_yap.gox:245:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:246:1
			zlog.Error("upload file error:", filename)
//line cmd/gopcomm/community_yap.gox:247:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:248:1
			return
		}
//line cmd/gopcomm/community_yap.gox:252:1
		dst, err := os.Create(filename)
//line cmd/gopcomm/community_yap.gox:253:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:254:1
			zlog.Error("create file error:", file)
//line cmd/gopcomm/community_yap.gox:255:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:256:1
			return
		}
//line cmd/gopcomm/community_yap.gox:258:1
		defer func() {
//line cmd/gopcomm/community_yap.gox:259:1
			file.Close()
//line cmd/gopcomm/community_yap.gox:260:1
			dst.Close()
//line cmd/gopcomm/community_yap.gox:261:1
			err = os.Remove(filename)
//line cmd/gopcomm/community_yap.gox:262:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:263:1
				zlog.Error("delete file error:", filename)
//line cmd/gopcomm/community_yap.gox:264:1
				return
			}
		}()
//line cmd/gopcomm/community_yap.gox:269:1
		_, err = io.Copy(dst, file)
//line cmd/gopcomm/community_yap.gox:270:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:271:1
			zlog.Error("copy file errer:", filename)
//line cmd/gopcomm/community_yap.gox:272:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:273:1
			return
		}
//line cmd/gopcomm/community_yap.gox:275:1
		bytes, err := os.ReadFile(filename)
//line cmd/gopcomm/community_yap.gox:276:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:277:1
			zlog.Error("read file errer:", filename)
//line cmd/gopcomm/community_yap.gox:278:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:279:1
			return
		}
//line cmd/gopcomm/community_yap.gox:281:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:282:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:283:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:288:1
		uid, err := this.community.GetUserId(token.Value)
//line cmd/gopcomm/community_yap.gox:289:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:290:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:297:1
		id, err := this.community.SaveMedia(context.Background(), uid, bytes)
//line cmd/gopcomm/community_yap.gox:298:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:299:1
			zlog.Error("save file", err.Error())
//line cmd/gopcomm/community_yap.gox:300:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:301:1
			return
		}
//line cmd/gopcomm/community_yap.gox:305:1
		ctx.JSON(200, id)
	})
//line cmd/gopcomm/community_yap.gox:308:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:309:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:310:1
	this.trans = translation.New(os.Getenv("NIUTRANS_API_KEY"), "", "")
//line cmd/gopcomm/community_yap.gox:312:1
	zlog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:313:1
	this.Run(endpoint)
}
func main() {
	yap.Gopt_App_Main(new(community))
}
