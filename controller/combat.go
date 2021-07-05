package controller

import (
	"lets-fight/combat"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CombatController(c *gin.Context) {
	fighters := c.QueryArray("fighters[]")
	if len(fighters) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No. of fighters is less than 2"})
		return
	}

	combatResult, err := combat.New(fighters[0], fighters[1])

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, combatResult)
}
