package main

import (
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
//line cmd/gopcomm/community_yap.gox:21
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:21:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:22:1
	endpoint := os.Getenv("GOP_COMMUNITY_ENDPOINT")
//line cmd/gopcomm/community_yap.gox:23:1
	domain := os.Getenv("GOP_COMMUNITY_DOMAIN")
//line cmd/gopcomm/community_yap.gox:24:1
	logger, _ := zap.NewProduction()
//line cmd/gopcomm/community_yap.gox:25:1
	defer logger.Sync()
//line cmd/gopcomm/community_yap.gox:26:1
	zlog := logger.Sugar()
//line cmd/gopcomm/community_yap.gox:28:1
	this.Static__0("/")
//line cmd/gopcomm/community_yap.gox:30:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:31:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:32:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:33:1
		ctx.Yap__1("article", map[string]interface {
		}{"ID": id, "Title": article.Title, "Content": article.Content, "Tags": article.Tags, "Cover": article.Cover, "Ctime": article.Ctime, "User": article.User})
	})
//line cmd/gopcomm/community_yap.gox:43:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:44:1
		ctx.Yap__1("markdown", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:47:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:48:1
		articles, _, _ := this.community.ListArticle(todo, core.MarkBegin, 20)
//line cmd/gopcomm/community_yap.gox:49:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:54:1
	this.Post("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:55:1
		searchValue := ctx.Param("search")
//line cmd/gopcomm/community_yap.gox:56:1
		if searchValue == "" {
//line cmd/gopcomm/community_yap.gox:57:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "err": "value can not be ''."})
		}
//line cmd/gopcomm/community_yap.gox:62:1
		articles, _ := this.community.SearchArticle(todo, searchValue)
//line cmd/gopcomm/community_yap.gox:63:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:68:1
	this.Get("/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:69:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:70:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:71:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:76:1
		uid, err := this.community.GetUserId(token.Value)
//line cmd/gopcomm/community_yap.gox:77:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:78:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:83:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:87:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:88:1
			if
//line cmd/gopcomm/community_yap.gox:88:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:90:1
				zlog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:91:1
				ctx.Json__1(map[string]interface {
				}{"code": 403, "err": err.Error()})
			}
//line cmd/gopcomm/community_yap.gox:96:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:103:1
			ctx.Yap__1("edit", article)
		}
//line cmd/gopcomm/community_yap.gox:105:1
		ctx.Json__1(map[string]interface {
		}{"code": 400, "err": "id is must."})
	})
//line cmd/gopcomm/community_yap.gox:111:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:112:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:113:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:114:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:115:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:120:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:127:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:129:1
		trans := ctx.Param("trans")
//line cmd/gopcomm/community_yap.gox:130:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:131:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:132:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:134:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:135:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:136:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:141:1
		uid, err := this.community.GetUserId(token.Value)
//line cmd/gopcomm/community_yap.gox:142:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:143:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:149:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags")}, Content: mdData, HtmlData: htmlData}
//line cmd/gopcomm/community_yap.gox:160:1
		id, _ = this.community.PutArticle(todo, uid, trans, article)
//line cmd/gopcomm/community_yap.gox:161:1
		article.ID = id
//line cmd/gopcomm/community_yap.gox:162:1
		ctx.Yap__1("edit", *article)
	})
//line cmd/gopcomm/community_yap.gox:166:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:168:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:169:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:170:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:175:1
		uid, err := this.community.GetUserId(token.Value)
//line cmd/gopcomm/community_yap.gox:176:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:177:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:182:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:183:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:184:1
		id, _ := this.community.SaveHtml(todo, uid, htmlData, mdData)
//line cmd/gopcomm/community_yap.gox:186:1
		transData, err := this.trans.TranslateMarkdownText(mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:187:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:188:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:193:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "id": id, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:200:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:201:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:203:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:205:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:207:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:208:1
		file, header, err := ctx.FormFile("file")
//line cmd/gopcomm/community_yap.gox:209:1
		filename := header.Filename
//line cmd/gopcomm/community_yap.gox:211:1
		ctx.ParseMultipartForm(10 << 20)
//line cmd/gopcomm/community_yap.gox:213:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:214:1
			zlog.Error("upload file error:", filename)
//line cmd/gopcomm/community_yap.gox:215:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:216:1
			return
		}
//line cmd/gopcomm/community_yap.gox:220:1
		dst, err := os.Create(filename)
//line cmd/gopcomm/community_yap.gox:221:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:222:1
			zlog.Error("create file error:", file)
//line cmd/gopcomm/community_yap.gox:223:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:224:1
			return
		}
//line cmd/gopcomm/community_yap.gox:226:1
		defer func() {
//line cmd/gopcomm/community_yap.gox:227:1
			file.Close()
//line cmd/gopcomm/community_yap.gox:228:1
			dst.Close()
//line cmd/gopcomm/community_yap.gox:229:1
			err = os.Remove(filename)
//line cmd/gopcomm/community_yap.gox:230:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:231:1
				zlog.Error("delete file error:", filename)
//line cmd/gopcomm/community_yap.gox:232:1
				return
			}
		}()
//line cmd/gopcomm/community_yap.gox:237:1
		_, err = io.Copy(dst, file)
//line cmd/gopcomm/community_yap.gox:238:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:239:1
			zlog.Error("copy file errer:", filename)
//line cmd/gopcomm/community_yap.gox:240:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:241:1
			return
		}
//line cmd/gopcomm/community_yap.gox:243:1
		bytes, err := os.ReadFile(filename)
//line cmd/gopcomm/community_yap.gox:244:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:245:1
			zlog.Error("read file errer:", filename)
//line cmd/gopcomm/community_yap.gox:246:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:247:1
			return
		}
//line cmd/gopcomm/community_yap.gox:249:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:250:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:251:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:256:1
		uid, err := this.community.GetUserId(token.Value)
//line cmd/gopcomm/community_yap.gox:257:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:258:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:264:1
		id, err := this.community.SaveMedia(context.Background(), uid, bytes)
//line cmd/gopcomm/community_yap.gox:265:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:266:1
			zlog.Error("save file", err.Error())
//line cmd/gopcomm/community_yap.gox:267:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:268:1
			return
		}
//line cmd/gopcomm/community_yap.gox:272:1
		ctx.JSON(200, id)
	})
//line cmd/gopcomm/community_yap.gox:275:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:276:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:277:1
	this.trans = translation.New(os.Getenv("NIUTRANS_API_KEY"), "", "")
//line cmd/gopcomm/community_yap.gox:279:1
	zlog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:280:1
	this.Run(endpoint)
}
func main() {
	yap.Gopt_App_Main(new(community))
}
