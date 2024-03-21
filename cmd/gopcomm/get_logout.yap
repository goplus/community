import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/qiniu/x/log"
)


tokenCookie, err := Request.Cookie("token")
if err != nil {
	log.Error("remove token error:", err)
	return
}

tokenCookie.MaxAge = -1
http.SetCookie(ResponseWriter, tokenCookie)

refererURL, err := url.Parse(Request.Referer())
if err != nil {
	log.Infof("Error parsing Referer: %#v", err)
	return
}

refererPath := refererURL.Path
if refererURL.RawQuery != "" {
	refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
}

http.Redirect(ResponseWriter, Request, refererPath, http.StatusFound)
