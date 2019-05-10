 # cmdline-openid-client
It is a CLI to generate OpenID TD Token from an openID connect server.
### How to build the project
Use the go tool chain to build the binary.
```text
go build cmd/openid-client.go
```
### How to test
```text
./openid-client -help
Usage: openid-client
       This is a CLI to generate OpenID TD Token from an openID connect server. Create a service provider/application in the openID connect server with call back url : http://localhost:8080/callback and set below flags to get an ID token
Flags:
      --authzURL        OAuth2 authorization URL. Default value is https://localhost:9443/oauth2/authorize.
      --tokenURL        OAuth2 token URL. Default value is https://localhost:9443/oauth2/token
      --clientID        OAuth2 client ID. This is a mandatory flag.
      --clientSecret    OAuth2 client secret. This is a mandatory flag.
``` 
for more details.