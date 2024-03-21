import (
	"fmt"
	"net/http"
    "net/url"

	"github.com/qiniu/x/log"
)

code := URL.Query().Get("code")
state := URL.Query().Get("state")
token, err := community.GetOAuthToken(code, state)
if err != nil {
	log.Error("set token error:", err)
}
cookie := http.Cookie{
	Name:     "token",
	Value:    token.AccessToken,
	Path:     "/",
	MaxAge:   168 * 60 * 60,
	HttpOnly: true,
}
http.SetCookie(ResponseWriter, &cookie)

origin_path := URL.Query().Get("origin_path")
unurl, err := url.QueryUnescape(origin_path)
if err != nil {
	log.Info("Unurl error", err)
	unurl = "/"
}

http.Redirect(ResponseWriter, Request, unurl, http.StatusFound)
