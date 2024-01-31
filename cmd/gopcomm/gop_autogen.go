package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/goplus/yap"
	"context"
	"io"
	"net/http"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"golang.org/x/text/language"
	_ "github.com/joho/godotenv/autoload"
	"github.com/qiniu/x/xlog"
)

const layoutUS = "January 2, 2006"

type community struct {
	yap.App
	community *core.Community
	trans     *translation.Engine
}
//line cmd/gopcomm/community_yap.gox:28
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:28:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:29:1
	endpoint := os.Getenv("GOP_COMMUNITY_ENDPOINT")
//line cmd/gopcomm/community_yap.gox:30:1
	domain := os.Getenv("GOP_COMMUNITY_DOMAIN")
//line cmd/gopcomm/community_yap.gox:31:1
	xLog := xlog.New("")
//line cmd/gopcomm/community_yap.gox:35:1
	this.Static__0("/static")
//line cmd/gopcomm/community_yap.gox:37:1
	this.Get("/success", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:38:1
		ctx.Yap__1("2xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:41:1
	this.Get("/error", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:42:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:45:1
	this.Get("/failed", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:46:1
		ctx.Yap__1("5xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:49:1
	this.Get("/demo", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:50:1
		ctx.Yap__1("demo", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:53:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:56:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:57:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:58:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:59:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:60:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:61:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:65:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:66:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:67:1
		ctx.Yap__1("article", map[string]interface {
		}{"User": user, "ID": id, "Title": article.Title, "Content": article.HtmlUrl, "Tags": article.Tags, "Cover": article.Cover, "Mtime": article.Mtime.Format(layoutUS), "Author": article.User})
	})
//line cmd/gopcomm/community_yap.gox:80:1
	this.Get("/getArticle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:81:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:82:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:83:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:89:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:90:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:92:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:93:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:94:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:98:1
		// todo middleware
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:99:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:100:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:101:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:102:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:103:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:107:1
		items, _ := this.community.GetArticlesByUid(todo, id)
//line cmd/gopcomm/community_yap.gox:108:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": userClaim, "User": user, "Items": items})
	})
//line cmd/gopcomm/community_yap.gox:116:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:117:1
		ctx.Yap__1("edit", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:120:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:122:1
		from := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:123:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:126:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:127:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:128:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:129:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:130:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:131:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:135:1
		page, err := strconv.Atoi(from)
//line cmd/gopcomm/community_yap.gox:136:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:137:1
			page = 1
		}
//line cmd/gopcomm/community_yap.gox:140:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:141:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:142:1
			limitInt = 20
		}
//line cmd/gopcomm/community_yap.gox:145:1
		articles, next, _ := this.community.Articles(todo, page, limitInt, "")
//line cmd/gopcomm/community_yap.gox:146:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": user, "Items": articles, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:153:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:154:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:155:1
		if searchValue == "" {
//line cmd/gopcomm/community_yap.gox:156:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "err": "value can not be ''."})
		}
//line cmd/gopcomm/community_yap.gox:162:1
		from := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:163:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:164:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:165:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:166:1
			limitInt = 10
		}
//line cmd/gopcomm/community_yap.gox:168:1
		page, err := strconv.Atoi(from)
//line cmd/gopcomm/community_yap.gox:169:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:170:1
			page = 1
		}
//line cmd/gopcomm/community_yap.gox:174:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:175:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:176:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:177:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:178:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:179:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:183:1
		articles, _ := this.community.SearchArticle(todo, searchValue)
//line cmd/gopcomm/community_yap.gox:184:1
		articles, total, _ := this.community.Articles(todo, page, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:185:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": user, "Items": articles, "Total": total, "Value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:193:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:194:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:195:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:196:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:202:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:203:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:204:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:210:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:211:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:212:1
			if
//line cmd/gopcomm/community_yap.gox:212:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:213:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:214:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:216:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:217:1
			ctx.Yap__1("edit", article)
		}
	})
//line cmd/gopcomm/community_yap.gox:221:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:222:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:223:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:224:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:225:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:230:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:237:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:239:1
		trans := ctx.Param("trans")
//line cmd/gopcomm/community_yap.gox:240:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:241:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:242:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:244:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:245:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:246:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:251:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:252:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:253:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:260:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags")}, Content: mdData, HtmlData: htmlData}
//line cmd/gopcomm/community_yap.gox:271:1
		id, _ = this.community.PutArticle(todo, uid, trans, article)
