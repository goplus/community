package main

import (
	"os"
	"github.com/goplus/yap"
	"context"
	"io"
	"net/http"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
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
	this.Get("/success", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:32:1
		ctx.Yap__1("2xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:35:1
	this.Get("/error", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:36:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:39:1
	this.Get("/failed", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:40:1
		ctx.Yap__1("5xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:43:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:44:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:45:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:46:1
		ctx.Yap__1("article", map[string]interface {
		}{"ID": id, "Title": article.Title, "Content": article.Content, "Tags": article.Tags, "Cover": article.Cover, "Ctime": article.Ctime, "User": article.User})
	})
//line cmd/gopcomm/community_yap.gox:57:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:58:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:62:1
		// Get User Info
		var userClaim *casdoorsdk.Claims
//line cmd/gopcomm/community_yap.gox:63:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:64:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:65:1
			userClaim, err = this.community.GetUserClaim(token.Value)
//line cmd/gopcomm/community_yap.gox:66:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:67:1
				zlog.Error("get user claim error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:72:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": userClaim, "User": userClaim})
	})
//line cmd/gopcomm/community_yap.gox:79:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:80:1
		ctx.Yap__1("edit", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:83:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:87:1
		// Get User Info
		var userClaim *casdoorsdk.Claims
//line cmd/gopcomm/community_yap.gox:88:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:89:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:90:1
			userClaim, err = this.community.GetUserClaim(token.Value)
//line cmd/gopcomm/community_yap.gox:91:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:92:1
				zlog.Error("get user claim error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:97:1
		articles, _, _ := this.community.ListArticle(todo, core.MarkBegin, 20)
//line cmd/gopcomm/community_yap.gox:98:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": userClaim, "Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:104:1
	this.Post("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:105:1
		searchValue := ctx.Param("search")
//line cmd/gopcomm/community_yap.gox:106:1
		zlog.Infof("SearchValue: %+v", searchValue)
//line cmd/gopcomm/community_yap.gox:107:1
		if searchValue == "" {
//line cmd/gopcomm/community_yap.gox:108:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "err": "value can not be ''."})
		}
//line cmd/gopcomm/community_yap.gox:113:1
		articles, _ := this.community.SearchArticle(todo, searchValue)
//line cmd/gopcomm/community_yap.gox:114:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:119:1
	this.Get("/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:120:1
		_, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:121:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:122:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:127:1
		uid := "66de7a90-74c8-4641-a4dd-2c63f62350d0"
//line cmd/gopcomm/community_yap.gox:135:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:139:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:140:1
			if
//line cmd/gopcomm/community_yap.gox:140:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:142:1
				zlog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:143:1
				ctx.Json__1(map[string]interface {
				}{"code": 403, "err": err.Error()})
			}
//line cmd/gopcomm/community_yap.gox:148:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:155:1
			ctx.Yap__1("edit", article)
		}
//line cmd/gopcomm/community_yap.gox:157:1
		ctx.Json__1(map[string]interface {
		}{"code": 400, "err": "id is must."})
	})
//line cmd/gopcomm/community_yap.gox:163:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:164:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:165:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:166:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:167:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:172:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:179:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:181:1
		trans := ctx.Param("trans")
//line cmd/gopcomm/community_yap.gox:182:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:183:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:184:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:186:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:187:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:188:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:193:1
		uid, err := this.community.GetUserId(token.Value)
//line cmd/gopcomm/community_yap.gox:194:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:195:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:201:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags")}, Content: mdData, HtmlData: htmlData}
//line cmd/gopcomm/community_yap.gox:212:1
		id, _ = this.community.PutArticle(todo, uid, trans, article)
//line cmd/gopcomm/community_yap.gox:213:1
		article.ID = id
//line cmd/gopcomm/community_yap.gox:214:1
		ctx.Yap__1("edit", *article)
	})
//line cmd/gopcomm/community_yap.gox:218:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:220:1
		_, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:221:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:222:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:227:1
		uid := "66de7a90-74c8-4641-a4dd-2c63f62350d0"
//line cmd/gopcomm/community_yap.gox:235:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:236:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:237:1
		id, _ := this.community.SaveHtml(todo, uid, htmlData, mdData)
//line cmd/gopcomm/community_yap.gox:238:1
		zlog.Info("mdData", mdData)
//line cmd/gopcomm/community_yap.gox:240:1
		transData, err := this.trans.TranslateMarkdownText(mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:241:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:242:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:247:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "id": id, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:254:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:255:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:257:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:259:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:261:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:262:1
		file, header, err := ctx.FormFile("file")
//line cmd/gopcomm/community_yap.gox:263:1
		filename := header.Filename
//line cmd/gopcomm/community_yap.gox:265:1
		ctx.ParseMultipartForm(10 << 20)
//line cmd/gopcomm/community_yap.gox:267:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:268:1
			zlog.Error("upload file error:", filename)
//line cmd/gopcomm/community_yap.gox:269:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:270:1
			return
		}
//line cmd/gopcomm/community_yap.gox:274:1
		dst, err := os.Create(filename)
//line cmd/gopcomm/community_yap.gox:275:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:276:1
			zlog.Error("create file error:", file)
//line cmd/gopcomm/community_yap.gox:277:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:278:1
			return
		}
//line cmd/gopcomm/community_yap.gox:280:1
		defer func() {
//line cmd/gopcomm/community_yap.gox:281:1
			file.Close()
//line cmd/gopcomm/community_yap.gox:282:1
			dst.Close()
//line cmd/gopcomm/community_yap.gox:283:1
			err = os.Remove(filename)
//line cmd/gopcomm/community_yap.gox:284:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:285:1
				zlog.Error("delete file error:", filename)
//line cmd/gopcomm/community_yap.gox:286:1
				return
			}
		}()
//line cmd/gopcomm/community_yap.gox:291:1
		_, err = io.Copy(dst, file)
//line cmd/gopcomm/community_yap.gox:292:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:293:1
			zlog.Error("copy file errer:", filename)
//line cmd/gopcomm/community_yap.gox:294:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:295:1
			return
		}
//line cmd/gopcomm/community_yap.gox:297:1
		bytes, err := os.ReadFile(filename)
//line cmd/gopcomm/community_yap.gox:298:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:299:1
			zlog.Error("read file errer:", filename)
//line cmd/gopcomm/community_yap.gox:300:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:301:1
			return
		}
//line cmd/gopcomm/community_yap.gox:303:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:304:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:305:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:310:1
		uid, err := this.community.GetUserId(token.Value)
//line cmd/gopcomm/community_yap.gox:311:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:312:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:318:1
		id, err := this.community.SaveMedia(context.Background(), uid, bytes)
//line cmd/gopcomm/community_yap.gox:319:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:320:1
			zlog.Error("save file", err.Error())
//line cmd/gopcomm/community_yap.gox:321:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:322:1
			return
		}
//line cmd/gopcomm/community_yap.gox:326:1
		ctx.JSON(200, id)
	})
//line cmd/gopcomm/community_yap.gox:329:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:330:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:331:1
	this.trans = translation.New(os.Getenv("NIUTRANS_API_KEY"), "", "")
//line cmd/gopcomm/community_yap.gox:332:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:334:1
	zlog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:335:1
	this.Run(endpoint)
}
func main() {
	yap.Gopt_App_Main(new(community))
}
