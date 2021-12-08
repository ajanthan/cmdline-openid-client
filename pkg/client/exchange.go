package client

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	oidc "github.com/coreos/go-oidc"
)

func HandleCorpIdpExchangeFlow(clientID string, clientSecret string, existingIdToken string, provider oidc.Provider) string {

	params := url.Values{}
	params.Add("assertion", existingIdToken)
	params.Add("response_type", `token id_token`)
	params.Add("client_id", clientID)

	body := strings.NewReader(params.Encode())

	tokenEndPoint := strings.Replace(provider.Endpoint().TokenURL, "/token", "/exchange/corporateidp", 1)
	log.Println("Call IdP Token Exchange Endpoint: " + tokenEndPoint)
	req, err := http.NewRequest("POST", tokenEndPoint, body)
	if err != nil {
		return ""
	}
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	if resp.StatusCode == http.StatusOK {
		return bodyString
	} else {
		log.Fatal("Error from token exchange")
		log.Fatal(bodyString)
		return ""
	}
}
