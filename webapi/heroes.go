/*
 * mtStats Devour - Web API Package - Heroes
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
func (api *WebApi) GetHeroes(lang string) (heroes models.Heroes, err error) {

	// Prepare API Endpoint URL
	params := url.Values{}
	params.Add("key", api.key)
	params.Add("language", lang)

	apiurl := api.urls.Heroes + "?" + params.Encode()

	response, err := http.Get(apiurl)

	if err != nil || response.StatusCode != 200 {

		// Handle Error
		fmt.Printf("HTTP Error: %s", err)

		return heroes, err

	} else {

		// Success Get Response Content
		content, err := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			fmt.Printf("%s", err)
			fmt.Printf("JSON: ", response.Body)
		}

		// Decode JSON into Struct
		var data models.HeroesData
		err = json.Unmarshal(content, &data)
		if err != nil {
			fmt.Printf("Unmarshall Error: %s", err)
		}

		heroes = data.Result.Heroes

		return heroes, err

	}

}
