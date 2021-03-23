package service

import (
	"context"

	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ****************** //
// ** GRPC Service ** //
// ****************** //
// ExchangeArgs is Peer protobuf
type ExchangeArgs struct {
	Lobby []byte
	Peer  []byte
}

// ExchangeResponse is also Peer protobuf
type ExchangeResponse struct {
	Data []byte
}

// Service Struct
type ExchangeService struct {
	getUser   md.ReturnPeer
	syncLobby md.SyncLobby
}

// ^ Calls Invite on Remote Peer ^ //
func (ps *ExchangeService) ExchangeWith(ctx context.Context, args ExchangeArgs, reply *ExchangeResponse) error {
	// Peer Data
	remoteLobbyRef := &md.Lobby{}
	err := proto.Unmarshal(args.Lobby, remoteLobbyRef)
	if err != nil {
		return err
	}

	remotePeer := &md.Peer{}
	err = proto.Unmarshal(args.Peer, remotePeer)
	if err != nil {
		return err
	}

	// Update Peers with Lobby
	ps.syncLobby(remoteLobbyRef, remotePeer)

	// Return User Peer
	userPeer := ps.getUser()
	replyData, err := proto.Marshal(userPeer)
	if err != nil {
		return err
	}

	// Set Message data and call done
	reply.Data = replyData
	return nil
}
