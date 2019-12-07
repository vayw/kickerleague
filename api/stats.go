package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vayw/kickerleague/database"
	"github.com/vayw/kickerleague/game"
	"github.com/vayw/kickerleague/models"
)

func matchResults(c *gin.Context) {
	type Data struct {
		Num int `json:"num" binding:"required"`
	}

	type Result struct {
		Red    int
		Blue   int
		Lineup []game.Player
	}

	var data Data
	var matches []models.Match

	c.BindJSON(&data)

	result := make([]Result, data.Num)
	database.DBCon.Limit(data.Num).Order("ts desc").Find(&matches)

	var matchlineup []models.MatchData
	for index, i := range matches {
		result[index] = Result{}
		result[index].Red = i.Red_score
		result[index].Blue = i.Blue_score
		result[index].Lineup = make([]game.Player, 4)
		database.DBCon.Where("match_id = ?", i.ID).Find(&matchlineup)
		for pindex, j := range matchlineup {
			result[index].Lineup[pindex] = game.Player{j.PlayerID, j.Team, j.Position}
		}
	}

	c.JSON(200, result)
}
