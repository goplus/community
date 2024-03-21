import (
	"strconv"

	"github.com/goplus/community/internal/core"
	"github.com/qiniu/x/log"
)

token, err := Request.Cookie("token")
if err != nil {
	log.Error("get token error:", err)
	json {
		"code":  0,
		"msg":   "get token failed",
		"users": nil,
		"next":  1,
	}
}
casdoorUser, e := community.GetUser(token.Value)
if e != nil {
	log.Error("get user error:", err)
	json {
		"code":  0,
		"msg":   "get user failed",
		"users": nil,
		"next":  1,
	}
}
user := &core.User{
	Id: casdoorUser.Id,
}
userAuth, err := community.GetUserAuthById(user.Id)
if err != nil {
	log.Error("get current user auth error:", err)
}

if !userAuth.Status {
	json {
		"code":  0,
		"msg":   "auth failed",
		"items": nil,
		"total": 10,
		"next":  1,
	}
}

uid := Param("uid")
isPublic := Param("is_public")
isPublicBool, err := strconv.ParseBool(isPublic)
if err != nil {
	log.Error("parse bool error:", err)
	json {
		"code": 0,
		"msg":  "parse bool failed",
	}
}

res, err := community.UpdateUserPublicAuth(uid, isPublicBool)
if err != nil || !res {
	log.Info(err)
	json {
		"code": 0,
		"msg":  "update failed",
	}
}

json {
	"code": 200,
}
