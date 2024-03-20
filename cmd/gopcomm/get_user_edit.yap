import (
	"github.com/goplus/community/internal/core"
	gopaccountsdk "github.com/liuscraft/gop-casdoor-account-sdk"
	"github.com/qiniu/x/xlog"
	"net/http"
)

var (
	user      *core.User
)

xLog := xlog.New("")
var uid string
token, err := Request.Cookie("token")
if err == nil {
	user, err = community.GetUser(token.Value)
	if err != nil {
		xLog.Error("get user error:", err)
	}
	uid = user.Id
}
userClaim, err := community.GetUserClaim(uid)
if err != nil {
	xLog.Error("get current user error:", err)
}
gac, err := gopaccountsdk.GetClient(token.Value)
if err != nil {
	xLog.Info("gac", err)
	http.Redirect(ResponseWriter, Request, "/error", http.StatusTemporaryRedirect)
}
appInfo, err := community.GetApplicationInfo()
if err != nil {
	xLog.Error("get application info error:", err)
	http.Redirect(ResponseWriter, Request, "/error", http.StatusTemporaryRedirect)
}
providers := gac.GetProviders()

yap "user_edit", {
	"User":        user,
	"CurrentUser": userClaim,
	"Application": appInfo,
	"Binds":       gac.GetProviderBindStatus(),
	"casdoorUrl":  gopaccountsdk.GetCasdoorEndPoint(),
	"providers":   providers,
}
