package main

import (
	"fmt"
	"os"
	"github.com/goplus/yap"
	"net/http"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"go.uber.org/zap"
	_ "github.com/joho/godotenv/autoload"
)

type account struct {
	yap.App
}
//line account/cmd/account_yap.gox:11
func (this *account) MainEntry() {
//line account/cmd/account_yap.gox:11:1
	endpoint := os.Getenv("GOP_ACCOUNT_ENDPOINT")
//line account/cmd/account_yap.gox:12:1
	logger, _ := zap.NewProduction()
//line account/cmd/account_yap.gox:13:1
	defer logger.Sync()
//line account/cmd/account_yap.gox:14:1
	zlog := logger.Sugar()
//line account/cmd/account_yap.gox:16:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line account/cmd/account_yap.gox:17:1
		ctx.Yap__1("article", map[string]string{"id": ctx.Param("id")})
	})
//line account/cmd/account_yap.gox:21:1
	this.Get("/", func(ctx *yap.Context) {
//line account/cmd/account_yap.gox:22:1
		ctx.Yap__1("home", map[string]interface {
		}{})
	})
//line account/cmd/account_yap.gox:25:1
	this.Get("/login", func(ctx *yap.Context) {
//line account/cmd/account_yap.gox:26:1
		casdoorURL := "https://casdoor-community.qiniu.io"
//line account/cmd/account_yap.gox:27:1
		clientID := "49a8ac9729e314a05bf0"
//line account/cmd/account_yap.gox:28:1
		redirectURI := "http://localhost:8080/callback"
//line account/cmd/account_yap.gox:29:1
		ResponseType := "code"
//line account/cmd/account_yap.gox:30:1
		Scope := "read"
//line account/cmd/account_yap.gox:31:1
		State := "casdoor"
//line account/cmd/account_yap.gox:32:1
		loginURL := fmt.Sprintf("%s/login/oauth/authorize?client_id=%s&response_type=%s&redirect_uri=%s&scope=%s&state=%s", casdoorURL, clientID, ResponseType, redirectURI, Scope, State)
//line account/cmd/account_yap.gox:33:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, loginURL, http.StatusFound)
	})
//line account/cmd/account_yap.gox:36:1
	this.Get("/callback", func(ctx *yap.Context) {
//line account/cmd/account_yap.gox:37:1
		code := ctx.URL.Query().Get("code")
//line account/cmd/account_yap.gox:38:1
		state := ctx.URL.Query().Get("state")
//line account/cmd/account_yap.gox:39:1
		fmt.Println("--code and state--")
//line account/cmd/account_yap.gox:40:1
		fmt.Println("code:", code)
//line account/cmd/account_yap.gox:41:1
		fmt.Println("state:", state)
//line account/cmd/account_yap.gox:42:1
		token, err := casdoorsdk.GetOAuthToken(code, state)
//line account/cmd/account_yap.gox:43:1
		if err != nil {
//line account/cmd/account_yap.gox:44:1
			fmt.Println("err", err)
		}
//line account/cmd/account_yap.gox:46:1
		claim, err := casdoorsdk.ParseJwtToken(token.AccessToken)
//line account/cmd/account_yap.gox:47:1
		if err != nil {
//line account/cmd/account_yap.gox:48:1
			zlog.Error("err", err)
		}
//line account/cmd/account_yap.gox:50:1
		username := claim.User.Name
//line account/cmd/account_yap.gox:51:1
		ctx.Yap__1("callback", map[string]string{"username": username, "usertoken": token.AccessToken})
	})
//line account/cmd/account_yap.gox:57:1
	this.Get("/test", func(ctx *yap.Context) {
//line account/cmd/account_yap.gox:58:1
		cookie, _ := ctx.Request.Cookie("token")
//line account/cmd/account_yap.gox:59:1
		claim, err := casdoorsdk.ParseJwtToken(cookie.Value)
//line account/cmd/account_yap.gox:60:1
		if err != nil {
//line account/cmd/account_yap.gox:61:1
			zlog.Error("err", err)
		} else {
//line account/cmd/account_yap.gox:63:1
			fmt.Println("username:", claim.User.Name)
//line account/cmd/account_yap.gox:64:1
			fmt.Println("userId:", claim.User.Id)
//line account/cmd/account_yap.gox:65:1
			fmt.Println("password:", claim.User.Password)
		}
//line account/cmd/account_yap.gox:67:1
		ctx.Yap__1("test", map[string]interface {
		}{})
	})
//line account/cmd/account_yap.gox:70:1
	zlog.Info("Started in endpoint: ", endpoint)
//line account/cmd/account_yap.gox:71:1
	this.Run(endpoint)
}
func main() {
	yap.Gopt_App_Main(new(account))
}
