package main

import (
	"fmt"
	"github.com/goplus/yap"
	"github.com/goplus/community/internal/core"
	"github.com/spf13/viper"
)

type community struct {
	yap.App
	community *core.Community
}
//line cmd/gopcomm/community_yap.gox:11
func (this *community) MainEntry() {
//line cmd/gopcomm/community_yap.gox:11:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:12:1
		id := ctx.Param("id")
//line cmd/gopcomm/community_yap.gox:13:1
		article, _ := this.community.Article(nil, id)
//line cmd/gopcomm/community_yap.gox:15:1
		ctx.Yap__1("article", map[string]string{"ID": id, "Title": article.Title, "Body": article.HtmlText})
	})
//line cmd/gopcomm/community_yap.gox:21:1
	this.Get("/", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:22:1
		articles, _, _ := this.community.ListArticle(nil, core.MarkBegin, 20)
//line cmd/gopcomm/community_yap.gox:23:1
		ctx.Yap__1("home", map[string][]*core.ArticleEntry{"Items": articles})
	})
//line cmd/gopcomm/community_yap.gox:27:1
	this.Get("/edit", func(ctx *yap.Context) {
//line cmd/gopcomm/community_yap.gox:29:1
		ctx.Yap__1("edit", map[string]string{"ID": ctx.Param("id")})
	})
//line cmd/gopcomm/community_yap.gox:33:1
	this.Post("/commit", func(ctx *yap.Context) {
	})
//line cmd/gopcomm/community_yap.gox:37:1
	viper.SetConfigFile("./config.yaml")
//line cmd/gopcomm/community_yap.gox:38:1
	viper.AddConfigPath(".")
//line cmd/gopcomm/community_yap.gox:40:1
	if
//line cmd/gopcomm/community_yap.gox:40:1
	err := viper.ReadInConfig(); err != nil {
//line cmd/gopcomm/community_yap.gox:41:1
		fmt.Println(err)
//line cmd/gopcomm/community_yap.gox:42:1
		return
	}
//line cmd/gopcomm/community_yap.gox:45:1
	config := &core.Config{Mysql: core.Mysql{Ip: viper.GetString("mysql.ip"), Port: viper.GetInt("mysql.port"), Username: viper.GetString("mysql.username"), Password: viper.GetString("mysql.password"), Database: viper.GetString("mysql.database")}}
//line cmd/gopcomm/community_yap.gox:56:1
	this.community, _ = core.New(config)
//line cmd/gopcomm/community_yap.gox:57:1
	core.InitMysqlDB(config)
//line cmd/gopcomm/community_yap.gox:59:1
	fmt.Println("start")
//line cmd/gopcomm/community_yap.gox:60:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community))
}
