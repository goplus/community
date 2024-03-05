package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"

	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"github.com/goplus/yap"
	_ "github.com/joho/godotenv/autoload"
	gopaccountsdk "github.com/liuscraft/gop-casdoor-account-sdk"
	"github.com/qiniu/x/xlog"
	"golang.org/x/text/language"
)

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

//line cmd/gopcomm/community_yap.gox:30
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:30:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:31:1
	endpoint := os.Getenv("GOP_COMMUNITY_ENDPOINT")
//line cmd/gopcomm/community_yap.gox:32:1
	domain := os.Getenv("GOP_COMMUNITY_DOMAIN")
//line cmd/gopcomm/community_yap.gox:33:1
	xLog := xlog.New("")
//line cmd/gopcomm/community_yap.gox:37:1
	this.Static__0("/static")
//line cmd/gopcomm/community_yap.gox:39:1
	this.Get("/success", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:40:1
		ctx.Yap__1("2xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:43:1
	this.Get("/error", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:44:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:45:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:46:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:47:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:48:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:49:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:50:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:52:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:55:1
		ctx.Yap__1("4xx", map[string]interface {
		}{"UserId": userId, "User": user})
	})
//line cmd/gopcomm/community_yap.gox:61:1
	this.Get("/failed", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:62:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:63:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:64:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:65:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:66:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:67:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:68:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:70:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:73:1
		ctx.Yap__1("5xx", map[string]interface {
		}{"UserId": userId, "User": user})
	})
//line cmd/gopcomm/community_yap.gox:80:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:82:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:83:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:84:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:85:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:86:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:87:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:88:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:90:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:94:1
		articles, next, err := this.community.ListArticle(todo, core.MarkBegin, limitConst, "", labelConst)
//line cmd/gopcomm/community_yap.gox:95:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:96:1
			xLog.Error("get article error:", err)
		}
//line cmd/gopcomm/community_yap.gox:98:1
		articlesJson, err := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:99:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:100:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:102:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": string(articlesJson), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:113:1
	this.Get("/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:116:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:117:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:118:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:119:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:120:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:121:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:122:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:124:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:127:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:128:1
		article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:129:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:130:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:131:1
			return
		}
//line cmd/gopcomm/community_yap.gox:133:1
		likeState, err := this.community.ArticleLikeState(todo, userId, id)
//line cmd/gopcomm/community_yap.gox:134:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:135:1
			xLog.Error("article state err:", err)
//line cmd/gopcomm/community_yap.gox:136:1
			return
		}
//line cmd/gopcomm/community_yap.gox:138:1
		ip := this.community.GetClientIP(ctx.Request)
//line cmd/gopcomm/community_yap.gox:139:1
		this.community.ArticleLView(todo, id, ip, userId)
//line cmd/gopcomm/community_yap.gox:140:1
		articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:141:1
		ctx.Yap__1("article", map[string]interface {
		}{"UserId": userId, "User": user, "Article": string(articleJson), "LikeState": likeState})
	})
//line cmd/gopcomm/community_yap.gox:149:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:150:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:151:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:152:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:153:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:154:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:155:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:156:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:158:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:161:1
		ctx.Yap__1("edit", map[string]interface {
		}{"User": user, "UserId": userId})
	})
//line cmd/gopcomm/community_yap.gox:167:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:168:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:169:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:170:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:171:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:172:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:173:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:174:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:176:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:180:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:181:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:182:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:183:1
			editable, err := this.community.CanEditable(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:184:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:185:1
				xLog.Error("can editable error:", err)
//line cmd/gopcomm/community_yap.gox:186:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:188:1
			if !editable {
//line cmd/gopcomm/community_yap.gox:189:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:190:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:192:1
			article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:193:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:194:1
				xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:195:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:197:1
			articleJson, err := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:198:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:199:1
				xLog.Error("json marshal error:", err)
//line cmd/gopcomm/community_yap.gox:200:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:202:1
			ctx.Yap__1("edit", map[string]interface {
			}{"UserId": userId, "User": user, "Article": string(articleJson)})
		}
	})
