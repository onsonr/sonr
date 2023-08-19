package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CurrentUser returns the current user's details.
//
// @Summary Current user's details
// @Description Returns the current user's details.
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string "JWT Token"
// @Failure 400 {object} map[string]string "Missing JWT"
// @Failure 401 {object} map[string]string "No JWT cookie found"
// @Router /currentUser [get]
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

// GetHealth returns the health of the service.
//
// @Summary Health of the service
// @Description Returns the health of the service.
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string "Status message"
// @Router /getHealth [get]
func GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "im alive and well",
	})
}

// GetBlockHeight returns the current block height.
//
// @Summary Current block height
// @Description Returns the current block height.
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]int "Block Height"
// @Router /getBlockHeight [get]
func GetBlockHeight(c *gin.Context) {
	c.JSON(200, gin.H{
		"blockHeight": 0,
	})
}

// GetValidatorSet returns the current validator set.
//
// @Summary Current validator set
// @Description Returns the current validator set.
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]string "Validators"
// @Router /getValidatorSet [get]
func GetValidatorSet(c *gin.Context) {
	c.JSON(200, gin.H{
		"validators": []string{},
	})
}
