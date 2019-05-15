package main

import (
	"flag"
	"fmt"
	"github.com/ajanthan/cmdline-openid-client/pkg/client"
	"log"
)

func main() {
	const callbackURL = "http://localhost:8080/callback"
	flag.Usage = func() {
		fmt.Println("Usage: openid-client \n" +
			"       This is a CLI to generate OpenID TD Token from an openID connect server. Create a service provider/application in the openID connect server with call back url : " + callbackURL + " and set below flags to get an ID token\n" +
			"Flags:\n" +
			"      --authzURL        OAuth2 authorization URL. Default value is https://localhost:9443/oauth2/authorize.\n" +
			"      --tokenURL        OAuth2 token URL. Default value is https://localhost:9443/oauth2/token\n" +
			"      --clientID        OAuth2 client ID. This is a mandatory flag.\n" +
			"      --clientSecret    OAuth2 client secret. This is a mandatory flag.")
	}

	var authzEp = flag.String("authzURL", "https://localhost:9443/oauth2/authorize", "OAuth2 authorization URL")
	var tokenEp = flag.String("tokenURL", "https://localhost:9443/oauth2/token", "OAuth2 token URL")
	var clientID = flag.String("clientID", "client", "OAuth2 client ID")
	var clientSecret = flag.String("clientSecret", "clientSecret", "OAuth2 client secret")

	flag.Parse()
	if *clientID == "" {
		log.Fatal("clientID is required to run this command")
	} else if *clientSecret == "" {
		log.Fatal("clientID is required to run this command")
	}
	client.HandleOpenIDFlow(*clientID, *clientSecret, callbackURL,*authzEp,*tokenEp)
}
