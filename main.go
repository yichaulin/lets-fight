package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"lets-fight/controller"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/combat", controller.CombatController)
		api.GET("/health", func(c *gin.Context) {
			c.String(http.StatusOK, "Let's fight backend")
		})
	}

	r.Run()
}
