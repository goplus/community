package main

import (
	"fmt"
	"os"
	"github.com/goplus/yap"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"log"
	"net/http"
)

type CasdoorConfig struct {
	endPoint         string
	clientId         string
	clientSecret     string
	certificate      string
	organizationName string
	applicationName  string
}
type community struct {
	yap.App
}
//line login/service/community_yap.gox:11
func (this *community) MainEntry() {
//line login/service/community_yap.gox:11:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line login/service/community_yap.gox:12:1
		ctx.Yap__1("article", map[string]string{"id": ctx.Param("id")})
	})
//line login/service/community_yap.gox:16:1
	this.Get("/", func(ctx *yap.Context) {
//line login/service/community_yap.gox:17:1
		ctx.Yap__1("home", map[string]interface {
		}{})
	})
//line login/service/community_yap.gox:23:1
	this.Get("/login", func(ctx *yap.Context) {
//line login/service/community_yap.gox:24:1
		casdoorURL := "https://casdoor-community.qiniu.io"
//line login/service/community_yap.gox:25:1
		clientID := "49a8ac9729e314a05bf0"
//line login/service/community_yap.gox:26:1
		redirectURI := "http://localhost:8080/callback"
//line login/service/community_yap.gox:27:1
		ResponseType := "code"
//line login/service/community_yap.gox:28:1
		Scope := "read"
//line login/service/community_yap.gox:29:1
		State := "casdoor"
//line login/service/community_yap.gox:30:1
		loginURL := fmt.Sprintf("%s/login/oauth/authorize?client_id=%s&response_type=%s&redirect_uri=%s&scope=%s&state=%s", casdoorURL, clientID, ResponseType, redirectURI, Scope, State)
//line login/service/community_yap.gox:31:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, loginURL, http.StatusFound)
	})
//line login/service/community_yap.gox:34:1
	this.Get("/callback", func(ctx *yap.Context) {
//line login/service/community_yap.gox:36:1
		code := ctx.URL.Query().Get("code")
//line login/service/community_yap.gox:37:1
		state := ctx.URL.Query().Get("state")
//line login/service/community_yap.gox:38:1
		fmt.Println("--code and state--")
//line login/service/community_yap.gox:39:1
		fmt.Println("code:", code)
//line login/service/community_yap.gox:40:1
		fmt.Println("state:", state)
//line login/service/community_yap.gox:41:1
		token, err := casdoorsdk.GetOAuthToken(code, state)
//line login/service/community_yap.gox:42:1
		if err != nil {
//line login/service/community_yap.gox:43:1
			fmt.Println("err", err)
		}
//line login/service/community_yap.gox:45:1
		claim, err := casdoorsdk.ParseJwtToken(token.AccessToken)
//line login/service/community_yap.gox:46:1
		if err != nil {
//line login/service/community_yap.gox:47:1
			log.Fatal("err", err)
		}
//line login/service/community_yap.gox:49:1
		username := claim.User.Name
//line login/service/community_yap.gox:50:1
		ctx.Yap__1("callback", map[string]string{"username": username, "usertoken": token.AccessToken})
	})
//line login/service/community_yap.gox:59:1
	this.Get("/test", func(ctx *yap.Context) {
//line login/service/community_yap.gox:60:1
		cookie, _ := ctx.Request.Cookie("token")
//line login/service/community_yap.gox:61:1
		claim, err := casdoorsdk.ParseJwtToken(cookie.Value)
//line login/service/community_yap.gox:62:1
		if err != nil {
//line login/service/community_yap.gox:63:1
			log.Fatal("err", err)
		} else {
//line login/service/community_yap.gox:65:1
			fmt.Println("username:", claim.User.Name)
//line login/service/community_yap.gox:66:1
			fmt.Println("userId:", claim.User.Id)
//line login/service/community_yap.gox:67:1
			fmt.Println("password:", claim.User.Password)
		}
//line login/service/community_yap.gox:69:1
		ctx.Yap__1("test", map[string]interface {
		}{})
	})
//line login/service/community_yap.gox:74:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community))
}
//line login/service/community.gop:15:1
func init() {
//line login/service/community.gop:16:1
	conf := new(CasdoorConfig)
//line login/service/community.gop:18:1
	endPoint := conf.endPoint
//line login/service/community.gop:19:1
	clientID := conf.clientId
//line login/service/community.gop:20:1
	clientSecret := conf.clientSecret
//line login/service/community.gop:21:1
	certificate := conf.certificate
//line login/service/community.gop:22:1
	organizationName := conf.organizationName
//line login/service/community.gop:23:1
	applicationName := conf.applicationName
//line login/service/community.gop:25:1
	if endPoint == "" {
//line login/service/community.gop:26:1
		endPoint = os.Getenv("GOP_CASDOOR_ENDPOINT")
	}
//line login/service/community.gop:28:1
	if clientID == "" {
//line login/service/community.gop:29:1
		clientID = os.Getenv("GOP_CASDOOR_CLIENTID")
	}
//line login/service/community.gop:31:1
	if clientSecret == "" {
//line login/service/community.gop:32:1
		clientSecret = os.Getenv("GOP_CASDOOR_CLIENTSECRET")
	}
//line login/service/community.gop:34:1
	if certificate == "" {
//line login/service/community.gop:35:1
		certificate = os.Getenv("GOP_CASDOOR_CERTIFICATE")
	}
//line login/service/community.gop:37:1
	if organizationName == "" {
//line login/service/community.gop:38:1
		organizationName = os.Getenv("GOP_CASDOOR_ORGANIZATIONNAME")
	}
//line login/service/community.gop:40:1
	if applicationName == "" {
//line login/service/community.gop:41:1
		applicationName = os.Getenv("GOP_CASDOOR_APPLICATONNAME")
	}
//line login/service/community.gop:43:1
	casdoorsdk.InitConfig(endPoint, clientID, clientSecret, certificate, organizationName, applicationName)
}
