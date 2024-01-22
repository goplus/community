package core

import (
	"os"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

// Deprecated: use casdoorsdk instead
type User struct {
	ID       string
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

// Deprecated: use casdoorsdk instead
// GetUser return author
func (p *Community) GetUser(token string) (user *User, err error) {
	claim, err := casdoorsdk.ParseJwtToken(token)
	if err != nil {
		p.zlog.Error(err)
		return  &User{}, ErrNotExist
	}
	user = &User{
		Name:claim.Name,
		Avatar:claim.Avatar,
		ID:claim.Id,
	}
	return
}


// GetUserId return user id by token
func (p *Community) GetUserId(token string) (userId string, err error) {
	claim, err := casdoorsdk.ParseJwtToken(token)
	if err != nil {
		p.zlog.Error(err)
		return "", ErrNotExist
	}
	return claim.Id, nil
}

func (p *Community) GetUserClaim(token string) (claim *casdoorsdk.Claims, err error) {
	claim, err = casdoorsdk.ParseJwtToken(token)
	if err != nil {
		p.zlog.Error(err)
		return nil, ErrNotExist
	}

	return claim, nil
}
