package lobby

import (
	"context"

	// gorpc "github.com/libp2p/go-libp2p-gorpc"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ****************** //
// ** GRPC Service ** //
// ****************** //
// Argument is ExchangeMessage protobuf
type ExchangeArgs struct {
	Data []byte
}

// Reply is also ExchangeMessage protobuf
type ExchangeReply struct {
	Data []byte
}

// Service Struct
type ExchangeService struct {
	// Current Data
	peerData *md.Peer
}

// ^ Calls Exchange on Remote Peer ^ //
func (ex *ExchangeService) Exchange(ctx context.Context, args ExchangeArgs, reply *ExchangeReply) error {
	// Marshal Peer Data
	bytes, err := proto.Marshal(ex.peerData)
	if err != nil {
		return err
	}

	// Sends Peer Data
	reply.Data = bytes
	return nil
}
