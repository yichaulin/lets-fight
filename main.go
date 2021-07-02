package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"lets-fight/controller"
)

func main() {
	r := gin.Default()
	r.GET("/combat", controller.CombatController)
	r.GET("/welcome", func(c *gin.Context) {
		fighters := c.QueryArray("fighters[]")
		qqq := c.Query("qqq")

		c.String(http.StatusOK, "Hello %s %s", fighters, qqq)
	})
	r.Run()
}
