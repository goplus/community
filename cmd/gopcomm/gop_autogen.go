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
	"time"
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
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:47:1
	this.Get("/failed", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:48:1
		ctx.Yap__1("5xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:51:1
	this.Get("/signin", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:52:1
		ctx.Yap__1("signin", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:56:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:58:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:59:1
		userId := ""
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
			} else {
//line cmd/gopcomm/community_yap.gox:66:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:70:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, "")
//line cmd/gopcomm/community_yap.gox:71:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:72:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:83:1
	this.Get("/page/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:86:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:87:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:88:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:89:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:90:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:91:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:95:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:96:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:97:1
		articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:98:1
		ctx.Yap__1("article", map[string]interface {
		}{"User": user, "Article": strings.Replace(string(articleJson), `\"`, `"`, -1)})
	})
//line cmd/gopcomm/community_yap.gox:104:1
	this.Get("/page/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:105:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:106:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:107:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:108:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:109:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:110:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:113:1
		ctx.Yap__1("edit", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:118:1
	this.Get("/page/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:119:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:121:1
		_, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:122:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:123:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:126:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:127:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:128:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:129:1
			if
//line cmd/gopcomm/community_yap.gox:129:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:130:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:131:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:133:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:134:1
			articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:135:1
			ctx.Yap__1("edit", map[string]interface {
			}{"User": user, "Article": strings.Replace(string(articleJson), `\"`, `"`, -1)})
		}
	})
//line cmd/gopcomm/community_yap.gox:145:1
	this.Get("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:146:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:147:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:148:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:154:1
	this.Delete("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:155:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:156:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:157:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:158:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:159:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:160:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:165:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:166:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:167:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:172:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:179:1
	this.Get("/api/articles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:180:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:181:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:182:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:184:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:185:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:186:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:189:1
		articles, next, _ := this.community.ListArticle(todo, from, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:190:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:198:1
	this.Get("/api/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:199:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:202:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:203:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:204:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:205:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:206:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:207:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:208:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:210:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:214:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue)
//line cmd/gopcomm/community_yap.gox:215:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:216:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Value": searchValue, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:225:1
	this.Post("/api/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:228:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:229:1
		content := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:230:1
		title := ctx.Param("title")
//line cmd/gopcomm/community_yap.gox:231:1
		tags := ctx.Param("tags")
//line cmd/gopcomm/community_yap.gox:232:1
		abstract := ctx.Param("abstract")
//line cmd/gopcomm/community_yap.gox:233:1
		trans, _ := strconv.ParseBool(ctx.Param("trans"))
//line cmd/gopcomm/community_yap.gox:235:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:236:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:237:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:238:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:243:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:244:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:245:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:246:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:251:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: title, UId: uid, Cover: ctx.Param("cover"), Tags: tags, Abstract: abstract}, Content: content, Trans: trans}
//line cmd/gopcomm/community_yap.gox:265:1
		if trans {
//line cmd/gopcomm/community_yap.gox:266:1
			article, _ = this.community.TranslateArticle(todo, article)
		}
//line cmd/gopcomm/community_yap.gox:269:1
		id, err = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:270:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:271:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "add failed"})
		}
//line cmd/gopcomm/community_yap.gox:276:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:286:1
	this.Get("/page/user/account", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:287:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:288:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:289:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:290:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:291:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:292:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:296:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:297:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:299:1
		ctx.Yap__1("user_account", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:304:1
	this.Get("/page/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:305:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:307:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:308:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:309:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:313:1
		// todo middleware
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:314:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:315:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:316:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:317:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:318:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:322:1
		items, _ := this.community.GetArticlesByUid(todo, id)
//line cmd/gopcomm/community_yap.gox:323:1
		userClaimJson, _ := json.Marshal(&userClaim)
//line cmd/gopcomm/community_yap.gox:324:1
		itemsJson, _ := json.Marshal(&items)
//line cmd/gopcomm/community_yap.gox:325:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": strings.Replace(string(userClaimJson), `\"`, `"`, -1), "User": user, "Items": strings.Replace(string(itemsJson), `\"`, `"`, -1), "UserId": user.Id})
	})
//line cmd/gopcomm/community_yap.gox:341:1
	this.Get("/api/medias", func(ctx *yap.Context) {
	})
//line cmd/gopcomm/community_yap.gox:358:1
	this.Delete("/api/medias/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:359:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:360:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:361:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:362:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:363:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:364:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:369:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:370:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:371:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:376:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:384:1
	this.Get("/api/translations/:id", func(ctx *yap.Context) {
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
//line cmd/gopcomm/community_yap.gox:399:1
	this.Post("/api/translations", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:416:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:417:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, "auto", language.English)
//line cmd/gopcomm/community_yap.gox:418:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:419:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:424:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:431:1
	this.Get("/api/medias/url/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:432:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:433:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:434:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:435:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:436:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:441:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:447:1
	this.Get("/api/video-with-caption/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:449:1
		time.Sleep(5 * time.Second)
//line cmd/gopcomm/community_yap.gox:450:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:451:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:452:1
		m := make(map[string]string, 2)
//line cmd/gopcomm/community_yap.gox:453:1
		format, err := this.community.GetMediaType(todo, id)
//line cmd/gopcomm/community_yap.gox:454:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:455:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:460:1
		if format == "video/mp4" {
//line cmd/gopcomm/community_yap.gox:461:1
			subtitle, err := this.community.GetVideoSubtitle(todo, id)
//line cmd/gopcomm/community_yap.gox:462:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:463:1
				if err != nil {
//line cmd/gopcomm/community_yap.gox:464:1
					ctx.Json__1(map[string]interface {
					}{"code": 500, "err": err.Error()})
				}
//line cmd/gopcomm/community_yap.gox:469:1
				return
			}
//line cmd/gopcomm/community_yap.gox:471:1
			m["subtitle"] = domain + subtitle
		}
//line cmd/gopcomm/community_yap.gox:473:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:474:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:475:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:480:1
		m["fileKey"] = htmlUrl
//line cmd/gopcomm/community_yap.gox:481:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": m})
	})
//line cmd/gopcomm/community_yap.gox:487:1
	this.Post("/api/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:488:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:494:1
	this.Get("/api/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:499:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:500:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:501:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:502:1
			return
		}
//line cmd/gopcomm/community_yap.gox:505:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:506:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:507:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:510:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:512:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:513:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:516:1
	this.Get("/api/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:517:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:518:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:519:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:522:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:523:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:524:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:525:1
			return
		}
//line cmd/gopcomm/community_yap.gox:528:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:529:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:530:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:533:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:536:1
	this.Get("/api/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:537:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:538:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:539:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:541:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:542:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:543:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:544:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:545:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:548:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:551:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:552:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:553:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:556:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:557:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:560:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:563:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:564:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:565:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:566:1
				if
//line cmd/gopcomm/community_yap.gox:566:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:567:1
					xLog.Error(err)
//line cmd/gopcomm/community_yap.gox:568:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:572:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
