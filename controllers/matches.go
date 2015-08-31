/*
 * mtStats Devour - Controllers Package - Match History by Sequence
 */

package controllers

import (
	"database/sql"
	"fmt"
	"github.com/mtstats/devour/models"
	"github.com/mtstats/devour/webapi"
)

// Fetch Matches in Sequence from the Steam API and save them to the DB.
func ContMatchHistorySeq(db *sql.DB, api webapi.WebApi, lastSeqNum *int) (err error) {

	// Get Matches from API
	matches, err := api.GetMatchHistorySeq(lastSeqNum)
	if err != nil {
		return err
	}

	// Filter Matches for Ranked and Tournament
	/*for i := len(matches) - 1; i >= 0; i-- {

		switch matches[i].LobbyType {
		case 2:
			// Tournament
		case 7:
			// Ranked Match
		default:
			matches = append(matches[:i], matches[i+1:]...)
		}

	}*/

	// Save Matches to DB
	err = models.SaveSeqMatches(db, matches)
	if err != nil {
		return err
	}

	if len(matches) > 0 {
		lastItem := len(matches) - 1
		*lastSeqNum = matches[lastItem].MatchSeqNum
	}

	return err

}

func GetLastSeqNum(db *sql.DB) (last int, err error) {

	err = db.QueryRow("SELECT seq_num FROM tbl_matches ORDER BY seq_num DESC LIMIT 1").Scan(&last)
	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("No matches found.")
	case err != nil:
		fmt.Printf(err.Error())
	}

	return last, err

}