//line cmd/gopcomm/community_yap.gox:210:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:211:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:212:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:213:1
		if label == "" {
//line cmd/gopcomm/community_yap.gox:214:1
			label = "article"
		}
//line cmd/gopcomm/community_yap.gox:218:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:219:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:220:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:221:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:222:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:223:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:224:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:226:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:230:1
		articles, next, err := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue, label)
//line cmd/gopcomm/community_yap.gox:231:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:232:1
			xLog.Error("get article error:", err)
		}
//line cmd/gopcomm/community_yap.gox:234:1
		articlesJson, err := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:235:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:236:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:238:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": string(articlesJson), "Value": searchValue, "Next": next, "Tab": label})
	})
//line cmd/gopcomm/community_yap.gox:251:1
	this.Get("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:252:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:253:1
		article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:254:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:255:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:256:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get article failed"})
		}
//line cmd/gopcomm/community_yap.gox:261:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:267:1
	this.Delete("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:268:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:269:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:270:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:271:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:272:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:273:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:278:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:279:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:280:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:285:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:292:1
	this.Get("/api/articles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:293:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:294:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:295:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:296:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:298:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:299:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:300:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:303:1
		articles, next, err := this.community.ListArticle(todo, from, limitInt, searchValue, label)
//line cmd/gopcomm/community_yap.gox:304:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:305:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:306:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get article failed"})
		}
//line cmd/gopcomm/community_yap.gox:311:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:319:1
	this.Post("/api/article/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:320:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:321:1
		content := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:322:1
		title := ctx.Param("title")
//line cmd/gopcomm/community_yap.gox:323:1
		tags := ctx.Param("tags")
//line cmd/gopcomm/community_yap.gox:324:1
		abstract := ctx.Param("abstract")
//line cmd/gopcomm/community_yap.gox:325:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:326:1
		trans, err := strconv.ParseBool(ctx.Param("trans"))
//line cmd/gopcomm/community_yap.gox:327:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:328:1
			xLog.Error("parse bool error:", err)
		}
//line cmd/gopcomm/community_yap.gox:331:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:332:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:333:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:334:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:339:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:340:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:341:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:342:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:347:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: title, UId: uid, Cover: ctx.Param("cover"), Tags: tags, Abstract: abstract, Label: label}, Content: content, Trans: trans, Vtt_id: ctx.Param("vtt_id")}
//line cmd/gopcomm/community_yap.gox:363:1
		if trans {
//line cmd/gopcomm/community_yap.gox:364:1
			article, _ = this.community.TranslateArticle(todo, article)
		}
//line cmd/gopcomm/community_yap.gox:367:1
		id, err = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:368:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:369:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "add failed"})
		}
//line cmd/gopcomm/community_yap.gox:374:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
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
		userClaimJson, err := json.Marshal(&userClaim)
//line cmd/gopcomm/community_yap.gox:408:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:409:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:411:1
		itemsJson, err := json.Marshal(&items)
//line cmd/gopcomm/community_yap.gox:412:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:413:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:415:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": string(userClaimJson), "User": user, "Items": string(itemsJson), "UserId": userId, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:425:1
	this.Get("/user/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:426:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:427:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:428:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:429:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:430:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:431:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:432:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:434:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:437:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:438:1
			xLog.Error("get token error:", err)
//line cmd/gopcomm/community_yap.gox:439:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:441:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:442:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:443:1
			xLog.Error("get uid error:", err)
//line cmd/gopcomm/community_yap.gox:444:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:446:1
		userClaim, err := this.community.GetUserClaim(uid)
//line cmd/gopcomm/community_yap.gox:447:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:448:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:450:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:451:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:452:1
			xLog.Info("gac", err)
//line cmd/gopcomm/community_yap.gox:453:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:460:1
		appInfo, err := this.community.GetApplicationInfo()
//line cmd/gopcomm/community_yap.gox:461:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:462:1
			xLog.Error("get application info error:", err)
//line cmd/gopcomm/community_yap.gox:463:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:465:1
		appInfoStr, err := json.Marshal(*appInfo)
//line cmd/gopcomm/community_yap.gox:466:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:467:1
			xLog.Error("json marshal error:", err)
//line cmd/gopcomm/community_yap.gox:468:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:470:1
		binds, err := json.Marshal(gac.GetProviderBindStatus())
//line cmd/gopcomm/community_yap.gox:471:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:472:1
			xLog.Error("json marshal error:", err)
//line cmd/gopcomm/community_yap.gox:473:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:475:1
		currentUser, err := json.Marshal(userClaim)
//line cmd/gopcomm/community_yap.gox:476:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:477:1
			xLog.Error("json marshal error:", err)
//line cmd/gopcomm/community_yap.gox:478:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:481:1
		ctx.Yap__1("user_edit", map[string]interface {
		}{"UserId": userId, "User": user, "CurrentUser": string(currentUser), "Application": string(appInfoStr), "Binds": string(binds)})
	})
