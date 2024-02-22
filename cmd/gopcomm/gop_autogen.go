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
		ctx.Yap__1("2xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:42:1
	this.Get("/error", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:43:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:46:1
	this.Get("/failed", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:47:1
		ctx.Yap__1("5xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:50:1
	this.Get("/demo", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:51:1
		ctx.Yap__1("demo", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:54:1
	this.Get("/signin", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:55:1
		ctx.Yap__1("signin", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:58:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:61:1
		// todo middleware
		// Get User Info
		var user *core.User
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
			}
		}
//line cmd/gopcomm/community_yap.gox:70:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:71:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:72:1
		articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:73:1
		ctx.Yap__1("article", map[string]interface {
		}{"User": user, "Article": strings.Replace(string(articleJson), `\"`, `"`, -1)})
	})
//line cmd/gopcomm/community_yap.gox:79:1
	this.Get("/getArticle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:80:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:81:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:82:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:88:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:89:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:91:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:92:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:93:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:97:1
		// todo middleware
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:98:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:99:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:100:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:101:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:102:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:105:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:106:1
		if user != nil {
//line cmd/gopcomm/community_yap.gox:107:1
			userId = user.Id
		}
//line cmd/gopcomm/community_yap.gox:110:1
		items, _ := this.community.GetArticlesByUid(todo, id)
//line cmd/gopcomm/community_yap.gox:111:1
		userClaimJson, _ := json.Marshal(&userClaim)
//line cmd/gopcomm/community_yap.gox:112:1
		itemsJson, _ := json.Marshal(&items)
//line cmd/gopcomm/community_yap.gox:113:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": strings.Replace(string(userClaimJson), `\"`, `"`, -1), "User": user, "Items": strings.Replace(string(itemsJson), `\"`, `"`, -1), "UserId": userId})
	})
//line cmd/gopcomm/community_yap.gox:121:1
	this.Get("/userUnlink", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:122:1
		pv := ctx.Param("pv")
//line cmd/gopcomm/community_yap.gox:123:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:124:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:125:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/login", http.StatusTemporaryRedirect)
//line cmd/gopcomm/community_yap.gox:126:1
			return
		}
//line cmd/gopcomm/community_yap.gox:128:1
		switch pv {
//line cmd/gopcomm/community_yap.gox:129:1
		case "Twitter":
//line cmd/gopcomm/community_yap.gox:130:1
		case "Facebook":
//line cmd/gopcomm/community_yap.gox:131:1
		case "Github":
//line cmd/gopcomm/community_yap.gox:132:1
		case "WeChat":
//line cmd/gopcomm/community_yap.gox:133:1
		default:
//line cmd/gopcomm/community_yap.gox:134:1
			pv = ""
		}
//line cmd/gopcomm/community_yap.gox:136:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:137:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:138:1
			gac.UnLink(pv)
		}
//line cmd/gopcomm/community_yap.gox:140:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, "/userEdit", http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:142:1
	this.Get("/userEdit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:143:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:144:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:145:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:147:1
		gac, err := gopaccountsdk.GetClient(token.Value)
//line cmd/gopcomm/community_yap.gox:148:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:149:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:151:1
		appInfo, _ := this.community.GetApplicationInfo()
//line cmd/gopcomm/community_yap.gox:152:1
		appInfoStr, _ := json.Marshal(*appInfo)
//line cmd/gopcomm/community_yap.gox:153:1
		binds, _ := json.Marshal(gac.GetProviderBindStatus())
//line cmd/gopcomm/community_yap.gox:154:1
		ctx.Yap__1("user_edit", map[string]interface {
		}{"User": gac.GetUserSimple(), "Application": string(appInfoStr), "Binds": string(binds)})
	})
//line cmd/gopcomm/community_yap.gox:161:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:162:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:163:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:164:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:165:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:166:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:167:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:170:1
		ctx.Yap__1("edit", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:175:1
	this.Get("/delete", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:176:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:177:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:178:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:179:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:180:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:181:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:186:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:187:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:188:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:193:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:200:1
	this.Get("/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:201:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:202:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:203:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:204:1
		fromInt, err := strconv.Atoi(from)
//line cmd/gopcomm/community_yap.gox:205:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:206:1
			fromInt = 0
		}
//line cmd/gopcomm/community_yap.gox:208:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:209:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:210:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:211:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:213:1
		files, next, err := this.community.ListMediaByUserId(todo, uid, format, fromInt, limitInt)
//line cmd/gopcomm/community_yap.gox:214:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:215:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "next": next, "err": err.Error()})
		} else {
//line cmd/gopcomm/community_yap.gox:221:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "next": next, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:229:1
	this.Get("/delMedia", func(ctx *yap.Context) {
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
		err = this.community.DelMedia(todo, uid, id)
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
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:256:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:257:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:258:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:259:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:260:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:261:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:262:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:264:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:268:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, "")
//line cmd/gopcomm/community_yap.gox:269:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:270:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:278:1
	this.Get("/get", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:279:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:280:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:281:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:283:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:284:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:285:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:288:1
		articles, next, _ := this.community.ListArticle(todo, from, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:289:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:297:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:298:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:307:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:308:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:309:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:310:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:311:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:312:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:313:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:315:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:319:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue)
//line cmd/gopcomm/community_yap.gox:320:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:321:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Value": searchValue, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:330:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:331:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:332:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:333:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:334:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:335:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:336:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:340:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:341:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:342:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:343:1
			if
//line cmd/gopcomm/community_yap.gox:343:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:344:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:345:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:347:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:348:1
			articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:349:1
			ctx.Yap__1("edit", map[string]interface {
			}{"User": user, "Article": strings.Replace(string(articleJson), `\"`, `"`, -1)})
		}
	})
//line cmd/gopcomm/community_yap.gox:356:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:357:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:358:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:359:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:360:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:365:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:372:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:373:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:374:1
		content := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:375:1
		title := ctx.Param("title")
//line cmd/gopcomm/community_yap.gox:376:1
		tags := ctx.Param("tags")
//line cmd/gopcomm/community_yap.gox:377:1
		abstract := ctx.Param("abstract")
//line cmd/gopcomm/community_yap.gox:378:1
		trans, _ := strconv.ParseBool(ctx.Param("trans"))
//line cmd/gopcomm/community_yap.gox:380:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:381:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:382:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:383:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:388:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:389:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:390:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:391:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:396:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: title, UId: uid, Cover: ctx.Param("cover"), Tags: tags, Abstract: abstract}, Content: content, Trans: trans}
//line cmd/gopcomm/community_yap.gox:410:1
		if trans {
//line cmd/gopcomm/community_yap.gox:411:1
			article, _ = this.community.TranslateArticle(todo, article)
		}
//line cmd/gopcomm/community_yap.gox:414:1
		id, err = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:415:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:416:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "add failed"})
		}
