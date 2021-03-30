package messaging

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tgbv/telnyx-golang/config"
)

/*
*	retrieves a message by ID
 */
func (M *Messaging) Get(id string) (map[string]interface{}, error) {

	// build request
	req, _ := http.NewRequest("GET", config.API_V2+"/messages/"+id, nil)
	req.Header.Add("Authorization", "Bearer "+M.Config.Api.V2)
	req.Header.Add("Accept", "application/json")

	// make request
	res, err := M.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// read body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// unmarshal body
	unmarshaled := map[string]interface{}{}
	err = json.Unmarshal(body, &unmarshaled)
	if err != nil {
		return nil, err
	}

	// check status code and return accordingly
	if res.StatusCode == 200 {
		return unmarshaled["data"].(map[string]interface{}), nil
	} else {
		return nil, e(res.StatusCode, body)
	}
}
