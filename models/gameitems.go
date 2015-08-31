/*
 * mtStats Devour - Models Package - Game Items
 */

package models

import (
	"database/sql"
	"fmt"
)

// Game Item Data
type GameItemsData struct {
	Result struct {
		Items  []GameItem `json:"items"`
		Status int        `json:"status"`
	} `json:"result"`
}

type GameItems []GameItem

type GameItem struct {
	Cost          int    `json:"cost"`
	ID            int    `json:"id"`
	LocalizedName string `json:"localized_name"`
	Name          string `json:"name"`
	Recipe        int    `json:"recipe"`
	SecretShop    int    `json:"secret_shop"`
	SideShop      int    `json:"side_shop"`
}

// Uses GameItem data to produce a URL to an GameItems image.
func (i *GameItem) GetGameItemImageURL() string {

	// Remove unwanted characters "item_"
	name := i.Name[5:len(i.Name)]

	// Set image size
	size := "lg.png"

	// Produce URL to image
	url := fmt.Sprintf("http://cdn.dota2.com/apps/dota2/images/items/%s_%s", name, size)

	return url

}

// Save Game Items to the Database
func SaveGameItems(db *sql.DB, items GameItems, lang string) (err error) {

	// Prepare the Game Items Statement
	//stmt, err := db.Prepare("INSERT tbl_items SET id_item=?, name=?, cost=?, secret_shop=?, side_shop=?, recipe=?, localized_name=? ON DUPLICATE KEY UPDATE name=?, cost=?, secret_shop=?, side_shop=?, recipe=?, localized_name=?")
	stmt, err := db.Prepare(`
		INSERT
			tbl_items
		SET
			id_item=?,
			name=?,
			cost=?,
			secret_shop=?,
			side_shop=?,
			recipe=?,
			localized_name=?,
			language=?
		ON DUPLICATE KEY UPDATE
			name=?,
			cost=?,
			secret_shop=?,
			side_shop=?,
			recipe=?,
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
	for _, i := range items {

		if _, err = tx.Stmt(stmt).Exec(i.ID, i.Name, i.Cost, i.SecretShop, i.SideShop, i.Recipe, i.LocalizedName, lang, i.Name, i.Cost, i.SecretShop, i.SideShop, i.Recipe, i.LocalizedName, lang); err != nil {
			fmt.Println(err)
			return err
		}

	}

	return

}
