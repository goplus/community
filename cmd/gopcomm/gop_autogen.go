package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/goplus/yap"
	"context"
	"net/http"
	"net/url"
	"io"
	"regexp"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	_ "github.com/joho/godotenv/autoload"
	gopaccountsdk "github.com/liuscraft/gop-casdoor-account-sdk"
	"github.com/qiniu/x/xlog"
	"golang.org/x/text/language"
	"github.com/gabriel-vasile/mimetype"
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
}
//line cmd/gopcomm/community_yap.gox:33
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:33:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:34:1
	endpoint := os.Getenv("GOP_COMMUNITY_ENDPOINT")
//line cmd/gopcomm/community_yap.gox:35:1
	domain := os.Getenv("GOP_COMMUNITY_DOMAIN")
//line cmd/gopcomm/community_yap.gox:36:1
	xLog := xlog.New("")
//line cmd/gopcomm/community_yap.gox:40:1
	this.Static__0("/static")
//line cmd/gopcomm/community_yap.gox:42:1
	this.Get("/success", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:43:1
		ctx.Yap__1("2xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:46:1
	this.Get("/error", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:47:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:48:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:49:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:50:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:51:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:52:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:53:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:55:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:58:1
		ctx.Yap__1("4xx", map[string]interface {
		}{"UserId": userId, "User": user})
	})
//line cmd/gopcomm/community_yap.gox:64:1
	this.Get("/failed", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:65:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:66:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:67:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:68:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:69:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:70:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:71:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:73:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:76:1
		ctx.Yap__1("5xx", map[string]interface {
		}{"UserId": userId, "User": user})
	})
//line cmd/gopcomm/community_yap.gox:83:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:85:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:86:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:87:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:88:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:89:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:90:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:91:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:93:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:97:1
		articles, next, err := this.community.ListArticle(todo, core.MarkBegin, limitConst, "", labelConst)
//line cmd/gopcomm/community_yap.gox:98:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:99:1
			xLog.Error("get article error:", err)
		}
//line cmd/gopcomm/community_yap.gox:102:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:103:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:105:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": articles, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:116:1
	this.Get("/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:119:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:120:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:121:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:122:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:123:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:124:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:125:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:127:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:130:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:131:1
		platform := ctx.Param("platform")
//line cmd/gopcomm/community_yap.gox:132:1
		ip := this.community.GetClientIP(ctx.Request)
//line cmd/gopcomm/community_yap.gox:133:1
		this.community.ArticleLView(todo, id, ip, userId, platform)
//line cmd/gopcomm/community_yap.gox:134:1
		article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:135:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:136:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:137:1
			return
		}
//line cmd/gopcomm/community_yap.gox:139:1
		likeState, err := this.community.ArticleLikeState(todo, userId, id)
//line cmd/gopcomm/community_yap.gox:140:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:141:1
			xLog.Error("article state err:", err)
//line cmd/gopcomm/community_yap.gox:142:1
			return
		}
//line cmd/gopcomm/community_yap.gox:145:1
		ctx.Yap__1("article", map[string]interface {
		}{"UserId": userId, "User": user, "Article": article, "LikeState": likeState})
	})
//line cmd/gopcomm/community_yap.gox:153:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:154:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:155:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:156:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:157:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:158:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:159:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:160:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:162:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:165:1
		ctx.Yap__1("edit", map[string]interface {
		}{"User": user, "UserId": userId})
	})
//line cmd/gopcomm/community_yap.gox:171:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:172:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:173:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:174:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:175:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:176:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:177:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:178:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:180:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:184:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:185:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:186:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:187:1
			editable, err := this.community.CanEditable(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:188:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:189:1
				xLog.Error("can editable error:", err)
//line cmd/gopcomm/community_yap.gox:190:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:192:1
			if !editable {
//line cmd/gopcomm/community_yap.gox:193:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:194:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:196:1
			article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:197:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:198:1
				xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:199:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:202:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:203:1
				xLog.Error("json marshal error:", err)
//line cmd/gopcomm/community_yap.gox:204:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:206:1
			ctx.Yap__1("edit", map[string]interface {
			}{"UserId": userId, "User": user, "Article": article})
		}
	})
