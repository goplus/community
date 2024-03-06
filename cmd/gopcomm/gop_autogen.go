// Code generated by gop (Go+); DO NOT EDIT.

package main

import (
	"context"
	"fmt"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"github.com/goplus/yap"
	_ "github.com/joho/godotenv/autoload"
	"github.com/liuscraft/gop-casdoor-account-sdk"
	"github.com/qiniu/x/xlog"
	"golang.org/x/text/language"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
)

const _ = true
const (
	layoutUS   = "January 2, 2006"
	limitConst = 10
	labelConst = "article"
)

type community struct {
	yap.App
	community *core.Community
	trans     *translation.Engine
}
//line cmd/gopcomm/community_yap.gox:29
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:29:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:30:1
	endpoint := os.Getenv("GOP_COMMUNITY_ENDPOINT")
//line cmd/gopcomm/community_yap.gox:31:1
	domain := os.Getenv("GOP_COMMUNITY_DOMAIN")
//line cmd/gopcomm/community_yap.gox:32:1
	xLog := xlog.New("")
//line cmd/gopcomm/community_yap.gox:36:1
	this.Static__0("/static")
//line cmd/gopcomm/community_yap.gox:38:1
	this.Get("/success", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:39:1
		ctx.Yap__1("2xx", map[string]interface{}{})
	})
//line cmd/gopcomm/community_yap.gox:42:1
	this.Get("/error", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:43:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:44:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:45:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:46:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:47:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:48:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:49:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:51:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:54:1
		ctx.Yap__1("4xx", map[string]interface{}{"UserId": userId, "User": user})
	})
//line cmd/gopcomm/community_yap.gox:60:1
	this.Get("/failed", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:61:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:62:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:63:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:64:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:65:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:66:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:67:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:69:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:72:1
		ctx.Yap__1("5xx", map[string]interface{}{"UserId": userId, "User": user})
	})
//line cmd/gopcomm/community_yap.gox:79:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:81:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:82:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:83:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:84:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:85:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:86:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:87:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:89:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:93:1
		articles, next, err := this.community.ListArticle(todo, core.MarkBegin, limitConst, "", labelConst)
//line cmd/gopcomm/community_yap.gox:94:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:95:1
			xLog.Error("get article error:", err)
		}
//line cmd/gopcomm/community_yap.gox:98:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:99:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:101:1
		ctx.Yap__1("home", map[string]interface{}{"UserId": userId, "User": user, "Items": articles, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:112:1
	this.Get("/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:115:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:116:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:117:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:118:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:119:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:120:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:121:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:123:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:126:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:127:1
		article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:128:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:129:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:130:1
			return
		}
//line cmd/gopcomm/community_yap.gox:132:1
		likeState, err := this.community.ArticleLikeState(todo, userId, id)
//line cmd/gopcomm/community_yap.gox:133:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:134:1
			xLog.Error("article state err:", err)
//line cmd/gopcomm/community_yap.gox:135:1
			return
		}
//line cmd/gopcomm/community_yap.gox:137:1
		ip := this.community.GetClientIP(ctx.Request)
//line cmd/gopcomm/community_yap.gox:138:1
		this.community.ArticleLView(todo, id, ip, userId)
//line cmd/gopcomm/community_yap.gox:140:1
		ctx.Yap__1("article", map[string]interface{}{"UserId": userId, "User": user, "Article": article, "LikeState": likeState})
	})
//line cmd/gopcomm/community_yap.gox:148:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:149:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:150:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:151:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:152:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:153:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:154:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:155:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:157:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:160:1
		ctx.Yap__1("edit", map[string]interface{}{"User": user, "UserId": userId})
	})
//line cmd/gopcomm/community_yap.gox:166:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:167:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:168:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:169:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:170:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:171:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:172:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:173:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:175:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:179:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:180:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:181:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:182:1
			editable, err := this.community.CanEditable(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:183:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:184:1
				xLog.Error("can editable error:", err)
//line cmd/gopcomm/community_yap.gox:185:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:187:1
			if !editable {
//line cmd/gopcomm/community_yap.gox:188:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:189:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:191:1
			article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:192:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:193:1
				xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:194:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:197:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:198:1
				xLog.Error("json marshal error:", err)
//line cmd/gopcomm/community_yap.gox:199:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:201:1
			ctx.Yap__1("edit", map[string]interface{}{"UserId": userId, "User": user, "Article": article})
		}
	})
//line cmd/gopcomm/community_yap.gox:209:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:210:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:211:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:212:1
		if label == "" {
//line cmd/gopcomm/community_yap.gox:213:1
			label = "article"
		}
//line cmd/gopcomm/community_yap.gox:217:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:218:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:219:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:220:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:221:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:222:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:223:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:225:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:229:1
		articles, next, err := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue, label)
