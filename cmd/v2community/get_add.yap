import (
	"github.com/goplus/community/internal/core"
)

var (
	community *core.Community
	user      *core.User
)

token, err := Request.Cookie("token")
if err == nil {
	user, _ = community.GetUser(token.Value)
}
yap "edit", {
	"User": user,
}
