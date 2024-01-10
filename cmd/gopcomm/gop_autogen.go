package main

import (
	"fmt"
	"github.com/goplus/yap"
	"github.com/goplus/community/internal/core"
)

type community struct {
	yap.App
}
//line cmd/gopcomm/community_yap.gox:3
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:3:1
	this.Get("/article/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:4:1
		param := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:5:1
		fmt.Println("Visiting article " + param)
//line cmd/gopcomm/community_yap.gox:6:1
		s := core.NewArticleService()
//line cmd/gopcomm/community_yap.gox:7:1
		info, _ := s.GetArticle(param)
//line cmd/gopcomm/community_yap.gox:9:1
		ctx.Yap__1("article", map[string]string{"id": info.Id, "title": info.Title, "content": info.Content})
	})
//line cmd/gopcomm/community_yap.gox:15:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:16:1
		ctx.Yap__1("home", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:19:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:20:1
		ctx.Yap__1("user", map[string]string{"id": ctx.Param("id")})
	})
//line cmd/gopcomm/community_yap.gox:25:1
	fmt.Println("Community server running on :8080")
//line cmd/gopcomm/community_yap.gox:26:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community))
}
