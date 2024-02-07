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
//line cmd/gopcomm/community_yap.gox:56:1
	this.Get("/signin", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:57:1
		ctx.Yap__1("signin", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:60:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:63:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:64:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:65:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:66:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:67:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:68:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:72:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:73:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:74:1
=======
//line cmd/gopcomm/community_yap.gox:53:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:56:1
		// todo middleware
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:57:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:58:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:59:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:60:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:61:1
				xLog.Error("get user error:", err)
			}
		}
//line cmd/gopcomm/community_yap.gox:65:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:66:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:67:1
		ctx.Yap__1("article", map[string]interface {
		}{"User": user, "ID": id, "Title": article.Title, "Content": article.Content, "HtmlUrl": article.HtmlUrl, "Tags": article.Tags, "Cover": article.Cover, "Mtime": article.Mtime.Format(layoutUS), "Author": article.User})
	})
//line cmd/gopcomm/community_yap.gox:87:1
	this.Get("/getArticle/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:88:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:89:1
		article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:90:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": article})
	})
//line cmd/gopcomm/community_yap.gox:96:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:97:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:99:1
		userClaim, err := this.community.GetUserClaim(id)
//line cmd/gopcomm/community_yap.gox:100:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:101:1
			xLog.Error("get current user error:", err)
		}
//line cmd/gopcomm/community_yap.gox:105:1
		// todo middleware
		// get user by token
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
//line cmd/gopcomm/community_yap.gox:114:1
		items, _ := this.community.GetArticlesByUid(todo, id)
//line cmd/gopcomm/community_yap.gox:115:1
		userClaimJson, _ := json.Marshal(&userClaim)
//line cmd/gopcomm/community_yap.gox:116:1
		itemsJson, _ := json.Marshal(&items)
//line cmd/gopcomm/community_yap.gox:117:1
		ctx.Yap__1("user", map[string]interface {
		}{"Id": id, "CurrentUser": strings.Replace(string(userClaimJson), `\"`, `"`, -1), "User": user, "Items": strings.Replace(string(itemsJson), `\"`, `"`, -1)})
	})
//line cmd/gopcomm/community_yap.gox:125:1
	this.Get("/add", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:126:1
		ctx.Yap__1("edit", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:129:1
	this.Get("/delete", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:130:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:131:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:132:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:133:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:134:1
			xLog.Error("token parse error")
//line cmd/gopcomm/community_yap.gox:135:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:140:1
		err = this.community.DeleteArticle(todo, uid, id)
//line cmd/gopcomm/community_yap.gox:141:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:142:1
			ctx.Json__1(map[string]interface {
			}{"code": 0, "err": "delete failed"})
		} else {
//line cmd/gopcomm/community_yap.gox:147:1
=======
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
//line cmd/gopcomm/community_yap.gox:154:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:156:1
		// Get User Info
		var user *core.User
//line cmd/gopcomm/community_yap.gox:157:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:158:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:159:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:160:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:161:1
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
//line cmd/gopcomm/community_yap.gox:165:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, "")
//line cmd/gopcomm/community_yap.gox:166:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:167:1
//line cmd/gopcomm/community_yap.gox:208:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, "")
//line cmd/gopcomm/community_yap.gox:209:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:210:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:174:1
	this.Get("/get", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:175:1
		from := ctx.Param("from")
//line cmd/gopcomm/community_yap.gox:176:1
		limit := ctx.Param("limit")
//line cmd/gopcomm/community_yap.gox:177:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:179:1
		limitInt, err := strconv.Atoi(limit)
//line cmd/gopcomm/community_yap.gox:180:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:181:1
			limitInt = limitConst
		}
//line cmd/gopcomm/community_yap.gox:184:1
		articles, next, _ := this.community.ListArticle(todo, from, limitInt, searchValue)
//line cmd/gopcomm/community_yap.gox:186:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "items": articles, "next": next, "value": searchValue})
	})
//line cmd/gopcomm/community_yap.gox:194:1
	this.Get("/search", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:195:1
		searchValue := ctx.Param("value")
//line cmd/gopcomm/community_yap.gox:196:1
		if searchValue == "" {
//line cmd/gopcomm/community_yap.gox:197:1
			ctx.Json__1(map[string]interface {
			}{"code": 400, "err": "value can not be ''."})
		}
//line cmd/gopcomm/community_yap.gox:204:1
		// todo middleware
		var user *core.User
//line cmd/gopcomm/community_yap.gox:205:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:206:1
		if err == nil {
//line cmd/gopcomm/community_yap.gox:207:1
			user, err = this.community.GetUser(token.Value)
//line cmd/gopcomm/community_yap.gox:208:1
			if err != nil {
//line cmd/gopcomm/community_yap.gox:209:1
=======
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
//line cmd/gopcomm/community_yap.gox:213:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue)
//line cmd/gopcomm/community_yap.gox:214:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:215:1
//line cmd/gopcomm/community_yap.gox:259:1
		articles, next, _ := this.community.ListArticle(todo, core.MarkBegin, limitConst, searchValue)
//line cmd/gopcomm/community_yap.gox:260:1
		articlesJson, _ := json.Marshal(&articles)
//line cmd/gopcomm/community_yap.gox:261:1
		ctx.Yap__1("home", map[string]interface {
		}{"UserId": userId, "User": user, "Items": strings.Replace(string(articlesJson), `\"`, `"`, -1), "Value": searchValue, "Next": next})
	})
//line cmd/gopcomm/community_yap.gox:223:1
	this.Get("/edit/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:224:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:225:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:226:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:232:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:233:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:234:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:240:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:241:1
		if id != "" {
//line cmd/gopcomm/community_yap.gox:242:1
			if
//line cmd/gopcomm/community_yap.gox:242:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
//line cmd/gopcomm/community_yap.gox:243:1
				xLog.Error("no permissions")
//line cmd/gopcomm/community_yap.gox:244:1
				http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
			}
//line cmd/gopcomm/community_yap.gox:246:1
			article, _ := this.community.Article(todo, id)
//line cmd/gopcomm/community_yap.gox:247:1
			ctx.Yap__1("edit", article)
		}
	})
