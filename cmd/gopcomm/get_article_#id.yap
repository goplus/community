import (
	c "context"
	"github.com/goplus/community/internal/core"
	"github.com/qiniu/x/xlog"
)

var (
	user      *core.User
)

// Get User Info
uid := ""
xLog := xlog.New("")
todo := c.TODO()
token, err := Request.Cookie("token")
if err == nil {
	user, err = community.GetUser(token.Value)
	if err != nil {
		xLog.Error("get user error:", err)
		return
	}
	uid = user.Id
}

id := param("id")
platform := Param("platform")
ip := community.GetClientIP(Request)
community.ArticleLView(todo, id, ip, uid, platform)
article, err := community.Article(todo, id)
if err != nil {
	xLog.Error("get article error:", err)
	return
}
likeState, err := community.ArticleLikeState(todo, uid, id)
if err != nil {
	xLog.Error("article state err:", err)
	return
}

// articleJson, _ := json.Marshal(&article)
yap "article", {
	"User":      user,
	"Article":   article,
	"LikeState": likeState,
}
