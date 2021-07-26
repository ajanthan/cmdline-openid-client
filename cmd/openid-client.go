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
			"      --ias             IAS. Default is https://<yourtenant>.accounts.ondemand.com\n" +
			"      --clientID        IAS client ID. This is a mandatory flag.\n" +
			"      --clientSecret    IAS client secret. This is an optional flag.\n")
	}

	var iasEp = flag.String("ias", "", "IAS Base URL")
	var clientID = flag.String("clientID", "", "IAS client ID")
	var clientSecret = flag.String("clientSecret", "", "IAS client secret")

	flag.Parse()
	if *clientID == "" {
		log.Fatal("clientID is required to run this command")
	} else if *iasEp == "" {
		log.Fatal("IAS is required to run this command")
	}
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, *iasEp)
	if err != nil {
		log.Fatal(err)
	}
	client.HandleOpenIDFlow(*clientID, *clientSecret, callbackURL, *provider)
}
