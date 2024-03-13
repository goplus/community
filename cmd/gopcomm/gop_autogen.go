package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"github.com/goplus/yap"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	gopaccountsdk "github.com/liuscraft/gop-casdoor-account-sdk"
	"github.com/qiniu/x/xlog"
	"golang.org/x/text/language"
)

const (
	limitConst      = 10
	mediaLimitConst = 8
	firstConst      = "1"
	labelConst      = "article"
)

type community struct {
	yap.App
	community *core.Community
	trans     *translation.Engine
	domain    string
	htmlUrl   string
}

//line cmd/gopcomm/community_yap.gox:37
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:37:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:38:1
	xLog := xlog.New("")
//line cmd/gopcomm/community_yap.gox:42:1
	this.Static__0("/static")
//line cmd/gopcomm/community_yap.gox:44:1
	this.Get("/success", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:45:1
		ctx.Yap__1("2xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:48:1
	this.Get("/error", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:49:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:50:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:51:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:52:1
			user, _ = this.community.GetUser(token.Value)
		}
//line cmd/gopcomm/community_yap.gox:55:1
		ctx.Yap__1("4xx", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:60:1
	this.Get("/failed", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:61:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:62:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:63:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:64:1
			user, _ = this.community.GetUser(token.Value)
		}
//line cmd/gopcomm/community_yap.gox:67:1
		ctx.Yap__1("5xx", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:73:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:75:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:76:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:77:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:78:1
			user, _ = this.community.GetUser(token.Value)
		}
//line cmd/gopcomm/community_yap.gox:81:1
		articles, next, err := this.community.ListArticle(todo, core.MarkBegin, limitConst, "", labelConst)
//line cmd/gopcomm/community_yap.gox:82:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:83:1
			xLog.Error("get article error:", err)
		}
//line cmd/gopcomm/community_yap.gox:86:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:87:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:89:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": user, "Items": articles, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:99:1
	this.Get("/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:102:1
		uid := ""
//line cmd/gopcomm/community_yap.gox:104:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:105:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:106:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:107:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:108:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:109:1
				uid = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:113:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:114:1
		platform := ctx.Param("platform")
//line cmd/gopcomm/community_yap.gox:115:1
		ip := this.community.GetClientIP(ctx.Request)
//line cmd/gopcomm/community_yap.gox:116:1
		this.community.ArticleLView(todo, id, ip, uid, platform)
//line cmd/gopcomm/community_yap.gox:117:1
		article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:118:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:119:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:120:1
			return
		}
//line cmd/gopcomm/community_yap.gox:122:1
		likeState, err := this.community.ArticleLikeState(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:123:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:124:1
			xLog.Error("article state err:", err)
//line cmd/gopcomm/community_yap.gox:125:1
			return
		}
//line cmd/gopcomm/community_yap.gox:128:1
		ctx.Yap__1("article", map[string]interface {
		}{"User": user, "Article": article, "LikeState": likeState})
	})
//line cmd/gopcomm/community_yap.gox:135:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:136:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:137:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:138:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:139:1
			user, _ = this.community.GetUser(token.Value)
		}
//line cmd/gopcomm/community_yap.gox:141:1
		ctx.Yap__1("edit", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:146:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:147:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:148:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:149:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:150:1
			user, _ = this.community.GetUser(token.Value)
		}
//line cmd/gopcomm/community_yap.gox:153:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:154:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:155:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:156:1
			editable, err := this.community.CanEditable(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:157:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:158:1
				xLog.Error("can editable error:", err)
//line cmd/gopcomm/community_yap.gox:159:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:161:1
			if !editable {
//line cmd/gopcomm/community_yap.gox:162:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:163:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:165:1
			article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:166:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:167:1
				xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:168:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:171:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:172:1
				xLog.Error("json marshal error:", err)
//line cmd/gopcomm/community_yap.gox:173:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:175:1
			ctx.Yap__1("edit", map[string]interface {
			}{"User": user, "Article": article})
		}
	})
//line cmd/gopcomm/community_yap.gox:182:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:183:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:184:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:185:1
		if label == "" {
//line cmd/gopcomm/community_yap.gox:186:1
			label = "article"
		}
//line cmd/gopcomm/community_yap.gox:190:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:191:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:192:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:193:1
			user, _ = this.community.GetUser(token.Value)
		}
//line cmd/gopcomm/community_yap.gox:196:1
		articles, next, err := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue, label)
