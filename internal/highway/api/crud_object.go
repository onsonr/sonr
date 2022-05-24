package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kataras/go-events"
	metrics "github.com/sonr-io/sonr/internal/highway/x/prometheus"
	t "github.com/sonr-io/sonr/types"
	ot "github.com/sonr-io/sonr/x/object/types"
)

// @Summary Create Object
// @Schemes
// @Description CreateObject creates a Object for a registered application
// @Tags Object
// @Produce json
// @Param 		 data body ot.MsgCreateObject true "Parameters"
// @Success      200  {object}  ot.MsgCreateObjectResponse
// @Failure      500  {string}  message
// @Router /v1/object/create [post]
func (s *HighwayServer) CreateObject(c *gin.Context) {
	// Unmarshal the request body
	var req ot.MsgCreateObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})

		return
	}

	// Broadcast the Transaction
	resp, err := s.Cosmos.BroadcastCreateObject(&req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})

		return
	}

	defer events.Emit(metrics.ON_OBJECT_ADD, "")

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"what_is": resp.WhatIs,
	})
}

// @Summary Update Object
// @Schemes
// @Description UpdateObject updates and object reference for a registered application
// @Tags Object
// @Produce json
// @Param 		 data body ot.MsgUpdateObject true "Parameters"
// @Success      200  {object}  ot.MsgUpdateObjectResponse
// @Failure      500  {string}  message
// @Router /v1/object/update [post]
func (s *HighwayServer) UpdateObject(c *gin.Context) {
	// Unmarshal the request body
	var req ot.MsgUpdateObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
		return
	}

	// Broadcast the Transaction
	resp, err := s.Cosmos.BroadcastUpdateObject(&req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"what_is": resp.WhatIs,
	})
}

// @Summary  Deactivate Object
// @Schemes
// @Description DeactivateObject disables a Object for a registered application
// @Tags Object
// @Produce json
// @Param 		 data body ot.MsgDeactivateObject true "Parameters"
// @Success      200  {object}  ot.MsgDeactivateObjectResponse
// @Failure      500  {string}  message
// @Router /v1/object/deactivate [post]
func (s *HighwayServer) DeactivateObject(c *gin.Context) {
	// Unmarshal the request body
	var req ot.MsgDeactivateObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
		return
	}

	// Deactivate the bucket
	resp, err := s.Cosmos.BroadcastDeactivateObject(&req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer events.Emit(metrics.ON_OBJECT_REMOVE, "")

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
	})
}
