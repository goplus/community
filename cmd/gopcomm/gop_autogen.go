package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
  "strconv"
	"os"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"github.com/goplus/yap"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

const layoutUS = "January 2, 2006"

type community struct {
	yap.App
	community *core.Community
	trans     *translation.Engine
}

//line cmd/gopcomm/community_yap.gox:23
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:28:1
	todo := context.TODO()
//line cmd/gopcomm/community_yap.gox:29:1
	endpoint := os.Getenv("GOP_COMMUNITY_ENDPOINT")
//line cmd/gopcomm/community_yap.gox:30:1
	domain := os.Getenv("GOP_COMMUNITY_DOMAIN")
//line cmd/gopcomm/community_yap.gox:31:1
	logger, _ := zap.NewProduction()
//line cmd/gopcomm/community_yap.gox:32:1
	defer logger.Sync()
//line cmd/gopcomm/community_yap.gox:33:1
	zlog := logger.Sugar()
//line cmd/gopcomm/community_yap.gox:35:1
	this.Static__0("/")
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
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:52:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:53:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:54:1
		zlog.Info(token)
//line cmd/gopcomm/community_yap.gox:55:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:56:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:57:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:58:1
				zlog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:62:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:63:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:64:1
		ctx.Yap__1("article", map[string]interface {
		}{"User": user, "ID": id, "Title": article.Title, "Content": article.HtmlUrl, "Tags": article.Tags, "Cover": article.Cover, "Mtime": article.Mtime.Format(layoutUS), "Author": article.User})
	})
//line cmd/gopcomm/community_yap.gox:77:1
	this.Get("/getArticle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:78:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:79:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:80:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:86:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:87:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:89:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:90:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:91:1
			zlog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:95:1
		// todo middleware
		// get user by token
		var user *core.User
//line cmd/gopcomm/community_yap.gox:96:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:97:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:98:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:99:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:100:1
				zlog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:104:1
		items, _ := this.community.GetArticlesByUid(todo, id)
//line cmd/gopcomm/community_yap.gox:105:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": userClaim, "User": user, "Items": items})
	})
//line cmd/gopcomm/community_yap.gox:113:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:114:1
		ctx.Yap__1("edit", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:117:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:118:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:119:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:122:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:123:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:124:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:125:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:126:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:127:1
				zlog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:131:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:132:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:133:1
			limitInt = 20
		}
//line cmd/gopcomm/community_yap.gox:136:1
		articles, next, _ := this.community.ListArticle(todo, from, limitInt)
//line cmd/gopcomm/community_yap.gox:137:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": user, "Items": articles, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:144:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:145:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:146:1
		if searchValue == "" {
//line cmd/gopcomm/community_yap.gox:147:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "err": "value can not be ''."})
		}
//line cmd/gopcomm/community_yap.gox:153:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:154:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:155:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:156:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:157:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:158:1
				zlog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:161:1
		articles, _ := this.community.SearchArticle(todo, searchValue)
//line cmd/gopcomm/community_yap.gox:162:1
		ctx.Yap__1("home", map[string]interface {
		}{"User": user, "Items": articles, "Value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:169:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:170:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:171:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:172:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:178:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:179:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:180:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:186:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:187:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:188:1
			if
//line cmd/gopcomm/community_yap.gox:188:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:189:1
				zlog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:190:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:192:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:193:1
			ctx.Yap__1("edit", article)
		}
	})
//line cmd/gopcomm/community_yap.gox:197:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:198:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:199:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:200:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:201:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:206:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:213:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:215:1
		trans := ctx.Param("trans")
//line cmd/gopcomm/community_yap.gox:216:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:217:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:218:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:220:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:221:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:222:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:227:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:228:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:229:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:236:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags")}, Content: mdData, HtmlData: htmlData}
//line cmd/gopcomm/community_yap.gox:247:1
		id, _ = this.community.PutArticle(todo, uid, trans, article)
//line cmd/gopcomm/community_yap.gox:248:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:256:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:258:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:259:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:260:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:265:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:266:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:267:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:273:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:274:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:275:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:277:1
		transData, err := this.trans.TranslateMarkdownText(mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:278:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:279:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:284:1
		id, _ = this.community.SaveHtml(todo, uid, htmlData, mdData, id)
//line cmd/gopcomm/community_yap.gox:285:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "id": id, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:292:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:293:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:295:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:297:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:300:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:301:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:302:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:303:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:304:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:305:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:310:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:316:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:317:1
		file, header, err := ctx.FormFile("file")
//line cmd/gopcomm/community_yap.gox:318:1
		filename := header.Filename
//line cmd/gopcomm/community_yap.gox:320:1
		ctx.ParseMultipartForm(10 << 20)
//line cmd/gopcomm/community_yap.gox:322:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:323:1
			zlog.Error("upload file error:", filename)
//line cmd/gopcomm/community_yap.gox:324:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:325:1
			return
		}
//line cmd/gopcomm/community_yap.gox:329:1
		dst, err := os.Create(filename)
//line cmd/gopcomm/community_yap.gox:330:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:331:1
			zlog.Error("create file error:", file)
//line cmd/gopcomm/community_yap.gox:332:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:333:1
			return
		}
//line cmd/gopcomm/community_yap.gox:335:1
		defer func() {
//line cmd/gopcomm/community_yap.gox:336:1
			file.Close()
//line cmd/gopcomm/community_yap.gox:337:1
			dst.Close()
//line cmd/gopcomm/community_yap.gox:338:1
			err = os.Remove(filename)
//line cmd/gopcomm/community_yap.gox:339:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:340:1
				zlog.Error("delete file error:", filename)
//line cmd/gopcomm/community_yap.gox:341:1
				return
			}
		}()
//line cmd/gopcomm/community_yap.gox:346:1
		_, err = io.Copy(dst, file)
//line cmd/gopcomm/community_yap.gox:347:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:348:1
			zlog.Error("copy file errer:", filename)
//line cmd/gopcomm/community_yap.gox:349:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:350:1
			return
		}
//line cmd/gopcomm/community_yap.gox:352:1
		bytes, err := os.ReadFile(filename)
//line cmd/gopcomm/community_yap.gox:353:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:354:1
			zlog.Error("read file errer:", filename)
//line cmd/gopcomm/community_yap.gox:355:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:356:1
			return
		}
//line cmd/gopcomm/community_yap.gox:358:1
		token, err := ctx.Request.Cookie("token")
//line cmd/gopcomm/community_yap.gox:359:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:360:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:365:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:366:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:367:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:372:1
		id, err := this.community.SaveMedia(context.Background(), uid, bytes)
//line cmd/gopcomm/community_yap.gox:373:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:374:1
			zlog.Error("save file", err.Error())
//line cmd/gopcomm/community_yap.gox:375:1
			ctx.JSON(500, err.Error())
//line cmd/gopcomm/community_yap.gox:376:1
			return
		}
//line cmd/gopcomm/community_yap.gox:380:1
		ctx.JSON(200, id)
	})
//line cmd/gopcomm/community_yap.gox:383:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:384:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:385:1
	this.trans = translation.New(os.Getenv("NIUTRANS_API_KEY"), "", "")
//line cmd/gopcomm/community_yap.gox:386:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:388:1
	zlog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:389:1
	this.Run(endpoint)
}
func main() {
	yap.Gopt_App_Main(new(community))
}
