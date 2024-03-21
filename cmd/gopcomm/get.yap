import (
	c "context"
	"github.com/goplus/community/internal/core"
	"github.com/qiniu/x/xlog"
)

const (
	limitConst = 10
	labelConst = "article"
)

xLog := xlog.New("")
todo := c.TODO()
var user *core.User
token, err := Request.Cookie("token")
if err == nil {
	user, _ = community.GetUser(token.Value)
}

// Get Article Info
articles, next, err := community.ListArticle(todo, core.MarkBegin, limitConst, "", labelConst)
if err != nil {
	xLog.Error("get article error:", err)
}

// articlesJson, err := json.Marshal(&articles)
if err != nil {
	xLog.Error("json marshal error:", err)
}
yap "home", {
	"User":  user,
	"Items": articles,
	"Next":  next,
}
