package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/vayw/kickerleague/docs"
	"github.com/vayw/kickerleague/player"
)

// @title Swagger Example API
// @version 1.0
// @description This is
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/vayw/kickerleague

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @BasePath /api/

// @x-extension-openapi {"example": "value on a json format"}

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

	docs.SwaggerInfo.Title = "Swagger Kickerleague API"
	docs.SwaggerInfo.Version = "1.0"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	http.ListenAndServe(addr, router)
}