//line cmd/gopcomm/community_yap.gox:214:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:215:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:216:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:217:1
		if label == "" {
//line cmd/gopcomm/community_yap.gox:218:1
			label = "article"
		}
//line cmd/gopcomm/community_yap.gox:222:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:223:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:224:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:225:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:226:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:227:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:228:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:230:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:234:1
		articles, next, err := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue, label)
//line cmd/gopcomm/community_yap.gox:235:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:236:1
			xLog.Error("get article error:", err)
		}
//line cmd/gopcomm/community_yap.gox:239:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:240:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:242:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": articles, "Value": searchValue, "Next": next, "Tab": label})
	})
//line cmd/gopcomm/community_yap.gox:255:1
	this.Get("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:256:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:257:1
		article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:258:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:259:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:260:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get article failed"})
		}
//line cmd/gopcomm/community_yap.gox:265:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:271:1
	this.Delete("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:272:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:273:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:274:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:275:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:276:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:277:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:282:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:283:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:284:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:289:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:296:1
	this.Get("/api/articles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:297:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:298:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:299:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:300:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:302:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:303:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:304:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:307:1
		articles, next, err := this.community.ListArticle(todo, from, limitInt, searchValue, label)
//line cmd/gopcomm/community_yap.gox:308:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:309:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:310:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get article failed"})
		}
//line cmd/gopcomm/community_yap.gox:315:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:323:1
	this.Post("/api/article/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:324:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:325:1
		content := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:326:1
		title := ctx.Param("title")
//line cmd/gopcomm/community_yap.gox:327:1
		tags := ctx.Param("tags")
//line cmd/gopcomm/community_yap.gox:328:1
		abstract := ctx.Param("abstract")
//line cmd/gopcomm/community_yap.gox:329:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:330:1
		trans, err := strconv.ParseBool(ctx.Param("trans"))
//line cmd/gopcomm/community_yap.gox:331:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:332:1
			xLog.Error("parse bool error:", err)
		}
//line cmd/gopcomm/community_yap.gox:335:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:336:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:337:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:338:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:343:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:344:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:345:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:346:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:351:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: title, UId: uid, Cover: ctx.Param("cover"), Tags: tags, Abstract: abstract, Label: label}, Content: content, Trans: trans, VttId: ctx.Param("vtt_id")}
//line cmd/gopcomm/community_yap.gox:367:1
		if trans {
//line cmd/gopcomm/community_yap.gox:368:1
			article, _ = this.community.TranslateArticle(todo, article)
		}
//line cmd/gopcomm/community_yap.gox:371:1
		id, err = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:372:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:373:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:374:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "add failed"})
		}
//line cmd/gopcomm/community_yap.gox:379:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:388:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:389:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:391:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:392:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:393:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:396:1
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:397:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:398:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:399:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:400:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:401:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:402:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:404:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:408:1
		items, total, err := this.community.GetArticlesByUid(todo, id, firstConst, mediaLimitConst)
//line cmd/gopcomm/community_yap.gox:409:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:410:1
			xLog.Error("get article list error:", err)
		}
//line cmd/gopcomm/community_yap.gox:412:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": userClaim, "User": user, "Items": items, "UserId": userId, "Total": total})
	})
//line cmd/gopcomm/community_yap.gox:422:1
	this.Get("/user/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:423:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:424:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:425:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:426:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:427:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:428:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:429:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:431:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:434:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:435:1
			xLog.Error("get token error:", err)
//line cmd/gopcomm/community_yap.gox:436:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:438:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:439:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:440:1
			xLog.Error("get uid error:", err)
//line cmd/gopcomm/community_yap.gox:441:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:443:1
		userClaim, err := this.community.GetUserClaim(uid)
//line cmd/gopcomm/community_yap.gox:444:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:445:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:447:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:448:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:449:1
			xLog.Info("gac", err)
//line cmd/gopcomm/community_yap.gox:450:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:452:1
		appInfo, err := this.community.GetApplicationInfo()
//line cmd/gopcomm/community_yap.gox:453:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:454:1
			xLog.Error("get application info error:", err)
//line cmd/gopcomm/community_yap.gox:455:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:457:1
		ctx.Yap__1("user_edit", map[string]interface {
		}{"UserId": userId, "User": user, "CurrentUser": userClaim, "Application": appInfo, "Binds": gac.GetProviderBindStatus()})
	})
