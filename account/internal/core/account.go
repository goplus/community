package core

import (
	"os"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type CasdoorConfig struct {
	endPoint         string
	clientId         string
	clientSecret     string
	certificate      string
	organizationName string
	applicationName  string
}

func init() {
	conf := new(CasdoorConfig)

	endPoint := conf.endPoint
	clientID := conf.clientId
	clientSecret := conf.clientSecret
	certificate := conf.certificate
	organizationName := conf.organizationName
	applicationName := conf.applicationName
	//TODO Post insertion of environmental variables
	if endPoint == "" {
		endPoint = os.Getenv("GOP_CASDOOR_ENDPOINT")
	}
	if clientID == "" {
		clientID = os.Getenv("GOP_CASDOOR_CLIENTID")
	}
	if clientSecret == "" {
		clientSecret = os.Getenv("GOP_CASDOOR_CLIENTSECRET")
	}
	if certificate == "" {
		certificate = os.Getenv("GOP_CASDOOR_CERTIFICATE")
	}
	if organizationName == "" {
		organizationName = os.Getenv("GOP_CASDOOR_ORGANIZATIONNAME")
	}
	if applicationName == "" {
		applicationName = os.Getenv("GOP_CASDOOR_APPLICATONNAME")
	}
	casdoorsdk.InitConfig(endPoint, clientID, clientSecret, certificate, organizationName, applicationName)
}
