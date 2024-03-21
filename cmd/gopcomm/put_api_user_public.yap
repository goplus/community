import (
	"github.com/goplus/community/internal/core"
	"github.com/qiniu/x/xlog"
	"strconv"
)

var (
	user      *core.User
)

xLog := xlog.New("")
token, err := Request.Cookie("token")
if err == nil {
	casdoorUser, _ := community.GetUser(token.Value)
	user = &core.User{
		Id: casdoorUser.Id,
	}
} else {
	xLog.Error("get token error:", err)
	json {
		"code":  0,
		"msg":   "get token failed",
		"users": nil,
		"next":  1,
	}
}

// Get current user auth Info by id
userAuth, err := community.GetUserAuthById(user.Id)
if err != nil {
	xLog.Error("get current user auth error:", err)
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
	xLog.Error("parse bool error:", err)
	json {
		"code": 0,
		"msg":  "parse bool failed",
	}
}

res, err := community.UpdateUserPublicAuth(uid, isPublicBool)
if err != nil || !res {
	xLog.Info(err)
	json {
		"code": 0,
		"msg":  "update failed",
	}
}

json {
	"code": 200,
}
