/*
 * mtStats Devour - Controllers Package - Game Items
 */

package controllers

import (
	"database/sql"
	"github.com/mtstats/devour/models"
	"github.com/mtstats/devour/webapi"
)

// Fetch all Game Items from the Steam API and save them to the DB.
func GetGameItems(db *sql.DB, api webapi.WebApi) (err error) {

	lang := "en_us"

	items, err := api.GetGameItems(lang)
	if err != nil {
		return err
	}

	err = models.SaveGameItems(db, items, lang)
	if err != nil {
		return err
	}

	return err

}
