import (
	c "context"

	"github.com/goplus/community/internal/core"
	"github.com/qiniu/x/log"
)

var user *core.User

id := param("id")
userClaim, err := community.GetUserClaim(id)
if err != nil {
	log.Error("get current user error:", err)
	json {
		"code": 0,
		"err":  err.Error(),
	}
}

token, err := Request.Cookie("token")
if token != nil {
	user, err = community.GetUser(token.Value)
	if err != nil {
		log.Error("get user error:", err)
		json {
			"code":  0,
			"msg":   "get user failed",
			"users": nil,
			"next":  1,
		}
	}
}

// get article list published by uid
items, total, err := community.GetArticlesByUid(c.TODO(), id, firstConst, mediaLimitConst)
if err != nil {
	log.Error("get article list error:", err)
}

isAdmin, err := community.IsAdmin(id)
if err != nil {
	log.Error("check admin error:", err)
}
yap "user", {
	"Id":          id,
	"CurrentUser": userClaim,
	"User":        user,
	"Items":       items,
	"Total":       total,
	"IsAdmin":     isAdmin,
}
