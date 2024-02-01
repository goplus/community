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
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"golang.org/x/text/language"
	_ "github.com/joho/godotenv/autoload"
	"github.com/qiniu/x/xlog"
)

const layoutUS = "January 2, 2006"

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
	this.Get("/demo", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:52:1
		ctx.Yap__1("demo", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:55:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:58:1
		// todo middleware
		// Get User Info
		var user *core.User
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
			}
		}
//line cmd/gopcomm/community_yap.gox:67:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:68:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:69:1
		ctx.Yap__1("article", map[string]interface {
		}{"User": user, "ID": id, "Title": article.Title, "Content": article.HtmlUrl, "Tags": article.Tags, "Cover": article.Cover, "Mtime": article.Mtime.Format(layoutUS), "Author": article.User})
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
//line cmd/gopcomm/community_yap.gox:100:1
		// todo middleware
		// get user by token
		var user *core.User
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
			}
		}
//line cmd/gopcomm/community_yap.gox:109:1
		items, _ := this.community.GetArticlesByUid(todo, id)
//line cmd/gopcomm/community_yap.gox:110:1
		userClaimJson, _ := json.Marshal(&userClaim)
//line cmd/gopcomm/community_yap.gox:111:1
		itemsJson, _ := json.Marshal(&items)
//line cmd/gopcomm/community_yap.gox:112:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": strings.Replace(string(userClaimJson), `\"`, `"`, -1), "User": user, "Items": strings.Replace(string(itemsJson), `\"`, `"`, -1)})
	})
//line cmd/gopcomm/community_yap.gox:120:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:121:1
		ctx.Yap__1("edit", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:124:1
	this.Get("/delete", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:125:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:126:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:127:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:128:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:129:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:130:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:135:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:136:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:137:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:142:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:149:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:151:1
		from := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:152:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:155:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:156:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:157:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:158:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:159:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:160:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:164:1
		page, err := strconv.Atoi(from)
//line cmd/gopcomm/community_yap.gox:165:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:166:1
			page = 1
		}
//line cmd/gopcomm/community_yap.gox:169:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:170:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:171:1
			limitInt = 20
		}
//line cmd/gopcomm/community_yap.gox:175:1
		articles, total, _ := this.community.Articles(todo, page, limitInt, "")
//line cmd/gopcomm/community_yap.gox:176:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": user, "Items": articles, "Page": page, "TotalPage": (total + limitInt - 1) / limitInt, "Total": total})
	})
//line cmd/gopcomm/community_yap.gox:185:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:186:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:187:1
		if searchValue == "" {
//line cmd/gopcomm/community_yap.gox:188:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "err": "value can not be ''."})
		}
//line cmd/gopcomm/community_yap.gox:194:1
		from := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:195:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:196:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:197:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:198:1
			limitInt = 10
		}
//line cmd/gopcomm/community_yap.gox:200:1
		page, err := strconv.Atoi(from)
//line cmd/gopcomm/community_yap.gox:201:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:202:1
			page = 1
		}
//line cmd/gopcomm/community_yap.gox:206:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:207:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:208:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:209:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:210:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:211:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:215:1
		articles, total, _ := this.community.Articles(todo, page, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:216:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": user, "Items": articles, "Value": searchValue, "Page": page, "TotalPage": (total + limitInt - 1) / limitInt, "Total": total})
	})
//line cmd/gopcomm/community_yap.gox:226:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:227:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:228:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:229:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:235:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:236:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:237:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:243:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:244:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:245:1
			if
//line cmd/gopcomm/community_yap.gox:245:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:246:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:247:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:249:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:250:1
			ctx.Yap__1("edit", article)
		}
	})
//line cmd/gopcomm/community_yap.gox:254:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:255:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:256:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:257:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:258:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:263:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:270:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:272:1
		trans := ctx.Param("trans")
//line cmd/gopcomm/community_yap.gox:273:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:274:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:275:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:277:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:278:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:279:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:284:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:285:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:286:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:293:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags"), Abstract: ctx.Param("abstract")}, Content: mdData, HtmlData: htmlData}
//line cmd/gopcomm/community_yap.gox:305:1
		id, _ = this.community.PutArticle(todo, uid, trans, article)
//line cmd/gopcomm/community_yap.gox:306:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:314:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:316:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:317:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:318:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:323:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:324:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:325:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:331:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:332:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:333:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:335:1
		transData, err := this.trans.TranslateMarkdownText(mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:336:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:337:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:342:1
		id, _ = this.community.SaveHtml(todo, uid, htmlData, mdData, id)
//line cmd/gopcomm/community_yap.gox:343:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "id": id, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:350:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:351:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:353:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:355:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:358:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:359:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:360:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:361:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:362:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:363:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:368:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:374:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:375:1
		core.UploadFile(ctx, this.community)
	})
//line cmd/gopcomm/community_yap.gox:378:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:383:1
		redirectURL := fmt.Sprintf("%s/%s", ctx.Request.Referer(), "callback")
//line cmd/gopcomm/community_yap.gox:385:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:386:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:390:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:391:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:392:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:393:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:397:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("/"), http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:400:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:401:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:402:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:403:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:408:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("/"), http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:411:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:412:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:413:1
	this.trans = translation.New(os.Getenv("NIUTRANS_API_KEY"), "", "")
//line cmd/gopcomm/community_yap.gox:414:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:417:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:418:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:421:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:424:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:426:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:427:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:428:1
				if
//line cmd/gopcomm/community_yap.gox:428:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:429:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:433:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
