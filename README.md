# telnyx-golang
Wrapper for Telnyx API written in Go

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
```go
import tyx "github.com/tgbv/telnyx-golang/pkg"

Telnyx := tyx.init(map[string]interface{}{
  // Telnyx V1 token
  // See https://portal.telnyx.com/#/app/api-tokens
  "v1":   "V1_TOKEN",
  
  // Telnyx V2 token
  // See https://portal.telnyx.com/#/app/api-keys
  "v2":   "V2_TOKEN",
  
  // Telnyx user account email
  // Get from https://portal.telnyx.com/#/app/account/general
  "user": "USER_EMAIL",
})
```
### Sending a message 

```go
res, _ := Telnyx.Messaging.send(map[string]interface{}{

  // Max 13 characters, non numeric 
  // OR the phone number assigned to the messaging profile (see below)
  "from": "me",

  // In international format
  "to": "+1234567890",

  // Max 160 characters, preferably GSM-7 encoded 
  "text": "Hello dear!",
  
  // Can be any profile, as long as you own it and it's activated
  // See https://portal.telnyx.com/#/app/messaging
  "profile": "MESSAGING_PROFILE",
})


fmt.Println(reflect.Type(res), res)
// outputs map[string]interface{}, Telnyx API response
```
### Receiving a message
1. In order to receive messages you need to setup a webhook server
2. Update webhook server address (host/port) in Telnyx portal: https://portal.telnyx.com/#/app/messaging
3. In case your machine is behind a router/proxy you must have your server listening port forwarded. Try https://canyouseeme.org/ for a fast check

```go
// Initialize webhoock server with a host and a port
wh, _ := Telnyx.Messaging.InitWebhook(":8080")
```
You can then bind callbacks to the webhook server instance. They will all be called once Telnyx contacts your server in case it receives a message on your behalf
```go
wh.PushCb(func(r *http.Request) {
  body, _ := ioutil.ReadAll(r.Body)

  data := map[string]interface{}{}
  _ = json.Unmarshal(body, &data)

  fmt.Println(r.RequestURI, data)
  // outputs request URI and response data from Telnyx
})
```

Then start the server
```go
err := wh.StartServer()
```
You can dynamically push callbacks while webhook server is started as well. Action is mutex protected so don't worry about data racing.

### Number verification
A beta feature of Telnyx. You can set it up here: https://portal.telnyx.com/#/app/verify/profiles

In order to send a message verification code to a number you can use this:
```go
out, err := Tyx.Verify.Send(map[string]string{
  "number":  "+1234567890",
  "profile": "MESSAGING_VERIFICATION_PROFILE",
})
```

In order to check a verification code sent to a certain number:
```go
out, err := Tyx.Verify.Check("+1234567890", "some code")
```

### API Reference
--------------
- #### Query number information
```go
out, err := Tyx.Numbers.Lookup("+1234567890")
```

- #### Retrieve messaging profiles
```go
out, err := Tyx.Messaging.GetProfiles()
```

- #### Request MDR generation
```go
out, err := Tyx.Messaging.GenMDR(map[string]string{
  "start_time": "2021-03-29T00:00:00+00:00",
})
```

- #### Get MDR data by its ID
```go
out, err := Tyx.Messaging.GetMDR("xxxx-xxxx....")
```

- #### Delete MDR by ID
```go
out, err := Tyx.Messaging.DelMDR("xxxx-xxxx....")
```
