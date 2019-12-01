package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vayw/kickerleague/player"
)

func Server(addr string) {
	router := gin.Default()
	router.GET("/api/players", func(c *gin.Context) {
		list, _ := player.PlayerList()
		c.JSON(200, list)
	})

	router.POST("/api/match/start", func(c *gin.Context) {
		//var lineup map[int]game.Player
		type R struct {
			cmd string
		}
		var q R
		c.BindJSON(&q)
		fmt.Println(q)
	})

	router.POST("/api/match/end", func(c *gin.Context) {
		list, _ := player.PlayerList()
		c.JSON(200, list)
	})

	http.ListenAndServe(addr, router)
}
