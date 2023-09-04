package handler

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gin-gonic/gin"

	databasepb "github.com/sonrhq/core/types/highway/database/v1"
)


// DatabaseHandler is the handler for the authentication service
type DatabaseHandler struct {
	cctx client.Context
}

// Health returns the health of the service.
//
// @Summary Health of the service
// @Description Returns the health of the service.
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string "Status message"
// @Router /getHealth [get]
func (a *DatabaseHandler) Health(ctx context.Context, req *databasepb.HealthRequest) (*databasepb.HealthResponse, error) {
	return &databasepb.HealthResponse{
		Ok: true,
	}, nil
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
