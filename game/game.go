package game

import (
	"fmt"

	"github.com/vayw/kickerleague/database"
	"github.com/vayw/kickerleague/models"
)

type MatchInfo struct {
	Match models.Match
	RD    models.MatchData
	RF    models.MatchData
	BD    models.MatchData
	BF    models.MatchData
	Goals []Goal
}

type Positions struct {
	Defender string
	Forward  string
}

type LineUp struct {
	Red  Positions
	Blue Positions
}

func NewMatch(positions Positions) (models.Match, error) {
	database.ConnectDB()
	defer database.DBCon.Close()

	var temp_player models.Player
	var match MatchInfo

	for _, t := range []string{"Red", "Blue"} {
		for _, p := range []string{"Defender", "Forward"} {
			res := db.Where(&models.Player{Name: positions[t][p]}).First(&temp_player)
			if res.Error == nil {
				match[p] = models.MatchData{PlayerID: temp_player["ID"], Position: p, Team: t}
			} else {
				fmt.Println(res.Error)
			}
		}
	}

	//match = models.Match{Red_score: 0, Blue_score: 0, Winner: "None", TS: time.Unix(1e9, 0).UTC()}
	//return match nil
}
