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
		token, err := core.GetToken(ctx)
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
		token, err := core.GetToken(ctx)
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
		userClaimJson, _ := json.Marshal(&userClaim)
//line cmd/gopcomm/community_yap.gox:109:1
		itemsJson, _ := json.Marshal(&items)
//line cmd/gopcomm/community_yap.gox:110:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": strings.Replace(string(userClaimJson), `\"`, `"`, -1), "User": user, "Items": strings.Replace(string(itemsJson), `\"`, `"`, -1)})
	})
//line cmd/gopcomm/community_yap.gox:118:1
	this.Get("/userEdit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:119:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:120:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:121:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:122:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:123:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:124:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:128:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:129:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:131:1
		ctx.Yap__1("user_edit", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:136:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:137:1
		ctx.Yap__1("edit", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:140:1
	this.Get("/delete", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:141:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:142:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:143:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:144:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:145:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:146:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:151:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:152:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:153:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:158:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:165:1
	this.Get("/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:166:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:167:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:168:1
		files, err := this.community.ListMediaByUserId(todo, uid, format)
//line cmd/gopcomm/community_yap.gox:169:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:170:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get media failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:175:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:182:1
	this.Get("/delMedia", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:183:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:184:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:185:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:186:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:187:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:188:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:193:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:194:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:195:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:200:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:207:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:209:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:210:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:211:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:212:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:213:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:214:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:215:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:217:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:221:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, "")
//line cmd/gopcomm/community_yap.gox:222:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:223:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:231:1
	this.Get("/get", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:232:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:233:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:234:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:236:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:237:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:238:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:241:1
		articles, next, _ := this.community.ListArticle(todo, from, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:242:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:250:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:251:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:252:1
		if searchValue == "" {
//line cmd/gopcomm/community_yap.gox:253:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "err": "value can not be ''."})
		}
//line cmd/gopcomm/community_yap.gox:260:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:261:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:262:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:263:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:264:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:265:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:266:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:268:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:272:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue)
//line cmd/gopcomm/community_yap.gox:273:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:274:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Value": searchValue, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:283:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:284:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:285:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:286:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:292:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:293:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:294:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:300:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:301:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:302:1
			if
//line cmd/gopcomm/community_yap.gox:302:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:303:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:304:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:306:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:307:1
			ctx.Yap__1("edit", article)
		}
	})
//line cmd/gopcomm/community_yap.gox:311:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:312:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:313:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:314:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:315:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:320:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:327:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:329:1
		trans := ctx.Param("trans")
//line cmd/gopcomm/community_yap.gox:330:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:331:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:332:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:334:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:335:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:336:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:341:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:342:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:343:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:350:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags"), Abstract: ctx.Param("abstract")}, Content: mdData, HtmlData: htmlData}
//line cmd/gopcomm/community_yap.gox:362:1
		id, _ = this.community.PutArticle(todo, uid, trans, article)
//line cmd/gopcomm/community_yap.gox:363:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:371:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:373:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:374:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:375:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:380:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:381:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:382:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:388:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:389:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:390:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:392:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:393:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:394:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:399:1
		id, _ = this.community.SaveHtml(todo, uid, htmlData, mdData, id)
//line cmd/gopcomm/community_yap.gox:400:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "id": id, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:407:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:408:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:410:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:412:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:415:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:416:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:417:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:418:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:419:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:420:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:425:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:431:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:432:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:435:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:440:1
		xLog.Info("ctx.Request.Referer()", ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:441:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:442:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:443:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:444:1
			return
		}
//line cmd/gopcomm/community_yap.gox:447:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:448:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:449:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:452:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:454:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:455:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:458:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:459:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:460:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:461:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:464:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:465:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:466:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:467:1
			return
		}
//line cmd/gopcomm/community_yap.gox:470:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:471:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:472:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:475:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:478:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:479:1
		xLog.Info("ctx.URL", ctx.URL)
//line cmd/gopcomm/community_yap.gox:480:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:481:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:482:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:485:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:486:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:487:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:488:1
			xLog.Info("Unurlerror", err)
//line cmd/gopcomm/community_yap.gox:489:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:494:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:497:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:498:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:499:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:502:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:503:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:506:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:509:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:511:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:512:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:513:1
				if
//line cmd/gopcomm/community_yap.gox:513:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:514:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:518:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
