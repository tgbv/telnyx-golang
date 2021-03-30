package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	tyx "github.com/tgbv/telnyx-golang/pkg"
)

const TARGET_PHONE = "+40731704463"
const V1_TOKEN = "g7V8JstvT6azk5yKi7VjoA"
const V2_TOKEN = "KEY01787CBA781BAE8414560A9421E7E34D_jYcYEw29htUYXzXN1J9kWH"
const USER_EMAIL = "support@crypto-scam.io"

const MESSAGING_PROFILE = "4001787c-ba7a-480d-8d37-0c79fa2b52ab"
const MESSAGING_VERIFICATION_PROFILE = "49000178-7d07-ad7c-9ed4-3f9592e3f025"

func main() {

	// initialize telnyx
	//
	Tyx := tyx.Init(map[string]string{
		"v1":   V1_TOKEN,
		"v2":   V2_TOKEN,
		"user": USER_EMAIL,
	})

	// initialize webhoock
	//
	wh, err := Tyx.Messaging.InitWebhook(":8080")
	if err != nil {
		fmt.Println(err)
	}

	// start webhock server
	go wh.StartServer()

	// we can add callbacks to the webhock server while it's started
	// as a sidenote, webhock instance can be accessed from messaging directly
	// such as..
	Tyx.Messaging.WebHook.PushCb(func(r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)

		data := map[string]interface{}{}
		_ = json.Unmarshal(body, &data)

		fmt.Println(r.RequestURI, data)
	})

	// we can delete callbacks based on their index (the position order in which they were added)
	_, _ = wh.DelCb(0)

	// s := ""
	// fmt.Scanf("\r\n", &s)

	// lookup number
	//
	out1, err := Tyx.Numbers.Lookup(TARGET_PHONE)
	if err != nil {
		panic(err)
	}
	fmt.Println(out1)

	// retrieve messaging profiles
	//
	out2, err := Tyx.Messaging.GetProfiles()
	if err != nil {
		panic(err)
	}
	fmt.Println(out2)

	// sends a message to a number
	//
	out3, err := Tyx.Messaging.Send(map[string]interface{}{
		"from":    "me",
		"to":      TARGET_PHONE,
		"text":    "Hello baby!",
		"profile": MESSAGING_PROFILE,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out3)

	// retrieve a message status from Telnyx, based on its ID
	// this code creates a callback and pushes it into webhock callbacks queue, so when a message will be received it will get it's ID via Messaging.Get(...)
	//
	Tyx.Messaging.WebHook.PushCb(func(r *http.Request) {
		data := map[string]interface{}{}

		body, _ := ioutil.ReadAll(r.Body)
		_ = json.Unmarshal(body, &data)

		id := data["data"].(map[string]interface{})["id"].(string)

		out4, err := Tyx.Messaging.Get(id)
		if err != nil {
			panic(err)
		}
		fmt.Println("Message detains:", out4)
	})

	// start generating MDR
	//
	out5, err := Tyx.Messaging.GenMDR(map[string]string{
		"start_time": "2021-03-29T00:00:00+00:00",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("MDR generation status:", out5)

	// retrieve MDR from telnyx by ID
	//
	out6, err := Tyx.Messaging.GetMDR(out5["id"].(string))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("MDR:", out6)

	// retrieve user's MDRs
	//
	out7, err := Tyx.Messaging.GetMDRs()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("MDRs:", out7)

	// delete a MDR
	//
	out8, err := Tyx.Messaging.DelMDR(out5["id"].(string))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("MDR deletion status:", out8)

	// begin sending a verification message to a number
	//
	out9, err := Tyx.Verify.Send(map[string]string{
		"number":  TARGET_PHONE,
		"profile": MESSAGING_VERIFICATION_PROFILE,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Status:", out9)

	// check verification message code
	//
	out10, err := Tyx.Verify.Check(TARGET_PHONE, "some code")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Status:", out10)
}
