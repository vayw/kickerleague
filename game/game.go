package game

import (
	"fmt"
	"time"

	"github.com/vayw/kickerleague/database"
	"github.com/vayw/kickerleague/models"
)

type MatchInfo struct {
	Match  models.Match
	Lineup LineUp
	Goals  []MGoal
	TS     time.Time
}

type MGoal struct {
	TS     time.Time
	Scorer int
}

type Positions struct {
	Defender string
	Forward  string
}

type LineUp struct {
	Red  Positions
	Blue Positions
}

func NewMatch(lineup Positions) (string, error) {
	database.ConnectDB()
	defer database.DBCon.Close()

	var temp_player models.Player
	var match MatchInfo

	for _, t := range []string{"Red", "Blue"} {
		for _, p := range []string{"Defender", "Forward"} {
			res := db.Where(&models.Player{Name: lineup[t][p]}).First(&temp_player)
			if res.Error == nil {
				match[p] = models.MatchData{PlayerID: temp_player["ID"], Position: p, Team: t}
			} else {
				fmt.Println(res.Error)
			}
		}
	}

	return "3482", nil
}
