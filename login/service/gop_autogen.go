package main

import (
	"fmt"
	"github.com/goplus/yap"
	"log"
	"net/http"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type community struct {
	yap.App
}
//line login/login/service/community_yap.gox:15
func (this *community) MainEntry() {
//line login/login/service/community_yap.gox:15:1
	endPoint := "https://casdoor-community.qiniu.io"
//line login/login/service/community_yap.gox:16:1
	clientID := "49a8ac9729e314a05bf0"
//line login/login/service/community_yap.gox:17:1
	clientSecret := "9ff5d5be617618c92180c668ff28562271f310b4"
//line login/login/service/community_yap.gox:18:1
	certificate := "-----BEGIN CERTIFICATE-----\nMIIE3TCCAsWgAwIBAgIDAeJAMA0GCSqGSIb3DQEBCwUAMCgxDjAMBgNVBAoTBWFk\nbWluMRYwFAYDVQQDEw1jZXJ0LWJ1aWx0LWluMB4XDTI0MDExMDA5MjUxNVoXDTQ0\nMDExMDA5MjUxNVowKDEOMAwGA1UEChMFYWRtaW4xFjAUBgNVBAMTDWNlcnQtYnVp\nbHQtaW4wggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQDZ6k0hJ5aCOgHI\nkLY+doPSYI/WB0F36saOKtGolxA1j6p99sJZTVRDocsseQj6EGIXKdRi58H5ncfG\nIjE7ePoKVI+U4e67VtE1Y5fmI6Dg740tSh7q1+IJkCHZ06dz24irzgVDWMiEl5eQ\n+ekq5ebN0zCIK/MFE0QEraWp3jX3MXiiBshnZgbI8TMyQCD0CNR/OiB+9VEzUlQO\nBQAGefzOZkj7NxV9USnVu5TsQBvEzL43hlrQ5c5NgbRg8VGKluByNCniSftRP08H\nUebc2nFc0cbV9vkyCmR9gHAsFyJBH8593q5sahXoUaZYcwpu7aieH9dUqSitsOm8\n5gGA/Aueude29kKrFSDwzdg5l1qYDOdpnPsJ2282EmM+/V8EG0S8+y4sWSFDo6km\n6JZhi0ULHjJI0kCjDzykrpoEgzroN1Zfdk33MH2EecUBirnJ0d+t+VOzB2z7O/Ve\n+bKlniZs9SEIVQCIStOTxXGx6gmMZyXQJzTaaBfN6Y2N/IXAimrbgaAXwtTJAimk\nCRT13sMDpdGBCMzWAG4NVl2mTT+NrbbYT2GS3MgqBLZCLcbWb6zXb0WUbmn/fmMl\nbJO4ilzTOm1cmZ9W5NyZFexiOFFt1YadMqcE9tttDLiBIsG3v4p4mhUDY19GUHWI\nrM651V/nP9WqM79N6zYJGc/hf1kpcwIDAQABoxAwDjAMBgNVHRMBAf8EAjAAMA0G\nCSqGSIb3DQEBCwUAA4ICAQCpBBW8zfljoCNBTra7pf5KSrwriDt+v+svkF3NStnQ\nz4VhSRTJtmhrMPhIOTFPNd1spmtqux6ph8Xdom3vnk5ZyZAOo9RefFmaiBjJ5BL8\nvD6KEj1f0tMs2HzgVeJe8fMKO9OC2Ltz0f8bzZiIFDySwdTaya6XgCjNgSyN+3hW\nQRklKtljSWsTxtUfvdPr5r4j62f13Jvm5rjrHk9YH0/lvKG+JVLoLFiJTntPZNdL\niCo7o7nu2xuSJo4ds3cgKfAdbRTU9smyLTNSdaGSI4r02sLvXDVEx99hquGGB8FO\nQg9mqmP83QsuXdNot23vcWoh+1jCZ/KIu77x2DqwtZn8Lh77V0lO18pQYeLC1QSc\nfnK6gliAmB86HeYSxUscimTXOtdRNNopuik4x36V1q3ZislEXOUgiGJvx/rGVicC\nTVXs8+5eexakcGteA3FDp95RCVFBedMGzjXPEbnNs2th4DPUQZvOLUtCB1BZwfE8\nEXkJGv/WFKqdQszfKl3hFxocmZppciLV+rTFMnZNNSXcIZIn1n7kEPy/2Teq0ObI\n7FqI+/gZu2Shw82RBcVtrTfRqY51S/a/lB9Ofu2kGnR05XzVX3gtJMx1GVjMhNck\npOXpMCdMP22XpVFd/uKXdCd9tGGJUd5C+ZvcsThY9susU/LilAL9WQbhH1amJDsG\nbQ==\n-----END CERTIFICATE-----"
//line login/login/service/community_yap.gox:19:1
	organizationName := "built-in"
//line login/login/service/community_yap.gox:20:1
	applicationName := "application_ot2837"
//line login/login/service/community_yap.gox:21:1
	casdoorsdk.InitConfig(endPoint, clientID, clientSecret, certificate, organizationName, applicationName)
//line login/login/service/community_yap.gox:24:1
	this.Get("/p/:id", func(ctx *yap.Context) {
//line login/login/service/community_yap.gox:25:1
		ctx.Yap__1("article", map[string]string{"id": ctx.Param("id")})
	})
//line login/login/service/community_yap.gox:29:1
	this.Get("/", func(ctx *yap.Context) {
//line login/login/service/community_yap.gox:30:1
		ctx.Yap__1("home", map[string]interface {
		}{})
	})
//line login/login/service/community_yap.gox:36:1
	this.Get("/login", func(ctx *yap.Context) {
//line login/login/service/community_yap.gox:37:1
		casdoorURL := "https://casdoor-community.qiniu.io"
//line login/login/service/community_yap.gox:38:1
		clientID := "49a8ac9729e314a05bf0"
//line login/login/service/community_yap.gox:39:1
		redirectURI := "http://localhost:8080/callback"
//line login/login/service/community_yap.gox:40:1
		ResponseType := "code"
//line login/login/service/community_yap.gox:41:1
		Scope := "read"
//line login/login/service/community_yap.gox:42:1
		State := "casdoor"
//line login/login/service/community_yap.gox:43:1
		loginURL := fmt.Sprintf("%s/login/oauth/authorize?client_id=%s&response_type=%s&redirect_uri=%s&scope=%s&state=%s", casdoorURL, clientID, ResponseType, redirectURI, Scope, State)
//line login/login/service/community_yap.gox:45:1
		http.Redirect(ctx.ResponseWriter, ctx.Request, loginURL, http.StatusFound)
	})
//line login/login/service/community_yap.gox:49:1
	this.Get("/callback", func(ctx *yap.Context) {
//line login/login/service/community_yap.gox:51:1
		code := ctx.URL.Query().Get("code")
//line login/login/service/community_yap.gox:52:1
		state := ctx.URL.Query().Get("state")
//line login/login/service/community_yap.gox:53:1
		fmt.Println("--code and state--")
//line login/login/service/community_yap.gox:54:1
		fmt.Println("code:", code)
//line login/login/service/community_yap.gox:55:1
		fmt.Println("state:", state)
//line login/login/service/community_yap.gox:57:1
		token, err := casdoorsdk.GetOAuthToken(code, state)
//line login/login/service/community_yap.gox:58:1
		if err != nil {
//line login/login/service/community_yap.gox:59:1
			fmt.Println("err", err)
		}
//line login/login/service/community_yap.gox:62:1
		claim, err := casdoorsdk.ParseJwtToken(token.AccessToken)
//line login/login/service/community_yap.gox:63:1
		if err != nil {
//line login/login/service/community_yap.gox:64:1
			log.Fatal("err", err)
		} else {
		}
//line login/login/service/community_yap.gox:68:1
		fmt.Println("============================================================")
//line login/login/service/community_yap.gox:69:1
		username := claim.User.Name
//line login/login/service/community_yap.gox:70:1
		ctx.Yap__1("callback", map[string]string{"username": username, "usertoken": token.AccessToken})
	})
//line login/login/service/community_yap.gox:79:1
	this.Get("/test", func(ctx *yap.Context) {
//line login/login/service/community_yap.gox:80:1
		cookie, _ := ctx.Request.Cookie("token")
//line login/login/service/community_yap.gox:83:1
		claim, err := casdoorsdk.ParseJwtToken(cookie.Value)
//line login/login/service/community_yap.gox:84:1
		if err != nil {
//line login/login/service/community_yap.gox:85:1
			log.Fatal("err", err)
		} else {
//line login/login/service/community_yap.gox:87:1
			fmt.Println(claim.User.Name)
		}
//line login/login/service/community_yap.gox:89:1
		ctx.Yap__1("test", map[string]interface {
		}{})
	})
//line login/login/service/community_yap.gox:94:1
	this.Run__1(":8080")
}
func main() {
	yap.Gopt_App_Main(new(community))
}
