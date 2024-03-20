import (
	"github.com/goplus/community/internal/core"
)

var (
	user *core.User
)


token, err := Request.Cookie("token")
if err == nil {
	user, _ = community.GetUser(token.Value)
}

yap "5xx", {
	"User": user,
}
