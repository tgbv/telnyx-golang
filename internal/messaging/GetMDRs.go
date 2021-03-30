package messaging

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tgbv/telnyx-golang/config"
)

/*
*	retrieve all MDRs
 */
func (M *Messaging) GetMDRs() ([]map[string]interface{}, error) {
	// build request
	req, _ := http.NewRequest("GET", config.API_V1+"/reporting/mdr_requests/", nil)
	req.Header.Add("x-api-token", M.Config.Api.V1)
	req.Header.Add("x-api-user", M.Config.Api.User)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")

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
	unmarshaled := make([]map[string]interface{}, 0)
	err = json.Unmarshal(body, &unmarshaled)
	if err != nil {
		return nil, err
	}

	// check status code and return accordingly
	if res.StatusCode == 200 {
		return unmarshaled, nil
	} else {
		return nil, e(res.StatusCode, body)
	}
}