//line cmd/gopcomm/community_yap.gox:230:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:231:1
			xLog.Error("get article error:", err)
		}
//line cmd/gopcomm/community_yap.gox:234:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:235:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:237:1
		ctx.Yap__1("home", map[string]interface{}{"UserId": userId, "User": user, "Items": articles, "Value": searchValue, "Next": next, "Tab": label})
	})
//line cmd/gopcomm/community_yap.gox:250:1
	this.Get("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:251:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:252:1
		article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:253:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:254:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:255:1
			ctx.Json__1(map[string]interface{}{"code": 0, "err": "get article failed"})
		}
//line cmd/gopcomm/community_yap.gox:260:1
		ctx.Json__1(map[string]interface{}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:266:1
	this.Delete("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:267:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:268:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:269:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:270:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:271:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:272:1
			ctx.Json__1(map[string]interface{}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:277:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:278:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:279:1
			ctx.Json__1(map[string]interface{}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:284:1
			ctx.Json__1(map[string]interface{}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:291:1
	this.Get("/api/articles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:292:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:293:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:294:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:295:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:297:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:298:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:299:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:302:1
		articles, next, err := this.community.ListArticle(todo, from, limitInt, searchValue, label)
//line cmd/gopcomm/community_yap.gox:303:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:304:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:305:1
			ctx.Json__1(map[string]interface{}{"code": 0, "err": "get article failed"})
		}
//line cmd/gopcomm/community_yap.gox:310:1
		ctx.Json__1(map[string]interface{}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:318:1
	this.Post("/api/article/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:319:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:320:1
		content := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:321:1
		title := ctx.Param("title")
//line cmd/gopcomm/community_yap.gox:322:1
		tags := ctx.Param("tags")
//line cmd/gopcomm/community_yap.gox:323:1
		abstract := ctx.Param("abstract")
//line cmd/gopcomm/community_yap.gox:324:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:325:1
		trans, err := strconv.ParseBool(ctx.Param("trans"))
//line cmd/gopcomm/community_yap.gox:326:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:327:1
			xLog.Error("parse bool error:", err)
		}
//line cmd/gopcomm/community_yap.gox:330:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:331:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:332:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:333:1
			ctx.Json__1(map[string]interface{}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:338:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:339:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:340:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:341:1
			ctx.Json__1(map[string]interface{}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:346:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: title, UId: uid, Cover: ctx.Param("cover"), Tags: tags, Abstract: abstract, Label: label}, Content: content, Trans: trans, Vtt_id: ctx.Param("vtt_id")}
//line cmd/gopcomm/community_yap.gox:362:1
		if trans {
//line cmd/gopcomm/community_yap.gox:363:1
			article, _ = this.community.TranslateArticle(todo, article)
		}
//line cmd/gopcomm/community_yap.gox:366:1
		id, err = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:367:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:368:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:369:1
			ctx.Json__1(map[string]interface{}{"code": 0, "err": "add failed"})
		}
//line cmd/gopcomm/community_yap.gox:374:1
		ctx.Json__1(map[string]interface{}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:383:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:384:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:386:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:387:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:388:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:391:1
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:392:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:393:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:394:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:395:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:396:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:397:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:399:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:403:1
		items, next, err := this.community.GetArticlesByUid(todo, id, core.MarkBegin, limitConst)
//line cmd/gopcomm/community_yap.gox:404:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:405:1
			xLog.Error("get article list error:", err)
		}
//line cmd/gopcomm/community_yap.gox:407:1
		ctx.Yap__1("user", map[string]interface{}{"Id": id, "CurrentUser": userClaim, "User": user, "Items": items, "UserId": userId, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:417:1
	this.Get("/user/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:418:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:419:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:420:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:421:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:422:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:423:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:424:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:426:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:429:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:430:1
			xLog.Error("get token error:", err)
//line cmd/gopcomm/community_yap.gox:431:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:433:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:434:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:435:1
			xLog.Error("get uid error:", err)
//line cmd/gopcomm/community_yap.gox:436:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:438:1
		userClaim, err := this.community.GetUserClaim(uid)
//line cmd/gopcomm/community_yap.gox:439:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:440:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:442:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:443:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:444:1
			xLog.Info("gac", err)
//line cmd/gopcomm/community_yap.gox:445:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:447:1
		appInfo, err := this.community.GetApplicationInfo()
//line cmd/gopcomm/community_yap.gox:448:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:449:1
			xLog.Error("get application info error:", err)
//line cmd/gopcomm/community_yap.gox:450:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:452:1
		ctx.Yap__1("user_edit", map[string]interface{}{"UserId": userId, "User": user, "CurrentUser": userClaim, "Application": appInfo, "Binds": gac.GetProviderBindStatus()})
	})
//line cmd/gopcomm/community_yap.gox:464:1
	this.Put("/api/user", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:465:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:466:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:467:1
		user := &core.UserInfo{Id: uid, Name: ctx.Param("name"), Birthday: ctx.Param("birthday"), Gender: ctx.Param("gender"), Phone: ctx.Param("phone"), Email: ctx.Param("email"), Avatar: ctx.Param("avatar"), Owner: ctx.Param("owner"), DisplayName: ctx.Param("displayName")}
//line cmd/gopcomm/community_yap.gox:479:1
		_, err = this.community.UpdateUser(user)
//line cmd/gopcomm/community_yap.gox:480:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:481:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:482:1
			ctx.Json__1(map[string]interface{}{"code": 0, "msg": "update failed"})
		}
//line cmd/gopcomm/community_yap.gox:487:1
		ctx.Json__1(map[string]int{"code": 200})
	})
//line cmd/gopcomm/community_yap.gox:492:1
	this.Get("/api/user/unlink", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:493:1
		pv := ctx.Param("pv")
//line cmd/gopcomm/community_yap.gox:494:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:495:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:496:1
			ctx.Json__1(map[string]int{"code": 200})
//line cmd/gopcomm/community_yap.gox:499:1
			return
		}
//line cmd/gopcomm/community_yap.gox:501:1
		switch pv {
//line cmd/gopcomm/community_yap.gox:502:1
		case "Twitter":
//line cmd/gopcomm/community_yap.gox:503:1
		case "Facebook":
//line cmd/gopcomm/community_yap.gox:504:1
		case "Github":
//line cmd/gopcomm/community_yap.gox:505:1
		case "WeChat":
//line cmd/gopcomm/community_yap.gox:506:1
		default:
//line cmd/gopcomm/community_yap.gox:507:1
			pv = ""
		}
//line cmd/gopcomm/community_yap.gox:509:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:510:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:511:1
			gac.UnLink(pv)
		}
//line cmd/gopcomm/community_yap.gox:513:1
		ctx.Json__1(map[string]int{"code": 200})
	})
//line cmd/gopcomm/community_yap.gox:518:1
	this.Get("/api/user/:id/articles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:519:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:520:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:521:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:523:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:524:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:525:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:527:1
		items, next, err := this.community.GetArticlesByUid(todo, id, from, limitInt)
//line cmd/gopcomm/community_yap.gox:528:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:529:1
			xLog.Error("get article list error:", err)
//line cmd/gopcomm/community_yap.gox:530:1
			ctx.Json__1(map[string]interface{}{"code": 0, "err": err.Error(), "total": 0})
		}
//line cmd/gopcomm/community_yap.gox:536:1
		ctx.Json__1(map[string]interface{}{"code": 200, "items": items, "next": next})
	})