//line cmd/gopcomm/community_yap.gox:469:1
	this.Put("/api/user", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:470:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:471:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:472:1
		user := &core.UserInfo{Id: uid, Name: ctx.Param("name"), Birthday: ctx.Param("birthday"), Gender: ctx.Param("gender"), Phone: ctx.Param("phone"), Email: ctx.Param("email"), Avatar: ctx.Param("avatar"), Owner: ctx.Param("owner"), DisplayName: ctx.Param("displayName")}
//line cmd/gopcomm/community_yap.gox:484:1
		_, err = this.community.UpdateUser(user)
//line cmd/gopcomm/community_yap.gox:485:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:486:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:487:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "msg": "update failed"})
		}
//line cmd/gopcomm/community_yap.gox:492:1
		ctx.Json__1(map[string]int{"code": 200})
	})
//line cmd/gopcomm/community_yap.gox:497:1
	this.Get("/api/user/unlink", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:498:1
		pv := ctx.Param("pv")
//line cmd/gopcomm/community_yap.gox:499:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:500:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:501:1
			ctx.Json__1(map[string]int{"code": 200})
//line cmd/gopcomm/community_yap.gox:504:1
			return
		}
//line cmd/gopcomm/community_yap.gox:506:1
		switch pv {
//line cmd/gopcomm/community_yap.gox:507:1
		case "Twitter":
//line cmd/gopcomm/community_yap.gox:508:1
		case "Facebook":
//line cmd/gopcomm/community_yap.gox:509:1
		case "Github":
//line cmd/gopcomm/community_yap.gox:510:1
		case "WeChat":
//line cmd/gopcomm/community_yap.gox:511:1
		default:
//line cmd/gopcomm/community_yap.gox:512:1
			pv = ""
		}
//line cmd/gopcomm/community_yap.gox:514:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:515:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:516:1
			gac.UnLink(pv)
		}
//line cmd/gopcomm/community_yap.gox:518:1
		ctx.Json__1(map[string]int{"code": 200})
	})
//line cmd/gopcomm/community_yap.gox:523:1
	this.Get("/api/user/:id/articles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:524:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:525:1
		page := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:526:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:528:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:529:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:530:1
			limitInt = mediaLimitConst
		}
//line cmd/gopcomm/community_yap.gox:532:1
		items, total, err := this.community.GetArticlesByUid(todo, id, page, limitInt)
//line cmd/gopcomm/community_yap.gox:533:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:534:1
			xLog.Error("get article list error:", err)
//line cmd/gopcomm/community_yap.gox:535:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error(), "total": 0})
		}
//line cmd/gopcomm/community_yap.gox:541:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": items, "total": total})
	})
//line cmd/gopcomm/community_yap.gox:548:1
	this.Get("/api/user/:id/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:549:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:550:1
		uid := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:551:1
		page := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:552:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:554:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:555:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:556:1
			limitInt = mediaLimitConst
		}
//line cmd/gopcomm/community_yap.gox:558:1
		pageInt, err := strconv.Atoi(page)
//line cmd/gopcomm/community_yap.gox:559:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:560:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "total": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:566:1
		files, total, err := this.community.ListMediaByUserId(todo, uid, format, pageInt, limitInt)
//line cmd/gopcomm/community_yap.gox:567:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:568:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "total": total, "err": err.Error()})
		} else {
//line cmd/gopcomm/community_yap.gox:574:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "total": total, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:585:1
	this.Delete("/api/media/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:586:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:587:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:588:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:589:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:590:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:591:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:596:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:597:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:598:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:603:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:610:1
	this.Get("/api/translation/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:611:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:612:1
		article, err := this.community.GetTranslateArticle(todo, id)
//line cmd/gopcomm/community_yap.gox:614:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:615:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:616:1
			ctx.Json__1(map[string]int{"code": 0})
		}
//line cmd/gopcomm/community_yap.gox:620:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "content": article.Content, "tags": article.Tags, "title": article.Title})
	})
//line cmd/gopcomm/community_yap.gox:628:1
	this.Post("/api/translation", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:629:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:630:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, "auto", language.English)
