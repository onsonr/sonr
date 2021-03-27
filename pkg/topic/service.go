package topic

import (
	"context"
	"time"

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
	InvReply []byte
	Peer     []byte
}

// Service Struct
type TopicService struct {
	// Methods
	GetUser   dt.ReturnPeer
	SyncLobby dt.SyncLobby

	// Current Data
	call   TopicHandler
	respCh chan *md.AuthReply
	invite *md.AuthInvite
}

// ^ Calls Invite on Remote Peer ^ //
func (tm *TopicManager) Exchange(id peer.ID, pb []byte) error {
	// Initialize RPC
	exchClient := rpc.NewClient(tm.host, tm.protocol)
	var reply TopicServiceResponse
	var args TopicServiceArgs

	// Set Args
	args.Lobby = tm.Lobby.Buffer()
	args.Peer = pb

	// Call to Peer
	err := exchClient.Call(id, "TopicService", "ExchangeWith", args, &reply)
	if err != nil {
		return err
	}

	// Received Message
	remotePeer := &md.Peer{}
	err = proto.Unmarshal(reply.Peer, remotePeer)

	// Send Error
	if err != nil {
		return err
	}

	// Update Peer with new data
	tm.Lobby.Add(remotePeer)
	return nil
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
	reply.Peer = replyData
	return nil
}

// ^ Invite: Handles User sent AuthInvite Response ^
func (tm *TopicManager) Invite(id peer.ID, inv *md.AuthInvite, p *md.Peer, cf *sf.FileItem) error {
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(inv)
	if err != nil {
		return err
	}

	// Initialize Data
	rpcClient := rpc.NewClient(tm.host, tm.protocol)
	var reply TopicServiceResponse
	var args TopicServiceArgs
	args.Invite = msgBytes

	// Call to Peer
	done := make(chan *rpc.Call, 1)
	err = rpcClient.Go(id, "TopicService", "InviteWith", args, &reply, done)

	// Await Response
	call := <-done
	if call.Error != nil {
		return err
	}
	tm.callback.OnReply(id, p, reply.InvReply)
	return nil
}

// ^ Calls Invite on Remote Peer ^ //
func (ts *TopicService) InviteWith(ctx context.Context, args TopicServiceArgs, reply *TopicServiceResponse) error {
	// Received Message
	receivedMessage := md.AuthInvite{}
	err := proto.Unmarshal(args.Invite, &receivedMessage)
	if err != nil {
		return err
	}

	// Set Current Message
	ts.invite = &receivedMessage

	// Send Callback
	ts.call.OnInvite(args.Invite)

	// Hold Select for Invite Type
	select {
	// Received Auth Channel Message
	case m := <-ts.respCh:
		// Convert Protobuf to bytes
		msgBytes, err := proto.Marshal(m)
		if err != nil {
			return err
		}

		// Set Message data and call done
		reply.InvReply = msgBytes
		ctx.Done()
		return nil
		// Context is Done
	case <-ctx.Done():
		return nil
	}
}

// ^ RespondToInvite to an Invitation ^ //
func (n *TopicManager) RespondToInvite(decision bool, fs *sf.FileSystem, p *md.Peer, c *md.Contact) {
	// Check Decision
	if decision {
		n.callback.OnReceiveTransfer(n.service.invite, fs)
	}

	// @ Pass Contact Back
	if n.service.invite.Payload == md.Payload_CONTACT {
		// Create Accept Response
		resp := &md.AuthReply{
			IsRemote: n.service.invite.IsRemote,
			From:     p,
			Type:     md.AuthReply_Contact,
			Card: &md.TransferCard{
				// SQL Properties
				Payload:  md.Payload_CONTACT,
				Received: int32(time.Now().Unix()),
				Preview:  p.Profile.Picture,
				Platform: p.Platform,

				// Transfer Properties
				Status: md.TransferCard_REPLY,

				// Owner Properties
				Username:  p.Profile.Username,
				FirstName: p.Profile.FirstName,
				LastName:  p.Profile.LastName,

				// Data Properties
				Contact: c,
			},
		}
		// Send to Channel
		n.service.respCh <- resp
	} else {
		// Create Accept Response
		resp := &md.AuthReply{
			IsRemote: n.service.invite.IsRemote,
			From:     p,
			Type:     md.AuthReply_Transfer,
			Decision: decision,
		}
		// Send to Channel
		n.service.respCh <- resp
	}
}
