package main

import (
	tyx "github.com/tgbv/telnyx-golang/pkg"
)

func main() {

	Tyx := tyx.Init(map[string]string{
		"v1":   "g7V8JstvT6azk5yKi7VjoA",
		"v2":   "KEY01787CBA781BAE8414560A9421E7E34D_jYcYEw29htUYXzXN1J9kWH",
		"user": "support@crypto-scam.io",
	})

	// out, err := Tyx.Numbers.Lookup("+40731704463")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(out)

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

	// out, err := Tyx.Messaging.Get("21478b39-7dd6-4caa-b36d-bf2d12807d55")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(out)

	// out, err := Tyx.Messaging.GenMDR(map[string]string{
	// 	"start_time": "2021-03-29T00:00:00+00:00",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(out)

	// out, err := Tyx.Messaging.GetMDR("949c0a89-90bf-4b5f-824f-69b64fbd8ac7")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(out)

	// out, err := Tyx.Messaging.GetMDRs()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(out)

	// out, err := Tyx.Messaging.DelMDR("949c0a89-90bf-4b5f-824f-69b64fbd8ac7")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(out)

	// out, err := Tyx.Verify.Send(map[string]string{
	// 	"number":  "+40731704463",
	// 	"profile": "49000178-7d07-ad7c-9ed4-3f9592e3f025",
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(out)

	// out, err := Tyx.Verify.Check("+40731704463", "91553")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(out)

	/*
		pseudo demo

		Telnyx.Messaging.Listen(func(map[string]string){
			// do stuff with received message
		})



	*/
}