//line cmd/gopcomm/community_yap.gox:197:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:198:1
			xLog.Error("get article error:", err)
		}
//line cmd/gopcomm/community_yap.gox:201:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:202:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:204:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": user, "Items": articles, "Value": searchValue, "Next": next, "Tab": label})
	})
//line cmd/gopcomm/community_yap.gox:216:1
	this.Get("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:217:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:218:1
		article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:219:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:220:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:221:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get article failed"})
		}
//line cmd/gopcomm/community_yap.gox:226:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:232:1
	this.Delete("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:233:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:234:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:235:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:236:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:237:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:238:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:243:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:244:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:245:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:250:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:257:1
	this.Get("/api/articles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:258:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:259:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:260:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:261:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:263:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:264:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:265:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:268:1
		articles, next, err := this.community.ListArticle(todo, from, limitInt, searchValue, label)
//line cmd/gopcomm/community_yap.gox:269:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:270:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:271:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get article failed"})
		}
//line cmd/gopcomm/community_yap.gox:276:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:284:1
	this.Post("/api/article/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:285:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:286:1
		content := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:287:1
		title := ctx.Param("title")
//line cmd/gopcomm/community_yap.gox:288:1
		tags := ctx.Param("tags")
//line cmd/gopcomm/community_yap.gox:289:1
		abstract := ctx.Param("abstract")
//line cmd/gopcomm/community_yap.gox:290:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:291:1
		trans, err := strconv.ParseBool(ctx.Param("trans"))
//line cmd/gopcomm/community_yap.gox:292:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:293:1
			xLog.Error("parse bool error:", err)
		}
//line cmd/gopcomm/community_yap.gox:296:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:297:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:298:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:299:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:304:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:305:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:306:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:307:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:312:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: title, UId: uid, Cover: ctx.Param("cover"), Tags: tags, Abstract: abstract, Label: label}, Content: content, Trans: trans, VttId: ctx.Param("vtt_id")}
//line cmd/gopcomm/community_yap.gox:328:1
		if trans {
//line cmd/gopcomm/community_yap.gox:329:1
			article, _ = this.community.TranslateArticle(todo, article)
		}
//line cmd/gopcomm/community_yap.gox:332:1
		id, err = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:333:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:334:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:335:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "add failed"})
		}
//line cmd/gopcomm/community_yap.gox:340:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:349:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:350:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:352:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:353:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:354:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:357:1
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:358:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:359:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:360:1
			user, _ = this.community.GetUser(token.Value)
		}
//line cmd/gopcomm/community_yap.gox:363:1
		items, total, err := this.community.GetArticlesByUid(todo, id, firstConst, mediaLimitConst)
//line cmd/gopcomm/community_yap.gox:364:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:365:1
			xLog.Error("get article list error:", err)
		}
//line cmd/gopcomm/community_yap.gox:368:1
		isAdmin, err := this.community.IsAdmin(id)
//line cmd/gopcomm/community_yap.gox:369:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:370:1
			xLog.Error("check admin error:", err)
		}
//line cmd/gopcomm/community_yap.gox:372:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": userClaim, "User": user, "Items": items, "Total": total, "IsAdmin": isAdmin})
	})
//line cmd/gopcomm/community_yap.gox:382:1
	this.Get("/user/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:383:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:384:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:385:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:386:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:387:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:388:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:391:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:392:1
			xLog.Error("get token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:394:1
		user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:395:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:396:1
			xLog.Info("get user error:", err)
//line cmd/gopcomm/community_yap.gox:397:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:399:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:400:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:401:1
			xLog.Error("get uid error:", err)
//line cmd/gopcomm/community_yap.gox:402:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:404:1
		userClaim, err := this.community.GetUserClaim(uid)
//line cmd/gopcomm/community_yap.gox:405:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:406:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:408:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:409:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:410:1
			xLog.Info("gac", err)
//line cmd/gopcomm/community_yap.gox:411:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:413:1
		appInfo, err := this.community.GetApplicationInfo()
//line cmd/gopcomm/community_yap.gox:414:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:415:1
			xLog.Error("get application info error:", err)
//line cmd/gopcomm/community_yap.gox:416:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:418:1
		ctx.Yap__1("user_edit", map[string]interface {
		}{"User": user, "CurrentUser": userClaim, "Application": appInfo, "Binds": gac.GetProviderBindStatus()})
	})
