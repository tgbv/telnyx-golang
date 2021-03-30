package messaging

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tgbv/telnyx-golang/config"
)

/*
*	starts generating a MDR
 */
func (M *Messaging) GenMDR(param map[string]string) (map[string]interface{}, error) {
	// make post body json
	postBody, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	// build request
	req, _ := http.NewRequest("POST", config.API_V1+"/reporting/mdr_requests", bytes.NewBuffer(postBody))
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
	unmarshaled := map[string]interface{}{}
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