//line cmd/gopcomm/community_yap.gox:493:1
	this.Put("/api/user", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:494:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:495:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:496:1
		user := &core.UserInfo{Id: uid, Name: ctx.Param("name"), Birthday: ctx.Param("birthday"), Gender: ctx.Param("gender"), Phone: ctx.Param("phone"), Email: ctx.Param("email"), Avatar: ctx.Param("avatar"), Owner: ctx.Param("owner"), DisplayName: ctx.Param("displayName")}
//line cmd/gopcomm/community_yap.gox:508:1
		_, err = this.community.UpdateUser(user)
//line cmd/gopcomm/community_yap.gox:509:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:510:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:511:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "msg": "update failed"})
		}
//line cmd/gopcomm/community_yap.gox:516:1
		ctx.Json__1(map[string]int{"code": 200})
	})
//line cmd/gopcomm/community_yap.gox:521:1
	this.Get("/api/user/unlink", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:522:1
		pv := ctx.Param("pv")
//line cmd/gopcomm/community_yap.gox:523:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:524:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:525:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/login", http.StatusTemporaryRedirect)
//line cmd/gopcomm/community_yap.gox:526:1
			return
		}
//line cmd/gopcomm/community_yap.gox:528:1
		switch pv {
//line cmd/gopcomm/community_yap.gox:529:1
		case "Twitter":
//line cmd/gopcomm/community_yap.gox:530:1
		case "Facebook":
//line cmd/gopcomm/community_yap.gox:531:1
		case "Github":
//line cmd/gopcomm/community_yap.gox:532:1
		case "WeChat":
//line cmd/gopcomm/community_yap.gox:533:1
		default:
//line cmd/gopcomm/community_yap.gox:534:1
			pv = ""
		}
//line cmd/gopcomm/community_yap.gox:536:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:537:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:538:1
			gac.UnLink(pv)
		}
//line cmd/gopcomm/community_yap.gox:540:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, "/user/edit", http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:543:1
	this.Get("/api/user/:id/articles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:544:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:545:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:546:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:548:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:549:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:550:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:552:1
		items, next, err := this.community.GetArticlesByUid(todo, id, from, limitInt)
//line cmd/gopcomm/community_yap.gox:553:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:554:1
			xLog.Error("get article list error:", err)
//line cmd/gopcomm/community_yap.gox:555:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error(), "total": 0})
		}
//line cmd/gopcomm/community_yap.gox:561:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": items, "next": next})
	})
//line cmd/gopcomm/community_yap.gox:568:1
	this.Get("/api/user/:id/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:569:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:570:1
		uid := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:571:1
		page := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:572:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:574:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:575:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:576:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:578:1
		pageInt, err := strconv.Atoi(page)
//line cmd/gopcomm/community_yap.gox:579:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:580:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "total": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:586:1
		files, total, err := this.community.ListMediaByUserId(todo, uid, format, pageInt, limitInt)
//line cmd/gopcomm/community_yap.gox:587:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:588:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "total": total, "err": err.Error()})
		} else {
//line cmd/gopcomm/community_yap.gox:594:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "total": total, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:605:1
	this.Delete("/api/media/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:606:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:607:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:608:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:609:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:610:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:611:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:616:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:617:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:618:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:623:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:630:1
	this.Get("/api/translation/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:631:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:632:1
		article, err := this.community.GetTranslateArticle(todo, id)
//line cmd/gopcomm/community_yap.gox:633:1
		xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:634:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:635:1
			ctx.Json__1(map[string]int{"code": 0})
		}
//line cmd/gopcomm/community_yap.gox:639:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "content": article.Content, "tags": article.Tags, "title": article.Title})
	})
