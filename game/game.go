package game

import "github.com/vayw/kickerleague/models"

type MatchInfo struct {
	Match models.Match
}

func NewMatch() (models.Match, error) {
	//match = models.Match{Red_score: 0, Blue_score: 0, Winner: "None", TS: time.Unix(1e9, 0).UTC()}
	//return match nil
}
