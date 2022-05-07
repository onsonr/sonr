package core

import (
	context "context"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ctv1 "github.com/sonr-io/sonr/internal/blockchain/x/channel/types"
	otv1 "github.com/sonr-io/sonr/internal/blockchain/x/object/types"
	t "github.com/sonr-io/sonr/types"
	ct "go.buf.build/grpc/go/sonr-io/blockchain/channel"
	v1 "go.buf.build/grpc/go/sonr-io/highway/v1"
	"google.golang.org/protobuf/proto"
)

// CreateChannel creates a new channel.
func (s *HighwayServer) CreateChannel(ctx context.Context, req *ct.MsgCreateChannel) (*ct.MsgCreateChannelResponse, error) {
	// Verify that channel fields are not nil
	if req.GetObjectToRegister().GetFields() == nil {
		return nil, errors.New("object to register must have fields")
	}

	// Broadcast the message
	res, err := s.Cosmos.BroadcastCreateChannel(ctv1.NewMsgCreateChannelFromBuf(req))
	if err != nil {
		return nil, err
	}

	// Create the Channel
	ch, err := ctv1.NewChannel(ctx, s.Host, res.GetHowIs().GetChannel())
	if err != nil {
		return nil, err
	}

	// Add to the list of Channels
	s.channels[res.GetHowIs().GetDid()] = ch
	return &ct.MsgCreateChannelResponse{
		Code:    res.Code,
		Message: res.Message,
		HowIs:   ctv1.NewHowIsToBuf(res.HowIs),
	}, nil
}

// @Summary Create Channel
// @Schemes
// @Description CreateChannel creates a specified channel for a registered application
// @Tags Channel
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /channel/create [post]
func (s *HighwayServer) CreateChannelHTTP(c *gin.Context) {
	// Unmarshal the request body
	var req ct.MsgCreateChannel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
	}

	// Create the channel
	resp, err := s.GRPCClient.CreateChannel(s.ctx, &req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Message,
		"how_is":  ctv1.NewHowIsFromBuf(resp.HowIs),
	})
}

// UpdateChannel updates a channel.
func (s *HighwayServer) UpdateChannel(ctx context.Context, req *ct.MsgUpdateChannel) (*ct.MsgUpdateChannelResponse, error) {
	resp, err := s.Cosmos.BroadcastUpdateChannel(ctv1.NewMsgUpdateChannelFromBuf(req))
	if err != nil {
		return nil, err
	}
	log.Println(resp.String())
	return &ct.MsgUpdateChannelResponse{
		Code:    resp.Code,
		Message: resp.Message,
	}, nil
}

// @Summary Update Channel
// @Schemes
// @Description ListenChannel puts a Channel into a listening state registered application
// @Tags Channel
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /channel/update [post]
func (s *HighwayServer) UpdateChannelHTTP(c *gin.Context) {
	// Unmarshal the request body
	var req ct.MsgUpdateChannel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
	}

	// Update the channel
	resp, err := s.GRPCClient.UpdateChannel(s.ctx, &req)
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

// ListenChannel listens to a channel.
func (s *HighwayServer) ListenChannel(req *v1.MsgListenChannel, stream v1.Highway_ListenChannelServer) error {
	// Initialize the channel listener
	opChan := make(chan *ctv1.ChannelMessage)
	ch, ok := s.channels[req.GetDid()]
	if !ok {
		return ErrInvalidQuery
	}
	go ch.Listen(opChan)

	// Listen to the channel
	for {
		select {
		case op := <-opChan:
			// Send the operation
			err := stream.Send(&ct.ChannelMessage{
				PeerDid:  op.GetPeerDid(),
				Did:      op.GetDid(),
				Object:   otv1.NewObjectDocToBuf(op.Object),
				Metadata: op.GetMetadata(),
			})
			if err != nil {
				return err
			}
		case <-stream.Context().Done():
			return nil
		}
	}
}

// @Summary Listen Channel
// @Schemes
// @Description ListenChannel puts a Channel into a listening state registered application
// @Tags Channel
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /channel/listen [post]
func (s *HighwayServer) ListenChannelHTTP(c *gin.Context) {
	// Unmarshal the request body
	var req v1.MsgListenChannel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
	}

	// Setup concurrent stream
	opChan := make(chan *ctv1.ChannelMessage)
	ch, ok := s.channels[req.GetDid()]
	if !ok {
		c.JSON(http.StatusBadRequest, t.ErrRequestBody.Error())
	}
	go ch.Listen(opChan)

	// Listen to the channel
	c.Stream(func(io.Writer) bool {
		// Get the next operation
		op, ok := <-opChan
		if !ok {
			return false
		}

		// Create ChannelMessage
		msg := &ct.ChannelMessage{
			PeerDid:  op.GetPeerDid(),
			Did:      op.GetDid(),
			Object:   otv1.NewObjectDocToBuf(op.Object),
			Metadata: op.GetMetadata(),
		}

		// Marshal the proto message
		data, err := proto.Marshal(msg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return false
		}

		// Send the message
		c.Writer.Write(data)
		return true
	})
}

// DeactivateChannel disables a Channel for a registered application
func (s *HighwayServer) DeactivateChannel(ctx context.Context, req *ct.MsgDeactivateChannel) (*ct.MsgDeactivateChannelResponse, error) {
	resp, err := s.Cosmos.BroadcastDeactivateChannel(ctv1.NewMsgDeactivateChannelFromBuf(req))
	if err != nil {
		return nil, err
	}
	log.Println(resp.String())
	return &ct.MsgDeactivateChannelResponse{
		Code:    resp.Code,
		Message: resp.Message,
	}, nil
}

// @Summary Deactivate Channel
// @Schemes
// @Description DeactivateChannel disables a Channel for a registered application
// @Tags Channel
// @Produce json
// @Success      200  {string}  message
// @Failure      500  {string}  message
// @Router /channel/deactivate [post]
func (s *HighwayServer) DeactivateChannelHTTP(c *gin.Context) {
	// Unmarshal the request body
	var req ct.MsgDeactivateChannel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": t.ErrRequestBody.Error(),
		})
	}

	// Deactivate the bucket
	resp, err := s.GRPCClient.DeactivateChannel(s.ctx, &req)
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