//line cmd/gopcomm/community_yap.gox:429:1
	this.Put("/api/user", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:430:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:431:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:432:1
		user := &core.UserInfo{Id: uid, Name: ctx.Param("name"), Birthday: ctx.Param("birthday"), Gender: ctx.Param("gender"), Phone: ctx.Param("phone"), Email: ctx.Param("email"), Avatar: ctx.Param("avatar"), Owner: ctx.Param("owner"), DisplayName: ctx.Param("displayName")}
//line cmd/gopcomm/community_yap.gox:444:1
		_, err = this.community.UpdateUser(user)
//line cmd/gopcomm/community_yap.gox:445:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:446:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:447:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "msg": "update failed"})
		}
//line cmd/gopcomm/community_yap.gox:452:1
		ctx.Json__1(map[string]int{"code": 200})
	})
//line cmd/gopcomm/community_yap.gox:457:1
	this.Get("/api/users", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:459:1
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:460:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:461:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:462:1
			casdoorUser, _ := this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:463:1
			user = &core.User{Id: casdoorUser.Id}
		} else {
//line cmd/gopcomm/community_yap.gox:467:1
			xLog.Error("get token error:", err)
//line cmd/gopcomm/community_yap.gox:468:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "msg": "get token failed", "users": nil, "next": 1})
		}
//line cmd/gopcomm/community_yap.gox:476:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:477:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:478:1
		fromInt, err := strconv.Atoi(from)
//line cmd/gopcomm/community_yap.gox:479:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:480:1
			fromInt = 1
		}
//line cmd/gopcomm/community_yap.gox:482:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:483:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:484:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:488:1
		userAuth, err := this.community.GetUserAuthById(user.Id)
//line cmd/gopcomm/community_yap.gox:489:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:490:1
			xLog.Error("get current user auth error:", err)
		}
//line cmd/gopcomm/community_yap.gox:493:1
		if !userAuth.Status {
//line cmd/gopcomm/community_yap.gox:494:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "msg": "auth failed", "items": nil, "total": 10, "next": 1})
		}
//line cmd/gopcomm/community_yap.gox:503:1
		users, next, err := this.community.ListPageUsers(fromInt, limitInt)
//line cmd/gopcomm/community_yap.gox:504:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:505:1
			xLog.Error("get users error:", err)
//line cmd/gopcomm/community_yap.gox:506:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "msg": "get users failed", "items": nil, "total": 10, "next": 1})
		}
//line cmd/gopcomm/community_yap.gox:515:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "msg": "ok", "items": users, "total": len(users), "next": next})
	})
//line cmd/gopcomm/community_yap.gox:524:1
	this.Put("/api/user/public", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:526:1
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:527:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:528:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:529:1
			casdoorUser, _ := this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:530:1
			user = &core.User{Id: casdoorUser.Id}
		} else {
//line cmd/gopcomm/community_yap.gox:534:1
			xLog.Error("get token error:", err)
//line cmd/gopcomm/community_yap.gox:535:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "msg": "get token failed", "users": nil, "next": 1})
		}
//line cmd/gopcomm/community_yap.gox:544:1
		userAuth, err := this.community.GetUserAuthById(user.Id)
//line cmd/gopcomm/community_yap.gox:545:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:546:1
			xLog.Error("get current user auth error:", err)
		}
//line cmd/gopcomm/community_yap.gox:549:1
		if !userAuth.Status {
//line cmd/gopcomm/community_yap.gox:550:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "msg": "auth failed", "items": nil, "total": 10, "next": 1})
		}
//line cmd/gopcomm/community_yap.gox:559:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:560:1
		isPublic := ctx.Param("is_public")
//line cmd/gopcomm/community_yap.gox:561:1
		isPublicBool, err := strconv.ParseBool(isPublic)
//line cmd/gopcomm/community_yap.gox:562:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:563:1
			xLog.Error("parse bool error:", err)
//line cmd/gopcomm/community_yap.gox:564:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "msg": "parse bool failed"})
		}
//line cmd/gopcomm/community_yap.gox:570:1
		res, err := this.community.UpdateUserPublicAuth(uid, isPublicBool)
//line cmd/gopcomm/community_yap.gox:571:1
		if err != nil || !res {
//line cmd/gopcomm/community_yap.gox:572:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:573:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "msg": "update failed"})
		}
//line cmd/gopcomm/community_yap.gox:579:1
		ctx.Json__1(map[string]int{"code": 200})
	})
