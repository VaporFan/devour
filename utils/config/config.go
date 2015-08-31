/*
 * mtStats Devour - Config Package
 */

package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	AppName     string `json:"name"`
	AppURL      string `json:"url"`
	AppVersion  string `json:"version"`
	Environment string `json:"environment"`

	SteamAPI SteamAPI `json:"steam_api"`

	Database Database `json:"database"`
}

type SteamAPI struct {
	EconomySchema        string `json:"economy_schema"`
	GameItems            string `json:"game_items"`
	Heroes               string `json:"heroes"`
	ItemRarities         string `json:"item_rarities"`
	LeagueListing        string `json:"league_listing"`
	LiveLeagueGames      string `json:"live_league_games"`
	MatchDetails         string `json:"match_details"`
	MatchHistory         string `json:"match_history"`
	MatchHistorySequence string `json:"match_history_sequence"`
	PlayerSummaries      string `json:"player_summaries"`
	TeamInfo             string `json:"team_info"`
	Key                  string `json:"key"`
}

type Database struct {
	Host string `json:"host"`
	Name string `json:"name"`
	Pass string `json:"password"`
	Port string `json:"port"`
	User string `json:"username"`
}

// Load the config file
// config.json
func LoadConfig() (c Configuration) {

	f, err := os.Open("config.json")
	defer f.Close()
	decoder := json.NewDecoder(f)
	if err != nil {
		fmt.Println("Error: Could not find the config file.")
		os.Exit(1)
	}

	err = decoder.Decode(&c)
	if err != nil {
		fmt.Println("Error: Could not decode the config file.")
		os.Exit(1)
	}

	return c

}