//line cmd/gopcomm/community_yap.gox:272:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:280:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:282:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:283:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:284:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:289:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:290:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:291:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:297:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:298:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:299:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:301:1
		transData, err := this.trans.TranslateMarkdownText(mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:302:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:303:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:308:1
		id, _ = this.community.SaveHtml(todo, uid, htmlData, mdData, id)
//line cmd/gopcomm/community_yap.gox:309:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "id": id, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:316:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:317:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:319:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:321:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:324:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:325:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:326:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:327:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:328:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:329:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:334:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:340:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:341:1
		file, header, err := ctx.FormFile("file")
//line cmd/gopcomm/community_yap.gox:342:1
		filename := header.Filename
//line cmd/gopcomm/community_yap.gox:344:1
		ctx.ParseMultipartForm(10 << 20)
//line cmd/gopcomm/community_yap.gox:346:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:347:1
			xLog.Error("upload file error:", filename)
//line cmd/gopcomm/community_yap.gox:348:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:349:1
			return
		}
//line cmd/gopcomm/community_yap.gox:353:1
		dst, err := os.Create(filename)
//line cmd/gopcomm/community_yap.gox:354:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:355:1
			xLog.Error("create file error:", file)
//line cmd/gopcomm/community_yap.gox:356:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:357:1
			return
		}
//line cmd/gopcomm/community_yap.gox:359:1
		defer func() {
//line cmd/gopcomm/community_yap.gox:360:1
			file.Close()
//line cmd/gopcomm/community_yap.gox:361:1
			dst.Close()
//line cmd/gopcomm/community_yap.gox:362:1
			err = os.Remove(filename)
//line cmd/gopcomm/community_yap.gox:363:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:364:1
				xLog.Error("delete file error:", filename)
//line cmd/gopcomm/community_yap.gox:365:1
				return
			}
		}()
//line cmd/gopcomm/community_yap.gox:370:1
		_, err = io.Copy(dst, file)
//line cmd/gopcomm/community_yap.gox:371:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:372:1
			xLog.Error("copy file errer:", filename)
//line cmd/gopcomm/community_yap.gox:373:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:374:1
			return
		}
//line cmd/gopcomm/community_yap.gox:376:1
		bytes, err := os.ReadFile(filename)
//line cmd/gopcomm/community_yap.gox:377:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:378:1
			xLog.Error("read file errer:", filename)
//line cmd/gopcomm/community_yap.gox:379:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:380:1
			return
		}
//line cmd/gopcomm/community_yap.gox:382:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:383:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:384:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:389:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:390:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:391:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:396:1
		id, err := this.community.SaveMedia(context.Background(), uid, bytes)
//line cmd/gopcomm/community_yap.gox:397:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:398:1
			xLog.Error("save file", err.Error())
//line cmd/gopcomm/community_yap.gox:399:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:400:1
			return
		}
//line cmd/gopcomm/community_yap.gox:404:1
		ctx.JSON(200, id)
	})
//line cmd/gopcomm/community_yap.gox:407:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:412:1
		redirectURL := fmt.Sprintf("%s/%s", ctx.Request.Referer(), "callback")
//line cmd/gopcomm/community_yap.gox:414:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:415:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:418:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:419:1
		tokenCookie, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:420:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:421:1
			xLog.Error("get token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:425:1
		tokenCookie.MaxAge = -1
//line cmd/gopcomm/community_yap.gox:426:1
		http.SetCookie(ctx.ResponseWriter, tokenCookie)
//line cmd/gopcomm/community_yap.gox:429:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("http://localhost:8080"), http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:432:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:433:1
		code := ctx.URL.Query().Get("code")
//line cmd/gopcomm/community_yap.gox:434:1
		state := ctx.URL.Query().Get("state")
//line cmd/gopcomm/community_yap.gox:436:1
		token, error := this.community.GetAccessToken(code, state)
//line cmd/gopcomm/community_yap.gox:437:1
		if error != nil {
//line cmd/gopcomm/community_yap.gox:438:1
			xLog.Error("err", error)
		}
//line cmd/gopcomm/community_yap.gox:441:1
		cookie := http.Cookie{Name: "token", Value: token.AccessToken, Path: "/", MaxAge: 3600}
//line cmd/gopcomm/community_yap.gox:447:1
		http.SetCookie(ctx.ResponseWriter, &cookie)
//line cmd/gopcomm/community_yap.gox:451:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("http://localhost:8080"), http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:454:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:455:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:456:1
	this.trans = translation.New(os.Getenv("NIUTRANS_API_KEY"), "", "")
//line cmd/gopcomm/community_yap.gox:457:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:460:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:461:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:464:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:467:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:469:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:470:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:471:1
				if
//line cmd/gopcomm/community_yap.gox:471:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:472:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:476:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
