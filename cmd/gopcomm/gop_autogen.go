package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/goplus/yap"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	_ "github.com/joho/godotenv/autoload"
	gopaccountsdk "github.com/liuscraft/gop-casdoor-account-sdk"
	"github.com/qiniu/x/xlog"
	"golang.org/x/text/language"
)

const (
	layoutUS   = "January 2, 2006"
	limitConst = 10
)

type community_yap struct {
	yap.App
	community *core.Community
	trans     *translation.Engine
}
//line cmd/gopcomm/community_yap.gox:28
func (this *community_yap) MainEntry() {
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
	this.Get("/signin", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:54:1
		ctx.Yap__1("signin", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:57:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:60:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:61:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:62:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:63:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:64:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:65:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:66:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:68:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:72:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:73:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:74:1
		articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:75:1
		ctx.Yap__1("article", map[string]interface {
		}{"UserId": userId, "User": user, "Article": string(articleJson)})
	})
//line cmd/gopcomm/community_yap.gox:82:1
	this.Get("/getArticle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:83:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:84:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:85:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:91:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:92:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:94:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:95:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:96:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:99:1
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:100:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:101:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:102:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:103:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:104:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:105:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:107:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:111:1
		items, next, _ := this.community.GetArticlesByUid(todo, id, core.MarkBegin, limitConst)
//line cmd/gopcomm/community_yap.gox:112:1
		userClaimJson, _ := json.Marshal(&userClaim)
//line cmd/gopcomm/community_yap.gox:113:1
		itemsJson, _ := json.Marshal(&items)
//line cmd/gopcomm/community_yap.gox:114:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": string(userClaimJson), "User": user, "Items": string(itemsJson), "UserId": userId, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:123:1
	this.Get("/userUnlink", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:124:1
		pv := ctx.Param("pv")
//line cmd/gopcomm/community_yap.gox:125:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:126:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:127:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/login", http.StatusTemporaryRedirect)
//line cmd/gopcomm/community_yap.gox:128:1
			return
		}
//line cmd/gopcomm/community_yap.gox:130:1
		switch pv {
//line cmd/gopcomm/community_yap.gox:131:1
		case "Twitter":
//line cmd/gopcomm/community_yap.gox:132:1
		case "Facebook":
//line cmd/gopcomm/community_yap.gox:133:1
		case "Github":
//line cmd/gopcomm/community_yap.gox:134:1
		case "WeChat":
//line cmd/gopcomm/community_yap.gox:135:1
		default:
//line cmd/gopcomm/community_yap.gox:136:1
			pv = ""
		}
//line cmd/gopcomm/community_yap.gox:138:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:139:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:140:1
			gac.UnLink(pv)
		}
//line cmd/gopcomm/community_yap.gox:142:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, "/userEdit", http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:144:1
	this.Get("/userEdit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:145:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:146:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:147:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:149:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:150:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:151:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:153:1
		appInfo, _ := this.community.GetApplicationInfo()
//line cmd/gopcomm/community_yap.gox:154:1
		appInfoStr, _ := json.Marshal(*appInfo)
//line cmd/gopcomm/community_yap.gox:155:1
		binds, _ := json.Marshal(gac.GetProviderBindStatus())
//line cmd/gopcomm/community_yap.gox:156:1
		user := gac.GetUserSimple()
//line cmd/gopcomm/community_yap.gox:157:1
		ctx.Yap__1("user_edit", map[string]interface {
		}{"UserId": user.Id, "User": user, "Application": string(appInfoStr), "Binds": string(binds)})
	})
//line cmd/gopcomm/community_yap.gox:165:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:166:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:167:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:168:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:169:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:170:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:171:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:172:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:174:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:177:1
		ctx.Yap__1("edit", map[string]interface {
		}{"User": user, "UserId": userId})
	})
//line cmd/gopcomm/community_yap.gox:183:1
	this.Get("/delete", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:184:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:185:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:186:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:187:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:188:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:189:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:194:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:195:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:196:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:201:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:208:1
	this.Get("/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:209:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:210:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:211:1
		page := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:212:1
		pageInt, err := strconv.Atoi(page)
//line cmd/gopcomm/community_yap.gox:213:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:214:1
			pageInt = 1
		}
//line cmd/gopcomm/community_yap.gox:216:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:217:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:218:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:219:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:221:1
		files, count, err := this.community.ListMediaByUserId(todo, uid, format, pageInt, limitInt)
//line cmd/gopcomm/community_yap.gox:222:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:223:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "count": count, "err": err.Error()})
		} else {
//line cmd/gopcomm/community_yap.gox:229:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "count": count, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:236:1
	this.Get("/delMedia", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:237:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:238:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:239:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:240:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:241:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:242:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:247:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:248:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:249:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:254:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:261:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:263:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:264:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:265:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:266:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:267:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:268:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:269:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:271:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:275:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, "")
//line cmd/gopcomm/community_yap.gox:276:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:277:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": string(articlesJson), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:285:1
	this.Get("/get", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:286:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:287:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:288:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:290:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:291:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:292:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:295:1
		articles, next, _ := this.community.ListArticle(todo, from, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:296:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:304:1
	this.Get("/userArticles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:305:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:306:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:307:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:309:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:310:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:311:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:313:1
		items, next, _ := this.community.GetArticlesByUid(todo, id, from, limitInt)
//line cmd/gopcomm/community_yap.gox:314:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": items, "next": next})
	})
//line cmd/gopcomm/community_yap.gox:321:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:322:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:331:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:332:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:333:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:334:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:335:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:336:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:337:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:339:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:343:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue)
//line cmd/gopcomm/community_yap.gox:344:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:345:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": string(articlesJson), "Value": searchValue, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:354:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:355:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:356:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:357:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:358:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:359:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:360:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:361:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:363:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:367:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:368:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:369:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:370:1
			if
//line cmd/gopcomm/community_yap.gox:370:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:371:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:372:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:374:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:375:1
			articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:376:1
			ctx.Yap__1("edit", map[string]interface {
			}{"UserId": userId, "User": user, "Article": string(articleJson)})
		}
	})
