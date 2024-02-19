package core

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

// get community application information
func (a *Community) GetApplicationInfo() (*casdoorsdk.Application, error) {
	a2, err := casdoorsdk.GetApplication("application_x8aevk")
	if err != nil {
		log.Println(err)
	}
	return a2, err
}

func (p *Community) GetAccountBinds(token string) map[string]bool {
	claim, err := casdoorsdk.ParseJwtToken(token)
	if err != nil {
		p.xLog.Error(err)
		return map[string]bool{}
	}
	u, err := casdoorsdk.GetUser(claim.Name)
	if err != nil {
		return map[string]bool{}
	}

	return map[string]bool{
		"twitter":  len(u.Twitter) > 0,
		"github":   len(u.GitHub) > 0,
		"facebook": len(u.Facebook) > 0,
		"wechat":   len(u.WeChat) > 0,
	}
}

func (p *Community) UnLink(token, provider string) bool {
	claim, err := casdoorsdk.ParseJwtToken(token)
	if err != nil {
		p.xLog.Error(err)
		return false
	}
	u, err := casdoorsdk.GetUser(claim.Name)
	if err != nil {
		log.Println(err)
		return false
	}
	postData, _ := json.Marshal(map[string]interface{}{
		"providerType": provider,
		"user":         u,
	})

	var response casdoorsdk.Response
	targetUrl := casdoorsdk.GetUrl("unlink", map[string]string{
		"access_token": token,
	})
	req, _ := http.NewRequest("POST", targetUrl, bytes.NewReader(postData))

	req.Header.Add("Content-Type", "application/json")
	resultBody, _ := http.DefaultClient.Do(req)
	body, err := io.ReadAll(resultBody.Body)
	if err != nil {
		return false
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false
	}

	if response.Status != "ok" {
		log.Println(response.Msg)
		return false
	}
	return true
}
