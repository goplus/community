package main

import (
	"fmt"
	"github.com/goplus/yap"
	"github.com/goplus/community/internal/core"
)

type community struct {
	yap.App
}
// TODO: Config Init
//
//line cmd/gopcomm/community_yap.gox:4
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:4:1
	config := &core.Config{core.ArticleConfig{}}
//line cmd/gopcomm/community_yap.gox:7:1
	community := core.New(config)
//line cmd/gopcomm/community_yap.gox:9:1
	this.Get("/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:10:1
		param := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:11:1
		fmt.Println("Visiting article " + param)
//line cmd/gopcomm/community_yap.gox:13:1
		h := community.GetArticleHandler()
//line cmd/gopcomm/community_yap.gox:14:1
		info, _ := h.GetArticle(param)
//line cmd/gopcomm/community_yap.gox:16:1
		ctx.Yap__1("article", map[string]string{"id": info.Id, "title": info.Title, "content": info.Content})
	})
//line cmd/gopcomm/community_yap.gox:22:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:23:1
		ctx.Yap__1("home", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:26:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:27:1
		ctx.Yap__1("user", map[string]string{"id": ctx.Param("id")})
	})
//line cmd/gopcomm/community_yap.gox:32:1
	fmt.Println("Community server running on :8080")
//line cmd/gopcomm/community_yap.gox:33:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community))
}
