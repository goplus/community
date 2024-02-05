package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/goplus/yap"
	"context"
	"encoding/json"
	"net/http"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"golang.org/x/text/language"
	_ "github.com/joho/godotenv/autoload"
	"github.com/qiniu/x/xlog"
)

const (
	layoutUS   = "January 2, 2006"
	limitConst = 10
)

type community struct {
	yap.App
	community *core.Community
	trans     *translation.Engine
}
//line cmd/gopcomm/community_yap.gox:31
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:31:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:32:1
	endpoint := os.Getenv("GOP_COMMUNITY_ENDPOINT")
//line cmd/gopcomm/community_yap.gox:33:1
	domain := os.Getenv("GOP_COMMUNITY_DOMAIN")
//line cmd/gopcomm/community_yap.gox:34:1
	xLog := xlog.New("")
//line cmd/gopcomm/community_yap.gox:38:1
	this.Static__0("/static")
//line cmd/gopcomm/community_yap.gox:40:1
	this.Get("/success", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:41:1
		ctx.Yap__1("2xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:44:1
	this.Get("/error", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:45:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:48:1
	this.Get("/failed", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:49:1
		ctx.Yap__1("5xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:52:1
	this.Get("/demo", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:53:1
		ctx.Yap__1("demo", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:56:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:59:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:60:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:61:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:62:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:63:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:64:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:68:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:69:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:70:1
		ctx.Yap__1("article", map[string]interface {
		}{"User": user, "ID": id, "Title": article.Title, "Content": article.HtmlUrl, "Tags": article.Tags, "Cover": article.Cover, "Mtime": article.Mtime.Format(layoutUS), "Author": article.User})
	})
//line cmd/gopcomm/community_yap.gox:83:1
	this.Get("/getArticle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:84:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:85:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:86:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:92:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:93:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:95:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:96:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:97:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:101:1
		// todo middleware
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:102:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:103:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:104:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:105:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:106:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:110:1
		items, _ := this.community.GetArticlesByUid(todo, id)
//line cmd/gopcomm/community_yap.gox:111:1
		userClaimJson, _ := json.Marshal(&userClaim)
//line cmd/gopcomm/community_yap.gox:112:1
		itemsJson, _ := json.Marshal(&items)
//line cmd/gopcomm/community_yap.gox:113:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": strings.Replace(string(userClaimJson), `\"`, `"`, -1), "User": user, "Items": strings.Replace(string(itemsJson), `\"`, `"`, -1)})
	})
//line cmd/gopcomm/community_yap.gox:121:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:122:1
		ctx.Yap__1("edit", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:125:1
	this.Get("/delete", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:126:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:127:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:128:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:129:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:130:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:131:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:136:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:137:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:138:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:143:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:150:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:152:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:153:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:154:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:155:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:156:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:157:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:161:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, "")
//line cmd/gopcomm/community_yap.gox:162:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:163:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:170:1
	this.Get("/get", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:171:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:172:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:173:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:175:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:176:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:177:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:180:1
		articles, next, _ := this.community.ListArticle(todo, from, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:182:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:190:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:191:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:192:1
		if searchValue == "" {
//line cmd/gopcomm/community_yap.gox:193:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "err": "value can not be ''."})
		}
//line cmd/gopcomm/community_yap.gox:200:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:201:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:202:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:203:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:204:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:205:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:209:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue)
//line cmd/gopcomm/community_yap.gox:210:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:211:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Value": searchValue, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:219:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:220:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:221:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:222:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:228:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:229:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:230:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:236:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:237:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:238:1
			if
//line cmd/gopcomm/community_yap.gox:238:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:239:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:240:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:242:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:243:1
			ctx.Yap__1("edit", article)
		}
	})
//line cmd/gopcomm/community_yap.gox:247:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:248:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:249:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:250:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:251:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:256:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:263:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:265:1
		trans := ctx.Param("trans")
//line cmd/gopcomm/community_yap.gox:266:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:267:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:268:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:270:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:271:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:272:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:277:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:278:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:279:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:286:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags"), Abstract: ctx.Param("abstract")}, Content: mdData, HtmlData: htmlData}
//line cmd/gopcomm/community_yap.gox:298:1
		id, _ = this.community.PutArticle(todo, uid, trans, article)
//line cmd/gopcomm/community_yap.gox:299:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:307:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:309:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:310:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:311:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:316:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:317:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:318:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:324:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:325:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:326:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:328:1
		transData, err := this.trans.TranslateMarkdownText(mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:329:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:330:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:335:1
		id, _ = this.community.SaveHtml(todo, uid, htmlData, mdData, id)
//line cmd/gopcomm/community_yap.gox:336:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "id": id, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:343:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:344:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:346:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:348:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:351:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:352:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:353:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:354:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:355:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:356:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:361:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:367:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:368:1
		core.UploadFile(ctx, this.community)
	})
//line cmd/gopcomm/community_yap.gox:371:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:376:1
		redirectURL := fmt.Sprintf("%s/%s", ctx.Request.Referer(), "callback")
//line cmd/gopcomm/community_yap.gox:378:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:379:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:383:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:384:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:385:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:386:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:390:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("/"), http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:393:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:394:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:395:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:396:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:401:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("/"), http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:404:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:405:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:406:1
	this.trans = translation.New(os.Getenv("NIUTRANS_API_KEY"), "", "")
//line cmd/gopcomm/community_yap.gox:407:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:410:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:411:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:414:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:417:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:419:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:420:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:421:1
				if
//line cmd/gopcomm/community_yap.gox:421:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:422:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:426:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
