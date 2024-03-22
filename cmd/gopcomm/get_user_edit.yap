import (
	"net/http"

	"github.com/goplus/community/internal/core"
	gopaccountsdk "github.com/liuscraft/gop-casdoor-account-sdk"
	"github.com/qiniu/x/log"
)

var user *core.User
var uid string

token, err := Request.Cookie("token")
if token != nil {
	user, err = community.GetUser(token.Value)
	if err != nil {
		log.Error("get user error:", err)
		json {
			"code":  0,
			"msg":   "get user failed",
			"users": nil,
			"next":  1,
		}
	}
	uid = user.Id
}
userClaim, err := community.GetUserClaim(uid)
if err != nil {
	log.Error("get current user error:", err)
}
gac, err := gopaccountsdk.GetClient(token.Value)
if err != nil {
	log.Info("gac", err)
	http.Redirect(ResponseWriter, Request, "/error", http.StatusTemporaryRedirect)
}
appInfo, err := community.GetApplicationInfo()
if err != nil {
	log.Error("get application info error:", err)
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