//line cmd/gopcomm/community_yap.gox:584:1
	this.Get("/api/user/unlink", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:585:1
		pv := ctx.Param("pv")
//line cmd/gopcomm/community_yap.gox:586:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:587:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:588:1
			ctx.Json__1(map[string]int{"code": 200})
//line cmd/gopcomm/community_yap.gox:591:1
			return
		}
//line cmd/gopcomm/community_yap.gox:593:1
		switch pv {
//line cmd/gopcomm/community_yap.gox:594:1
		case "Twitter":
//line cmd/gopcomm/community_yap.gox:595:1
		case "Facebook":
//line cmd/gopcomm/community_yap.gox:596:1
		case "Github":
//line cmd/gopcomm/community_yap.gox:597:1
		case "WeChat":
//line cmd/gopcomm/community_yap.gox:598:1
		default:
//line cmd/gopcomm/community_yap.gox:599:1
			pv = ""
		}
//line cmd/gopcomm/community_yap.gox:601:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:602:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:603:1
			gac.UnLink(pv)
		}
//line cmd/gopcomm/community_yap.gox:605:1
		ctx.Json__1(map[string]int{"code": 200})
	})
//line cmd/gopcomm/community_yap.gox:610:1
	this.Get("/api/user/:id/articles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:611:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:612:1
		page := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:613:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:615:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:616:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:617:1
			limitInt = mediaLimitConst
		}
//line cmd/gopcomm/community_yap.gox:619:1
		items, total, err := this.community.GetArticlesByUid(todo, id, page, limitInt)
//line cmd/gopcomm/community_yap.gox:620:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:621:1
			xLog.Error("get article list error:", err)
//line cmd/gopcomm/community_yap.gox:622:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error(), "total": 0})
		}
//line cmd/gopcomm/community_yap.gox:628:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": items, "total": total})
	})
//line cmd/gopcomm/community_yap.gox:635:1
	this.Get("/api/user/:id/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:636:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:637:1
		uid := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:638:1
		page := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:639:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:641:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:642:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:643:1
			limitInt = mediaLimitConst
		}
//line cmd/gopcomm/community_yap.gox:645:1
		pageInt, err := strconv.Atoi(page)
//line cmd/gopcomm/community_yap.gox:646:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:647:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "total": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:653:1
		files, total, err := this.community.ListMediaByUserId(todo, uid, format, pageInt, limitInt)
//line cmd/gopcomm/community_yap.gox:654:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:655:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "total": total, "err": err.Error()})
		} else {
//line cmd/gopcomm/community_yap.gox:661:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "total": total, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:672:1
	this.Delete("/api/media/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:673:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:674:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:675:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:676:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:677:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:678:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:683:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:684:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:685:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:690:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:697:1
	this.Get("/api/translation/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:698:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:699:1
		article, err := this.community.GetTranslateArticle(todo, id)
//line cmd/gopcomm/community_yap.gox:701:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:702:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:703:1
			ctx.Json__1(map[string]int{"code": 0})
		}
//line cmd/gopcomm/community_yap.gox:707:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "content": article.Content, "tags": article.Tags, "title": article.Title})
	})
//line cmd/gopcomm/community_yap.gox:715:1
	this.Post("/api/translation", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:716:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:717:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, "auto", language.English)
//line cmd/gopcomm/community_yap.gox:718:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:719:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:724:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:730:1
	this.Get("/api/media/:id/url", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:731:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:732:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:733:1
		htmlUrl := fmt.Sprintf("%s%s", this.domain, fileKey)
//line cmd/gopcomm/community_yap.gox:734:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:735:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:740:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:746:1
	this.Get("/api/video/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:747:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:748:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:749:1
		m := make(map[string]string, 4)
//line cmd/gopcomm/community_yap.gox:750:1
		format, err := this.community.GetMediaType(todo, id)
//line cmd/gopcomm/community_yap.gox:751:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:752:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:757:1
		htmlUrl := fmt.Sprintf("%s%s", this.domain, fileKey)
//line cmd/gopcomm/community_yap.gox:758:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:759:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:764:1
		m["fileKey"] = htmlUrl
//line cmd/gopcomm/community_yap.gox:765:1
		m["type"] = format
//line cmd/gopcomm/community_yap.gox:766:1
		m["subtitle"] = this.domain + id
//line cmd/gopcomm/community_yap.gox:767:1
		m["status"] = "0"
//line cmd/gopcomm/community_yap.gox:768:1
		match, _ := regexp.MatchString("^video", format)
//line cmd/gopcomm/community_yap.gox:769:1
		if match {
//line cmd/gopcomm/community_yap.gox:770:1
			subtitle, status, err := this.community.GetVideoSubtitle(todo, id)
//line cmd/gopcomm/community_yap.gox:771:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:772:1
				ctx.Json__1(map[string]interface {
				}{"code": 200, "url": m})
			}
//line cmd/gopcomm/community_yap.gox:777:1
			if status == "1" {
//line cmd/gopcomm/community_yap.gox:778:1
				m["subtitle"] = this.domain + subtitle
			}
//line cmd/gopcomm/community_yap.gox:780:1
			m["status"] = status
		}
//line cmd/gopcomm/community_yap.gox:782:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": m})
	})
