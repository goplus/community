package main

import "github.com/goplus/yap"

type community struct {
	yap.App
}
//line cmd/gopcomm/community_yap.gox:1
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:1:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:2:1
		ctx.Yap__1("article", map[string]string{"id": ctx.Param("id")})
	})
//line cmd/gopcomm/community_yap.gox:6:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:7:1
		ctx.Yap__1("home", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:10:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community))
}
