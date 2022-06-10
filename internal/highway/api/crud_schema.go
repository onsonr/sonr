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
	var req st.MsgCreateSchema
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
		return
	}

	resp, err := s.Cosmos.BroadcastCreateSchema(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	schemaDef := st.Schema{
		Did:    resp.Schema.Did,
		Label:  resp.Schema.Label,
		Fields: resp.Schema.Fields,
	}

	cid, err := s.ipfsProtocol.PutObjectSchema(&schemaDef)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"Schema":  resp.Schema,
		"cid":     cid.String(),
	})
}