//line cmd/gopcomm/community_yap.gox:788:1
	this.Post("/api/media", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:789:1
		xLog := xlog.New("")
//line cmd/gopcomm/community_yap.gox:790:1
		file, header, err := ctx.FormFile("file")
//line cmd/gopcomm/community_yap.gox:791:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:792:1
			xLog.Error("upload file error:", header)
//line cmd/gopcomm/community_yap.gox:793:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:794:1
			return
		}
//line cmd/gopcomm/community_yap.gox:796:1
		filename := header.Filename
//line cmd/gopcomm/community_yap.gox:797:1
		err = ctx.ParseMultipartForm(10 << 20)
//line cmd/gopcomm/community_yap.gox:798:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:799:1
			xLog.Error("upload file error:", filename)
//line cmd/gopcomm/community_yap.gox:800:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:801:1
			return
		}
//line cmd/gopcomm/community_yap.gox:804:1
		dst, err := os.Create(filename)
//line cmd/gopcomm/community_yap.gox:805:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:806:1
			xLog.Error("create file error:", file)
//line cmd/gopcomm/community_yap.gox:807:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:808:1
			return
		}
//line cmd/gopcomm/community_yap.gox:810:1
		defer func() {
//line cmd/gopcomm/community_yap.gox:811:1
			file.Close()
//line cmd/gopcomm/community_yap.gox:812:1
			dst.Close()
//line cmd/gopcomm/community_yap.gox:813:1
			err = os.Remove(filename)
//line cmd/gopcomm/community_yap.gox:814:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:815:1
				xLog.Error("delete file error:", filename)
//line cmd/gopcomm/community_yap.gox:816:1
				return
			}
		}()
//line cmd/gopcomm/community_yap.gox:820:1
		_, err = io.Copy(dst, file)
//line cmd/gopcomm/community_yap.gox:821:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:822:1
			xLog.Error("copy file errer:", filename)
//line cmd/gopcomm/community_yap.gox:823:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:824:1
			return
		}
//line cmd/gopcomm/community_yap.gox:827:1
		bytes, err := os.ReadFile(filename)
//line cmd/gopcomm/community_yap.gox:828:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:829:1
			xLog.Error("read file errer:", filename)
//line cmd/gopcomm/community_yap.gox:830:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:831:1
			return
		}
//line cmd/gopcomm/community_yap.gox:834:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:835:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:836:1
			ctx.JSON(500, err.Error())
		}
//line cmd/gopcomm/community_yap.gox:839:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:840:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:841:1
			ctx.JSON(500, err.Error())
		}
//line cmd/gopcomm/community_yap.gox:843:1
		ext := mimetype.Detect(bytes).String()
//line cmd/gopcomm/community_yap.gox:844:1
		id, err := this.community.SaveMedia(context.Background(), uid, bytes, ext)
//line cmd/gopcomm/community_yap.gox:845:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:846:1
			xLog.Error("save file", err.Error())
//line cmd/gopcomm/community_yap.gox:847:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:848:1
			return
		}
//line cmd/gopcomm/community_yap.gox:852:1
		if strings.Contains(ext, "video") {
//line cmd/gopcomm/community_yap.gox:854:1
			err := this.community.NewVideoTask(context.Background(), uid, strconv.FormatInt(id, 10))
//line cmd/gopcomm/community_yap.gox:855:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:856:1
				xLog.Error("start video task error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:860:1
		ctx.JSON(200, id)
	})
