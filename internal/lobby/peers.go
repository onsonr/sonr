package lobby

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ****************** //
// ** GRPC Service ** //
// ****************** //
// PeerArgs is Peer protobuf
type PeerArgs struct {
	Data []byte
}

// PeerResponse is also Peer protobuf
type PeerResponse struct {
	Data []byte
}

// Service Struct
type PeerService struct {
	getUser    md.ReturnPeer
	updatePeer md.UpdatePeer
}

// ^ Calls Invite on Remote Peer ^ //
func (ps *PeerService) ExchangeWith(ctx context.Context, args PeerArgs, reply *PeerResponse) error {
	// Peer Data
	remotePeer := &md.Peer{}
	err := proto.Unmarshal(args.Data, remotePeer)
	if err != nil {
		return err
	}

	// Update Peers
	ps.updatePeer(remotePeer)

	// Set Current Message
	userPeer := ps.getUser()

	// Convert Protobuf to bytes
	replyData, err := proto.Marshal(userPeer)
	if err != nil {
		return err
	}

	// Set Message data and call done
	reply.Data = replyData
	return nil
}

// ^ Calls Invite on Remote Peer ^ //
func (lob *Lobby) Exchange(id peer.ID) {
	// Get Peer Data
	userPeer := lob.call.Peer()
	msgBytes, err := proto.Marshal(userPeer)
	if err != nil {
		lob.call.Error(err)
	}

	// Initialize RPC
	rpcClient := gorpc.NewClient(lob.host, protocol.ID("/sonr/lobby/peers"))
	var reply PeerResponse
	var args PeerArgs
	args.Data = msgBytes

	// Call to Peer
	done := make(chan *gorpc.Call, 1)
	err = rpcClient.Go(id, "PeerService", "ExchangeWith", args, &reply, done)

	// Await Response
	call := <-done
	if call.Error != nil {
		lob.call.Error(err)
	}

	// Received Message
	remotePeer := &md.Peer{}
	err = proto.Unmarshal(reply.Data, remotePeer)
	if err != nil {
		// Send Error
		lob.call.Error(err)
	}

	// Update Peers
	lob.updatePeer(remotePeer)
}
