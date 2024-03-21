import (
	"fmt"
	"github.com/qiniu/x/xlog"
	"net/http"
	"net/url"
)

xLog := xlog.New("")
code := URL.Query().Get("code")
state := URL.Query().Get("state")
token, err := community.GetOAuthToken(code, state)
if err != nil {
	xLog.Error("set token error:", err)
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
	xLog.Info("Unurl error", err)
	unurl = "/"
}

http.Redirect(ResponseWriter, Request, unurl, http.StatusFound)
