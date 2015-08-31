/*
 * mtStats Devour - Models Package - Matches
 */

package models

import (
	"database/sql"
	"fmt"
	"time"
)

// Dota 2 Match Data
type MatchData struct {
	Result struct {
		Matches []Match `json:"matches"`
		Status  int     `json:"status"`
	} `json:"result"`
}

type Matches []Match

type Match struct {
	BarracksStatusDire    int         `json:"barracks_status_dire"`
	BarracksStatusRadiant int         `json:"barracks_status_radiant"`
	Cluster               int         `json:"cluster"`
	DireCaptain           int         `json:"dire_captain"`
	DireLogo              int         `json:"dire_logo"`
	DireName              string      `json:"dire_name"`
	DireTeamComplete      int         `json:"dire_team_complete"`
	DireTeamID            int         `json:"dire_team_id"`
	Duration              int         `json:"duration"`
	Engine                int         `json:"engine"`
	FirstBloodTime        int         `json:"first_blood_time"`
	GameMode              int         `json:"game_mode"`
	HumanPlayers          int         `json:"human_players"`
	Leagueid              int         `json:"leagueid"`
	LobbyType             int         `json:"lobby_type"`
	MatchID               int         `json:"match_id"`
	MatchSeqNum           int         `json:"match_seq_num"`
	NegativeVotes         int         `json:"negative_votes"`
	PicksBans             []PicksBans `json:"picks_bans"` // See picksbans.go
	Players               []Player    `json:"players"`
	PositiveVotes         int         `json:"positive_votes"`
	RadiantCaptain        int         `json:"radiant_captain"`
	RadiantLogo           int         `json:"radiant_logo"`
	RadiantName           string      `json:"radiant_name"`
	RadiantTeamComplete   int         `json:"radiant_team_complete"`
	RadiantTeamID         int         `json:"radiant_team_id"`
	RadiantWin            bool        `json:"radiant_win"`
	StartTime             int64       `json:"start_time"`
	TowerStatusDire       int         `json:"tower_status_dire"`
	TowerStatusRadiant    int         `json:"tower_status_radiant"`
}

// Save Sequence Matches to the Database
func SaveSeqMatches(db *sql.DB, matches Matches) (err error) {

	// Prepare the Match Statement
	//mStmt, err := db.Prepare("INSERT tbl_matches SET id_match=?, seq_num=?, start_time=?, lobby_type=?, radiant_team_id=?, dire_team_id=?")
	mStmt, err := db.Prepare(`
		INSERT
			tbl_matches
		SET
			id_match=?,
			seq_num=?,
			start_time=?,
			lobby_type=?,
			radiant_team_id=?,
			dire_team_id=?,
			radiant_win=?,
			duration=?,
			tower_status_radiant=?,
			tower_status_dire=?,
			barracks_status_radiant=?,
			barracks_status_dire=?,
			cluster=?,
			first_blood_time=?,
			human_players=?,
			league_id=?,
			positive_votes=?,
			negative_votes=?,
			game_mode=?,
			engine=?
	`)
	defer mStmt.Close()
	if err != nil {
		return err
	}

	// Prepare Player Statement
	//pStmt, err := db.Prepare("INSERT tbl_players SET account_id=?, player_slot=?, hero_id=?, id_match=?")
	pStmt, err := db.Prepare(`
		INSERT
			tbl_players
		SET
			account_id=?,
			player_slot=?,
			hero_id=?,
			item_0=?,
			item_1=?,
			item_2=?,
			item_3=?,
			item_4=?,
			item_5=?,
			kills=?,
			deaths=?,
			assists=?,
			leaver_status=?,
			gold=?,
			last_hits=?,
			denies=?,
			gold_per_min=?,
			xp_per_min=?,
			gold_spent=?,
			hero_damage=?,
			tower_damage=?,
			hero_healing=?,
			level=?,
			id_match=?
	`)
	defer pStmt.Close()
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// defer the closing, whether Commit or Rollback
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Loop all matches and execute the SQL statement
	for _, m := range matches {

		// Filter out Matches we dont want
		if m.HumanPlayers < 10 || m.Duration < 300 {
			continue
		}

		switch m.LobbyType {
		case 1:
			// Get Public Matchmaking Matches
		case 2:
			// Get Tournament Matches
		case 5:
			// Get Team Matches
		case 6:
			// Get Solo Matches
		case 7:
			// Get Ranked Matches
		default:
			continue
		}

		startStamp := time.Unix(m.StartTime, 0)

		// Execute Match Statement
		if _, err = tx.Stmt(mStmt).Exec(m.MatchID, m.MatchSeqNum, startStamp, m.LobbyType, m.RadiantTeamID, m.DireTeamID, m.RadiantWin, m.Duration, m.TowerStatusRadiant, m.TowerStatusDire, m.BarracksStatusRadiant, m.BarracksStatusDire, m.Cluster, m.FirstBloodTime, m.HumanPlayers, m.Leagueid, m.PositiveVotes, m.NegativeVotes, m.GameMode, m.Engine); err != nil {
			fmt.Println(err)
			return err
		}

		// Loop all players and execute the SQL statement
		for _, p := range m.Players {

			if _, err := tx.Stmt(pStmt).Exec(p.AccountID, p.PlayerSlot, p.HeroID, p.Item0, p.Item1, p.Item2, p.Item3, p.Item4, p.Item5, p.Kills, p.Deaths, p.Assists, p.LeaverStatus, p.Gold, p.LastHits, p.Denies, p.GoldPerMin, p.XpPerMin, p.GoldSpent, p.HeroDamage, p.TowerDamage, p.HeroHealing, p.Level, m.MatchID); err != nil {

				fmt.Println(err)
				return err

			}

		}

	}

	return err

}
