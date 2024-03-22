import (
	"github.com/goplus/community/internal/core"
	"github.com/qiniu/x/log"
)

token, err := Request.Cookie("token")
uid, err := community.ParseJwtToken(token.Value)
user := &core.UserInfo{
	Id:          uid,
	Name:        param("name"),
	Birthday:    param("birthday"),
	Gender:      param("gender"),
	Phone:       param("phone"),
	Email:       param("email"),
	Avatar:      param("avatar"),
	Owner:       param("owner"),
	DisplayName: param("displayName"),
}

_, err = community.UpdateUser(user)
if err != nil {
	log.Info(err)
	json {
		"code": 0,
		"msg":  "update failed",
	}
}
json {
	"code": 200,
}
