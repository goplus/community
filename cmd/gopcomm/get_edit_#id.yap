import (
	c "context"
	"github.com/goplus/community/internal/core"
	"github.com/qiniu/x/xlog"
	"net/http"
)

var (
	user      *core.User
)

xLog := xlog.New("")
todo := c.TODO()
token, err := Request.Cookie("token")
if err == nil {
	user, _ = community.GetUser(token.Value)
}

uid := user.Id
id := param("id")
if id != "" {
	editable, err := community.CanEditable(todo, uid, id)
	if err != nil {
		xLog.Error("can editable error:", err)
		http.Redirect(ResponseWriter, Request, "/error", http.StatusTemporaryRedirect)
	}
	if !editable {
		xLog.Error("no permissions")
		http.Redirect(ResponseWriter, Request, "/error", http.StatusTemporaryRedirect)
	}
	article, err := community.Article(todo, id)
	if err != nil {
		xLog.Error("get article error:", err)
		http.Redirect(ResponseWriter, Request, "/error", http.StatusTemporaryRedirect)
	}
	// articleJson, err := json.Marshal(&article)
	if err != nil {
		xLog.Error("json marshal error:", err)
		http.Redirect(ResponseWriter, Request, "/error", http.StatusTemporaryRedirect)
	}
	yap "edit", {
		"User":    user,
		"Article": article,
	}
}
