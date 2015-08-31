/*
 * mtStats Devour - Models Package - Picks and Bans
 */

package models

import (
//"database/sql"
//"fmt"
)

// Picks and Bans Data
// Part of struct Match
type PicksBans struct {
	HeroID int  `json:"hero_id"`
	IsPick bool `json:"is_pick"`
	Order  int  `json:"order"`
	Team   int  `json:"team"`
}
