/*
 * mtStats Devour - Web API Package - Game Items
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

// Get all Game Items
func (api *WebApi) GetGameItems(lang string) (items models.GameItems, err error) {

	// Prepare API Endpoint URL
	params := url.Values{}
	params.Add("key", api.key)
	params.Add("language", lang)

	apiurl := api.urls.GameItems + "?" + params.Encode()

	response, err := http.Get(apiurl)

	if err != nil || response.StatusCode != 200 {

		// Handle Error
		fmt.Printf("HTTP Error: %s", err)

		return items, err

	} else {

		// Success Get Response Content
		content, err := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			fmt.Printf("%s", err)
			fmt.Printf("JSON: ", response.Body)
		}

		// Decode JSON into Struct
		var data models.GameItemsData
		err = json.Unmarshal(content, &data)
		if err != nil {
			fmt.Printf("Unmarshall Error: %s", err)
		}

		items = data.Result.Items

		return items, err

	}

}
