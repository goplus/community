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
//line cmd/gopcomm/community_yap.gox:124:1
	this.Get("/userEdit", func(ctx *yap.Context) {
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
		if err != nil {
//line cmd/gopcomm/community_yap.gox:135:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:137:1
		ctx.Yap__1("user_edit", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:142:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:143:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:144:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:145:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:146:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:147:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:148:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:151:1
		ctx.Yap__1("edit", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:156:1
	this.Get("/delete", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:157:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:158:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:159:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:160:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:161:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:162:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:167:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:168:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:169:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:174:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:181:1
	this.Get("/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:182:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:183:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:184:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:185:1
		fromInt, err := strconv.Atoi(from)
//line cmd/gopcomm/community_yap.gox:186:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:187:1
			fromInt = 0
		}
//line cmd/gopcomm/community_yap.gox:189:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:190:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:191:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:192:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:194:1
		files, next, err := this.community.ListMediaByUserId(todo, uid, format, fromInt, limitInt)
//line cmd/gopcomm/community_yap.gox:195:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:196:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "next": next, "err": err.Error()})
		} else {
//line cmd/gopcomm/community_yap.gox:202:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "next": next, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:210:1
	this.Get("/delMedia", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:211:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:212:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:213:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:214:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:215:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:216:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:221:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:222:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:223:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:228:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:235:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:237:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:238:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:239:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:240:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:241:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:242:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:243:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:245:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:249:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, "")
//line cmd/gopcomm/community_yap.gox:250:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:251:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:259:1
	this.Get("/get", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:260:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:261:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:262:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:264:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:265:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:266:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:269:1
		articles, next, _ := this.community.ListArticle(todo, from, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:270:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:278:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:279:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:288:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:289:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:290:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:291:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:292:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:293:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:294:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:296:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:300:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue)
//line cmd/gopcomm/community_yap.gox:301:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:302:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Value": searchValue, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:311:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:312:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:313:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:314:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:315:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:316:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:317:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:321:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:322:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:323:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:324:1
			if
//line cmd/gopcomm/community_yap.gox:324:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:325:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:326:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:328:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:329:1
			articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:330:1
			ctx.Yap__1("edit", map[string]interface {
			}{"User": user, "Article": strings.Replace(string(articleJson), `\"`, `"`, -1)})
		}
	})
//line cmd/gopcomm/community_yap.gox:337:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:338:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:339:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:340:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:341:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:346:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:353:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:356:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:357:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:358:1
		transData := ctx.Param("trans_data")
//line cmd/gopcomm/community_yap.gox:360:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:361:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:362:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:363:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:368:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:369:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:370:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:371:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:377:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags"), Abstract: ctx.Param("abstract")}, Content: mdData, TransContent: transData}
//line cmd/gopcomm/community_yap.gox:389:1
		id, _ = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:390:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:398:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:415:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:419:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:420:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:421:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:426:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:438:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:439:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:441:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:443:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:446:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:447:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:448:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:449:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:450:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:451:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:456:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:462:1
	this.Get("/getVideoAndSubtitle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:463:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:464:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:465:1
		m := make(map[string]string, 2)
//line cmd/gopcomm/community_yap.gox:466:1
		format, err := this.community.GetMediaType(todo, id)
//line cmd/gopcomm/community_yap.gox:467:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:468:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:473:1
		if format == "video/mp4" {
//line cmd/gopcomm/community_yap.gox:474:1
			subtitle, err := this.community.GetVideoSubtitle(todo, id)
//line cmd/gopcomm/community_yap.gox:475:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:476:1
				if err != nil {
//line cmd/gopcomm/community_yap.gox:477:1
					ctx.Json__1(map[string]interface {
					}{"code": 500, "err": err.Error()})
				}
//line cmd/gopcomm/community_yap.gox:482:1
				return
			}
//line cmd/gopcomm/community_yap.gox:484:1
			m["subtitle"] = domain + subtitle
		}
//line cmd/gopcomm/community_yap.gox:486:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:487:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:488:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:493:1
		m["fileKey"] = htmlUrl
//line cmd/gopcomm/community_yap.gox:494:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": m})
	})
//line cmd/gopcomm/community_yap.gox:500:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:501:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:504:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:509:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:510:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:511:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:512:1
			return
		}
//line cmd/gopcomm/community_yap.gox:515:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:516:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:517:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:520:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:522:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:523:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:526:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:527:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:528:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:529:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:532:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:533:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:534:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:535:1
			return
		}
//line cmd/gopcomm/community_yap.gox:538:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:539:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:540:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:543:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:546:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:547:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:548:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:549:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:551:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:552:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:553:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:554:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:555:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:558:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:561:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:562:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:563:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:566:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:567:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:570:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:573:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:575:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:576:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:577:1
				if
//line cmd/gopcomm/community_yap.gox:577:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:578:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:582:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community_yap))
}