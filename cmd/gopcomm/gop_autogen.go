package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"github.com/goplus/yap"
	_ "github.com/joho/godotenv/autoload"
	"github.com/qiniu/x/xlog"
	"golang.org/x/text/language"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const _ = true
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
		}{"User": user, "ID": id, "Title": article.Title, "Content": article.Content, "Tags": article.Tags, "Cover": article.Cover, "Mtime": article.Mtime.Format(layoutUS), "Author": article.User})
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
		}{"Id": id, "CurrentUser": strings.Replace(string(userClaimJson), `\"`, `"`, -1), "User": user, "Items": strings.Replace(string(itemsJson), `\"`, `"`, -1), "UserId": user.Id})
	})
//line cmd/gopcomm/community_yap.gox:123:1
	this.Get("/userUnlink", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:124:1
		pv := ctx.Param("pv")
//line cmd/gopcomm/community_yap.gox:125:1
		switch pv {
//line cmd/gopcomm/community_yap.gox:126:1
		case "Twitter":
//line cmd/gopcomm/community_yap.gox:127:1
		case "Facebook":
//line cmd/gopcomm/community_yap.gox:128:1
		case "Github":
//line cmd/gopcomm/community_yap.gox:129:1
		case "WeChat":
//line cmd/gopcomm/community_yap.gox:130:1
		default:
//line cmd/gopcomm/community_yap.gox:131:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/userEdit", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:133:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:134:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:135:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/userEdit", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:138:1
		this.community.UnLink(token.Value, pv)
//line cmd/gopcomm/community_yap.gox:139:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, "/userEdit", http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:141:1
	this.Get("/userEdit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:142:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:143:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:144:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:145:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:146:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:147:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:151:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:152:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:154:1
		appInfo, err := this.community.GetApplicationInfo()
//line cmd/gopcomm/community_yap.gox:155:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:156:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:158:1
		appInfoStr, _ := json.Marshal(*appInfo)
//line cmd/gopcomm/community_yap.gox:159:1
		ctx.Yap__1("user_edit", map[string]interface {
		}{"User": user, "Application": string(appInfoStr), "Binds": this.community.GetAccountBinds(token.Value)})
	})
//line cmd/gopcomm/community_yap.gox:166:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:167:1
		var user *core.User
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
			}
		}
//line cmd/gopcomm/community_yap.gox:175:1
		ctx.Yap__1("edit", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:180:1
	this.Get("/delete", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:181:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:182:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:183:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:184:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:185:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:186:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:191:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:192:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:193:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:198:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:205:1
	this.Get("/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:206:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:207:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:208:1
		files, err := this.community.ListMediaByUserId(todo, uid, format)
//line cmd/gopcomm/community_yap.gox:209:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:210:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get media failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:215:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:222:1
	this.Get("/delMedia", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:223:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:224:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:225:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:226:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:227:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:228:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:233:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:234:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:235:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:240:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:247:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:249:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:250:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:251:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:252:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:253:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:254:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:255:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:257:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:261:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, "")
//line cmd/gopcomm/community_yap.gox:262:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:263:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:271:1
	this.Get("/get", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:272:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:273:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:274:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:276:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:277:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:278:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:281:1
		articles, next, _ := this.community.ListArticle(todo, from, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:282:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:290:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:291:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:300:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:301:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:302:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:303:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:304:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:305:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:306:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:308:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:312:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue)
//line cmd/gopcomm/community_yap.gox:313:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:314:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Value": searchValue, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:323:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:324:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:325:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:326:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:327:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:328:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:329:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:333:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:334:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:335:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:336:1
			if
//line cmd/gopcomm/community_yap.gox:336:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:337:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:338:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:340:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:341:1
			articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:342:1
			ctx.Yap__1("edit", map[string]interface {
			}{"User": user, "Article": strings.Replace(string(articleJson), `\"`, `"`, -1)})
		}
	})
//line cmd/gopcomm/community_yap.gox:349:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:350:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:351:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:352:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:353:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:358:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:365:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:368:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:369:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:370:1
		transData := ctx.Param("trans_data")
//line cmd/gopcomm/community_yap.gox:372:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:373:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:374:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:375:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:380:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:381:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:382:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:383:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:389:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags"), Abstract: ctx.Param("abstract")}, Content: mdData, TransContent: transData}
//line cmd/gopcomm/community_yap.gox:401:1
		id, _ = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:402:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:410:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:427:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:431:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:432:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:433:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:438:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:450:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:451:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:453:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:455:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:458:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:459:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:460:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:461:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:462:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:463:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:468:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:474:1
	this.Get("/getVideoAndSubtitle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:475:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:476:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:477:1
		m := make(map[string]string, 2)
//line cmd/gopcomm/community_yap.gox:478:1
		format, err := this.community.GetMediaType(todo, id)
//line cmd/gopcomm/community_yap.gox:479:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:480:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:485:1
		if format == "video/mp4" {
//line cmd/gopcomm/community_yap.gox:486:1
			subtitle, err := this.community.GetVideoSubtitle(todo, id)
//line cmd/gopcomm/community_yap.gox:487:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:488:1
				if err != nil {
//line cmd/gopcomm/community_yap.gox:489:1
					ctx.Json__1(map[string]interface {
					}{"code": 500, "err": err.Error()})
				}
//line cmd/gopcomm/community_yap.gox:494:1
				return
			}
//line cmd/gopcomm/community_yap.gox:496:1
			m["subtitle"] = domain + subtitle
		}
//line cmd/gopcomm/community_yap.gox:498:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:499:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:500:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:505:1
		m["fileKey"] = htmlUrl
//line cmd/gopcomm/community_yap.gox:506:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": m})
	})
//line cmd/gopcomm/community_yap.gox:512:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:513:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:516:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:521:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:522:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:523:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:524:1
			return
		}
//line cmd/gopcomm/community_yap.gox:527:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:528:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:529:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:532:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:534:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:535:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:538:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:539:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:540:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:541:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:544:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:545:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:546:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:547:1
			return
		}
//line cmd/gopcomm/community_yap.gox:550:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:551:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:552:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:555:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:558:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:559:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:560:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:561:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:563:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:564:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:565:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:566:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:567:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:570:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:573:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:574:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:575:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:578:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:579:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:582:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:585:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:587:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:588:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:589:1
				if
//line cmd/gopcomm/community_yap.gox:589:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:590:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:594:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
//line cmd/gopcomm/community_yap.gox:585:1
	yap.Gopt_App_Main(new(community))
}
