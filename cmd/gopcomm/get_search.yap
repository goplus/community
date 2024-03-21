import (
	c "context"
	"github.com/goplus/community/internal/core"
	"github.com/qiniu/x/xlog"
)

var (
	user      *core.User
)

xLog := xlog.New("")
todo := c.TODO()
searchValue := param("value")
label := param("label")
if label == "" {
	label = "article"
}

token, err := Request.Cookie("token")
if err == nil {
	user, _ = community.GetUser(token.Value)
}

articles, next, err := community.ListArticle(todo, core.MarkBegin, limitConst, searchValue, label)
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
	"Value": searchValue,
	"Next":  next,
	"Tab":   label,
}
