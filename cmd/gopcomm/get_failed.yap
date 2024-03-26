import (
	"github.com/goplus/community/internal/core"
	"github.com/qiniu/x/log"
)

var user *core.User

token, err := Request.Cookie("token")
if token != nil {
	user, err = community.GetUser(token.Value)
	if err != nil {
		log.Error("get user error")
		json {
			"code": 0,
			"err":  err.Error(),
		}
	}
}

yap "5xx", {
	"User": user,
}
