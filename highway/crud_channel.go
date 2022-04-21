package highway

import (
	context "context"
	"errors"

	ctv1 "github.com/sonr-io/blockchain/x/channel/types"
	ot "github.com/sonr-io/blockchain/x/object/types"
	"github.com/sonr-io/core/channel"
	ct "go.buf.build/grpc/go/sonr-io/blockchain/channel"
)

// CreateChannel creates a new channel.
func (s *HighwayServer) CreateChannel(ctx context.Context, req *ct.MsgCreateChannel) (*ct.MsgCreateChannelResponse, error) {
	// Verify that channel fields are not nil
	if req.GetObjectToRegister().GetFields() == nil {
		return nil, errors.New("object to register must have fields")
	}

	// Create ctv1 message to broadcast
	var fields map[string]*ot.TypeField
	fields = make(map[string]*ot.TypeField, len(req.GetObjectToRegister().GetFields()))
	for _, f := range req.GetObjectToRegister().GetFields() {
		fields[f.GetLabel()] = &ot.ObjectField{
			Label: f.GetLabel(),
			Type:  ot.ObjectFieldType(f.GetType()),
			Did:   f.GetDid(),
		}
	}

	// Build Transaction
	tx := &ctv1.MsgCreateChannel{
		Creator: req.GetCreator(),
		Label:   req.GetLabel(),
		ObjectToRegister: &ot.ObjectDoc{
			Label:       req.GetObjectToRegister().GetLabel(),
			Description: req.GetObjectToRegister().GetDescription(),
			Did:         req.GetObjectToRegister().GetDid(),
			BucketDid:   req.GetObjectToRegister().GetBucketDid(),
			Fields:      fields,
		},
	}

	// Broadcast the message
	res, err := s.cosmos.BroadcastCreateChannel(tx)
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
	return nil, ErrMethodUnimplemented
}

// UpdateChannel updates a channel.
func (s *HighwayServer) UpdateChannel(ctx context.Context, req *ct.MsgUpdateChannel) (*ct.MsgUpdateChannelResponse, error) {
	return nil, ErrMethodUnimplemented
}

// // ListenChannel listens to a channel.
// func (s *HighwayServer) ListenChannel(req *ct.ListenChannelRequest, stream v1.HighwayService_ListenChannelServer) error {
// 	// Find channel by DID
// 	ch, ok := s.channels[req.GetDid()]
// 	if !ok {
// 		return ErrInvalidQuery
// 	}

// 	// Listen to the channel
// 	chListen, err := ch.Listen()
// 	if err != nil {
// 		return err
// 	}

// 	// Listen to the channel
// 	for {
// 		select {
// 		case msg := <-chListen:
// 			// Send peer to client
// 			if err := stream.Send(&v1.ListenChannelResponse{
// 				Message: msg.GetData(),
// 			}); err != nil {
// 				return err
// 			}
// 		case <-stream.Context().Done():
// 			return nil
// 		}
// 	}
// }
