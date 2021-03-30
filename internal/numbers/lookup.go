package numbers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tgbv/telnyx-golang/config"
)

/*
*	get information of a specific number
	returns error on any possible case
*/
func (N *Numbers) Lookup(nr string) (map[string]interface{}, error) {

	// create request
	req, _ := http.NewRequest("GET", config.API_V2+"/number_lookup/"+nr, nil)
	req.Header.Add("Authorization", "Bearer "+N.Config.Api.V2)
	req.Header.Add("Accept", "application/json")

	// make request
	res, err := N.HttpClient.Do(req)
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
