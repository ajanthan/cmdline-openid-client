package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"
)

type httpServer struct {
	server         *http.Server
	code           string
	shutdownSignal chan string
}

func (h *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	if code != "" {
		h.code = code
		fmt.Fprintln(w, "Login is sucessfull, You may close the browser and goto commandline")
	} else {
		fmt.Fprintln(w, "Login is not sucessfull, You may close the browser and try again")
	}
	h.shutdownSignal <- "shutdown"
}

func main() {
	const callbackURL = "http://localhost:8080/callback"
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

	httpServer := &httpServer{}
	httpServer.shutdownSignal = make(chan string)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	httpServer.server = s
	http.Handle("/callback", httpServer)
	authzURL, authzURLParseError := url.Parse(*authzEp)

	if authzURLParseError != nil {
		log.Fatal(authzURLParseError)
	}
	query := authzURL.Query()
	query.Set("response_type", "code")
	query.Set("scope", "openid")
	query.Set("client_id", *clientID)
	query.Set("redirect_uri", callbackURL)
	authzURL.RawQuery = query.Encode()

	cmd := exec.Command("open", authzURL.String())
	cmdErorr := cmd.Start()
	if cmdErorr != nil {
		log.Fatal(authzURLParseError)
	}

	go func() {
		s.ListenAndServe()
	}()

	<-httpServer.shutdownSignal
	httpServer.server.Shutdown(context.Background())
	log.Println("Authorization code is ", httpServer.code)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	vals := url.Values{}
	vals.Set("grant_type", "authorization_code")
	vals.Set("code", httpServer.code)
	vals.Set("redirect_uri", callbackURL)
	req, requestError := http.NewRequest("POST", *tokenEp, strings.NewReader(vals.Encode()))
	if requestError != nil {
		log.Fatal(requestError)
	}
	req.SetBasicAuth(*clientID, *clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, clientError := client.Do(req)
	if clientError != nil {
		log.Fatal(clientError)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	if result != nil {
		jsonStr, marshalError := json.Marshal(result)
		if marshalError != nil {
			log.Fatal(marshalError)
		}
		log.Println(string(jsonStr))
	} else {
		log.Println("Error while getting ID token")
	}

}
