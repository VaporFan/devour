/*
 * mtStats Devour - Web API Package - Match History by SeqNum
 */

package webapi

import (
	"encoding/json"
	"fmt"
	"github.com/mtstats/devour/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Get Match History by Sequence Number
func (api *WebApi) GetMatchHistorySeq(startFrom *int) (matches models.Matches, err error) {

	// Prepare API Endpoint URL
	params := url.Values{}
	params.Add("key", api.key)
	params.Add("start_at_match_seq_num", strconv.Itoa(*startFrom+1))

	apiurl := api.urls.MatchHistorySequence + "?" + params.Encode()

	response, err := http.Get(apiurl)

	if err != nil || response.StatusCode != 200 {

		// Handle Error
		fmt.Printf("HTTP Error: %s\n", err)

		return matches, err

	} else {

		// Success Get Response Content
		content, err := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			fmt.Println(err)
			fmt.Printf("JSON: ", response.Body)
		}

		// Decode JSON into Struct
		var data models.MatchData
		err = json.Unmarshal(content, &data)
		if err != nil {
			fmt.Printf("Unmarshall Error: %s", err)
		}

		matches = data.Result.Matches

		return matches, err

	}

}
