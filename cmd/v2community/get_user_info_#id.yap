import (
	c "context"
	"github.com/goplus/community/internal/core"
	"github.com/qiniu/x/xlog"
)

const (
	mediaLimitConst = 8
	firstConst      = "1"
)

var (
	user      *core.User
)

xLog := xlog.New("")
todo := c.TODO()
id := param("id")

// Get current User Info by id
userClaim, err := community.GetUserClaim(id)
if err != nil {
	xLog.Error("get current user error:", err)
}

// get user by token
token, err := Request.Cookie("token")
if err == nil {
	user, _ = community.GetUser(token.Value)
}

// get article list published by uid
items, total, err := community.GetArticlesByUid(todo, id, firstConst, mediaLimitConst)
if err != nil {
	xLog.Error("get article list error:", err)
}

// check admin
isAdmin, err := community.IsAdmin(id)
if err != nil {
	xLog.Error("check admin error:", err)
}
yap "user", {
	"Id":          id,
	"CurrentUser": userClaim,
	"User":        user,
	"Items":       items,
	"Total":       total,
	"IsAdmin":     isAdmin,
}
