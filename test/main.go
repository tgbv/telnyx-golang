package main

import (
	"fmt"

	tyx "github.com/tgbv/telnyx-golang/pkg"
)

func main() {

	Tyx := tyx.Init(map[string]string{
		"v1":   "g7V8JstvT6azk5yKi7VjoA",
		"v2":   "KEY01787CBA781BAE8414560A9421E7E34D_jYcYEw29htUYXzXN1J9kWH",
		"user": "support@crypto-scam.io",
	})

	out, err := Tyx.Numbers.Lookup("+40731704463")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)

	// out, err := Tyx.Messaging.GetProfiles()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(out)

	// out, err := Tyx.Messaging.Send(map[string]interface{}{
	// 	"from":    "me",
	// 	"to":      "+40731704463",
	// 	"text":    "Hello baby!",
	// 	"profile": "4001787c-ba7a-480d-8d37-0c79fa2b52ab",
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(out)

	/*
		pseudo demo

		Telnyx := tyx.Init(map[string]string{
			"v1": "g7V8JstvT6azk5yKi7VjoA",
			"v2": "KEY01787CBA781BAE8414560A9421E7E34D_jYcYEw29htUYXzXN1J9kWH",
			"user": "support@crypto-scam.io",
		})

		Telnyx.Messaging.Send(map[string]string{
			"from": "me",
			"to":"+40731704463",
			"text": "Hello baby!",
			"profile": "4001787c-ba7a-480d-8d37-0c79fa2b52ab",
		})
		Telnyx.Messaging.Listen(func(map[string]string){
			// do stuff with received message
		})

		Telnyx.Messaging.Get("21478b39-7dd6-4caa-b36d-bf2d12807d55")
		Telnyx.Messaging.GenMDR(map[string]string{
			"start_time": "2021-03-29T00:00:00+00:00",
		})
		Telnyx.Messaging.GetMDR("mdr id")
		Telnyx.Messaging.GetMDRs()

	*/
}
