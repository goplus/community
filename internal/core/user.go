package core

import (
	"github.com/goplus/yap"
	"net/http"
	"os"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

// Deprecated: use casdoorsdk instead
type User struct {
	Id       string
	Name     string
	Password string
	Avatar   string
}

type UserClaim casdoorsdk.Claims
type UserInfo casdoorsdk.User

// Init casdoor parser
func CasdoorConfigInit() {
	endPoint := os.Getenv("GOP_CASDOOR_ENDPOINT")
	clientID := os.Getenv("GOP_CASDOOR_CLIENTID")
	clientSecret := os.Getenv("GOP_CASDOOR_CLIENTSECRET")
	certificate := os.Getenv("GOP_CASDOOR_CERTIFICATE")
	organizationName := os.Getenv("GOP_CASDOOR_ORGANIZATIONNAME")
	applicationName := os.Getenv("GOP_CASDOOR_APPLICATONNAME")

	casdoorsdk.InitConfig(endPoint, clientID, clientSecret, certificate, organizationName, applicationName)
}

// GetUser return author by token
func (p *Community) GetUser(token string) (user *User, err error) {
	claim, err := casdoorsdk.ParseJwtToken(token)
	if err != nil {
		p.xLog.Error(err)
		return &User{}, ErrNotExist
	}
	user = &User{
		Name:   claim.Name,
		Avatar: claim.Avatar,
		Id:     claim.Id,
	}
	return
}

// ParseJwtToken return user id by token
func (p *Community) ParseJwtToken(token string) (userId string, err error) {
	claim, err := casdoorsdk.ParseJwtToken(token)
	if err != nil {
		p.xLog.Error(err)
		return "", ErrNotExist
	}
	return claim.Id, nil
}

// GetUserClaim get user（full） by token
func (p *Community) GetUserClaim(uid string) (claim *casdoorsdk.User, err error) {
	claim, err = casdoorsdk.GetUserByUserId(uid)
	if err != nil {
		p.xLog.Error(err)
		return &casdoorsdk.User{}, ErrNotExist
	}
	return
}

// GetUserById get user by uid
func (p *Community) GetUserById(uid string) (user *User, err error) {
	claim, err := casdoorsdk.GetUserByUserId(uid)
	if err != nil {
		p.xLog.Error(err)
		return &User{}, ErrNotExist
	}
	user = &User{
		Name:   claim.Name,
		Avatar: claim.Avatar,
		Id:     claim.Id,
	}
	return
}

// UpdateUserById update user by uid
func (p *Community) UpdateUserById(uid string, user *casdoorsdk.User) (res bool, err error) {
	res, err = casdoorsdk.UpdateUserById(uid, user)
	return
}

func SetToken(ctx *yap.Context) (err error) {
	code := ctx.URL.Query().Get("code")
	state := ctx.URL.Query().Get("state")
	token, err := casdoorsdk.GetOAuthToken(code, state)
	if err != nil {
		return err
	}
	cookie := http.Cookie{
		Name:   "token",
		Value:  token.AccessToken,
		Path:   "/",
		MaxAge: 3600,
	}
	http.SetCookie(ctx.ResponseWriter, &cookie)
	return err
}

func GetToken(ctx *yap.Context) (token *http.Cookie, err error) {
	tokenCookie, err := ctx.Request.Cookie("token")
	if err != nil {
		return tokenCookie, err
	}
	return tokenCookie, err
}

func RemoveToken(ctx *yap.Context) (err error) {
	tokenCookie, err := ctx.Request.Cookie("token")
	if err != nil {
		return err
	}
	// Delete token
	tokenCookie.MaxAge = -1
	http.SetCookie(ctx.ResponseWriter, tokenCookie)
	return err
}
