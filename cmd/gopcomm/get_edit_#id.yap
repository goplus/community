import (
	c "context"
	"net/http"

	"github.com/goplus/community/internal/core"
	"github.com/qiniu/x/log"
)

var user *core.User
uid := ""
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
	uid = user.Id
}

id := param("id")
if id != "" {
	editable, err := community.CanEditable(c.TODO(), uid, id)
	if err != nil {
		log.Error("can editable error:", err)
		http.Redirect(ResponseWriter, Request, "/error", http.StatusTemporaryRedirect)
	}
	if !editable {
		log.Error("no permissions")
		http.Redirect(ResponseWriter, Request, "/error", http.StatusTemporaryRedirect)
	}
	article, err := community.Article(c.TODO(), id)
	if err != nil {
		log.Error("get article error:", err)
		http.Redirect(ResponseWriter, Request, "/error", http.StatusTemporaryRedirect)
	}
	// articleJson, err := json.Marshal(&article)
	if err != nil {
		log.Error("json marshal error:", err)
		http.Redirect(ResponseWriter, Request, "/error", http.StatusTemporaryRedirect)
	}
	yap "edit", {
		"User":    user,
		"Article": article,
	}
}
