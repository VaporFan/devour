/*
 * mtStats Devour - Controllers Package - Heroes
 */

package controllers

import (
	"database/sql"
	"github.com/mtstats/devour/models"
	"github.com/mtstats/devour/webapi"
)

// Fetch all Heroes from the Steam API and save them to the DB.
func GetHeroes(db *sql.DB, api webapi.WebApi) (err error) {

	lang := "en_us"

	heroes, err := api.GetHeroes(lang)
	if err != nil {
		return err
	}

	err = models.SaveHeroes(db, heroes, lang)
	if err != nil {
		return err
	}

	return err

}
