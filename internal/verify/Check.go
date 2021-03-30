package verify

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tgbv/telnyx-golang/config"
)

/*
*	Verify a code sent using Send() method. Unlinke most methods, this one returns a bool. True when code matches, false otherwise.
 */
func (V *Verify) Check(nr string, code string) (bool, error) {
	// marshal params
	reqBody, err := json.Marshal(map[string]string{"code": code})
	if err != nil {
		return false, err
	}

	// create request
	req, _ := http.NewRequest("POST", config.API_V2+"/verifications/by_phone_number/"+nr+"/actions/verify", bytes.NewBuffer(reqBody))
	req.Header.Add("Authorization", "Bearer "+V.Config.Api.V2)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")

	// make request
	res, err := V.HttpClient.Do(req)
	if err != nil {
		return false, err
	}

	// read body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	// unserialize json body
	unmarshaled := map[string]interface{}{}
	err = json.Unmarshal(body, &unmarshaled)
	if err != nil {
		return false, err
	}

	// for ok status (200)
	if res.StatusCode == 200 {

		// if code is correct
		// disclaimer: i don't know what 'idiomatic' means
		if unmarshaled["data"].(map[string]interface{})["response_code"].(string) == "accepted" {
			return true, nil
		} else {
			return false, nil
		}

	} else

	// for incorrect status code
	{
		return false, e(res.StatusCode, body)
	}
}
