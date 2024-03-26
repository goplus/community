import (
	gopaccountsdk "github.com/liuscraft/gop-casdoor-account-sdk"
)

pv := param("pv")
token, err := Request.Cookie("token")
if err != nil {
	json {
		"code": 200,
	}
	return
}
switch pv {
case "Twitter":
case "Facebook":
case "Github":
case "WeChat":
default:
	pv = ""
}
gac, err := gopaccountsdk.GetClient(token.Value)
if err == nil {
	gac.UnLink(pv)
}
json {
	"code": 200,
}
