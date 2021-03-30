package messaging

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tgbv/telnyx-golang/config"
)

/*
*	retrieves the messaging profiles
 */
func (M *Messaging) GetProfiles() ([]map[string]interface{}, error) {
	// create request
	req, _ := http.NewRequest("GET", config.API_V1+"/messaging/profiles", nil)
	req.Header.Add("x-api-token", M.Config.Api.V1)
	req.Header.Add("x-api-user", M.Config.Api.User)
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

	// unserialize json body
	unmarshaled := make([]map[string]interface{}, 0)
	err = json.Unmarshal(body, &unmarshaled)
	if err != nil {
		return nil, err
	}

	// if status code 200
	if res.StatusCode == 200 {
		return unmarshaled, err
	} else
	// otherwise some shit happened
	{
		return nil, e(res.StatusCode, body)
	}

}
