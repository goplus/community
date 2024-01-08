package main

import (
	"fmt"
	"github.com/goplus/yap"
	"github.com/goplus/community/cmd/gopcomm/internal/config"
)

type community struct {
	yap.App
}
//line cmd/gopcomm/community_yap.gox:3
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:3:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:4:1
		ctx.Yap__1("article", map[string]string{"id": ctx.Param("id")})
	})
//line cmd/gopcomm/community_yap.gox:8:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:9:1
		ctx.Yap__1("home", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:12:1
	this.Get("/user/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:13:1
		ctx.Yap__1("user", map[string]string{"id": ctx.Param("id")})
	})
//line cmd/gopcomm/community_yap.gox:17:1
	this.Get("/login", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:18:1
		ctx.Yap__1("login", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:20:1
	this.Get("/register", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:21:1
		ctx.Yap__1("register", map[string]interface {
		}{})
	})
//line cmd/gopcomm/community_yap.gox:24:1
	config.ConfigInit("./config.yaml")
//line cmd/gopcomm/community_yap.gox:25:1
	fmt.Println(config.Config)
//line cmd/gopcomm/community_yap.gox:26:1
	fmt.Println("Community server running on :8080")
//line cmd/gopcomm/community_yap.gox:27:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community))
}