//line cmd/gopcomm/community_yap.gox:863:1
	this.Post("/api/caption/task", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:864:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:865:1
		vid := ctx.Param("vid")
//line cmd/gopcomm/community_yap.gox:867:1
		if uid == "" || vid == "" {
//line cmd/gopcomm/community_yap.gox:868:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Invalid param"})
		}
//line cmd/gopcomm/community_yap.gox:874:1
		if
//line cmd/gopcomm/community_yap.gox:874:1
		err := this.community.RetryCaptionGenerate(todo, uid, vid); err != nil {
//line cmd/gopcomm/community_yap.gox:875:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Request task error"})
		}
//line cmd/gopcomm/community_yap.gox:881:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "msg": "Ok"})
	})
//line cmd/gopcomm/community_yap.gox:890:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:895:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:896:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:897:1
			xLog.Infof("Error parsing Referer: %#v", err)
//line cmd/gopcomm/community_yap.gox:898:1
			return
		}
//line cmd/gopcomm/community_yap.gox:901:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:902:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:903:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:906:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "login/callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:908:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:909:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:912:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:913:1
		tokenCookie, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:914:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:915:1
			xLog.Error("remove token error:", err)
//line cmd/gopcomm/community_yap.gox:916:1
			return
		}
//line cmd/gopcomm/community_yap.gox:919:1
		tokenCookie.MaxAge = -1
//line cmd/gopcomm/community_yap.gox:920:1
		http.SetCookie(ctx.ResponseWriter, tokenCookie)
//line cmd/gopcomm/community_yap.gox:922:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:923:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:924:1
			xLog.Infof("Error parsing Referer: %#v", err)
//line cmd/gopcomm/community_yap.gox:925:1
			return
		}
//line cmd/gopcomm/community_yap.gox:928:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:929:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:930:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:933:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:936:1
	this.Get("/login/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:937:1
		code := ctx.URL.Query().Get("code")
//line cmd/gopcomm/community_yap.gox:938:1
		state := ctx.URL.Query().Get("state")
//line cmd/gopcomm/community_yap.gox:939:1
		token, err := this.community.GetOAuthToken(code, state)
//line cmd/gopcomm/community_yap.gox:940:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:941:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:943:1
		cookie := http.Cookie{Name: "token", Value: token.AccessToken, Path: "/", MaxAge: 168 * 60 * 60, HttpOnly: true}
//line cmd/gopcomm/community_yap.gox:950:1
		http.SetCookie(ctx.ResponseWriter, &cookie)
//line cmd/gopcomm/community_yap.gox:952:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:953:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:954:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:955:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:956:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:959:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:962:1
	this.Post("/api/article/like", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:963:1
		articleId := ctx.Param("articleId")
//line cmd/gopcomm/community_yap.gox:964:1
		articleIdInt, err := strconv.Atoi(articleId)
//line cmd/gopcomm/community_yap.gox:965:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:966:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:971:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:972:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:973:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:974:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:979:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:980:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:981:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:986:1
		res, err := this.community.ArticleLike(todo, articleIdInt, uid)
//line cmd/gopcomm/community_yap.gox:987:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:988:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:993:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": res})
	})
//line cmd/gopcomm/community_yap.gox:1000:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:1001:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:1004:1
	configFile := ".env"
//line cmd/gopcomm/community_yap.gox:1005:1
	flag.StringVar(&configFile, "config", ".env", "Path to the config file")
//line cmd/gopcomm/community_yap.gox:1006:1
	flag.Parse()
//line cmd/gopcomm/community_yap.gox:1007:1
	if
//line cmd/gopcomm/community_yap.gox:1007:1
	err := godotenv.Load(configFile); err != nil {
//line cmd/gopcomm/community_yap.gox:1008:1
		xLog.Error(err)
	}
//line cmd/gopcomm/community_yap.gox:1011:1
	conf := core.NewConfigFromEnv()
//line cmd/gopcomm/community_yap.gox:1012:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:1013:1
	this.domain = conf.QiNiuConfig.Domain
//line cmd/gopcomm/community_yap.gox:1014:1
	endpoint := conf.AppConfig.EndPoint
//line cmd/gopcomm/community_yap.gox:1016:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:1019:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:1021:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:1022:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:1024:1
				if
//line cmd/gopcomm/community_yap.gox:1024:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:1025:1
					xLog.Error(err)
//line cmd/gopcomm/community_yap.gox:1026:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:1030:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
