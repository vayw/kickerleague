package game

import (
	"errors"
	"fmt"
	"time"

	"github.com/vayw/kickerleague/database"
	"github.com/vayw/kickerleague/models"
)

type Player struct {
	ID       int    `json:"id"`
	Team     string `json:"team"`
	Position string `json:"position"`
}

type Result struct {
	Red    int
	Blue   int
	Winner string
}

func NewMatch(lineup []Player) (int, error) {
	match := models.Match{Red_score: 0, Blue_score: 0, Winner: "None", TS: time.Now()}
	if len(lineup) != 4 {
		return 0, errors.New("you must provide 4 players")
	}

	m := make(map[int]bool)
	for _, i := range lineup {
		if _, ok := m[i.ID]; !ok {
			m[i.ID] = true
		} else {
			return 0, errors.New("there should not be the same player more than once")
		}
	}

	tx := database.DBCon.Begin()

	var matchid int
	if err := tx.Create(&match).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	matchid = match.ID

	for _, i := range lineup {
		matchdata := models.MatchData{PlayerID: i.ID, Position: i.Position, Team: i.Team, MatchID: matchid}
		if err := tx.Create(&matchdata).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
	return matchid, nil
}

func Score(scorer int, matchid int, auto bool) error {
	goal := models.Goal{PlayerID: scorer, MatchID: matchid, Auto: auto, TS: time.Now()}
	if err := database.DBCon.Create(&goal).Error; err != nil {
		return err
	}
	return nil
}

func EndMatch(matchid int) (Result, error) {
	var match models.Match

	if err := database.DBCon.First(&match, matchid).Error; err != nil {
		return Result{}, err
	}

	var red_score int
	var blue_score int
	var winner string

	rows, err := database.DBCon.Table("goals").Select("goals.auto, match_data.team").Joins("inner join match_data on match_data.player_id = goals.player_id and match_data.match_id = goals.match_id").Where("goals.match_id = ?", match.ID).Rows()
	if err != nil {
		return Result{}, err
	}
	var team string
	var auto bool
	for rows.Next() {
		if err := rows.Scan(&auto, &team); err != nil {
			fmt.Println(err)
		}
		switch team {
		case "red", "Red":
			if !auto {
				red_score += 1
			} else {
				blue_score += 1
			}
		case "blue", "Blue":
			if !auto {
				blue_score += 1
			} else {
				red_score += 1
			}
		}
	}

	switch {
	case red_score > blue_score:
		winner = "Red"
	case red_score < blue_score:
		winner = "Blue"
	default:
		winner = "Draw"
	}

	match.Red_score = red_score
	match.Blue_score = blue_score
	match.Winner = winner
	database.DBCon.Save(&match)

	res := Result{red_score, blue_score, winner}
	return res, err
}
