package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	t "github.com/sonr-io/sonr/types"
	btt "github.com/sonr-io/sonr/x/bucket/types"
)

// @Summary Create Bucket
// @Schemes
// @Description CreateBucket creates a new bucket for a registered application via HTTP.
// @Tags Bucket
// @Produce json
// @Success      200  {string}  cid
// @Failure      500  {string}  message
// @Router /v1/bucket/create [post]
func (s *HighwayServer) CreateBucket(c *gin.Context) {
	// Unmarshal the request body
	var req btt.MsgCreateBucket
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
	}
	resp, err := s.Cosmos.BroadcastCreateBucket(&req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"code":     resp.Code,
		"message":  resp.Message,
		"which_is": resp.WhichIs,
	})
}

// @Summary Update Bucket
// @Schemes
// @Description UpdateBucket updates a bucket for a registered application via HTTP.
// @Tags Bucket
// @Produce json
// @Success      200  {string}  cid
// @Failure      500  {string}  message
// @Router /v1/bucket/update [post]
func (s *HighwayServer) UpdateBucket(c *gin.Context) {
	// Unmarshal the request body
	var req btt.MsgUpdateBucket
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
	}

	resp, err := s.Cosmos.BroadcastUpdateBucket(&req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"code":     resp.Code,
		"message":  resp.Message,
		"which_is": resp.WhichIs,
	})
}

// @Summary Deactivate Bucket
// @Schemes
// @Description DeactivateBucket disables a bucket for a registered application via HTTP.
// @Tags Bucket
// @Produce json
// @Success      200  {string}  message
// @Failure      400  {string}  message
// @Failure      502  {string}  message
// @Router /v1/bucket/deactivate [post]
func (s *HighwayServer) DeactivateBucket(c *gin.Context) {
	// Unmarshal the request body
	var req btt.MsgDeactivateBucket
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
	}
	resp, err := s.Cosmos.BroadcastDeactivateBucket(&req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}
	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
	})
}
