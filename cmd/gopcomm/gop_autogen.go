package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"github.com/goplus/yap"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	logger, _ := zap.NewProduction()
//line cmd/gopcomm/community_yap.gox:34:1
	defer logger.Sync()
//line cmd/gopcomm/community_yap.gox:35:1
	zlog := logger.Sugar()
//line cmd/gopcomm/community_yap.gox:37:1
	this.Static__0("/")
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
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:54:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:55:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:56:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:57:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:58:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:59:1
				zlog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:63:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:64:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:65:1
		ctx.Yap__1("article", map[string]interface {
		}{"User": user, "ID": id, "Title": article.Title, "Content": article.HtmlUrl, "Tags": article.Tags, "Cover": article.Cover, "Mtime": article.Mtime.Format(layoutUS), "Author": article.User})
	})
//line cmd/gopcomm/community_yap.gox:78:1
	this.Get("/getArticle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:79:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:80:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:81:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:87:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:88:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:90:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:91:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:92:1
			zlog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:96:1
		// todo middleware
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:97:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:98:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:99:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:100:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:101:1
				zlog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:105:1
		items, _ := this.community.GetArticlesByUid(todo, id)
//line cmd/gopcomm/community_yap.gox:106:1
		userClaimJson, _ := json.Marshal(&userClaim)
//line cmd/gopcomm/community_yap.gox:107:1
		itemsJson, _ := json.Marshal(&items)
//line cmd/gopcomm/community_yap.gox:108:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": strings.Replace(string(userClaimJson), `\"`, `"`, -1), "User": user, "Items": strings.Replace(string(itemsJson), `\"`, `"`, -1)})
	})
//line cmd/gopcomm/community_yap.gox:116:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:117:1
		ctx.Yap__1("edit", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:120:1
	this.Get("/delete", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:121:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:122:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:123:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:124:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:125:1
			zlog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:126:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:131:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:132:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:133:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:138:1
			ctx.Json__1(map[string]interface {
			}{"code": 200, "msg": "delete success"})
		}
	})
//line cmd/gopcomm/community_yap.gox:145:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:147:1
		from := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:148:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:151:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:152:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:153:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:154:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:155:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:156:1
				zlog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:160:1
		page, err := strconv.Atoi(from)
//line cmd/gopcomm/community_yap.gox:161:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:162:1
			page = 1
		}
//line cmd/gopcomm/community_yap.gox:165:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:166:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:167:1
			limitInt = 20
		}
//line cmd/gopcomm/community_yap.gox:171:1
		articles, total, _ := this.community.Articles(todo, page, limitInt, "")
//line cmd/gopcomm/community_yap.gox:172:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": user, "Items": articles, "Page": page, "TotalPage": (total + limitInt - 1) / limitInt, "Total": total})
	})
//line cmd/gopcomm/community_yap.gox:181:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:182:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:183:1
		if searchValue == "" {
//line cmd/gopcomm/community_yap.gox:184:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "err": "value can not be ''."})
		}
//line cmd/gopcomm/community_yap.gox:190:1
		from := ctx.Param("page")
//line cmd/gopcomm/community_yap.gox:191:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:192:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:193:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:194:1
			limitInt = 10
		}
//line cmd/gopcomm/community_yap.gox:196:1
		page, err := strconv.Atoi(from)
//line cmd/gopcomm/community_yap.gox:197:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:198:1
			page = 1
		}
//line cmd/gopcomm/community_yap.gox:202:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:203:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:204:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:205:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:206:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:207:1
				zlog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:211:1
		articles, total, _ := this.community.Articles(todo, page, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:212:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": user, "Items": articles, "Value": searchValue, "Page": page, "TotalPage": (total + limitInt - 1) / limitInt, "Total": total})
	})
//line cmd/gopcomm/community_yap.gox:222:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:223:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:224:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:225:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:231:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:232:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:233:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:239:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:240:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:241:1
			if
//line cmd/gopcomm/community_yap.gox:241:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:242:1
				zlog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:243:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:245:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:246:1
			ctx.Yap__1("edit", article)
		}
	})
//line cmd/gopcomm/community_yap.gox:250:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:251:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:252:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:253:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:254:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:259:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:266:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:268:1
		trans := ctx.Param("trans")
//line cmd/gopcomm/community_yap.gox:269:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:270:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:271:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:273:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:274:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:275:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:280:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:281:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:282:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:289:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags"), Abstract: ctx.Param("abstract")}, Content: mdData, HtmlData: htmlData}
//line cmd/gopcomm/community_yap.gox:301:1
		id, _ = this.community.PutArticle(todo, uid, trans, article)
