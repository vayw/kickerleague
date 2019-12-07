package api

import (
	"fmt"

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
	c.BindJSON(&data)

	var matches []models.Match
	database.DBCon.Limit(data.Num).Order("ts desc").Find(&matches)

	for _, i := range matches {
		fmt.Println(i)
	}
	c.JSON(200, "OK")

}
