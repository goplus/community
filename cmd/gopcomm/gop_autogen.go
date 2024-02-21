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
	this.Get("/signin", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:51:1
		ctx.Yap__1("signin", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:55:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:57:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:58:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:59:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:60:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:61:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:62:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:63:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:65:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:69:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, "")
//line cmd/gopcomm/community_yap.gox:70:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:71:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:82:1
	this.Get("/page/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:85:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:86:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:87:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:88:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:89:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:90:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:94:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:95:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:96:1
		ctx.Yap__1("article", map[string]interface {
		}{"User": user, "ID": id, "Title": article.Title, "Content": article.Content, "Tags": article.Tags, "Cover": article.Cover, "Mtime": article.Mtime.Format(layoutUS), "Author": article.User})
	})
//line cmd/gopcomm/community_yap.gox:110:1
	this.Get("/page/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:111:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:112:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:113:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:114:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:115:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:116:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:119:1
		ctx.Yap__1("edit", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:124:1
	this.Get("/page/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:125:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:126:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:127:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:128:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:129:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:130:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:134:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:135:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:136:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:137:1
			if
//line cmd/gopcomm/community_yap.gox:137:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:138:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:139:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:141:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:142:1
			articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:143:1
			ctx.Yap__1("edit", map[string]interface {
			}{"User": user, "Article": strings.Replace(string(articleJson), `\"`, `"`, -1)})
		}
	})
//line cmd/gopcomm/community_yap.gox:153:1
	this.Get("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:154:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:155:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:156:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:162:1
	this.Delete("/api/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:163:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:164:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:165:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:166:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:167:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:168:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:173:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:174:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:175:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:180:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:187:1
	this.Get("/api/articles", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:188:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:189:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:190:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:192:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:193:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:194:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:197:1
		articles, next, _ := this.community.ListArticle(todo, from, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:198:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:206:1
	this.Get("/api/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:207:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:210:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:211:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:212:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:213:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:214:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:215:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:216:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:218:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:222:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue)
//line cmd/gopcomm/community_yap.gox:223:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:224:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Value": searchValue, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:233:1
	this.Post("/api/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:236:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:237:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:238:1
		transData := ctx.Param("trans_data")
//line cmd/gopcomm/community_yap.gox:240:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:241:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:242:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:243:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:248:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:249:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:250:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:251:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:257:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags"), Abstract: ctx.Param("abstract")}, Content: mdData, TransContent: transData}
//line cmd/gopcomm/community_yap.gox:269:1
		id, _ = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:270:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:280:1
	this.Get("/page/user/account", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:281:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:282:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:283:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:284:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:285:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:286:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:290:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:291:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:293:1
		ctx.Yap__1("user_account", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:298:1
	this.Get("/page/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:299:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:301:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:302:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:303:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:307:1
		// todo middleware
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:308:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:309:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:310:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:311:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:312:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:316:1
		items, _ := this.community.GetArticlesByUid(todo, id)
//line cmd/gopcomm/community_yap.gox:317:1
		userClaimJson, _ := json.Marshal(&userClaim)
//line cmd/gopcomm/community_yap.gox:318:1
		itemsJson, _ := json.Marshal(&items)
//line cmd/gopcomm/community_yap.gox:319:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": strings.Replace(string(userClaimJson), `\"`, `"`, -1), "User": user, "Items": strings.Replace(string(itemsJson), `\"`, `"`, -1), "UserId": user.Id})
	})
//line cmd/gopcomm/community_yap.gox:335:1
	this.Get("/api/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:336:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:337:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:338:1
		files, err := this.community.ListMediaByUserId(todo, uid, format)
//line cmd/gopcomm/community_yap.gox:339:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:340:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get media failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:345:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:352:1
	this.Delete("/api/medias/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:353:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:354:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:355:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:356:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:357:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:358:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:363:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:364:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:365:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:370:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:378:1
	this.Get("/api/translations/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:379:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:380:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:381:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:382:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:387:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:393:1
	this.Post("/api/translations", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:410:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:414:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:415:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:416:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:421:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:433:1
	this.Get("/api/medias/url/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:434:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:435:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:436:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:437:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:438:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:443:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:449:1
	this.Get("/api/video-with-caption/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:451:1
		time.Sleep(5 * time.Second)
//line cmd/gopcomm/community_yap.gox:452:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:453:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:454:1
		m := make(map[string]string, 2)
//line cmd/gopcomm/community_yap.gox:455:1
		format, err := this.community.GetMediaType(todo, id)
//line cmd/gopcomm/community_yap.gox:456:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:457:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:462:1
		if format == "video/mp4" {
//line cmd/gopcomm/community_yap.gox:463:1
			subtitle, err := this.community.GetVideoSubtitle(todo, id)
//line cmd/gopcomm/community_yap.gox:464:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:465:1
				if err != nil {
//line cmd/gopcomm/community_yap.gox:466:1
					ctx.Json__1(map[string]interface {
					}{"code": 500, "err": err.Error()})
				}
//line cmd/gopcomm/community_yap.gox:471:1
				return
			}
//line cmd/gopcomm/community_yap.gox:473:1
			m["subtitle"] = domain + subtitle
		}
//line cmd/gopcomm/community_yap.gox:475:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:476:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:477:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:482:1
		m["fileKey"] = htmlUrl
//line cmd/gopcomm/community_yap.gox:483:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": m})
	})
//line cmd/gopcomm/community_yap.gox:489:1
	this.Post("/api/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:490:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:496:1
	this.Get("/api/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:501:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:502:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:503:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:504:1
			return
		}
//line cmd/gopcomm/community_yap.gox:507:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:508:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:509:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:512:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:514:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:515:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:518:1
	this.Get("/api/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:519:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:520:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:521:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:524:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:525:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:526:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:527:1
			return
		}
//line cmd/gopcomm/community_yap.gox:530:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:531:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:532:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:535:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:538:1
	this.Get("/api/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:539:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:540:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:541:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:543:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:544:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:545:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:546:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:547:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:550:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:553:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:554:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:555:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:558:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:559:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:562:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:565:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:566:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:567:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:568:1
				if
//line cmd/gopcomm/community_yap.gox:568:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:569:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:573:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
