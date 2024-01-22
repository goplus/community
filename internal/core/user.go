package core

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type User struct {
	ID       string
	Name     string
	Password string
	Avatar   string
}

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


func (p *Community) GetUserId(token string) (userId string, err error) {
	claim, err := casdoorsdk.ParseJwtToken(token)
	if err != nil {
		p.zlog.Error(err)
		return "", ErrNotExist
	}
	return claim.Id, nil
}