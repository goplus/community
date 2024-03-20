import (
	gopaccountsdk "github.com/liuscraft/gop-casdoor-account-sdk"
)

token, err := Request.Cookie("token")
if err != nil {
	json {
		"code": 500,
	}
	return
}
gac, err := gopaccountsdk.GetClient(token.Value)
if err != nil {
	json {
		"code": 500,
	}
	return
}
qrData, ticket, err := gac.GetProviderWeChatQRCode("provider_wechat")
if err != nil {
	json {
		"code": 500,
	}
	return
}
json {
	"data":  qrData,
	"data2": ticket,
	"code":  200,
}
