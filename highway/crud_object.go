package highway

import (
	context "context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	otv1 "github.com/sonr-io/blockchain/x/object/types"
	ot "go.buf.build/grpc/go/sonr-io/blockchain/object"
)

// CreateObject creates a new object.
func (s *HighwayServer) CreateObject(ctx context.Context, req *ot.MsgCreateObject) (*ot.MsgCreateObjectResponse, error) {
	// Broadcast the Transaction
	resp, err := s.cosmos.BroadcastCreateObject(otv1.NewMsgCreateObjectFromBuf(req))
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

// CreateBucketHTTP creates a new bucket via HTTP.
func (s *HighwayServer) CreateObjectHTTP(c *gin.Context) {
	// Unmarshal the request body
	var req ot.MsgCreateObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrRequestBody.Error(),
		})
	}

	// Create the bucket
	resp, err := s.grpcClient.CreateObject(s.ctx, &req)
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
	resp, err := s.cosmos.BroadcastUpdateObject(otv1.NewMsgUpdateObjectFromBuf(req))
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

// CreateBucketHTTP creates a new bucket via HTTP.
func (s *HighwayServer) UpdateObjectHTTP(c *gin.Context) {
	// Unmarshal the request body
	var req ot.MsgUpdateObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrRequestBody.Error(),
		})
	}

	// Create the bucket
	resp, err := s.grpcClient.UpdateObject(s.ctx, &req)
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
	resp, err := s.cosmos.BroadcastDeactivateObject(otv1.NewMsgDeactivateObjectFromBuf(req))
	if err != nil {
		return nil, err
	}
	logger.Infof(resp.String())
	return &ot.MsgDeactivateObjectResponse{
		Code:    resp.Code,
		Message: resp.Message,
	}, nil
}

// DeactivateBucketHTTP disables a bucket for a registered application via HTTP.
func (s *HighwayServer) DeactivateObjectlHTTP(c *gin.Context) {
	// Unmarshal the request body
	var req ot.MsgDeactivateObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": ErrRequestBody.Error(),
		})
	}

	// Deactivate the bucket
	resp, err := s.grpcClient.DeactivateObject(s.ctx, &req)
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
