package main

import (
	"fmt"
	"strconv"
	"github.com/goplus/yap"
	"context"
	"flag"
	"net/http"
	"net/url"
	"regexp"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	_ "github.com/joho/godotenv/autoload"
	"github.com/joho/godotenv"
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
//line cmd/gopcomm/community_yap.gox:34
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:34:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:35:1
	xLog := xlog.New("")
//line cmd/gopcomm/community_yap.gox:39:1
	this.Static__0("/static")
//line cmd/gopcomm/community_yap.gox:41:1
	this.Get("/success", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:42:1
		ctx.Yap__1("2xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:45:1
	this.Get("/error", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:46:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:47:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:48:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:49:1
			user, _ = this.community.GetUser(token.Value)
		}
//line cmd/gopcomm/community_yap.gox:52:1
		ctx.Yap__1("4xx", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:57:1
	this.Get("/failed", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:58:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:59:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:60:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:61:1
			user, _ = this.community.GetUser(token.Value)
		}
//line cmd/gopcomm/community_yap.gox:64:1
		ctx.Yap__1("5xx", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:70:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:72:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:73:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:74:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:75:1
			user, _ = this.community.GetUser(token.Value)
		}
//line cmd/gopcomm/community_yap.gox:78:1
		articles, next, err := this.community.ListArticle(todo, core.MarkBegin, limitConst, "", labelConst)
//line cmd/gopcomm/community_yap.gox:79:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:80:1
			xLog.Error("get article error:", err)
		}
//line cmd/gopcomm/community_yap.gox:83:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:84:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:86:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": user, "Items": articles, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:96:1
	this.Get("/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:99:1
		uid := ""
//line cmd/gopcomm/community_yap.gox:101:1
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
				uid = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:110:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:111:1
		platform := ctx.Param("platform")
//line cmd/gopcomm/community_yap.gox:112:1
		ip := this.community.GetClientIP(ctx.Request)
//line cmd/gopcomm/community_yap.gox:113:1
		this.community.ArticleLView(todo, id, ip, uid, platform)
//line cmd/gopcomm/community_yap.gox:114:1
		article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:115:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:116:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:117:1
			return
		}
//line cmd/gopcomm/community_yap.gox:119:1
		likeState, err := this.community.ArticleLikeState(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:120:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:121:1
			xLog.Error("article state err:", err)
//line cmd/gopcomm/community_yap.gox:122:1
			return
		}
//line cmd/gopcomm/community_yap.gox:125:1
		ctx.Yap__1("article", map[string]interface {
		}{"User": user, "Article": article, "LikeState": likeState})
	})
//line cmd/gopcomm/community_yap.gox:132:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:133:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:134:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:135:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:136:1
			user, _ = this.community.GetUser(token.Value)
		}
//line cmd/gopcomm/community_yap.gox:138:1
		ctx.Yap__1("edit", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:143:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:144:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:145:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:146:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:147:1
			user, _ = this.community.GetUser(token.Value)
		}
//line cmd/gopcomm/community_yap.gox:150:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:151:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:152:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:153:1
			editable, err := this.community.CanEditable(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:154:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:155:1
				xLog.Error("can editable error:", err)
//line cmd/gopcomm/community_yap.gox:156:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:158:1
			if !editable {
//line cmd/gopcomm/community_yap.gox:159:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:160:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:162:1
			article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:163:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:164:1
				xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:165:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:168:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:169:1
				xLog.Error("json marshal error:", err)
//line cmd/gopcomm/community_yap.gox:170:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:172:1
			ctx.Yap__1("edit", map[string]interface {
			}{"User": user, "Article": article})
		}
	})
//line cmd/gopcomm/community_yap.gox:179:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:180:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:181:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:182:1
		if label == "" {
//line cmd/gopcomm/community_yap.gox:183:1
			label = "article"
		}
//line cmd/gopcomm/community_yap.gox:187:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:188:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:189:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:190:1
			user, _ = this.community.GetUser(token.Value)
		}
//line cmd/gopcomm/community_yap.gox:193:1
		articles, next, err := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue, label)
//line cmd/gopcomm/community_yap.gox:194:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:195:1
			xLog.Error("get article error:", err)
		}
//line cmd/gopcomm/community_yap.gox:198:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:199:1
			xLog.Error("json marshal error:", err)
		}
//line cmd/gopcomm/community_yap.gox:201:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": user, "Items": articles, "Value": searchValue, "Next": next, "Tab": label})
	})
