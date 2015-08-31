/*
 * mtStats Devour - Models Package - Heroes
 */

package models

import (
	"database/sql"
	"fmt"
)

// Hero Data
type HeroesData struct {
	Result struct {
		Count  int    `json:"count"`
		Heroes []Hero `json:"heroes"`
		Status int    `json:"status"`
	} `json:"result"`
}

type Heroes []Hero

type Hero struct {
	ID            int    `json:"id"`
	LocalizedName string `json:"localized_name"`
	Name          string `json:"name"`
}

// Uses Hero data to produce a URL to a Heroes image.
func (h *Hero) GetHeroImageURL(imgsize string) (url string) {

	// Remove unwanted characters "npc_dota_hero_"
	name := h.Name[14:len(h.Name)]

	var size string

	// Set image size
	switch imgsize {

	case "full":
		size = "full.png"
	case "large":
		size = "lg.png"
	case "small":
		size = "sb.png"
	case "vertical":
		size = "vert.jpg"
	default:
		size = "lg.png"

	}

	// Produce URL to image
	url = fmt.Sprintf("http://cdn.dota2.com/apps/dota2/images/heroes/%s_%s", name, size)

	return url

}

// Save Game Items to the Database
func SaveHeroes(db *sql.DB, heroes Heroes, lang string) (err error) {

	// Prepare the Heroes Statement
	stmt, err := db.Prepare(`
		INSERT
			tbl_heroes
		SET
			id_hero=?,
			name=?,
			localized_name=?,
			language=?
		ON DUPLICATE KEY UPDATE
			name=?,
			localized_name=?,
			language=?
	`)
	if err != nil {
		fmt.Println(err)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		return
	}

	// defer the closing, whether Commit or Rollback
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Loop all items and execute the SQL statement
	for _, h := range heroes {

		if _, err = tx.Stmt(stmt).Exec(h.ID, h.Name, h.LocalizedName, lang, h.Name, h.LocalizedName, lang); err != nil {
			fmt.Println(err)
			return err
		}

	}

	return

}
