import (
	c "context"

	"github.com/goplus/community/internal/core"
	"github.com/qiniu/x/log"
)

searchValue := param("value")
label := param("label")
if label == "" {
	label = "article"
}
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

articles, next, err := community.ListArticle(c.TODO(), core.MarkBegin, limitConst, searchValue, label)
if err != nil {
	log.Error("get article error:", err)
}

if err != nil {
	log.Error("json marshal error:", err)
}
yap "home", {
	"User":  user,
	"Items": articles,
	"Value": searchValue,
	"Next":  next,
	"Tab":   label,
}