//line cmd/gopcomm/community_yap.gox:213:1
	this.Get("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:214:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:215:1
		article, err := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:216:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:217:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:218:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get article failed"})
		}
//line cmd/gopcomm/community_yap.gox:223:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:229:1
	this.Delete("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:230:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:231:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:232:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:233:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:234:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:235:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:240:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:241:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:242:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:247:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:254:1
	this.Get("/api/articles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:255:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:256:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:257:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:258:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:260:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:261:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:262:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:265:1
		articles, next, err := this.community.ListArticle(todo, from, limitInt, searchValue, label)
//line cmd/gopcomm/community_yap.gox:266:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:267:1
			xLog.Error("get article error:", err)
//line cmd/gopcomm/community_yap.gox:268:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get article failed"})
		}
//line cmd/gopcomm/community_yap.gox:273:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:281:1
	this.Post("/api/article/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:282:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:283:1
		content := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:284:1
		title := ctx.Param("title")
//line cmd/gopcomm/community_yap.gox:285:1
		tags := ctx.Param("tags")
//line cmd/gopcomm/community_yap.gox:286:1
		abstract := ctx.Param("abstract")
//line cmd/gopcomm/community_yap.gox:287:1
		label := ctx.Param("label")
//line cmd/gopcomm/community_yap.gox:288:1
		trans, err := strconv.ParseBool(ctx.Param("trans"))
//line cmd/gopcomm/community_yap.gox:289:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:290:1
			xLog.Error("parse bool error:", err)
		}
//line cmd/gopcomm/community_yap.gox:293:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:294:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:295:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:296:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:301:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:302:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:303:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:304:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:309:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: title, UId: uid, Cover: ctx.Param("cover"), Tags: tags, Abstract: abstract, Label: label}, Content: content, Trans: trans, Vtt_id: ctx.Param("vtt_id")}
//line cmd/gopcomm/community_yap.gox:325:1
		if trans {
//line cmd/gopcomm/community_yap.gox:326:1
			article, _ = this.community.TranslateArticle(todo, article)
		}
//line cmd/gopcomm/community_yap.gox:329:1
		id, err = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:330:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:331:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:332:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "add failed"})
		}
//line cmd/gopcomm/community_yap.gox:337:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:346:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:347:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:349:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:350:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:351:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:354:1
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:355:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:356:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:357:1
			user, _ = this.community.GetUser(token.Value)
		}
//line cmd/gopcomm/community_yap.gox:360:1
		items, total, err := this.community.GetArticlesByUid(todo, id, firstConst, mediaLimitConst)
//line cmd/gopcomm/community_yap.gox:361:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:362:1
			xLog.Error("get article list error:", err)
		}
//line cmd/gopcomm/community_yap.gox:364:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": userClaim, "User": user, "Items": items, "Total": total})
	})
//line cmd/gopcomm/community_yap.gox:373:1
	this.Get("/user/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:374:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:375:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:376:1
			xLog.Error("get token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:378:1
		user, err := this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:379:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:380:1
			xLog.Info("get user error:", err)
//line cmd/gopcomm/community_yap.gox:381:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:383:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:384:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:385:1
			xLog.Error("get uid error:", err)
//line cmd/gopcomm/community_yap.gox:386:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:388:1
		userClaim, err := this.community.GetUserClaim(uid)
//line cmd/gopcomm/community_yap.gox:389:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:390:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:392:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:393:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:394:1
			xLog.Info("gac", err)
//line cmd/gopcomm/community_yap.gox:395:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:397:1
		appInfo, err := this.community.GetApplicationInfo()
//line cmd/gopcomm/community_yap.gox:398:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:399:1
			xLog.Error("get application info error:", err)
//line cmd/gopcomm/community_yap.gox:400:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:402:1
		ctx.Yap__1("user_edit", map[string]interface {
		}{"User": user, "CurrentUser": userClaim, "Application": appInfo, "Binds": gac.GetProviderBindStatus()})
	})
//line cmd/gopcomm/community_yap.gox:413:1
	this.Put("/api/user", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:414:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:415:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:416:1
		user := &core.UserInfo{Id: uid, Name: ctx.Param("name"), Birthday: ctx.Param("birthday"), Gender: ctx.Param("gender"), Phone: ctx.Param("phone"), Email: ctx.Param("email"), Avatar: ctx.Param("avatar"), Owner: ctx.Param("owner"), DisplayName: ctx.Param("displayName")}
//line cmd/gopcomm/community_yap.gox:428:1
		_, err = this.community.UpdateUser(user)
//line cmd/gopcomm/community_yap.gox:429:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:430:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:431:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "msg": "update failed"})
		}
