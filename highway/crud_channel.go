package highway

import (
	context "context"

	"github.com/sonr-io/core/channel"
	ct "go.buf.build/grpc/go/sonr-io/sonr/channel"
)

// CreateChannel creates a new channel.
func (s *HighwayServer) CreateChannel(ctx context.Context, req *ct.MsgCreateChannel) (*ct.MsgCreateChannelResponse, error) {
	// Create the Channel
	ch, err := channel.New(ctx, s.node, req.Name)
	if err != nil {
		return nil, err
	}

	// Add to the list of Channels
	s.channels[req.Name] = ch
	return nil, ErrMethodUnimplemented
}

// ReadChannel reads a channel.
func (s *HighwayServer) ReadChannel(ctx context.Context, req *ct.MsgReadChannel) (*ct.MsgReadChannelResponse, error) {
	// Find channel by DID
	ch, ok := s.channels[req.GetDid()]
	if !ok {
		return nil, ErrInvalidQuery
	}

	// Read the channel
	peers := ch.Read()
	logger.Debugf("Read %d peers from channel %s", len(peers), peers)
	return &ct.MsgReadChannelResponse{
		// Peers: peers,
	}, nil
}

// UpdateChannel updates a channel.
func (s *HighwayServer) UpdateChannel(ctx context.Context, req *ct.MsgUpdateChannel) (*ct.MsgUpdateChannelResponse, error) {
	return nil, ErrMethodUnimplemented
}

// DeleteChannel deletes a channel.
func (s *HighwayServer) DeleteChannel(ctx context.Context, req *ct.MsgDeleteChannel) (*ct.MsgDeleteChannelResponse, error) {
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
