package core

import (
	context "context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	otv1 "github.com/sonr-io/sonr/internal/blockchain/x/object/types"
	t "github.com/sonr-io/sonr/types"
	ot "go.buf.build/grpc/go/sonr-io/blockchain/object"
)

// CreateObject creates a new object.
func (s *HighwayServer) CreateObject(ctx context.Context, req *ot.MsgCreateObject) (*ot.MsgCreateObjectResponse, error) {
	// Broadcast the Transaction
	resp, err := s.Cosmos.BroadcastCreateObject(otv1.NewMsgCreateObjectFromBuf(req))
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.String())

	// Upload Object Schema to IPFS
	cid, err := s.ipfsProtocol.PutObjectSchema(resp.GetWhatIs().GetObjectDoc())
	if err != nil {
		return nil, err
	}
	fmt.Println(cid)
	return &ot.MsgCreateObjectResponse{}, nil
}

// @Summary Create Object
// @Schemes
// @Description CreateObject creates a Object for a registered application
// @Tags Object
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /object/create [post]
func (s *HighwayServer) CreateObjectHTTP(c *gin.Context) {
	// Unmarshal the request body
	var req ot.MsgCreateObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
	}

	// Create the bucket
	resp, err := s.GRPCClient.CreateObject(s.ctx, &req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"what_is": otv1.NewWhatIsFromBuf(resp.WhatIs),
	})
}

// UpdateObject updates an object.
func (s *HighwayServer) UpdateObject(ctx context.Context, req *ot.MsgUpdateObject) (*ot.MsgUpdateObjectResponse, error) {
	// Broadcast the Transaction
	resp, err := s.Cosmos.BroadcastUpdateObject(otv1.NewMsgUpdateObjectFromBuf(req))
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.String())
	return &ot.MsgUpdateObjectResponse{
		Code:    resp.Code,
		Message: resp.Message,
		WhatIs:  otv1.NewWhatIsToBuf(resp.WhatIs),
	}, nil
}

// @Summary Update Object
// @Schemes
// @Description UpdateObject updates and object reference for a registered application
// @Tags Object
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /object/update [post]
func (s *HighwayServer) UpdateObjectHTTP(c *gin.Context) {
	// Unmarshal the request body
	var req ot.MsgUpdateObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
	}

	// Create the bucket
	resp, err := s.GRPCClient.UpdateObject(s.ctx, &req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"what_is": otv1.NewWhatIsFromBuf(resp.WhatIs),
	})
}

// DeactivateBucket disables a bucket for a registered application
func (s *HighwayServer) DeactivateObject(ctx context.Context, req *ot.MsgDeactivateObject) (*ot.MsgDeactivateObjectResponse, error) {
	resp, err := s.Cosmos.BroadcastDeactivateObject(otv1.NewMsgDeactivateObjectFromBuf(req))
	if err != nil {
		return nil, err
	}
	logger.Infof(resp.String())
	return &ot.MsgDeactivateObjectResponse{
		Code:    resp.Code,
		Message: resp.Message,
	}, nil
}

// @Summary  Deactivate Object
// @Schemes
// @Description DeactivateObject disables a Object for a registered application
// @Tags Object
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /object/deactivate [post]
func (s *HighwayServer) DeactivateObjectlHTTP(c *gin.Context) {
	// Unmarshal the request body
	var req ot.MsgDeactivateObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
	}

	// Deactivate the bucket
	resp, err := s.GRPCClient.DeactivateObject(s.ctx, &req)
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