//line cmd/gopcomm/community_yap.gox:302:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:310:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:312:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:313:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:314:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:319:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:320:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:321:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:327:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:328:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:329:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:331:1
		transData, err := this.trans.TranslateMarkdownText(mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:332:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:333:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:338:1
		id, _ = this.community.SaveHtml(todo, uid, htmlData, mdData, id)
//line cmd/gopcomm/community_yap.gox:339:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "id": id, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:346:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:347:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:349:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:351:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:354:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:355:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:356:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:357:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:358:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:359:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:364:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:370:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:371:1
		file, header, err := ctx.FormFile("file")
//line cmd/gopcomm/community_yap.gox:372:1
		filename := header.Filename
//line cmd/gopcomm/community_yap.gox:374:1
		ctx.ParseMultipartForm(10 << 20)
//line cmd/gopcomm/community_yap.gox:376:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:377:1
			zlog.Error("upload file error:", filename)
//line cmd/gopcomm/community_yap.gox:378:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:379:1
			return
		}
//line cmd/gopcomm/community_yap.gox:383:1
		dst, err := os.Create(filename)
//line cmd/gopcomm/community_yap.gox:384:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:385:1
			zlog.Error("create file error:", file)
//line cmd/gopcomm/community_yap.gox:386:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:387:1
			return
		}
//line cmd/gopcomm/community_yap.gox:389:1
		defer func() {
//line cmd/gopcomm/community_yap.gox:390:1
			file.Close()
//line cmd/gopcomm/community_yap.gox:391:1
			dst.Close()
//line cmd/gopcomm/community_yap.gox:392:1
			err = os.Remove(filename)
//line cmd/gopcomm/community_yap.gox:393:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:394:1
				zlog.Error("delete file error:", filename)
//line cmd/gopcomm/community_yap.gox:395:1
				return
			}
		}()
//line cmd/gopcomm/community_yap.gox:400:1
		_, err = io.Copy(dst, file)
//line cmd/gopcomm/community_yap.gox:401:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:402:1
			zlog.Error("copy file errer:", filename)
//line cmd/gopcomm/community_yap.gox:403:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:404:1
			return
		}
//line cmd/gopcomm/community_yap.gox:406:1
		bytes, err := os.ReadFile(filename)
//line cmd/gopcomm/community_yap.gox:407:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:408:1
			zlog.Error("read file errer:", filename)
//line cmd/gopcomm/community_yap.gox:409:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:410:1
			return
		}
//line cmd/gopcomm/community_yap.gox:412:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:413:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:414:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:419:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:420:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:421:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:426:1
		id, err := this.community.SaveMedia(context.Background(), uid, bytes)
//line cmd/gopcomm/community_yap.gox:427:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:428:1
			zlog.Error("save file", err.Error())
//line cmd/gopcomm/community_yap.gox:429:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:430:1
			return
		}
//line cmd/gopcomm/community_yap.gox:434:1
		ctx.JSON(200, id)
	})
//line cmd/gopcomm/community_yap.gox:437:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:439:1
		redirectURL := ctx.URL.Query().Get("redirect_url")
//line cmd/gopcomm/community_yap.gox:440:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:441:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:445:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:446:1
		code := ctx.URL.Query().Get("code")
//line cmd/gopcomm/community_yap.gox:447:1
		state := ctx.URL.Query().Get("state")
//line cmd/gopcomm/community_yap.gox:449:1
		token, error := this.community.GetAccessToken(code, state)
//line cmd/gopcomm/community_yap.gox:450:1
		if error != nil {
//line cmd/gopcomm/community_yap.gox:451:1
			zlog.Error("err", error)
		}
//line cmd/gopcomm/community_yap.gox:454:1
		cookie := http.Cookie{Name: "token", Value: token.AccessToken, Path: "/", MaxAge: 3600}
//line cmd/gopcomm/community_yap.gox:460:1
		http.SetCookie(ctx.ResponseWriter, &cookie)
//line cmd/gopcomm/community_yap.gox:464:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("%s?token=%s", endpoint, token.AccessToken), http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:467:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:468:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:469:1
	this.trans = translation.New(os.Getenv("NIUTRANS_API_KEY"), "", "")
//line cmd/gopcomm/community_yap.gox:470:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:477:1
	zlog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:480:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:482:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:483:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:484:1
				if
//line cmd/gopcomm/community_yap.gox:484:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:485:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:489:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
