# telnyx-golang
Wrapper for Telnyx API written in Go

**TIP**: use CTRL+F to lookup faster

## Features
- [x] Send messages
- [x] Receive messages
- [x] Request MDR generation
- [x] Get MDR info by ID
- [x] Delete MDR by ID
- [x] Any number information lookup
- [x] Code number verification (see https://portal.telnyx.com/#/app/verify/profiles)

## How to include

```
go get github.com/tgbv/telnyx-golang@d954fc469e692c1618c0c17609346d7c5b1d269c
```

### Telnyx initialization
```
import tyx "github.com/tgbv/telnyx-golang/pkg"

// Telnyx messaging verification profile
// get from https://portal.telnyx.com/#/app/verify/profiles
// const MESSAGING_VERIFICATION_PROFILE = ""

Telnyx := tyx.init(map[string]interface{}{
  // Telnyx V1 token.
  // see https://portal.telnyx.com/#/app/api-tokens
  "v1":   "V1_TOKEN",
  
  // Telnyx V2 token.
  // see https://portal.telnyx.com/#/app/api-keys
  "v2":   "V2_TOKEN",
  
  // Telnyx user account email.
  // get from https://portal.telnyx.com/#/app/account/general
  "user": "USER_EMAIL",
})
```
### Sending a message 

```
res, _ := Telnyx.Messaging.send(map[string]interface{}{

  // max 13 characters, non numeric 
  // OR the phone number assigned to the messaging profile (see below)
  "from": "me",

  // in international format
  "to": "+1234567890",

  // max 160 characters, GSM-7 encoded 
  "text": "Hello dear!",
  
  // can be any profile, as long as you own it and it's activated
  // see https://portal.telnyx.com/#/app/messaging
  "profile": "MESSAGING_PROFILE",
})


fmt.Println(reflect.Type(res), res)
// outputs map[string]interface{}, Telnyx API response
```
### Receiving a message
1. In order to receive messages you need to setup a webhook server
2. Update webhook server address (host/port) in Telnyx portal: https://portal.telnyx.com/#/app/messaging
3. In case your machine is behind a router/proxy you must have your server listening port forwarded. Try https://canyouseeme.org/ for a fast check

```
// initialize webhoock server with a host and a port
// the following listens on any incoming address on port 8080
wh, _ := Telnyx.Messaging.InitWebhook(":8080")
```
You can then bind callbacks to the webhook server instance. They will all be called once Telnyx contacts your server in case it receives a message on your behalf
```
wh.PushCb(func(r *http.Request) {
  body, _ := ioutil.ReadAll(r.Body)

  data := map[string]interface{}{}
  _ = json.Unmarshal(body, &data)

  fmt.Println(r.RequestURI, data)
  // outputs request URI and response data from Telnyx
})
```

Then start the server
```
err := wh.StartServer()
```
You can dynamically push callbacks while webhook server is started as well. Action is mutex protected so don't worry about data racing.
