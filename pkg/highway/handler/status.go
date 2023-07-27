package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CurrentUser returns the current user's details.
func CurrentUser(c *gin.Context) {
	jwt, err := c.Cookie("sonr-jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "no jwt cookie found"})
		return
	}
	if jwt == "" {
		c.JSON(400, gin.H{
			"error": "missing jwt",
		})
		return
	}
	c.JSON(200, gin.H{
		"jwt": jwt,
	})
}

func GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "im alive and well",
	})
}

func GetBlockHeight(c *gin.Context) {
	c.JSON(200, gin.H{
		"blockHeight": 0,
	})
}

func GetValidatorSet(c *gin.Context) {
	c.JSON(200, gin.H{
		"validators": []string{},
	})
}