//line cmd/gopcomm/community_yap.gox:436:1
		ctx.Json__1(map[string]int{"code": 200})
	})
//line cmd/gopcomm/community_yap.gox:441:1
	this.Get("/api/user/unlink", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:442:1
		pv := ctx.Param("pv")
//line cmd/gopcomm/community_yap.gox:443:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:444:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:445:1
			ctx.Json__1(map[string]int{"code": 200})
//line cmd/gopcomm/community_yap.gox:448:1
			return
		}
//line cmd/gopcomm/community_yap.gox:450:1
		switch pv {
//line cmd/gopcomm/community_yap.gox:451:1
		case "Twitter":
//line cmd/gopcomm/community_yap.gox:452:1
		case "Facebook":
//line cmd/gopcomm/community_yap.gox:453:1
		case "Github":
//line cmd/gopcomm/community_yap.gox:454:1
		case "WeChat":
//line cmd/gopcomm/community_yap.gox:455:1
		default:
//line cmd/gopcomm/community_yap.gox:456:1
			pv = ""
		}
//line cmd/gopcomm/community_yap.gox:458:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:459:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:460:1
			gac.UnLink(pv)
		}
//line cmd/gopcomm/community_yap.gox:462:1
		ctx.Json__1(map[string]int{"code": 200})
	})
//line cmd/gopcomm/community_yap.gox:467:1
	this.Get("/api/user/:id/articles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:468:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:469:1
		page := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:470:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:472:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:473:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:474:1
			limitInt = mediaLimitConst
		}
//line cmd/gopcomm/community_yap.gox:476:1
		items, total, err := this.community.GetArticlesByUid(todo, id, page, limitInt)
//line cmd/gopcomm/community_yap.gox:477:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:478:1
			xLog.Error("get article list error:", err)
//line cmd/gopcomm/community_yap.gox:479:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error(), "total": 0})
		}
//line cmd/gopcomm/community_yap.gox:485:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": items, "total": total})
	})
//line cmd/gopcomm/community_yap.gox:492:1
	this.Get("/api/user/:id/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:493:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:494:1
		uid := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:495:1
		page := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:496:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:498:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:499:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:500:1
			limitInt = mediaLimitConst
		}
//line cmd/gopcomm/community_yap.gox:502:1
		pageInt, err := strconv.Atoi(page)
//line cmd/gopcomm/community_yap.gox:503:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:504:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "total": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:510:1
		files, total, err := this.community.ListMediaByUserId(todo, uid, format, pageInt, limitInt)
//line cmd/gopcomm/community_yap.gox:511:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:512:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "total": total, "err": err.Error()})
		} else {
//line cmd/gopcomm/community_yap.gox:518:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "total": total, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:529:1
	this.Delete("/api/media/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:530:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:531:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:532:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:533:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:534:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:535:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:540:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:541:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:542:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:547:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:554:1
	this.Get("/api/translation/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:555:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:556:1
		article, err := this.community.GetTranslateArticle(todo, id)
//line cmd/gopcomm/community_yap.gox:558:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:559:1
			xLog.Info(err)
//line cmd/gopcomm/community_yap.gox:560:1
			ctx.Json__1(map[string]int{"code": 0})
		}
//line cmd/gopcomm/community_yap.gox:564:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "content": article.Content, "tags": article.Tags, "title": article.Title})
	})
//line cmd/gopcomm/community_yap.gox:572:1
	this.Post("/api/translation", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:573:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:574:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, "auto", language.English)
