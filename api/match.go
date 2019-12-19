package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vayw/kickerleague/game"
)

func matchStart(c *gin.Context) {
	type Data struct {
		Lineup []game.Player `json:"lineup" binding:"required"`
	}
	var data Data
	c.BindJSON(&data)
	matchid, err := game.NewMatch(data.Lineup)
	if err != nil {
		c.JSON(201, gin.H{"matchid": 0, "err": err.Error()})
	} else {
		c.JSON(201, gin.H{"matchid": matchid, "err": "nil"})
	}
}

func matchEnd(c *gin.Context) {
	type data struct {
		MatchID int `json:"matchid" binding:"required"`
	}
	var payload data
	c.BindJSON(&payload)
	res, err := game.EndMatch(payload.MatchID)
	if err == nil {
		c.JSON(200, res)
	} else {
		c.JSON(201, gin.H{"err": "it's not over, baby"})
	}
}

func matchScore(c *gin.Context) {
	type data struct {
		PID     int  `json:"pid" binding:"required"`
		MatchID int  `json:"matchid" binding:"required"`
		Auto    bool `json:"auto" binding:"required"`
	}
	var payload data
	c.BindJSON(&payload)
	err := game.Score(payload.PID, payload.MatchID, payload.Auto)
	if err != nil {
		c.JSON(200, gin.H{"err": err.Error()})
	} else {
		c.JSON(200, gin.H{"err": "nil"})
	}
}
