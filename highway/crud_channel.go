package highway

import (
	context "context"
	"errors"
	"log"

	ctv1 "github.com/sonr-io/blockchain/x/channel/types"
	otv1 "github.com/sonr-io/blockchain/x/object/types"
	"github.com/sonr-io/core/channel"
	ct "go.buf.build/sonr-io/grpc-gateway/sonr-io/blockchain/channel"
	v1 "go.buf.build/sonr-io/grpc-gateway/sonr-io/core/highway/v1"
)

// CreateChannel creates a new channel.
func (s *HighwayServer) CreateChannel(ctx context.Context, req *ct.MsgCreateChannel) (*ct.MsgCreateChannelResponse, error) {
	// Verify that channel fields are not nil
	if req.GetObjectToRegister().GetFields() == nil {
		return nil, errors.New("object to register must have fields")
	}

	// Broadcast the message
	res, err := s.cosmos.BroadcastCreateChannel(ctv1.NewMsgCreateChannelFromBuf(req))
	if err != nil {
		return nil, err
	}

	// Create the Channel
	ch, err := channel.New(ctx, s.node, res.GetHowIs().GetChannel())
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

// UpdateChannel updates a channel.
func (s *HighwayServer) UpdateChannel(ctx context.Context, req *ct.MsgUpdateChannel) (*ct.MsgUpdateChannelResponse, error) {
	resp, err := s.cosmos.BroadcastUpdateChannel(ctv1.NewMsgUpdateChannelFromBuf(req))
	if err != nil {
		return nil, err
	}
	log.Println(resp.String())
	return &ct.MsgUpdateChannelResponse{
		Code:    resp.Code,
		Message: resp.Message,
	}, nil
}

// ListenChannel listens to a channel.
func (s *HighwayServer) ListenChannel(req *v1.MsgListenChannel, stream v1.Highway_ListenChannelServer) error {
	// Find channel by DID
	ch, ok := s.channels[req.GetDid()]
	if !ok {
		return ErrInvalidQuery
	}

	// Listen to the channel
	for {
		select {
		case msg := <-ch.Listen():
			// Send peer to client
			if err := stream.Send(&ct.ChannelMessage{
				PeerDid:  msg.PeerDid,
				Did:      msg.Did,
				Object:   otv1.NewObjectDocToBuf(msg.Object),
				Metadata: msg.Metadata,
			}); err != nil {
				return err
			}
		case <-stream.Context().Done():
			return nil
		}
	}
}
