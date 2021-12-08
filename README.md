 # SAP IAS openid-commandline-client
It is a CLI to generate OpenID TD Token from an openID connect server, mainly created to test PKCE and Public Client support, like SAP IAS provides it. However any other OIDC provider can be used to get tokens.

### How to build the project

Use the go tool chain to build the binary.
```text
go build cmd/openid-client.go
```
### How to test
```text
./openid-client -help
Usage: openid-client
       This is a CLI to generate OpenID TD Token from an openID connect server. Create a service provider/application in the openID connect server with call back url : http://localhost:7000/callback and set below flags to get an ID token
Flags:
			   -issuer          IAS. Default is https://<yourtenant>.accounts.ondemand.com; XSUAA Default is: https://uaa.cf.eu10.hana.ondemand.com/oauth/token
			   -client_id       OIDC client ID. This is a mandatory flag.
			   -client_secret   OIDC client secret. This is an optional flag and only needed for confidential clients.
			   -scope           OIDC scope parameter. This is an optional flag, default is openid. If you set none the parameter scope will be omitted in request.
			   -refresh         Bool flag. Default false. If true, call refresh flow for the received id_token.
			   -idp_token       Bool flag. Default false. If true, call the OIDC IdP token exchange endpoint (IAS specific only) and return the response.
			   -h               Show this help
``` 
for more details.
