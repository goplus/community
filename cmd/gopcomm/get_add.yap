import (
	"github.com/qiniu/x/log"
	"github.com/goplus/community/internal/core"
)

var user *core.User
token, err := Request.Cookie("token")
if token!=nil{
    user, err = community.GetUser(token.Value)
    if err != nil {
    	log.Error("get user error")
    	json {
    		"code": 0,
    		"err":  err.Error(),
    	}
    }
}
yap "edit", {
	"User": user,
}