//line cmd/gopcomm/community_yap.gox:251:1
	this.Get("/getTrans", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:252:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:253:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:254:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:255:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:260:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:267:1
	this.Post("/commit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:269:1
		trans := ctx.Param("trans")
//line cmd/gopcomm/community_yap.gox:270:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:271:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:272:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:274:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:275:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:276:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:281:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:282:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:283:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:290:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags"), Abstract: ctx.Param("abstract")}, Content: mdData, HtmlData: htmlData}
//line cmd/gopcomm/community_yap.gox:302:1
		id, _ = this.community.PutArticle(todo, uid, trans, article)
//line cmd/gopcomm/community_yap.gox:303:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:311:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:313:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:314:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:315:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:320:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:321:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:322:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:328:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:329:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:330:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:332:1
		transData, err := this.trans.TranslateMarkdownText(mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:333:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:334:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:339:1
		id, _ = this.community.SaveHtml(todo, uid, htmlData, mdData, id)
//line cmd/gopcomm/community_yap.gox:340:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "id": id, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:347:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:348:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:350:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:352:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:355:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:356:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:357:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:358:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:359:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:360:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:365:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:371:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:372:1
		core.UploadFile(ctx, this.community)
	})
//line cmd/gopcomm/community_yap.gox:375:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:380:1
		redirectURL := fmt.Sprintf("%s/%s", ctx.Request.Referer(), "callback")
//line cmd/gopcomm/community_yap.gox:382:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:383:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:387:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:388:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:389:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:390:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:394:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("/"), http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:397:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:398:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:399:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:400:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:405:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("/"), http.StatusFound)
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
//line cmd/gopcomm/community_yap.gox:316:1
		trans := ctx.Param("trans")
//line cmd/gopcomm/community_yap.gox:317:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:318:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:319:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:321:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:322:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:323:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:328:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:329:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:330:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:337:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags"), Abstract: ctx.Param("abstract")}, Content: mdData, HtmlData: htmlData}
//line cmd/gopcomm/community_yap.gox:349:1
		id, _ = this.community.PutArticle(todo, uid, trans, article)
//line cmd/gopcomm/community_yap.gox:350:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": id})
	})
//line cmd/gopcomm/community_yap.gox:358:1
	this.Post("/translate", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:360:1
		token, err := core.GetToken(ctx)
//line cmd/gopcomm/community_yap.gox:361:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:362:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "no token"})
		}
//line cmd/gopcomm/community_yap.gox:367:1
		uid, err := this.community.ParseJwtToken(token.Value)
//line cmd/gopcomm/community_yap.gox:368:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:369:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:375:1
		mdData := ctx.Param("content")
//line cmd/gopcomm/community_yap.gox:376:1
		htmlData := ctx.Param("html")
//line cmd/gopcomm/community_yap.gox:377:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:392:1
		transData, err := this.community.TranslateMarkdownText(todo, mdData, language.Chinese, language.English)
//line cmd/gopcomm/community_yap.gox:393:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:381:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
//line cmd/gopcomm/community_yap.gox:386:1
		id, _ = this.community.SaveHtml(todo, uid, htmlData, mdData, id)
