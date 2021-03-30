package verify

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tgbv/telnyx-golang/config"
)

/*
*	Sends a verification code to a phone number, code which later can be verified using Check() method
 */
func (V *Verify) Send(param map[string]string) (map[string]interface{}, error) {
	// marshal params
	reqBody, err := json.Marshal(map[string]string{
		"phone_number":      param["number"],
		"verify_profile_id": param["profile"],
		"type":              "sms",
	})
	if err != nil {
		return nil, err
	}

	// create request
	req, _ := http.NewRequest("POST", config.API_V2+"/verifications", bytes.NewBuffer(reqBody))
	req.Header.Add("Authorization", "Bearer "+V.Config.Api.V2)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")

	// make request
	res, err := V.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// read body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// unserialize json body
	unmarshaled := map[string]interface{}{}
	err = json.Unmarshal(body, &unmarshaled)
	if err != nil {
		return nil, err
	}

	// for correct status code
	if res.StatusCode == 200 {
		return unmarshaled["data"].(map[string]interface{}), err
	} else

	// for incorrect status code
	{
		return nil, e(res.StatusCode, body)
	}
}