//line cmd/gopcomm/community_yap.gox:543:1
	this.Get("/api/user/:id/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:544:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:545:1
		uid := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:546:1
		page := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:547:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:549:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:550:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:551:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:553:1
		pageInt, err := strconv.Atoi(page)
//line cmd/gopcomm/community_yap.gox:554:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:555:1
			ctx.Json__1(map[string]interface{}{"code": 400, "total": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:561:1
		files, total, err := this.community.ListMediaByUserId(todo, uid, format, pageInt, limitInt)
//line cmd/gopcomm/community_yap.gox:562:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:563:1
			ctx.Json__1(map[string]interface{}{"code": 0, "total": total, "err": err.Error()})
		} else {
//line cmd/gopcomm/community_yap.gox:569:1
			ctx.Json__1(map[string]interface{}{"code": 200, "total": total, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:580:1
	this.Delete("/api/media/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:581:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:582:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:583:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:584:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:585:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:586:1
			ctx.Json__1(map[string]interface{}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:591:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:592:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:593:1
			ctx.Json__1(map[string]interface{}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:598:1
			ctx.Json__1(map[string]interface{}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:605:1
	this.Get("/api/translation/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:606:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:607:1
		article, err := this.community.GetTranslateArticle(todo, id)
//line cmd/gopcomm/community_yap.gox:609:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:610:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:611:1
			ctx.Json__1(map[string]int{"code": 0})
		}
//line cmd/gopcomm/community_yap.gox:615:1
		ctx.Json__1(map[string]interface{}{"code": 200, "content": article.Content, "tags": article.Tags, "title": article.Title})
	})
//line cmd/gopcomm/community_yap.gox:623:1
	this.Post("/api/translation", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:624:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:625:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, "auto", language.English)
//line cmd/gopcomm/community_yap.gox:626:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:627:1
			ctx.Json__1(map[string]interface{}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:632:1
		ctx.Json__1(map[string]interface{}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:638:1
	this.Get("/api/media/:id/url", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:639:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:640:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:641:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:642:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:643:1
			ctx.Json__1(map[string]interface{}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:648:1
		ctx.Json__1(map[string]interface{}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:654:1
	this.Get("/api/video/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:655:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:656:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:657:1
		m := make(map[string]string, 4)
//line cmd/gopcomm/community_yap.gox:658:1
		format, err := this.community.GetMediaType(todo, id)
//line cmd/gopcomm/community_yap.gox:659:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:660:1
			ctx.Json__1(map[string]interface{}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:665:1
		match, _ := regexp.MatchString("^video", format)
//line cmd/gopcomm/community_yap.gox:666:1
		if match {
//line cmd/gopcomm/community_yap.gox:667:1
			subtitle, status, err := this.community.GetVideoSubtitle(todo, id)
//line cmd/gopcomm/community_yap.gox:668:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:669:1
				if err != nil {
//line cmd/gopcomm/community_yap.gox:670:1
					ctx.Json__1(map[string]interface{}{"code": 500, "err": err.Error()})
				}
//line cmd/gopcomm/community_yap.gox:675:1
				return
			}
//line cmd/gopcomm/community_yap.gox:677:1
			m["subtitle"] = domain + subtitle
//line cmd/gopcomm/community_yap.gox:678:1
			m["status"] = status
		}
//line cmd/gopcomm/community_yap.gox:680:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:681:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:682:1
			ctx.Json__1(map[string]interface{}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:687:1
		m["fileKey"] = htmlUrl
//line cmd/gopcomm/community_yap.gox:688:1
		m["type"] = format
//line cmd/gopcomm/community_yap.gox:689:1
		ctx.Json__1(map[string]interface{}{"code": 200, "url": m})
	})
//line cmd/gopcomm/community_yap.gox:695:1
	this.Post("/api/media", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:696:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:699:1
	this.Post("/api/caption/task", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:700:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:701:1
		vid := ctx.Param("vid")
//line cmd/gopcomm/community_yap.gox:703:1
		if uid == "" || vid == "" {
//line cmd/gopcomm/community_yap.gox:704:1
			ctx.Json__1(map[string]interface{}{"code": 200, "msg": "Invalid param"})
		}
//line cmd/gopcomm/community_yap.gox:710:1
		if
//line cmd/gopcomm/community_yap.gox:710:1
		err := this.community.RetryCaptionGenerate(todo, uid, vid); err != nil {
//line cmd/gopcomm/community_yap.gox:711:1
			ctx.Json__1(map[string]interface{}{"code": 200, "msg": "Request task error"})
		}
//line cmd/gopcomm/community_yap.gox:717:1
		ctx.Json__1(map[string]interface{}{"code": 200, "msg": "Ok"})
	})
//line cmd/gopcomm/community_yap.gox:726:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:731:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:732:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:733:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:734:1
			return
		}
//line cmd/gopcomm/community_yap.gox:737:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:738:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:739:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:742:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "login/callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:744:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:745:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:748:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:749:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:750:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:751:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:754:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:755:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:756:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:757:1
			return
		}
//line cmd/gopcomm/community_yap.gox:760:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:761:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:762:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:765:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:768:1
	this.Get("/login/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:769:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:770:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:771:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:773:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:774:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:775:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:776:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:777:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:780:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:783:1
	this.Post("/api/article/like", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:784:1
		articleId := ctx.Param("articleId")
//line cmd/gopcomm/community_yap.gox:785:1
		articleIdInt, err := strconv.Atoi(articleId)
//line cmd/gopcomm/community_yap.gox:786:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:787:1
			ctx.Json__1(map[string]interface{}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:792:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:793:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:794:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:795:1
			ctx.Json__1(map[string]interface{}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:800:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:801:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:802:1
			ctx.Json__1(map[string]interface{}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:807:1
		res, err := this.community.ArticleLike(todo, articleIdInt, uid)
//line cmd/gopcomm/community_yap.gox:808:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:809:1
			ctx.Json__1(map[string]interface{}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:814:1
		ctx.Json__1(map[string]interface{}{"code": 200, "data": res})
	})
//line cmd/gopcomm/community_yap.gox:820:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:821:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:822:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:825:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:826:1
		ctx.Yap__1("4xx", map[string]interface{}{})
	})
//line cmd/gopcomm/community_yap.gox:829:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:832:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:834:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:835:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:837:1
				if
//line cmd/gopcomm/community_yap.gox:837:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:838:1
					xLog.Error(err)
//line cmd/gopcomm/community_yap.gox:839:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:843:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