//line cmd/gopcomm/community_yap.gox:384:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:385:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:386:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:387:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:388:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:393:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:400:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:401:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:402:1
		content := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:403:1
		title := ctx.Param("title")
//line cmd/gopcomm/community_yap.gox:404:1
		tags := ctx.Param("tags")
//line cmd/gopcomm/community_yap.gox:405:1
		abstract := ctx.Param("abstract")
//line cmd/gopcomm/community_yap.gox:406:1
		trans, _ := strconv.ParseBool(ctx.Param("trans"))
//line cmd/gopcomm/community_yap.gox:408:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:409:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:410:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:411:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:416:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:417:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:418:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:419:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:424:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: title, UId: uid, Cover: ctx.Param("cover"), Tags: tags, Abstract: abstract}, Content: content, Trans: trans}
//line cmd/gopcomm/community_yap.gox:438:1
		if trans {
//line cmd/gopcomm/community_yap.gox:439:1
			article, _ = this.community.TranslateArticle(todo, article)
		}
//line cmd/gopcomm/community_yap.gox:442:1
		id, err = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:443:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:444:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "add failed"})
		}
//line cmd/gopcomm/community_yap.gox:449:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:455:1
	this.Get("/tranArticle", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:456:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:458:1
		article, err := this.community.GetTranslateArticle(todo, id)
//line cmd/gopcomm/community_yap.gox:459:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:460:1
			ctx.Json__1(map[string]int{"code": 0})
		}
//line cmd/gopcomm/community_yap.gox:464:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "content": article.Content, "tags": article.Tags, "title": article.Title})
	})
//line cmd/gopcomm/community_yap.gox:473:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:474:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:475:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, "auto", language.English)
//line cmd/gopcomm/community_yap.gox:476:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:477:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:482:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:489:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:490:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:492:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:494:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:497:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:498:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:499:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:500:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:501:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:502:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:507:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:513:1
	this.Get("/getVideoAndSubtitle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:514:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:515:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:516:1
		m := make(map[string]string, 2)
//line cmd/gopcomm/community_yap.gox:517:1
		format, err := this.community.GetMediaType(todo, id)
//line cmd/gopcomm/community_yap.gox:518:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:519:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:524:1
		if format == "video/mp4" {
//line cmd/gopcomm/community_yap.gox:525:1
			subtitle, err := this.community.GetVideoSubtitle(todo, id)
//line cmd/gopcomm/community_yap.gox:526:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:527:1
				if err != nil {
//line cmd/gopcomm/community_yap.gox:528:1
					ctx.Json__1(map[string]interface {
					}{"code": 500, "err": err.Error()})
				}
//line cmd/gopcomm/community_yap.gox:533:1
				return
			}
//line cmd/gopcomm/community_yap.gox:535:1
			m["subtitle"] = domain + subtitle
		}
//line cmd/gopcomm/community_yap.gox:537:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:538:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:539:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:544:1
		m["fileKey"] = htmlUrl
//line cmd/gopcomm/community_yap.gox:545:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": m})
	})
//line cmd/gopcomm/community_yap.gox:551:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:552:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:555:1
	this.Post("/caption/task", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:556:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:557:1
		vid := ctx.Param("vid")
//line cmd/gopcomm/community_yap.gox:559:1
		if uid == "" || vid == "" {
//line cmd/gopcomm/community_yap.gox:560:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Invalid param"})
		}
//line cmd/gopcomm/community_yap.gox:566:1
		if
//line cmd/gopcomm/community_yap.gox:566:1
		err := this.community.RetryCaptionGenerate(todo, uid, vid); err != nil {
//line cmd/gopcomm/community_yap.gox:567:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Request task error"})
		}
//line cmd/gopcomm/community_yap.gox:573:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "msg": "Ok"})
	})
//line cmd/gopcomm/community_yap.gox:579:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:584:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:585:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:586:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:587:1
			return
		}
//line cmd/gopcomm/community_yap.gox:590:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:591:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:592:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:595:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:597:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:598:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:601:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:602:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:603:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:604:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:607:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:608:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:609:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:610:1
			return
		}
//line cmd/gopcomm/community_yap.gox:613:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:614:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:615:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:618:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:621:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:622:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:623:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:624:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:626:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:627:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:628:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:629:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:630:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:633:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:636:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:637:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:638:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:641:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:642:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:645:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:648:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:650:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:651:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:652:1
				if
//line cmd/gopcomm/community_yap.gox:652:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:653:1
					xLog.Error(err)
//line cmd/gopcomm/community_yap.gox:654:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:658:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community_yap))
}
