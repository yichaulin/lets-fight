package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"lets-fight/controller"
)

func main() {
	r := gin.Default()
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "https://lets-fight.maxisme.com"},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Origin"},
	})
	api := r.Group("/api")
	api.Use(corsMiddleware)
	{
		api.GET("/combat", controller.CombatController)
		api.GET("/health", func(c *gin.Context) {
			c.String(http.StatusOK, "Let's fight backend")
		})
	}

	r.Run()
}
