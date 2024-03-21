import (
	"fmt"
	"github.com/qiniu/x/xlog"
	"net/http"
	"net/url"
)

xLog := xlog.New("")
tokenCookie, err := Request.Cookie("token")
if err != nil {
	xLog.Error("remove token error:", err)
	return
}

// Delete token
tokenCookie.MaxAge = -1
http.SetCookie(ResponseWriter, tokenCookie)

refererURL, err := url.Parse(Request.Referer())
if err != nil {
	xLog.Infof("Error parsing Referer: %#v", err)
	return
}

refererPath := refererURL.Path
if refererURL.RawQuery != "" {
	refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
}

http.Redirect(ResponseWriter, Request, refererPath, http.StatusFound)