//line cmd/gopcomm/community_yap.gox:387:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "id": id, "data": transData})
	})
//line cmd/gopcomm/community_yap.gox:394:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:395:1
		mediaId := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:397:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
//line cmd/gopcomm/community_yap.gox:399:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, domain+fileKey, http.StatusTemporaryRedirect)
	})
//line cmd/gopcomm/community_yap.gox:402:1
	this.Get("/getMediaUrl/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:403:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:404:1
		fileKey, err := this.community.GetMediaUrl(todo, id)
//line cmd/gopcomm/community_yap.gox:405:1
		htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
//line cmd/gopcomm/community_yap.gox:406:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:407:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": "have no html media"})
		}
//line cmd/gopcomm/community_yap.gox:412:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "url": htmlUrl})
	})
//line cmd/gopcomm/community_yap.gox:418:1
	this.Post("/upload", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:432:1
		this.community.UploadFile(ctx)
	})
//line cmd/gopcomm/community_yap.gox:422:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:427:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:442:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:443:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:444:1
			return
		}
//line cmd/gopcomm/community_yap.gox:447:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:448:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:449:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:452:1
		redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "callback", url.QueryEscape(refererPath))
//line cmd/gopcomm/community_yap.gox:440:1
		loginURL := this.community.RedirectToCasdoor(redirectURL)
//line cmd/gopcomm/community_yap.gox:441:1
		ctx.Redirect(loginURL, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:445:1
	this.Get("/logout", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:446:1
		err := core.RemoveToken(ctx)
//line cmd/gopcomm/community_yap.gox:447:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:448:1
			xLog.Error("remove token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:451:1
		refererURL, err := url.Parse(ctx.Request.Referer())
//line cmd/gopcomm/community_yap.gox:452:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:453:1
			xLog.Info("Error parsing Referer: %v", err)
//line cmd/gopcomm/community_yap.gox:454:1
			return
		}
//line cmd/gopcomm/community_yap.gox:457:1
		refererPath := refererURL.Path
//line cmd/gopcomm/community_yap.gox:458:1
		if refererURL.RawQuery != "" {
//line cmd/gopcomm/community_yap.gox:459:1
			refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
		}
//line cmd/gopcomm/community_yap.gox:462:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, refererPath, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:465:1
	this.Get("/callback", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:466:1
		err := core.SetToken(ctx)
//line cmd/gopcomm/community_yap.gox:467:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:468:1
			xLog.Error("set token error:", err)
		}
//line cmd/gopcomm/community_yap.gox:470:1
		origin_path := ctx.URL.Query().Get("origin_path")
//line cmd/gopcomm/community_yap.gox:471:1
		unurl, err := url.QueryUnescape(origin_path)
//line cmd/gopcomm/community_yap.gox:472:1
		if err != nil {
//line cmd/gopcomm/community_yap.gox:473:1
			xLog.Info("Unurl error", err)
//line cmd/gopcomm/community_yap.gox:474:1
			unurl = "/"
		}
//line cmd/gopcomm/community_yap.gox:477:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, unurl, http.StatusFound)
	})
//line cmd/gopcomm/community_yap.gox:408:1
	conf := &core.Config{}
//line cmd/gopcomm/community_yap.gox:409:1
	this.community, _ = core.New(todo, conf)
//line cmd/gopcomm/community_yap.gox:410:1
	this.trans = translation.New(os.Getenv("NIUTRANS_API_KEY"), "", "")
//line cmd/gopcomm/community_yap.gox:411:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:414:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:415:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:418:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:421:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:423:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:424:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:425:1
				if
//line cmd/gopcomm/community_yap.gox:425:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:426:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:430:1
//line cmd/gopcomm/community_yap.gox:499:1
	core.CasdoorConfigInit()
//line cmd/gopcomm/community_yap.gox:502:1
	this.Handle("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:503:1
		ctx.Yap__1("4xx", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:506:1
	xLog.Info("Started in endpoint: ", endpoint)
//line cmd/gopcomm/community_yap.gox:509:1
	this.Run(endpoint, func(h http.Handler) http.Handler {
//line cmd/gopcomm/community_yap.gox:511:1
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//line cmd/gopcomm/community_yap.gox:512:1
			defer func() {
//line cmd/gopcomm/community_yap.gox:513:1
				if
//line cmd/gopcomm/community_yap.gox:513:1
				err := recover(); err != nil {
//line cmd/gopcomm/community_yap.gox:514:1
					http.Redirect(w, r, "/failed", http.StatusFound)
				}
			}()
//line cmd/gopcomm/community_yap.gox:518:1
			h.ServeHTTP(w, r)
		})
	})
}
func main() {
	yap.Gopt_App_Main(new(community))
}
