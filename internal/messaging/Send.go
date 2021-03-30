package messaging

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tgbv/telnyx-golang/config"
)

/*
*	send a message to random number
 */
func (M *Messaging) Send(param map[string]interface{}) (map[string]interface{}, error) {
	// make post body json
	postBody, err := json.Marshal(map[string]string{
		"from":                 param["from"].(string),
		"to":                   param["to"].(string),
		"text":                 param["text"].(string),
		"messaging_profile_id": param["profile"].(string),
	})
	if err != nil {
		return nil, err
	}

	// create request
	req, _ := http.NewRequest("POST", config.API_V2+"/messages", bytes.NewBuffer(postBody))
	req.Header.Add("Authorization", "Bearer "+M.Config.Api.V2)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")

	// make request
	res, err := M.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// read json body
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

	// if status code 200
	if res.StatusCode == 200 {
		return unmarshaled["data"].(map[string]interface{}), nil
	} else
	// otherwise some shit happened
	{
		return nil, e(res.StatusCode, body)
	}
}
