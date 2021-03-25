package topic

import (
	"context"
	"log"

	rpc "github.com/libp2p/go-libp2p-gorpc"

	"github.com/libp2p/go-libp2p-core/peer"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	dt "github.com/sonr-io/core/pkg/data"
	"google.golang.org/protobuf/proto"
)

// ExchangeArgs is Peer protobuf
type TopicServiceArgs struct {
	Lobby  []byte
	Peer   []byte
	Invite []byte
}

// ExchangeResponse is also Peer protobuf
type TopicServiceResponse struct {
	Data []byte
}

// Service Struct
type TopicService struct {
	// Methods
	GetUser   dt.ReturnPeer
	SyncLobby dt.SyncLobby

	// Current Data
	call   dt.NodeCallback
	respCh chan *md.AuthReply
	invite *md.AuthInvite
}

// ^ Calls Invite on Remote Peer ^ //
func (tm *TopicManager) Exchange(id peer.ID, pb []byte) {
	// Initialize RPC
	exchClient := rpc.NewClient(tm.host, tm.exchProtocol)
	var reply TopicServiceResponse
	var args TopicServiceArgs

	// Set Args
	args.Lobby = tm.Lobby.Buffer()
	args.Peer = pb

	// Call to Peer
	err := exchClient.Call(id, "TopicService", "ExchangeWith", args, &reply)
	if err != nil {
		tm.call.Error(err, "Exchange")
	}

	// Received Message
	remotePeer := &md.Peer{}
	err = proto.Unmarshal(reply.Data, remotePeer)

	// Send Error
	if err != nil {
		tm.call.Error(err, "Exchange")
	}

	// Update Peer with new data
	tm.Lobby.Add(remotePeer)
}

// ^ Calls Invite on Remote Peer ^ //
func (ts *TopicService) ExchangeWith(ctx context.Context, args TopicServiceArgs, reply *TopicServiceResponse) error {
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
	ts.SyncLobby(remoteLobbyRef, remotePeer)

	// Return User Peer
	userPeer := ts.GetUser()
	replyData, err := proto.Marshal(userPeer)
	if err != nil {
		return err
	}

	// Set Message data and call done
	reply.Data = replyData
	return nil
}

// ^ Invite: Handles User sent AuthInvite Response ^
func (tm *TopicManager) Invite(id peer.ID, inv *md.AuthInvite, p *md.Peer, cf *sf.ProcessedFile) {
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(inv)
	if err != nil {
		log.Println(err)
	}

	// Initialize Data
	rpcClient := rpc.NewClient(tm.host, tm.authProtocol)
	var reply TopicServiceResponse
	var args TopicServiceArgs
	args.Invite = msgBytes

	// Call to Peer
	done := make(chan *rpc.Call, 1)
	err = rpcClient.Go(id, "TopicService", "InviteWith", args, &reply, done)

	// Await Response
	call := <-done
	if call.Error != nil {
		tm.call.Error(err, "Request")
	}

	// Send Callback and Reset
	tm.call.Responded(reply.Data)

	// Check for File
	tm.cont(id, p, cf, reply.Data)
}

// ^ Calls Invite on Remote Peer ^ //
func (ts *TopicService) InviteWith(ctx context.Context, args TopicServiceArgs, reply *TopicServiceResponse) error {
	// Received Message
	receivedMessage := md.AuthInvite{}
	err := proto.Unmarshal(args.Invite, &receivedMessage)
	if err != nil {
		log.Println(err)
		return err
	}

	// Set Current Message
	ts.invite = &receivedMessage

	// Send Callback
	ts.call.Invited(args.Invite)

	// Hold Select for Invite Type
	select {
	// Received Auth Channel Message
	case m := <-ts.respCh:

		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(m)
		if err != nil {
			log.Println(err)
			return err
		}

		// Set Message data and call done
		reply.Data = msgBytes
		ctx.Done()
		return nil
		// Context is Done
	case <-ctx.Done():
		return nil
	}
}