//line cmd/gopcomm/community_yap.gox:631:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:632:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:637:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:643:1
	this.Get("/api/media/:id/url", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:644:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:645:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:646:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:647:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:648:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:653:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:659:1
	this.Get("/api/video/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:660:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:661:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:662:1
		m := make(map[string]string, 4)
//line cmd/gopcomm/community_yap.gox:663:1
		format, err := this.community.GetMediaType(todo, id)
//line cmd/gopcomm/community_yap.gox:664:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:665:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:670:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:671:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:672:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:677:1
		m["fileKey"] = htmlUrl
//line cmd/gopcomm/community_yap.gox:678:1
		m["type"] = format
//line cmd/gopcomm/community_yap.gox:679:1
		m["subtitle"] = domain + id
//line cmd/gopcomm/community_yap.gox:680:1
		m["status"] = "0"
//line cmd/gopcomm/community_yap.gox:681:1
		match, _ := regexp.MatchString("^video", format)
//line cmd/gopcomm/community_yap.gox:682:1
		if match {
//line cmd/gopcomm/community_yap.gox:683:1
			subtitle, status, err := this.community.GetVideoSubtitle(todo, id)
//line cmd/gopcomm/community_yap.gox:684:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:685:1
				ctx.Json__1(map[string]interface {
				}{"code": 200, "url": m})
			}
//line cmd/gopcomm/community_yap.gox:690:1
			if status == "1" {
//line cmd/gopcomm/community_yap.gox:691:1
				m["subtitle"] = domain + subtitle
			}
//line cmd/gopcomm/community_yap.gox:693:1
			m["status"] = status
		}
//line cmd/gopcomm/community_yap.gox:695:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": m})
	})
//line cmd/gopcomm/community_yap.gox:701:1
	this.Post("/api/media", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:702:1
		file, header, err := ctx.FormFile("file")
//line cmd/gopcomm/community_yap.gox:703:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:704:1
			xLog.Error("upload file error:", header)
//line cmd/gopcomm/community_yap.gox:705:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:706:1
			return
		}
//line cmd/gopcomm/community_yap.gox:708:1
		filename := header.Filename
//line cmd/gopcomm/community_yap.gox:709:1
		err = ctx.ParseMultipartForm(10 << 20)
//line cmd/gopcomm/community_yap.gox:710:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:711:1
			xLog.Error("upload file error:", filename)
//line cmd/gopcomm/community_yap.gox:712:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:713:1
			return
		}
//line cmd/gopcomm/community_yap.gox:716:1
		dst, err := os.Create(filename)
//line cmd/gopcomm/community_yap.gox:717:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:718:1
			xLog.Error("create file error:", file)
//line cmd/gopcomm/community_yap.gox:719:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:720:1
			return
		}
//line cmd/gopcomm/community_yap.gox:722:1
		defer func() {
//line cmd/gopcomm/community_yap.gox:723:1
			file.Close()
//line cmd/gopcomm/community_yap.gox:724:1
			dst.Close()
//line cmd/gopcomm/community_yap.gox:725:1
			err = os.Remove(filename)
//line cmd/gopcomm/community_yap.gox:726:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:727:1
				xLog.Error("delete file error:", filename)
//line cmd/gopcomm/community_yap.gox:728:1
				return
			}
		}()
//line cmd/gopcomm/community_yap.gox:732:1
		_, err = io.Copy(dst, file)
//line cmd/gopcomm/community_yap.gox:733:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:734:1
			xLog.Error("copy file errer:", filename)
//line cmd/gopcomm/community_yap.gox:735:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:736:1
			return
		}
//line cmd/gopcomm/community_yap.gox:739:1
		bytes, err := os.ReadFile(filename)
//line cmd/gopcomm/community_yap.gox:740:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:741:1
			xLog.Error("read file errer:", filename)
//line cmd/gopcomm/community_yap.gox:742:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:743:1
			return
		}
//line cmd/gopcomm/community_yap.gox:746:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:747:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:748:1
			ctx.JSON(500, err.Error())
		}
//line cmd/gopcomm/community_yap.gox:751:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:752:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:753:1
			ctx.JSON(500, err.Error())
		}
//line cmd/gopcomm/community_yap.gox:755:1
		ext := mimetype.Detect(bytes).String()
