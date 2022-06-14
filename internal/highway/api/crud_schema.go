package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	t "github.com/sonr-io/sonr/types"
	st "github.com/sonr-io/sonr/x/schema/types"
)

// @Summary Create Schema
// @Schemes
// @Description Creates an application data schema.
// @Tags Schema
// @Produce json
// @Success      200  {string}  MsgCreateSchemaResponse
// @Failure      500  {string}  message
// @Router /v1/ipfs/upload [post]
func (s *HighwayServer) CreateSchema(c *gin.Context) {
	var req st.SchemaDefinition
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
		return
	}
	cid, err := s.ipfsProtocol.PutObjectSchema(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	createSchema := st.MsgCreateSchema{
		Creator: req.Creator,
		Label:   req.Label,
		Cid:     cid.String(),
	}
	resp, err := s.Cosmos.BroadcastCreateSchema(&createSchema)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, st.MsgCreateSchemaResponse{
		Code:    200,
		Message: "Schema Created Sucessfully",
		WhatIs:  resp.WhatIs,
	})
}