//line cmd/gopcomm/community_yap.gox:575:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:576:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:581:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:587:1
	this.Get("/api/media/:id/url", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:588:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:589:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:590:1
		htmlUrl := fmt.Sprintf("%s%s", this.domain, fileKey)
//line cmd/gopcomm/community_yap.gox:591:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:592:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:597:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:603:1
	this.Get("/api/video/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:604:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:605:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:606:1
		m := make(map[string]string, 4)
//line cmd/gopcomm/community_yap.gox:607:1
		format, err := this.community.GetMediaType(todo, id)
//line cmd/gopcomm/community_yap.gox:608:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:609:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:614:1
		htmlUrl := fmt.Sprintf("%s%s", this.domain, fileKey)
//line cmd/gopcomm/community_yap.gox:615:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:616:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:621:1
		m["fileKey"] = htmlUrl
//line cmd/gopcomm/community_yap.gox:622:1
		m["type"] = format
//line cmd/gopcomm/community_yap.gox:623:1
		m["subtitle"] = this.domain + id
//line cmd/gopcomm/community_yap.gox:624:1
		m["status"] = "0"
//line cmd/gopcomm/community_yap.gox:625:1
		match, _ := regexp.MatchString("^video", format)
//line cmd/gopcomm/community_yap.gox:626:1
		if match {
//line cmd/gopcomm/community_yap.gox:627:1
			subtitle, status, err := this.community.GetVideoSubtitle(todo, id)
//line cmd/gopcomm/community_yap.gox:628:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:629:1
				ctx.Json__1(map[string]interface {
				}{"code": 200, "url": m})
			}
//line cmd/gopcomm/community_yap.gox:634:1
			if status == "1" {
//line cmd/gopcomm/community_yap.gox:635:1
				m["subtitle"] = this.domain + subtitle
			}
//line cmd/gopcomm/community_yap.gox:637:1
			m["status"] = status
		}
//line cmd/gopcomm/community_yap.gox:639:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": m})
	})
//line cmd/gopcomm/community_yap.gox:645:1
	this.Post("/api/media", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:646:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:649:1
	this.Post("/api/caption/task", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:650:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:651:1
		vid := ctx.Param("vid")
//line cmd/gopcomm/community_yap.gox:653:1
		if uid == "" || vid == "" {
//line cmd/gopcomm/community_yap.gox:654:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Invalid param"})
		}
//line cmd/gopcomm/community_yap.gox:660:1
		if
//line cmd/gopcomm/community_yap.gox:660:1
		err := this.community.RetryCaptionGenerate(todo, uid, vid); err != nil {
//line cmd/gopcomm/community_yap.gox:661:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Request task error"})
		}
//line cmd/gopcomm/community_yap.gox:667:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "msg": "Ok"})
	})
//line cmd/gopcomm/community_yap.gox:676:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:681:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:682:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:683:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:684:1
			return
		}
//line cmd/gopcomm/community_yap.gox:687:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:688:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:689:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:692:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "login/callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:694:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:695:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:698:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:699:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:700:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:701:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:704:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:705:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:706:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:707:1
			return
		}
//line cmd/gopcomm/community_yap.gox:710:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:711:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:712:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:715:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:718:1
	this.Get("/login/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:719:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:720:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:721:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:723:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:724:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:725:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:726:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:727:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:730:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:733:1
	this.Post("/api/article/like", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:734:1
		articleId := ctx.Param("articleId")
//line cmd/gopcomm/community_yap.gox:735:1
		articleIdInt, err := strconv.Atoi(articleId)
//line cmd/gopcomm/community_yap.gox:736:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:737:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:742:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:743:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:744:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:745:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:750:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:751:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:752:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:757:1
		res, err := this.community.ArticleLike(todo, articleIdInt, uid)
//line cmd/gopcomm/community_yap.gox:758:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:759:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:764:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": res})
	})
//line cmd/gopcomm/community_yap.gox:771:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:772:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:775:1
	configFile := ".env"
//line cmd/gopcomm/community_yap.gox:776:1
	flag.StringVar(&configFile, "config", ".env", "Path to the config file")
//line cmd/gopcomm/community_yap.gox:777:1
	flag.Parse()
//line cmd/gopcomm/community_yap.gox:778:1
	if
//line cmd/gopcomm/community_yap.gox:778:1
	err := godotenv.Load(configFile); err != nil {
//line cmd/gopcomm/community_yap.gox:779:1
		xLog.Error(err)
	}
//line cmd/gopcomm/community_yap.gox:782:1
	conf := core.NewConfigFromEnv()
//line cmd/gopcomm/community_yap.gox:783:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:784:1
	this.domain = conf.QiNiuConfig.Domain
//line cmd/gopcomm/community_yap.gox:785:1
	endpoint := conf.AppConfig.EndPoint
//line cmd/gopcomm/community_yap.gox:787:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:790:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:792:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:793:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:795:1
				if
//line cmd/gopcomm/community_yap.gox:795:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:796:1
					xLog.Error(err)
//line cmd/gopcomm/community_yap.gox:797:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:801:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