//line cmd/gopcomm/community_yap.gox:756:1
		id, err := this.community.SaveMedia(context.Background(), uid, bytes, ext)
//line cmd/gopcomm/community_yap.gox:757:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:758:1
			xLog.Error("save file", err.Error())
//line cmd/gopcomm/community_yap.gox:759:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:760:1
			return
		}
//line cmd/gopcomm/community_yap.gox:764:1
		if strings.Contains(ext, "video") {
//line cmd/gopcomm/community_yap.gox:766:1
			err := this.community.NewVideoTask(context.Background(), uid, strconv.FormatInt(id, 10))
//line cmd/gopcomm/community_yap.gox:767:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:768:1
				xLog.Error("start video task error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:772:1
		ctx.JSON(200, id)
	})
//line cmd/gopcomm/community_yap.gox:775:1
	this.Post("/api/caption/task", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:776:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:777:1
		vid := ctx.Param("vid")
//line cmd/gopcomm/community_yap.gox:779:1
		if uid == "" || vid == "" {
//line cmd/gopcomm/community_yap.gox:780:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Invalid param"})
		}
//line cmd/gopcomm/community_yap.gox:786:1
		if
//line cmd/gopcomm/community_yap.gox:786:1
		err := this.community.RetryCaptionGenerate(todo, uid, vid); err != nil {
//line cmd/gopcomm/community_yap.gox:787:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Request task error"})
		}
//line cmd/gopcomm/community_yap.gox:793:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "msg": "Ok"})
	})
//line cmd/gopcomm/community_yap.gox:802:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:807:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:808:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:809:1
			xLog.Infof("Error parsing Referer: %#v", err)
//line cmd/gopcomm/community_yap.gox:810:1
			return
		}
//line cmd/gopcomm/community_yap.gox:813:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:814:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:815:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:818:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "login/callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:820:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:821:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:824:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:825:1
		tokenCookie, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:826:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:827:1
			xLog.Error("remove token error:", err)
//line cmd/gopcomm/community_yap.gox:828:1
			return
		}
//line cmd/gopcomm/community_yap.gox:831:1
		tokenCookie.MaxAge = -1
//line cmd/gopcomm/community_yap.gox:832:1
		http.SetCookie(ctx.ResponseWriter, tokenCookie)
//line cmd/gopcomm/community_yap.gox:834:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:835:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:836:1
			xLog.Infof("Error parsing Referer: %#v", err)
//line cmd/gopcomm/community_yap.gox:837:1
			return
		}
//line cmd/gopcomm/community_yap.gox:840:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:841:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:842:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:845:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:848:1
	this.Get("/login/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:849:1
		code := ctx.URL.Query().Get("code")
//line cmd/gopcomm/community_yap.gox:850:1
		state := ctx.URL.Query().Get("state")
//line cmd/gopcomm/community_yap.gox:851:1
		token, err := this.community.GetOAuthToken(code, state)
//line cmd/gopcomm/community_yap.gox:852:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:853:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:855:1
		cookie := http.Cookie{Name: "token", Value: token.AccessToken, Path: "/", MaxAge: 3600}
//line cmd/gopcomm/community_yap.gox:861:1
		http.SetCookie(ctx.ResponseWriter, &cookie)
//line cmd/gopcomm/community_yap.gox:863:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:864:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:865:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:866:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:867:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:870:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:873:1
	this.Post("/api/article/like", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:874:1
		articleId := ctx.Param("articleId")
//line cmd/gopcomm/community_yap.gox:875:1
		articleIdInt, err := strconv.Atoi(articleId)
//line cmd/gopcomm/community_yap.gox:876:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:877:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:882:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:883:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:884:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:885:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:890:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:891:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:892:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:897:1
		res, err := this.community.ArticleLike(todo, articleIdInt, uid)
//line cmd/gopcomm/community_yap.gox:898:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:899:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:904:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": res})
	})
//line cmd/gopcomm/community_yap.gox:910:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:911:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:912:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:915:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:916:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:919:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:922:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:924:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:925:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:927:1
				if
//line cmd/gopcomm/community_yap.gox:927:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:928:1
					xLog.Error(err)
//line cmd/gopcomm/community_yap.gox:929:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:933:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
