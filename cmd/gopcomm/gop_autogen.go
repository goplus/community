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
//line cmd/gopcomm/community_yap.gox:32
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:32:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:33:1
	endpoint := os.Getenv("GOP_COMMUNITY_ENDPOINT")
//line cmd/gopcomm/community_yap.gox:34:1
	domain := os.Getenv("GOP_COMMUNITY_DOMAIN")
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
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:49:1
	this.Get("/failed", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:50:1
		ctx.Yap__1("5xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:53:1
	this.Get("/demo", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:54:1
		ctx.Yap__1("demo", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:57:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:60:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:61:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:62:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:63:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:64:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:65:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:69:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:70:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:71:1
		ctx.Yap__1("article", map[string]interface {
		}{"User": user, "ID": id, "Title": article.Title, "Content": article.Content, "HtmlUrl": article.HtmlUrl, "Tags": article.Tags, "Cover": article.Cover, "Mtime": article.Mtime.Format(layoutUS), "Author": article.User})
	})
//line cmd/gopcomm/community_yap.gox:85:1
	this.Get("/getArticle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:86:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:87:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:88:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:94:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:95:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:97:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:98:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:99:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:103:1
		// todo middleware
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:104:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:105:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:106:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:107:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:108:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:112:1
		items, _ := this.community.GetArticlesByUid(todo, id)
//line cmd/gopcomm/community_yap.gox:113:1
		userClaimJson, _ := json.Marshal(&userClaim)
//line cmd/gopcomm/community_yap.gox:114:1
		itemsJson, _ := json.Marshal(&items)
//line cmd/gopcomm/community_yap.gox:115:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": strings.Replace(string(userClaimJson), `\"`, `"`, -1), "User": user, "Items": strings.Replace(string(itemsJson), `\"`, `"`, -1)})
	})
//line cmd/gopcomm/community_yap.gox:123:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:124:1
		ctx.Yap__1("edit", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:127:1
	this.Get("/delete", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:128:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:129:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:130:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:131:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:132:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:133:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:138:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:139:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:140:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:145:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:152:1
	this.Get("/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:153:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:154:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:155:1
		files, err := this.community.ListMediaByUserId(todo, uid, format)
//line cmd/gopcomm/community_yap.gox:156:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:157:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get media failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:162:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:169:1
	this.Get("/delMedia", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:170:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:171:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:172:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:173:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:174:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:175:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:180:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:181:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:182:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:187:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:194:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:196:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:197:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:198:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:199:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:200:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:201:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:202:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:204:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:208:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, "")
//line cmd/gopcomm/community_yap.gox:209:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:210:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:218:1
	this.Get("/get", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:219:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:220:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:221:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:223:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:224:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:225:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:228:1
		articles, next, _ := this.community.ListArticle(todo, from, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:229:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:237:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:238:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:239:1
		if searchValue == "" {
//line cmd/gopcomm/community_yap.gox:240:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "err": "value can not be ''."})
		}
//line cmd/gopcomm/community_yap.gox:247:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:248:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:249:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:250:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:251:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:252:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:253:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:255:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:259:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue)
//line cmd/gopcomm/community_yap.gox:260:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:261:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Value": searchValue, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:270:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:271:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:272:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:273:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:279:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:280:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:281:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:287:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:288:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:289:1
			if
//line cmd/gopcomm/community_yap.gox:289:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:290:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:291:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:293:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:294:1
			ctx.Yap__1("edit", article)
		}
	})
//line cmd/gopcomm/community_yap.gox:298:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:299:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:300:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:301:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:302:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:307:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:314:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:317:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:318:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:321:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:322:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:323:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:324:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:329:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:330:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:331:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:332:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:338:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Tags: ctx.Param("tags"), Abstract: ctx.Param("abstract")}, Content: mdData}
//line cmd/gopcomm/community_yap.gox:350:1
		id, _ = this.community.PutArticle(todo, uid, "", article)
//line cmd/gopcomm/community_yap.gox:351:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:359:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:361:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:362:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:363:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:368:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:369:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:370:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:376:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:377:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:378:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:380:1
		transData, err := this.trans.TranslateMarkdownText(mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:381:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:382:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:387:1
		id, _ = this.community.SaveHtml(todo, uid, htmlData, mdData, id)
//line cmd/gopcomm/community_yap.gox:388:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "id": id, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:395:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:396:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:398:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:400:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:403:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:404:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:405:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:406:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:407:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:408:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:413:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:419:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:420:1
		core.UploadFile(ctx, this.community)
	})
//line cmd/gopcomm/community_yap.gox:423:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:428:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:429:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:430:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:431:1
			return
		}
//line cmd/gopcomm/community_yap.gox:434:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:435:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:436:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:439:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:441:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:442:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:446:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:447:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:448:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:449:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:452:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:453:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:454:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:455:1
			return
		}
//line cmd/gopcomm/community_yap.gox:458:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:459:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:460:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:463:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:466:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:467:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:468:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:469:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:471:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:472:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:473:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:474:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:475:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:478:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:481:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:482:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:483:1
	this.trans = translation.New(os.Getenv("NIUTRANS_API_KEY"), "", "")
//line cmd/gopcomm/community_yap.gox:484:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:487:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:488:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:491:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:494:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:496:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:497:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:498:1
				if
//line cmd/gopcomm/community_yap.gox:498:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:499:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:503:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
