package topic

import (
	"context"

	rpc "github.com/libp2p/go-libp2p-gorpc"

	"github.com/libp2p/go-libp2p-core/peer"
	se "github.com/sonr-io/core/internal/session"
	md "github.com/sonr-io/core/pkg/models"
	us "github.com/sonr-io/core/pkg/user"
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
	// Current Data
	call  TopicHandler
	lobby *md.Lobby
	peer  *md.Peer

	respCh chan *md.AuthReply
	invite *md.AuthInvite
}

// ^ Direct: Handles User sent AuthInvite Response on FlatMode ^
func (tm *TopicManager) Direct(id peer.ID, inv *md.AuthInvite) error {
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(inv)
	if err != nil {
		return err
	}

	// Initialize Data
	rpcClient := rpc.NewClient(tm.host.Host, K_SERVICE_PID)
	var reply TopicServiceResponse
	var args TopicServiceArgs
	args.Invite = msgBytes

	// Call to Peer
	err = rpcClient.Call(id, "TopicService", "DirectWith", args, &reply)
	if err != nil {
		return err
	}

	tm.topicHandler.OnReply(id, reply.InvReply, nil)
	return nil
}

// ^ Calls Invite on Remote Peer for Flat Mode and makes Direct Response ^ //
func (ts *TopicService) DirectWith(ctx context.Context, args TopicServiceArgs, reply *TopicServiceResponse) error {
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
	resp := ts.peer.SignReplyWithContact(ts.call.GetContact(), true, nil)

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(resp)
	if err != nil {
		return err
	}

	reply.InvReply = msgBytes
	return nil
}

// ^ Calls Invite on Remote Peer ^ //
func (tm *TopicManager) Exchange(id peer.ID, peerBuf []byte, lobBuf []byte) error {
	// Initialize RPC
	exchClient := rpc.NewClient(tm.host.Host, K_SERVICE_PID)
	var reply TopicServiceResponse
	var args TopicServiceArgs

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
	tm.Lobby.Add(remotePeer)
	tm.Refresh()
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
	ts.lobby.Sync(remoteLobbyRef, remotePeer)
	ts.call.OnRefresh(ts.lobby)

	// Set Message data and call done
	buf, err := ts.peer.Buffer()
	if err != nil {
		return err
	}
	reply.Peer = buf
	return nil
}

// ^ Invite: Handles User sent AuthInvite Response ^
func (tm *TopicManager) Invite(id peer.ID, inv *md.AuthInvite, session *se.Session) error {
	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(inv)
	if err != nil {
		return err
	}

	// Initialize Data
	rpcClient := rpc.NewClient(tm.host.Host, K_SERVICE_PID)
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
	tm.topicHandler.OnReply(id, reply.InvReply, session)
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
		reply.InvReply = msgBytes
		ctx.Done()
		return nil
		// Context is Done
	case <-ctx.Done():
		return nil
	}
}

// ^ RespondToInvite to an Invitation ^ //
func (n *TopicManager) RespondToInvite(req *md.RespondRequest, fs *us.FileSystem, p *md.Peer, c *md.Contact, ) {
	// Prepare Transfer
	if req.Decision {
		n.topicHandler.OnResponded(n.service.invite, p, fs)
	}

	// @ Pass Contact Back
	if n.service.invite.Payload == md.Payload_CONTACT {
		// Create Accept Response
		resp := p.SignReplyWithContact(c, n.service.invite.IsFlat, req)
		// Send to Channel
		n.service.respCh <- resp
	} else {
		// Create Accept Response
		resp := p.SignReply(req.Decision, req)

		// Send to Channel
		n.service.respCh <- resp
	}
}
