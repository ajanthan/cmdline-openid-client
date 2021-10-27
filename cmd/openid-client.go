package main

import (
	"flag"
	"fmt"
	"log"

	oidc "github.com/coreos/go-oidc"
	"github.com/strehle/cmdline-openid-client/pkg/client"
	"golang.org/x/net/context"
)

func main() {
	const callbackURL = "http://localhost:7000/callback"
	flag.Usage = func() {
		fmt.Println("Usage: openid-client \n" +
			"       This is a CLI to generate OpenID TD Token from an openID connect server. Create a service provider/application in the openID connect server with call back url : " + callbackURL + " and set below flags to get an ID token\n" +
			"Flags:\n" +
			"      -issuer           IAS. Default is https://<yourtenant>.accounts.ondemand.com; XSUAA Default is: https://uaa.cf.eu10.hana.ondemand.com/oauth/token\n" +
			"      -client_id        OIDC client ID. This is a mandatory flag.\n" +
			"      -client_secret    OIDC client secret. This is an optional flag and only needed for confidential clients.\n" +
			"      -refresh          Refresh the received id_token.\n" +
			"      -h                Show this help\n")
	}

	var issEndPoint = flag.String("issuer", "", "OIDC Issuer URI")
	var clientID = flag.String("client_id", "", "OIDC client ID")
	var clientSecret = flag.String("client_secret", "", "OIDC client secret")
	var doRefresh = flag.Bool("refresh", false, "Refresh the received id_token")

	flag.Parse()
	if *clientID == "" {
		log.Fatal("clientID is required to run this command")
	} else if *issEndPoint == "" {
		log.Fatal("Issuer is required to run this command")
	}
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, *issEndPoint)
	if err != nil {
		log.Fatal(err)
	}
	var refreshToken = client.HandleOpenIDFlow(*clientID, *clientSecret, callbackURL, *provider)
	if *doRefresh {
		if refreshToken == "" {
			log.Println("No refresh token received.")
			return
		}
		var newRefresh = client.HandleRefreshFlow(*clientID, *clientSecret, refreshToken, *provider)
		log.Println("Old refresh token: " + refreshToken)
		log.Println("New refresh token: " + newRefresh)
	}
}
