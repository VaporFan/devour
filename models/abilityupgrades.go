/*
 * mtStats Devour - Models Package - Abilities
 */

package models

import (
//"fmt"
)

// Ability Upgrade Data
// Part of struct Player
type AbilityUpgrade struct {
	Ability int `json:"ability"`
	Level   int `json:"level"`
	Time    int `json:"time"`
}
