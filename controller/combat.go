package controller

import (
	"lets-fight/combat"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CombatController(c *gin.Context) {
	fighters := c.QueryArray("fighters[]")

	combatResult, err := combat.New(fighters[0], fighters[1])

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, combatResult)
}
