import (
	"strconv"

	"github.com/qiniu/x/log"
	"github.com/goplus/community/internal/core"
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
		"msg":   "get token failed",
		"users": nil,
		"next":  1,
	}
}
user := &core.User{
	Id: casdoorUser.Id,
}
from := Param("from")
limit := Param("limit")
fromInt, err := strconv.Atoi(from)
if err != nil {
	fromInt = 1
}
limitInt, err := strconv.Atoi(limit)
if err != nil {
	limitInt = limitConst
}

// Get current user auth Info by id
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

users, next, err := community.ListPageUsers(fromInt, limitInt)
if err != nil {
	log.Error("get users error:", err)
	json {
		"code":  0,
		"msg":   "get users failed",
		"items": nil,
		"total": 10,
		"next":  1,
	}
}

json {
	"code":  200,
	"msg":   "ok",
	"items": users,
	"total": len(users),
	"next":  next,
}
