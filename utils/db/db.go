/*
 * mtStats Devour - Database Package
 */

package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mtstats/devour/utils/config"
	"os"
)

func OpenDb(c config.Configuration) (db *sql.DB) {

	var err error

	db, err = sql.Open("mysql", c.Database.User+":"+c.Database.Pass+"@/"+c.Database.Name)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Check connection is valid
	if err = db.Ping(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return db

}
