 # SAP IAS openid-commandline-client
It is a CLI to generate OpenID TD Token from an openID connect server with PKCE and Public Client support thus SAP IAS provides it.

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
      --ias             IAS. Should be https://<yourtenant>.accounts.ondemand.com
      --clientID        IAS client ID, set as public client.
      --clientSecret    IAS client secret. This is an optional flag.
``` 
for more details.
