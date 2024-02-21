package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"github.com/goplus/yap"
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
		articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:71:1
		ctx.Yap__1("article", map[string]interface {
		}{"User": user, "Article": strings.Replace(string(articleJson), `\"`, `"`, -1), "Content": article.Content})
	})

	this.Get("/tranArticle", func(ctx *yap.Context) {
		id := ctx.Param("id")
		article, err := this.community.GetTranslateArticle(todo, id)
		if err != nil {
			ctx.Json__1(map[string]interface {
			}{"code": 0})
		}
		ctx.Json__1(map[string]interface {
		}{"code": 200,
			"content": article.Content,
			"tags":    article.Tags,
			"title":   article.Title})
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
		userId := ""
		if user != nil {
			userId = user.Id
		}
//line cmd/gopcomm/community_yap.gox:112:1
		items, _ := this.community.GetArticlesByUid(todo, id)
//line cmd/gopcomm/community_yap.gox:113:1
		userClaimJson, _ := json.Marshal(&userClaim)
//line cmd/gopcomm/community_yap.gox:114:1
		itemsJson, _ := json.Marshal(&items)
//line cmd/gopcomm/community_yap.gox:115:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": strings.Replace(string(userClaimJson), `\"`, `"`, -1), "User": user, "Items": strings.Replace(string(itemsJson), `\"`, `"`, -1), "UserId": userId})
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
		files, err := this.community.ListMediaByUserId(todo, uid, format)
//line cmd/gopcomm/community_yap.gox:185:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:186:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "get media failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:191:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "items": files})
		}
	})
//line cmd/gopcomm/community_yap.gox:198:1
	this.Get("/delMedia", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:199:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:200:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:201:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:202:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:203:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:204:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:209:1
		err = this.community.DelMedia(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:210:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:211:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:216:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:223:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:225:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:226:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:227:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:228:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:229:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:230:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:231:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:233:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:237:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, "")
//line cmd/gopcomm/community_yap.gox:238:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:239:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:247:1
	this.Get("/get", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:248:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:249:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:250:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:252:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:253:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:254:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:257:1
		articles, next, _ := this.community.ListArticle(todo, from, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:258:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:266:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:267:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:276:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:277:1
		userId := ""
//line cmd/gopcomm/community_yap.gox:278:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:279:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:280:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:281:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:282:1
				xLog.Error("get user error:", err)
			} else {
//line cmd/gopcomm/community_yap.gox:284:1
				userId = user.Id
			}
		}
//line cmd/gopcomm/community_yap.gox:288:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue)
//line cmd/gopcomm/community_yap.gox:289:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:290:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Value": searchValue, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:299:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:300:1
		var user *core.User
//line cmd/gopcomm/community_yap.gox:301:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:302:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:303:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:304:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:305:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:309:1
		uid := user.Id
//line cmd/gopcomm/community_yap.gox:310:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:311:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:312:1
			if
//line cmd/gopcomm/community_yap.gox:312:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:313:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:314:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:316:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:317:1
			articleJson, _ := json.Marshal(&article)
//line cmd/gopcomm/community_yap.gox:318:1
			ctx.Yap__1("edit", map[string]interface {
			}{"User": user, "Article": strings.Replace(string(articleJson), `\"`, `"`, -1)})
		}
	})
//line cmd/gopcomm/community_yap.gox:325:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:326:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:327:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:328:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:329:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:334:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:341:1
	this.Post("/commit", func(ctx *yap.Context) {
		id := ctx.Param("id")
		mdData := ctx.Param("content")
		//line cmd/gopcomm/community_yap.gox:346:1
		title := ctx.Param("title")
		tags := ctx.Param("tags")
		abstract := ctx.Param("abstract")
		trans, _ := strconv.ParseBool(ctx.Param("trans"))
//line cmd/gopcomm/community_yap.gox:348:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:349:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:350:1
			xLog.Info("token", err)
//line cmd/gopcomm/community_yap.gox:351:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:356:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:357:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:358:1
			xLog.Info("uid", err)
//line cmd/gopcomm/community_yap.gox:359:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
		article := &core.Article{
			ArticleEntry: core.ArticleEntry{
				ID:       id,
				Title:    title,
				UId:      uid,
				Cover:    ctx.Param("cover"),
				Tags:     tags,
				Abstract: abstract,
			},
			Content: mdData,
			Trans:   trans,
		}

		if trans {
			article, _ = this.community.TranslateArticle(todo, article)
		}
		id, _ = this.community.PutArticle(todo, uid, article)
		//line cmd/gopcomm/community_yap.gox:378:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:386:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:403:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:407:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, "auto", language.English)
//line cmd/gopcomm/community_yap.gox:408:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:409:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:414:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:426:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:427:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:429:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:431:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:434:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:435:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:436:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:437:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:438:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:439:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:444:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:450:1
	this.Get("/getVideoAndSubtitle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:451:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:452:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:453:1
		m := make(map[string]string, 2)
//line cmd/gopcomm/community_yap.gox:454:1
		format, err := this.community.GetMediaType(todo, id)
//line cmd/gopcomm/community_yap.gox:455:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:456:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:461:1
		if format == "video/mp4" {
//line cmd/gopcomm/community_yap.gox:462:1
			subtitle, err := this.community.GetVideoSubtitle(todo, id)
//line cmd/gopcomm/community_yap.gox:463:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:464:1
				if err != nil {
//line cmd/gopcomm/community_yap.gox:465:1
					ctx.Json__1(map[string]interface {
					}{"code": 500, "err": err.Error()})
				}
//line cmd/gopcomm/community_yap.gox:470:1
				return
			}
//line cmd/gopcomm/community_yap.gox:472:1
			m["subtitle"] = domain + subtitle
		}
//line cmd/gopcomm/community_yap.gox:474:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:475:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:476:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:481:1
		m["fileKey"] = htmlUrl
//line cmd/gopcomm/community_yap.gox:482:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": m})
	})
//line cmd/gopcomm/community_yap.gox:488:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:489:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:492:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:497:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:498:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:499:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:500:1
			return
		}
//line cmd/gopcomm/community_yap.gox:503:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:504:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:505:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:508:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:510:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:511:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:514:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:515:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:516:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:517:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:520:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:521:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:522:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:523:1
			return
		}
//line cmd/gopcomm/community_yap.gox:526:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:527:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:528:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:531:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:534:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:535:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:536:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:537:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:539:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:540:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:541:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:542:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:543:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:546:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:549:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:550:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:551:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:554:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:555:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:558:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:561:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:563:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:564:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:565:1
				if
//line cmd/gopcomm/community_yap.gox:565:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:566:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:570:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
