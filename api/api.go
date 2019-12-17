package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vayw/kickerleague/game"
	"github.com/vayw/kickerleague/player"
)

func Server(addr string) {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/api/players", func(c *gin.Context) {
		list, _ := player.PlayerList()
		c.JSON(200, list)
	})

	router.POST("/api/match/start", func(c *gin.Context) {
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
	})

	router.POST("/api/match/end", func(c *gin.Context) {
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
	})

	router.POST("/api/match/score", func(c *gin.Context) {
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
	})

	router.POST("/api/stats/matchresults", matchResults)
	router.POST("/api/stats/ratings/goals", scorersTable)
	router.GET("/api/stats/overall", overallStats)
	router.GET("/api/stats/winrate", winRate)

	http.ListenAndServe(addr, router)
}
