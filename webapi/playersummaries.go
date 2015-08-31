/*
 * mtStats Devour - Web API Package - Player Summaries
 */

package webapi

import (
	"encoding/json"
	"fmt"
	"github.com/mtstats/devour/models"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Get all Heroes
func (api *WebApi) GetPlayerSummaries(ids []int) (data models.PlayerSummaryData) {

	// Prepare API Endpoint URL
	params := url.Values{}
	params.Add("key", api.key)

	apiurl := api.urls.PlayerSummaries + "?" + params.Encode()

	response, err := http.Get(apiurl)

	if err != nil || response.StatusCode != 200 {

		// Handle Error
		fmt.Printf("HTTP Error: %s", err)

		return

	} else {

		// Success Get Response Content
		content, err := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			fmt.Printf("%s", err)
			fmt.Printf("JSON: ", response.Body)
		}

		// Decode JSON into Struct
		err = json.Unmarshal(content, &data)
		if err != nil {
			fmt.Printf("Unmarshall Error: %s", err)
		}

		return data

	}

}
