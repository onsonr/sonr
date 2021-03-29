//nolint
package topic

import (
	"context"

	rpc "github.com/libp2p/go-libp2p-gorpc"

	"github.com/libp2p/go-libp2p-core/peer"
	md "github.com/sonr-io/core/internal/models"
	se "github.com/sonr-io/core/internal/session"
	us "github.com/sonr-io/core/internal/user"
)

// Service Struct
type TopicService struct {
	// Current Data
	call  TopicHandler
	lobby *md.Lobby
	peer  *md.Peer

	respCh chan *md.AuthReply
	invite *md.AuthInvite
}

// ^ Calls Invite on Remote Peer ^ //
func (tm *TopicManager) Exchange(id peer.ID, l *md.Lobby, p *md.Peer) error {
	// Initialize RPC
	exchClient := rpc.NewClient(tm.host.Host, K_SERVICE_PID)
	var reply md.TopicServiceResponse
	var args md.TopicServiceArgs

	// Set Args
	args.Lobby = l
	args.Peer = p

	// Call to Peer
	err := exchClient.Call(id, "TopicService", "ExchangeWith", args, &reply)
	if err != nil {
		return err
	}
	// Update Peer with new data
	tm.Lobby.Add(reply.Peer)
	tm.Refresh()
	return nil
}

// ^ Calls Invite on Remote Peer ^ //
func (ts *TopicService) ExchangeWith(ctx context.Context, args md.TopicServiceArgs, reply *md.TopicServiceResponse) error {
	// Update Peers with Lobby
	ts.lobby.Sync(args.Lobby, args.Peer)
	ts.call.OnRefresh(ts.lobby)

	// Set Message data and call done
	reply.Peer = ts.peer
	return nil
}

// ^ Invite: Handles User sent AuthInvite Response ^
func (tm *TopicManager) Invite(id peer.ID, inv *md.AuthInvite, session *se.Session) error {
	// Initialize Data
	rpcClient := rpc.NewClient(tm.host.Host, K_SERVICE_PID)
	var reply md.TopicServiceResponse
	var args md.TopicServiceArgs
	args.Invite = inv

	// Call to Peer
	done := make(chan *rpc.Call, 1)
	err := rpcClient.Go(id, "TopicService", "InviteWith", args, &reply, done)

	// Await Response
	call := <-done
	if call.Error != nil {
		return err
	}
	tm.topicHandler.OnReply(id, reply.Reply, session)
	return nil
}

// ^ Calls Invite on Remote Peer ^ //
func (ts *TopicService) InviteWith(ctx context.Context, args md.TopicServiceArgs, reply *md.TopicServiceResponse) error {
	// Set Current Message
	ts.invite = args.Invite

	// Send Callback
	ts.call.OnInvite(args.Invite)

	// Hold Select for Invite Type
	select {
	// Received Auth Channel Message
	case m := <-ts.respCh:
		// Set Message data and call done
		reply.Reply = m
		ctx.Done()
		return nil
		// Context is Done
	case <-ctx.Done():
		return nil
	}
}

// ^ RespondToInvite to an Invitation ^ //
func (n *TopicManager) RespondToInvite(decision bool, fs *us.FileSystem, p *md.Peer, c *md.Contact) {
	// Prepare Transfer
	if decision {
		n.topicHandler.OnResponded(n.service.invite, p, fs)
	}

	// @ Pass Contact Back
	if n.service.invite.Payload == md.Payload_CONTACT {
		// Create Accept Response
		resp := p.SignReplyWithContact(c)
		// Send to Channel
		n.service.respCh <- resp
	} else {
		// Create Accept Response
		resp := p.SignReply()

		// Send to Channel
		n.service.respCh <- resp
	}
}
