package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	t "github.com/sonr-io/sonr/types"
	ctv1 "github.com/sonr-io/sonr/x/channel/types"
)

// @Summary Create Channel
// @Schemes
// @Description CreateChannel creates a specified channel for a registered application
// @Tags Channel
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /v1/channel/create [post]
func (s *HighwayServer) CreateChannel(c *gin.Context) {
	// Unmarshal the request body
	var req ctv1.MsgCreateChannel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
	}

	// Verify that channel fields are not nil
	if req.GetObjectToRegister().GetFields() == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
	}

	// Broadcast the message
	resp, err := s.Cosmos.BroadcastCreateChannel(&req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"how_is":  resp.HowIs,
	})
}

// @Summary Update Channel
// @Schemes
// @Description ListenChannel puts a Channel into a listening state registered application
// @Tags Channel
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /v1/channel/update [post]
func (s *HighwayServer) UpdateChannel(c *gin.Context) {
	// Unmarshal the request body
	var req ctv1.MsgUpdateChannel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
	}

	resp, err := s.Cosmos.BroadcastUpdateChannel(&req)
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

// // @Summary Listen Channel
// // @Schemes
// // @Description ListenChannel puts a Channel into a listening state registered application
// // @Tags Channel
// // @Produce json
// // @Success      200  {string}  message
// // @Failure      500  {string}  message
// // @Router /channel/listen [post]
// func (s *HighwayServer) ListenChannel(c *gin.Context) {
// 	// Unmarshal the request body
// 	var req ctv1.MsgListenChannel
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": t.ErrRequestBody.Error(),
// 		})
// 	}

// 	// Setup concurrent stream
// 	opChan := make(chan *ctv1.ChannelMessage)
// 	ch, ok := s.channels[req.GetDid()]
// 	if !ok {
// 		c.JSON(http.StatusBadRequest, t.ErrRequestBody.Error())
// 	}
// 	go ch.Listen(opChan)

// 	// Listen to the channel
// 	c.Stream(func(io.Writer) bool {
// 		// Get the next operation
// 		op, ok := <-opChan
// 		if !ok {
// 			return false
// 		}

// 		// Create ChannelMessage
// 		msg := &ct.ChannelMessage{
// 			PeerDid:  op.GetPeerDid(),
// 			Did:      op.GetDid(),
// 			Object:   otv1.NewObjectDocToBuf(op.Object),
// 			Metadata: op.GetMetadata(),
// 		}

// 		// Marshal the proto message
// 		data, err := proto.Marshal(msg)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 			return false
// 		}

// 		// Send the message
// 		c.Writer.Write(data)
// 		return true
// 	})
// }

// @Summary Deactivate Channel
// @Schemes
// @Description DeactivateChannel disables a Channel for a registered application
// @Tags Channel
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /v1/channel/deactivate [post]
func (s *HighwayServer) DeactivateChannel(c *gin.Context) {
	// Unmarshal the request body
	var req ctv1.MsgDeactivateChannel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
	}

	resp, err := s.Cosmos.BroadcastDeactivateChannel(&req)
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
