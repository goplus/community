package main

import (
	"context"
	"fmt"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"github.com/goplus/yap"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/text/language"
	"io"
	"log"
	"net/http"
	"os"
)

type community struct {
	yap.App
	community *core.Community
	trans     *translation.Engine
}

//line cmd/gopcomm/community_yap.gox:21
func (this *community) MainEntry() {
	//line cmd/gopcomm/community_yap.gox:21:1
	todo := context.TODO()
	//line cmd/gopcomm/community_yap.gox:23:1
	this.Get("/p/:id", func(ctx *yap.Context) {
		//line cmd/gopcomm/community_yap.gox:24:1
		id := ctx.Param("id")
		//line cmd/gopcomm/community_yap.gox:25:1
		article, _ := this.community.Article(todo, id)
		//line cmd/gopcomm/community_yap.gox:26:1
		ctx.Yap__1("article", map[string]string{"ID": id, "Title": article.Title, "Content": article.Content})
	})
	//line cmd/gopcomm/community_yap.gox:32:1
	this.Get("/add", func(ctx *yap.Context) {
		//line cmd/gopcomm/community_yap.gox:33:1
		ctx.Yap__1("add", map[string]interface {
		}{})
	})
	//line cmd/gopcomm/community_yap.gox:36:1
	this.Get("/", func(ctx *yap.Context) {
		//line cmd/gopcomm/community_yap.gox:37:1
		articles, _, _ := this.community.ListArticle(todo, core.MarkBegin, 20)
		//line cmd/gopcomm/community_yap.gox:38:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
	//line cmd/gopcomm/community_yap.gox:42:1
	this.Get("/edit", func(ctx *yap.Context) {
		//line cmd/gopcomm/community_yap.gox:43:1
		cookie, err := ctx.Request.Cookie("user_id")
		//line cmd/gopcomm/community_yap.gox:44:1
		if err != nil {
			//line cmd/gopcomm/community_yap.gox:45:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
		//line cmd/gopcomm/community_yap.gox:50:1
		uid := cookie.Value
		//line cmd/gopcomm/community_yap.gox:51:1
		id := ctx.Param("id")
		//line cmd/gopcomm/community_yap.gox:52:1
		doc := map[string]string{"ID": id}
		//line cmd/gopcomm/community_yap.gox:55:1
		if id != "" {
			//line cmd/gopcomm/community_yap.gox:56:1
			if
			//line cmd/gopcomm/community_yap.gox:56:1
			editable, _ := this.community.CanEditable(todo, uid, id); !editable {
				//line cmd/gopcomm/community_yap.gox:58:1
				log.Println("no permissions")
				//line cmd/gopcomm/community_yap.gox:59:1
				return
			}
			//line cmd/gopcomm/community_yap.gox:61:1
			article, _ := this.community.Article(todo, id)
			//line cmd/gopcomm/community_yap.gox:62:1
			doc["Title"] = article.Title
			//line cmd/gopcomm/community_yap.gox:63:1
			doc["Content"] = article.Content
		}
		//line cmd/gopcomm/community_yap.gox:65:1
		ctx.Yap__1("edit", doc)
	})
	//line cmd/gopcomm/community_yap.gox:68:1
	this.Get("/getTrans", func(ctx *yap.Context) {
		//line cmd/gopcomm/community_yap.gox:69:1
		id := ctx.Param("id")
		//line cmd/gopcomm/community_yap.gox:70:1
		htmlUrl, err := this.community.TransHtmlUrl(todo, id)
		//line cmd/gopcomm/community_yap.gox:71:1
		if err != nil {
			//line cmd/gopcomm/community_yap.gox:72:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
		//line cmd/gopcomm/community_yap.gox:77:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "data": htmlUrl})
	})
	//line cmd/gopcomm/community_yap.gox:84:1
	this.Post("/commit", func(ctx *yap.Context) {
		//line cmd/gopcomm/community_yap.gox:86:1
		trans := ctx.Param("trans")
		//line cmd/gopcomm/community_yap.gox:87:1
		id := ctx.Param("id")
		//line cmd/gopcomm/community_yap.gox:88:1
		mdData := ctx.Param("content")
		//line cmd/gopcomm/community_yap.gox:89:1
		htmlData := ctx.Param("html")
		//line cmd/gopcomm/community_yap.gox:91:1
		cookie, err := ctx.Request.Cookie("user_id")
		//line cmd/gopcomm/community_yap.gox:92:1
		if err != nil {
			//line cmd/gopcomm/community_yap.gox:93:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
		//line cmd/gopcomm/community_yap.gox:98:1
		uid := cookie.Value
		//line cmd/gopcomm/community_yap.gox:100:1
		article := &core.Article{ArticleEntry: core.ArticleEntry{ID: id, Title: ctx.Param("title"), UId: uid, Cover: ctx.Param("cover"), Tags: ctx.Param("tags")}, Content: mdData, HtmlData: htmlData}
		//line cmd/gopcomm/community_yap.gox:111:1
		id, _ = this.community.PutArticle(todo, uid, trans, article)
		//line cmd/gopcomm/community_yap.gox:112:1
		article.ID = id
		//line cmd/gopcomm/community_yap.gox:113:1
		ctx.Yap__1("edit", *article)
	})
	//line cmd/gopcomm/community_yap.gox:117:1
	this.Post("/translate", func(ctx *yap.Context) {
		//line cmd/gopcomm/community_yap.gox:119:1
		cookie, err := ctx.Request.Cookie("user_id")
		//line cmd/gopcomm/community_yap.gox:120:1
		if err != nil {
			//line cmd/gopcomm/community_yap.gox:121:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
		//line cmd/gopcomm/community_yap.gox:126:1
		uid := cookie.Value
		//line cmd/gopcomm/community_yap.gox:127:1
		mdData := ctx.Param("content")
		//line cmd/gopcomm/community_yap.gox:128:1
		htmlData := ctx.Param("html")
		//line cmd/gopcomm/community_yap.gox:129:1
		id, _ := this.community.SaveHtml(todo, uid, htmlData, mdData)
		//line cmd/gopcomm/community_yap.gox:131:1
		transData, err := this.trans.TranslateMarkdownText(mdData, language.Chinese, language.English)
		//line cmd/gopcomm/community_yap.gox:132:1
		if err != nil {
			//line cmd/gopcomm/community_yap.gox:133:1
			ctx.Json__1(map[string]interface {
			}{"code": 500, "err": err.Error()})
		}
		//line cmd/gopcomm/community_yap.gox:138:1
		ctx.Json__1(map[string]interface {
		}{"code": 200, "id": id, "data": transData})
	})
	//line cmd/gopcomm/community_yap.gox:145:1
	this.Get("/getMedia/:id", func(ctx *yap.Context) {
		//line cmd/gopcomm/community_yap.gox:146:1
		mediaId := ctx.Param("id")
		//line cmd/gopcomm/community_yap.gox:148:1
		fileKey, _ := this.community.GetMediaUrl(context.Background(), mediaId)
		//line cmd/gopcomm/community_yap.gox:150:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, "qiniu demain"+fileKey, http.StatusTemporaryRedirect)
	})
	//line cmd/gopcomm/community_yap.gox:152:1
	this.Post("/upload", func(ctx *yap.Context) {
		//line cmd/gopcomm/community_yap.gox:153:1
		file, header, err := ctx.FormFile("file")
		//line cmd/gopcomm/community_yap.gox:154:1
		filename := header.Filename
		//line cmd/gopcomm/community_yap.gox:156:1
		ctx.ParseMultipartForm(10 << 20)
		//line cmd/gopcomm/community_yap.gox:158:1
		if err != nil {
			//line cmd/gopcomm/community_yap.gox:159:1
			log.Fatalln("upload file error:", filename)
			//line cmd/gopcomm/community_yap.gox:160:1
			ctx.JSON(500, err.Error())
			//line cmd/gopcomm/community_yap.gox:161:1
			return
		}
		//line cmd/gopcomm/community_yap.gox:165:1
		dst, err := os.Create(filename)
		//line cmd/gopcomm/community_yap.gox:166:1
		if err != nil {
			//line cmd/gopcomm/community_yap.gox:167:1
			log.Fatalln("create file error:", file)
			//line cmd/gopcomm/community_yap.gox:168:1
			ctx.JSON(500, err.Error())
			//line cmd/gopcomm/community_yap.gox:169:1
			return
		}
		//line cmd/gopcomm/community_yap.gox:171:1
		defer func() {
			//line cmd/gopcomm/community_yap.gox:172:1
			file.Close()
			//line cmd/gopcomm/community_yap.gox:173:1
			dst.Close()
			//line cmd/gopcomm/community_yap.gox:174:1
			err = os.Remove(filename)
			//line cmd/gopcomm/community_yap.gox:175:1
			if err != nil {
				//line cmd/gopcomm/community_yap.gox:176:1
				log.Fatalln("delete file error:", filename)
				//line cmd/gopcomm/community_yap.gox:177:1
				return
			}
		}()
		//line cmd/gopcomm/community_yap.gox:182:1
		_, err = io.Copy(dst, file)
		//line cmd/gopcomm/community_yap.gox:183:1
		if err != nil {
			//line cmd/gopcomm/community_yap.gox:184:1
			log.Fatalln("copy file errer:", filename)
			//line cmd/gopcomm/community_yap.gox:185:1
			ctx.JSON(500, err.Error())
			//line cmd/gopcomm/community_yap.gox:186:1
			return
		}
		//line cmd/gopcomm/community_yap.gox:188:1
		bytes, err := os.ReadFile(filename)
		//line cmd/gopcomm/community_yap.gox:189:1
		if err != nil {
			//line cmd/gopcomm/community_yap.gox:190:1
			log.Fatalln("read file errer:", filename)
			//line cmd/gopcomm/community_yap.gox:191:1
			ctx.JSON(500, err.Error())
			//line cmd/gopcomm/community_yap.gox:192:1
			return
		}
		//line cmd/gopcomm/community_yap.gox:194:1
		cookie, err := ctx.Request.Cookie("user_id")
		//line cmd/gopcomm/community_yap.gox:195:1
		if err != nil {
			//line cmd/gopcomm/community_yap.gox:196:1
			log.Fatalln("token不存在")
			//line cmd/gopcomm/community_yap.gox:197:1
			ctx.JSON(500, err.Error())
			//line cmd/gopcomm/community_yap.gox:198:1
			return
		}
		//line cmd/gopcomm/community_yap.gox:201:1
		id, err := this.community.SaveMedia(context.Background(), cookie.Value, bytes)
		//line cmd/gopcomm/community_yap.gox:202:1
		if err != nil {
			//line cmd/gopcomm/community_yap.gox:203:1
			log.Fatalln("save file", err.Error())
			//line cmd/gopcomm/community_yap.gox:204:1
			ctx.JSON(500, err.Error())
			//line cmd/gopcomm/community_yap.gox:205:1
			return
		}
		//line cmd/gopcomm/community_yap.gox:209:1
		ctx.JSON(200, id)
	})
	//line cmd/gopcomm/community_yap.gox:212:1
	conf := &core.Config{}
	//line cmd/gopcomm/community_yap.gox:213:1
	this.community, _ = core.New(todo, conf)
	//line cmd/gopcomm/community_yap.gox:214:1
	this.trans = translation.New(os.Getenv("NIUTRANS_API_KEY"), "", "")
	//line cmd/gopcomm/community_yap.gox:216:1
	s3 := os.Getenv("S3_BUCKET")
	//line cmd/gopcomm/community_yap.gox:217:1
	fmt.Println(s3)
	//line cmd/gopcomm/community_yap.gox:218:1
	fmt.Println("start")
	//line cmd/gopcomm/community_yap.gox:219:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community))
}
