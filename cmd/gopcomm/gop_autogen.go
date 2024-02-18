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
		}{"Id": id, "CurrentUser": strings.Replace(string(userClaimJson), `\"`, `"`, -1), "User": user, "Items": strings.Replace(string(itemsJson), `\"`, `"`, -1)})
	})
//line cmd/gopcomm/community_yap.gox:123:1
	this.Get("/userEdit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:124:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:125:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:126:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:127:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:128:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:129:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:133:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:134:1
			http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
		}
//line cmd/gopcomm/community_yap.gox:136:1
		ctx.Yap__1("user_edit", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:141:1
	this.Get("/add", func(ctx *yap.Context) {
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
//line cmd/gopcomm/community_yap.gox:150:1
		ctx.Yap__1("edit", map[string]*core.User{"User": user})
	})
//line cmd/gopcomm/community_yap.gox:155:1
	this.Get("/delete", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:156:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:157:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:158:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:159:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:160:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:161:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:166:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:167:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:168:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:173:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:180:1
	this.Get("/medias", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:181:1
		format := ctx.Param("format")
//line cmd/gopcomm/community_yap.gox:182:1
		uid := ctx.Param("uid")
//line cmd/gopcomm/community_yap.gox:183:1
		files, err := this.community.ListMediaByUserId(todo, uid, format)
//line cmd/gopcomm/community_yap.gox:184:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:185:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get media failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:190:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:197:1
	this.Get("/delMedia", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:198:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:199:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:200:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:201:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:202:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:203:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:208:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:209:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:210:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:215:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:222:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:224:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:225:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:226:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:227:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:228:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:229:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:230:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:232:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:236:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, "")
//line cmd/gopcomm/community_yap.gox:237:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:238:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:246:1
	this.Get("/get", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:247:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:248:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:249:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:251:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:252:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:253:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:256:1
		articles, next, _ := this.community.ListArticle(todo, from, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:257:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:265:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:266:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:267:1
		if searchValue == "" {
//line cmd/gopcomm/community_yap.gox:268:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "err": "value can not be ''."})
		}
//line cmd/gopcomm/community_yap.gox:275:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:276:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:277:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:278:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:279:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:280:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:281:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:283:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:287:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue)
//line cmd/gopcomm/community_yap.gox:288:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:289:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Value": searchValue, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:298:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:299:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:300:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:301:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:302:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:303:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:304:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:308:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:309:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:310:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:311:1
			if
//line cmd/gopcomm/community_yap.gox:311:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:312:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:313:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:315:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:316:1
			articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:317:1
			ctx.Yap__1("edit", map[string]interface {
			}{"User": user, "Article": strings.Replace(string(articleJson), `\"`, `"`, -1)})
		}
	})
//line cmd/gopcomm/community_yap.gox:324:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:325:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:326:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:327:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:328:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:333:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:340:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:343:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:344:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:345:1
		transData := ctx.Param("trans_data")
//line cmd/gopcomm/community_yap.gox:347:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:348:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:349:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:350:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:355:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:356:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:357:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:358:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:364:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags"), Abstract: ctx.Param("abstract")}, Content: mdData, TransContent: transData}
//line cmd/gopcomm/community_yap.gox:376:1
		id, _ = this.community.PutArticle(todo, uid, article)
//line cmd/gopcomm/community_yap.gox:377:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:385:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:402:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:406:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:407:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:408:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:413:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:425:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:426:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:428:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:430:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:433:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
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
	this.Get("/getVideoAndSubtitle/:id", func(ctx *yap.Context) {
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
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:488:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:491:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:496:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:497:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:498:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:499:1
			return
		}
//line cmd/gopcomm/community_yap.gox:502:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:503:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:504:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:507:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:509:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:510:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:513:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:514:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:515:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:516:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:519:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:520:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:521:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:522:1
			return
		}
//line cmd/gopcomm/community_yap.gox:525:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:526:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:527:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:530:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:533:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:534:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:535:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:536:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:538:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:539:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:540:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:541:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:542:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:545:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:548:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:549:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:550:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:553:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:554:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:557:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:560:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:562:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:563:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:564:1
				if
//line cmd/gopcomm/community_yap.gox:564:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:565:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:569:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
