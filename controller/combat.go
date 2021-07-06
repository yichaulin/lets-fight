package controller

import (
	"lets-fight/combat"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CombatController(c *gin.Context) {
	fightersParam := c.QueryArray("fighters[]")
	if len(fightersParam) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No. of fighters is less than 2"})
		return
	}
	var fighters [2]string
	copy(fighters[:], fightersParam[:2])
	combatResult, err := combat.New(fighters)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, combatResult)
}
