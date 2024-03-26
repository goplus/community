import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/qiniu/x/log"
)

refererURL, err := url.Parse(Request.Referer())
if err != nil {
	log.Infof("Error parsing Referer: %#v", err)
	return
}

refererPath := refererURL.Path
if refererURL.RawQuery != "" {
	refererPath = fmt.Sprintf("%s?%s", refererURL.Path, refererURL.RawQuery)
}

redirectURL := fmt.Sprintf("%s://%s/%s?origin_path=%s", refererURL.Scheme, refererURL.Host, "login/callback", url.QueryEscape(refererPath))

loginURL := community.RedirectToCasdoor(redirectURL)
Redirect loginURL, http.StatusFound
