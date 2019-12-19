package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vayw/kickerleague/player"
)

func Server(addr string) {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/api/players", func(c *gin.Context) {
		list, _ := player.PlayerList()
		c.JSON(200, list)
	})

	router.POST("/api/match/start", matchStart)
	router.POST("/api/match/end", matchEnd)
	router.POST("/api/match/score", matchScore)

	router.POST("/api/stats/matchresults", matchResults)
	router.POST("/api/stats/ratings/goals", scorersTable)
	router.GET("/api/stats/overall", overallStats)
	router.GET("/api/stats/winrate", winRate)

	http.ListenAndServe(addr, router)
}