//line cmd/gopcomm/community_yap.gox:647:1
	this.Post("/api/translation", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:648:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:649:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, "auto", language.English)
//line cmd/gopcomm/community_yap.gox:650:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:651:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:656:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:662:1
	this.Get("/api/media/:id/url", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:663:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:664:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:665:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:666:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:667:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:672:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:678:1
	this.Get("/api/video/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:679:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:680:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:681:1
		m := make(map[string]string, 3)
//line cmd/gopcomm/community_yap.gox:682:1
		format, err := this.community.GetMediaType(todo, id)
//line cmd/gopcomm/community_yap.gox:683:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:684:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:689:1
		match, _ := regexp.MatchString("^video", format)
//line cmd/gopcomm/community_yap.gox:690:1
		if match {
//line cmd/gopcomm/community_yap.gox:691:1
			subtitle, status, err := this.community.GetVideoSubtitle(todo, id)
//line cmd/gopcomm/community_yap.gox:692:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:693:1
				if err != nil {
//line cmd/gopcomm/community_yap.gox:694:1
					ctx.Json__1(map[string]interface {
					}{"code": 500, "err": err.Error()})
				}
//line cmd/gopcomm/community_yap.gox:699:1
				return
			}
//line cmd/gopcomm/community_yap.gox:701:1
			m["subtitle"] = domain + subtitle
//line cmd/gopcomm/community_yap.gox:702:1
			m["status"] = status
		}
//line cmd/gopcomm/community_yap.gox:704:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:705:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:706:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:711:1
		m["fileKey"] = htmlUrl
//line cmd/gopcomm/community_yap.gox:712:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": m})
	})
//line cmd/gopcomm/community_yap.gox:718:1
	this.Post("/api/media", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:719:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:722:1
	this.Post("/api/caption/task", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:723:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:724:1
		vid := ctx.Param("vid")
//line cmd/gopcomm/community_yap.gox:726:1
		if uid == "" || vid == "" {
//line cmd/gopcomm/community_yap.gox:727:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Invalid param"})
		}
//line cmd/gopcomm/community_yap.gox:733:1
		if
//line cmd/gopcomm/community_yap.gox:733:1
		err := this.community.RetryCaptionGenerate(todo, uid, vid); err != nil {
//line cmd/gopcomm/community_yap.gox:734:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Request task error"})
		}
//line cmd/gopcomm/community_yap.gox:740:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "msg": "Ok"})
	})
//line cmd/gopcomm/community_yap.gox:749:1
	this.Get("/login", func(ctx *yap.Context) {
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
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "login/callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:767:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:768:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:771:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:772:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:773:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:774:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:777:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:778:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:779:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:780:1
			return
		}
//line cmd/gopcomm/community_yap.gox:783:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:784:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:785:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:788:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:791:1
	this.Get("/login/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:792:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:793:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:794:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:796:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:797:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:798:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:799:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:800:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:803:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:806:1
	this.Post("/api/article/like", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:807:1
		articleId := ctx.Param("articleId")
//line cmd/gopcomm/community_yap.gox:808:1
		articleIdInt, err := strconv.Atoi(articleId)
//line cmd/gopcomm/community_yap.gox:809:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:810:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:815:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:816:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:817:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:818:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:823:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:824:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:825:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:830:1
		res, err := this.community.ArticleLike(todo, articleIdInt, uid)
//line cmd/gopcomm/community_yap.gox:831:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:832:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:837:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": res})
	})
//line cmd/gopcomm/community_yap.gox:843:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:844:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:845:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:848:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:849:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:852:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:855:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:857:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:858:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:859:1
				if
//line cmd/gopcomm/community_yap.gox:859:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:860:1
					xLog.Error(err)
//line cmd/gopcomm/community_yap.gox:861:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:865:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
