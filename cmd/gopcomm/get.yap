import (
	c "context"

	"github.com/goplus/community/internal/core"
	"github.com/qiniu/x/log"
)

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

// Get Article Info
articles, next, err := community.ListArticle(c.TODO(), core.MarkBegin, limitConst, "", labelConst)
if err != nil {
	log.Error("get article error:", err)
}

yap "home", {
	"User":  user,
	"Items": articles,
	"Next":  next,
}
