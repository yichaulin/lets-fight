package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"lets-fight/controller"
)

func main() {
	r := gin.Default()
	r.GET("/combat", controller.CombatController)
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Let's fight backend")
	})
	r.Run()
}