//line cmd/gopcomm/community_yap.gox:421:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:427:1
	this.Get("/tranArticle", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:428:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:430:1
		article, err := this.community.GetTranslateArticle(todo, id)
//line cmd/gopcomm/community_yap.gox:431:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:432:1
			ctx.Json__1(map[string]int{"code": 0})
		}
//line cmd/gopcomm/community_yap.gox:436:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "content": article.Content, "tags": article.Tags, "title": article.Title})
	})
//line cmd/gopcomm/community_yap.gox:445:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:446:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:447:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, "auto", language.English)
//line cmd/gopcomm/community_yap.gox:448:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:449:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:454:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:461:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:462:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:464:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:466:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:469:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:470:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:471:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:472:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:473:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:474:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:479:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:485:1
	this.Get("/getVideoAndSubtitle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:486:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:487:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:488:1
		m := make(map[string]string, 2)
//line cmd/gopcomm/community_yap.gox:489:1
		format, err := this.community.GetMediaType(todo, id)
//line cmd/gopcomm/community_yap.gox:490:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:491:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:496:1
		if format == "video/mp4" {
//line cmd/gopcomm/community_yap.gox:497:1
			subtitle, err := this.community.GetVideoSubtitle(todo, id)
//line cmd/gopcomm/community_yap.gox:498:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:499:1
				if err != nil {
//line cmd/gopcomm/community_yap.gox:500:1
					ctx.Json__1(map[string]interface {
					}{"code": 500, "err": err.Error()})
				}
//line cmd/gopcomm/community_yap.gox:505:1
				return
			}
//line cmd/gopcomm/community_yap.gox:507:1
			m["subtitle"] = domain + subtitle
		}
//line cmd/gopcomm/community_yap.gox:509:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:510:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:511:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:516:1
		m["fileKey"] = htmlUrl
//line cmd/gopcomm/community_yap.gox:517:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": m})
	})
//line cmd/gopcomm/community_yap.gox:523:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:524:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:527:1
	this.Post("/caption/task", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:528:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:529:1
		vid := ctx.Param("vid")
//line cmd/gopcomm/community_yap.gox:531:1
		if uid == "" || vid == "" {
//line cmd/gopcomm/community_yap.gox:532:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Invalid param"})
		}
//line cmd/gopcomm/community_yap.gox:538:1
		if
//line cmd/gopcomm/community_yap.gox:538:1
		err := this.community.RetryCaptionGenerate(todo, uid, vid); err != nil {
//line cmd/gopcomm/community_yap.gox:539:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "Request task error"})
		}
//line cmd/gopcomm/community_yap.gox:545:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "msg": "Ok"})
	})
//line cmd/gopcomm/community_yap.gox:551:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:556:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:557:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:558:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:559:1
			return
		}
//line cmd/gopcomm/community_yap.gox:562:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:563:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:564:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:567:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:569:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:570:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:573:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:574:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:575:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:576:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:579:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:580:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:581:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:582:1
			return
		}
//line cmd/gopcomm/community_yap.gox:585:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:586:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:587:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:590:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:593:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:594:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:595:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:596:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:598:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:599:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:600:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:601:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:602:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:605:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:608:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:609:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:610:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:613:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:614:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:617:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:620:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:622:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:623:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:624:1
				if
//line cmd/gopcomm/community_yap.gox:624:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:625:1
					xLog.Error(err)
//line cmd/gopcomm/community_yap.gox:626:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:630:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
