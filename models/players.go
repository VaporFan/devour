/*
 * mtStats Devour - Models Package - Players
 */

package models

import (
//"fmt"
)

// Dota 2 Player Data
// Part of struct Match
type Player struct {
	AbilityUpgrades []AbilityUpgrade `json:"ability_upgrades"`
	AccountID       int              `json:"account_id"`
	Assists         int              `json:"assists"`
	Deaths          int              `json:"deaths"`
	Denies          int              `json:"denies"`
	Gold            int              `json:"gold"`
	GoldPerMin      int              `json:"gold_per_min"`
	GoldSpent       int              `json:"gold_spent"`
	HeroDamage      int              `json:"hero_damage"`
	HeroHealing     int              `json:"hero_healing"`
	HeroID          int              `json:"hero_id"`
	Item0           int              `json:"item_0"`
	Item1           int              `json:"item_1"`
	Item2           int              `json:"item_2"`
	Item3           int              `json:"item_3"`
	Item4           int              `json:"item_4"`
	Item5           int              `json:"item_5"`
	Kills           int              `json:"kills"`
	LastHits        int              `json:"last_hits"`
	LeaverStatus    int              `json:"leaver_status"`
	Level           int              `json:"level"`
	PlayerSlot      int              `json:"player_slot"`
	PlayerID        int64            `json:"player_id"`
	TowerDamage     int              `json:"tower_damage"`
	XpPerMin        int              `json:"xp_per_min"`
}

type Players []Player

// Steam Player Data
type PlayerSummaryData struct {
	Response struct {
		PlayerSummaries []PlayerSummary `json:"players"`
	} `json:"response"`
}

type PlayerSummaries []PlayerSummary

type PlayerSummary struct {
	Avatar                   string `json:"avatar"`
	Avatarfull               string `json:"avatarfull"`
	Avatarmedium             string `json:"avatarmedium"`
	Communityvisibilitystate int    `json:"communityvisibilitystate"`
	Lastlogoff               int    `json:"lastlogoff"`
	Loccountrycode           string `json:"loccountrycode"`
	Locstatecode             string `json:"locstatecode"`
	Personaname              string `json:"personaname"`
	Personastate             int    `json:"personastate"`
	Personastateflags        int    `json:"personastateflags"`
	Primaryclanid            string `json:"primaryclanid"`
	Profilestate             int    `json:"profilestate"`
	Profileurl               string `json:"profileurl"`
	Realname                 string `json:"realname"`
	Steamid                  string `json:"steamid"`
	Timecreated              int    `json:"timecreated"`
}

func convertSteamId32to64(id32 int) int {

	id64 := id32 + 76561197960265728

	return id64

}

func convertSteamId64to32(id64 int) int {

	id32 := id64 - 76561197960265728

	return id32

}
