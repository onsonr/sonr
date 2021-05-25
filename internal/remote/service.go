package remote

import (
	"context"

	rpc "github.com/libp2p/go-libp2p-gorpc"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ExchangeArgs is Peer protobuf
type RemoteServiceArgs struct {
	Lobby  []byte
	Peer   []byte
	Invite []byte
	Link   []byte
}

// ExchangeResponse is also Peer protobuf
type RemoteServiceResponse struct {
	Reply []byte
	Peer  []byte
}

// Service Struct
type RemoteService struct {
	// Current Data
	call  ClientCallback
	lobby *md.Lobby
	user  *md.User

	respCh chan *md.AuthReply
	invite *md.AuthInvite
}

// ^ Flat: Handles User sent AuthInvite Response on FlatMode ^
func (tm *RemoteManager) Flat(id peer.ID, inv *md.AuthInvite) error {
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(inv)
	if err != nil {
		return err
	}

	// Initialize Data
	rpcClient := rpc.NewClient(tm.host.Host, K_SERVICE_PID)
	var reply RemoteServiceResponse
	var args RemoteServiceArgs
	args.Invite = msgBytes

	// Call to Peer
	err = rpcClient.Call(id, "TopicService", "FlatWith", args, &reply)
	if err != nil {
		return err
	}

	tm.topicHandler.OnReply(id, reply.Reply)
	return nil
}

// ^ Calls Invite on Remote Peer for Flat Mode and makes Direct Response ^ //
func (ts *RemoteService) FlatWith(ctx context.Context, args RemoteServiceArgs, reply *RemoteServiceResponse) error {
	// Received Message
	receivedMessage := md.AuthInvite{}
	err := proto.Unmarshal(args.Invite, &receivedMessage)
	if err != nil {
		return err
	}

	// Set Current Message and send Callback
	ts.invite = &receivedMessage
	ts.call.OnInvite(args.Invite)

	// Sign Contact Reply
	resp := ts.user.SignFlatReply(receivedMessage.GetFrom())

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(resp)
	if err != nil {
		return err
	}

	reply.Reply = msgBytes
	return nil
}

// ^ Calls Invite on Remote Peer ^ //
func (tm *RemoteManager) Exchange(id peer.ID, peerBuf []byte, lobBuf []byte) error {
	// Initialize RPC
	exchClient := rpc.NewClient(tm.host.Host, K_SERVICE_PID)
	var reply RemoteServiceResponse
	var args RemoteServiceArgs

	// Set Args
	args.Lobby = lobBuf
	args.Peer = peerBuf

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
	tm.lobby.Add(remotePeer)
	tm.Refresh()
	return nil
}

// ^ Calls Invite on Remote Peer ^ //
func (ts *RemoteService) ExchangeWith(ctx context.Context, args RemoteServiceArgs, reply *RemoteServiceResponse) error {
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
	ts.lobby.Sync(remoteLobbyRef, remotePeer)
	ts.call.OnRefresh(ts.lobby)

	// Set Message data and call done
	buf, err := ts.user.Peer.Buffer()
	if err != nil {
		return err
	}
	reply.Peer = buf
	return nil
}

// ^ Invite: Handles User sent AuthInvite Response ^
func (tm *RemoteManager) Invite(id peer.ID, inv *md.AuthInvite) error {
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(inv)
	if err != nil {
		return err
	}

	// Initialize Data
	rpcClient := rpc.NewClient(tm.host.Host, K_SERVICE_PID)
	var reply RemoteServiceResponse
	var args RemoteServiceArgs
	args.Invite = msgBytes

	// Call to Peer
	done := make(chan *rpc.Call, 1)
	err = rpcClient.Go(id, "TopicService", "InviteWith", args, &reply, done)

	// Await Response
	call := <-done
	if call.Error != nil {
		return err
	}
	tm.topicHandler.OnReply(id, reply.Reply)
	return nil
}

// ^ Calls Invite on Remote Peer ^ //
func (ts *RemoteService) InviteWith(ctx context.Context, args RemoteServiceArgs, reply *RemoteServiceResponse) error {
	// Received Message
	receivedMessage := md.AuthInvite{}
	err := proto.Unmarshal(args.Invite, &receivedMessage)
	if err != nil {
		return err
	}

	// Set Current Message and send Callback
	ts.invite = &receivedMessage
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
		reply.Reply = msgBytes
		ctx.Done()
		return nil
		// Context is Done
	case <-ctx.Done():
		return nil
	}
}

// ^ RespondToInvite to an Invitation ^ //
func (n *RemoteManager) RespondToInvite(rep *md.AuthReply) {
	// Prepare Transfer
	if rep.Decision {
		n.topicHandler.OnResponded(n.service.invite)
	}

	// @ Pass Contact Back
	// Send to Channel
	n.service.respCh <- rep
}
