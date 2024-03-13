package core

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"golang.org/x/oauth2"

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

var _ CasdoorSDKService = &CasdoorSDKServiceAdapter{}

// CasdoorSDKService is the interface for casdoor sdk and easy to mock
type CasdoorSDKService interface {
	GetUser(name string) (user *casdoorsdk.User, err error)
	ParseJwtToken(token string) (claims *casdoorsdk.Claims, err error)
	GetUserClaim(uid string) (claim *casdoorsdk.User, err error)
	GetUserById(uid string) (user *casdoorsdk.User, err error)
	UpdateUserById(uid string, user *UserInfo) (res bool, err error)
	UpdateUser(user *UserInfo) (res bool, err error)
	GetUserByUserId(uid string) (user *casdoorsdk.User, err error)
	GetOAuthToken(code string, state string) (token *oauth2.Token, err error)
	GetApplication(name string) (*casdoorsdk.Application, error)
}

// Adapter for mock
type CasdoorSDKServiceAdapter struct{}

func (c *CasdoorSDKServiceAdapter) GetUser(name string) (user *casdoorsdk.User, err error) {
	return casdoorsdk.GetUser(name)
}

func (c *CasdoorSDKServiceAdapter) ParseJwtToken(token string) (claims *casdoorsdk.Claims, err error) {
	return casdoorsdk.ParseJwtToken(token)
}

func (c *CasdoorSDKServiceAdapter) GetUserClaim(uid string) (claim *casdoorsdk.User, err error) {
	return casdoorsdk.GetUserByUserId(uid)
}

func (c *CasdoorSDKServiceAdapter) GetUserById(uid string) (user *casdoorsdk.User, err error) {
	return casdoorsdk.GetUserByUserId(uid)
}

func (c *CasdoorSDKServiceAdapter) UpdateUserById(uid string, user *UserInfo) (res bool, err error) {
	return casdoorsdk.UpdateUserById(uid, (*casdoorsdk.User)(user))
}

func (c *CasdoorSDKServiceAdapter) UpdateUser(user *UserInfo) (res bool, err error) {
	return casdoorsdk.UpdateUser((*casdoorsdk.User)(user))
}

func (c *CasdoorSDKServiceAdapter) GetUserByUserId(uid string) (user *casdoorsdk.User, err error) {
	return casdoorsdk.GetUserByUserId(uid)
}

func (c *CasdoorSDKServiceAdapter) GetOAuthToken(code string, state string) (token *oauth2.Token, err error) {
	return casdoorsdk.GetOAuthToken(code, state)
}

func (c *CasdoorSDKServiceAdapter) GetApplication(name string) (*casdoorsdk.Application, error) {
	return casdoorsdk.GetApplication(name)
}

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
	claim, err := p.CasdoorSDKService.ParseJwtToken(token)
	if err != nil {
		p.xLog.Error(err)
		return &User{}, ErrNotExist
	}
	user, err = p.GetUserById(claim.Id)
	userId := user.Id
	if user.Name == "wx_user_"+userId {
		rand.Seed(time.Now().UnixNano())
		name := "GOP" + strconv.Itoa(rand.Intn(100000000))
		claim.User.DisplayName = name
		_, err := casdoorsdk.UpdateUserById(claim.User.GetId(), &claim.User)
		if err != nil {
			p.xLog.Error(err.Error())
		} else {
			user.Name = name
		}
	}
	return
}

// ParseJwtToken return user id by token
func (p *Community) ParseJwtToken(token string) (userId string, err error) {
	claim, err := p.CasdoorSDKService.ParseJwtToken(token)
	if err != nil {
		p.xLog.Error("parse token err:", token)
		return "", ErrNotExist
	}
	return claim.Id, nil
}

// GetUserClaim get user（full） by token
func (p *Community) GetUserClaim(uid string) (claim *casdoorsdk.User, err error) {
	claim, err = p.CasdoorSDKService.GetUserByUserId(uid)
	if err != nil {
		p.xLog.Error(err)
		return &casdoorsdk.User{}, ErrNotExist
	}
	// claim.Properties["oauth_Wechat_extra"] = ""
	return
}

// GetUserById get user by uid
func (p *Community) GetUserById(uid string) (user *User, err error) {
	claim, err := p.CasdoorSDKService.GetUserByUserId(uid)
	if err != nil {
		p.xLog.Error(err)
		return &User{}, ErrNotExist
	}
	if claim == nil {
		user = &User{
			Name:   "unknown",
			Avatar: "",
			Id:     uid,
		}
	} else {
		user = &User{
			Name:   claim.DisplayName,
			Avatar: claim.Avatar,
			Id:     claim.Id,
		}
	}

	return
}

// UpdateUserById update user by uid
func (p *Community) UpdateUserById(uid string, user *UserInfo) (res bool, err error) {
	res, err = p.CasdoorSDKService.UpdateUserById(uid, user)
	return
}

func (p *Community) UpdateUser(user *UserInfo) (res bool, err error) {
	res, err = p.CasdoorSDKService.UpdateUser(user)
	return
}

func (p *Community) GetOAuthToken(code string, state string) (token *oauth2.Token, err error) {
	token, err = p.CasdoorSDKService.GetOAuthToken(code, state)
	return
}
