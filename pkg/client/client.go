package client

import (
	"context"
	"errors"
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	net "github.com/sonr-io/core/internal/host"
	txt "github.com/sonr-io/core/internal/textile"
	tpc "github.com/sonr-io/core/internal/topic"
	md "github.com/sonr-io/core/pkg/models"
)

// Interface: Main Client handles Networking/Identity/Streams
type Client interface {
	// Client Methods
	Connect(cr *md.ConnectionRequest, keys *md.KeyPair) *md.SonrError
	Bootstrap() (*tpc.Manager, *md.SonrError)
	Invite(invite *md.InviteRequest, t *tpc.Manager) *md.SonrError
	Mail(mr *md.MailRequest) *md.SonrError
	Update(t *tpc.Manager) *md.SonrError
	Close(t *tpc.Manager)

	// Topic Callbacks
	OnEvent(*md.LobbyEvent)
	OnRefresh(*md.Lobby)
	OnInvite([]byte)
	OnReply(id peer.ID, data []byte)
	OnResponded(inv *md.InviteRequest)
}

// Struct: Main Client handles Networking/Identity/Streams
type client struct {
	Client

	// Properties
	ctx     context.Context
	call    md.Callback
	user    *md.User
	session *md.Session
	request *md.ConnectionRequest

	// References
	Host    net.HostNode
	Textile txt.TextileNode
}

// ^ NewClient Initializes Node with Router ^
func NewClient(ctx context.Context, u *md.User, call md.Callback) Client {
	// Returns Storj Enabled Client
	return &client{
		ctx:  ctx,
		call: call,
		user: u,
	}
}

// @ Connects Host Node from Private Key
func (c *client) Connect(cr *md.ConnectionRequest, keys *md.KeyPair) *md.SonrError {
	// Set Request
	c.request = cr

	// Set Host
	hn, err := net.NewHost(c.ctx, cr, keys)
	if err != nil {
		return err
	}

	// Get MultiAddrs
	maddr, err := hn.MultiAddr()
	if err != nil {
		return err
	}

	// Set Peer
	err = c.user.NewPeer(hn.ID(), maddr)
	if err != nil {
		return err
	}

	// Set Host
	c.Host = hn

	// Check Textile Option
	if c.request.GetTextileOptions().GetEnabled() {
		// Create Textile Node
		txtNode, err := txt.NewTextile(c.Host, c.request, keys)
		if err != nil {
			c.call.OnError(err)
			return nil
		}

		// Set Node
		c.Textile = txtNode
	}
	return nil
}

// @ Begins Bootstrapping HostNode
func (c *client) Bootstrap() (*tpc.Manager, *md.SonrError) {
	// Bootstrap Host
	err := c.Host.Bootstrap()
	if err != nil {
		return nil, err
	}

	// Join Local
	if t, err := tpc.NewLocal(c.ctx, c.Host, c.user, c.user.GetRouter().LocalTopic, c); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// @ Invite Processes Data and Sends Invite to Peer
func (n *client) Invite(invite *md.InviteRequest, t *tpc.Manager) *md.SonrError {
	// Check for Peer
	if t.HasPeer(invite.To.Id.Peer) {
		// Initialize Session if transfer
		if invite.IsPayloadFile() {
			// Start New Session
			n.session = md.NewOutSession(n.user, invite, n.call)
		}

		// Get PeerID and Check error
		id, err := t.FindPeerInTopic(invite.To.Id.Peer)
		if err != nil {
			return md.NewPeerFoundError(err, invite.GetTo().GetId().GetPeer())
		}

		// Run Routine
		go func(inv *md.InviteRequest) {
			// Send Invite
			err = t.Invite(id, inv)
			if err != nil {
				n.call.OnError(md.NewError(err, md.ErrorMessage_TOPIC_RPC))
			}
		}(invite)
	} else {
		return md.NewErrorWithType(md.ErrorMessage_PEER_NOT_FOUND_INVITE)
	}
	return nil
}

// @ Handle a MailRequest from Node
func (c *client) Mail(mr *md.MailRequest) *md.SonrError {
	if mr.Method == md.MailRequest_READ {

	} else if mr.Method == md.MailRequest_SEND {

	}
	return md.NewError(errors.New("Invalid MailRequest Method"), md.ErrorMessage_HOST_TEXTILE)
}

// @ Update proximity/direction and Notify Lobby
func (n *client) Update(t *tpc.Manager) *md.SonrError {
	// Inform Lobby
	if err := t.Publish(n.user.Peer.NewUpdateEvent()); err != nil {
		return md.NewError(err, md.ErrorMessage_TOPIC_UPDATE)
	}
	return nil
}

// @ Close Ends All Network Communication
func (n *client) Close(t *tpc.Manager) {
	// Inform Lobby
	if err := t.Publish(n.user.Peer.NewUpdateEvent()); err != nil {
		log.Println(md.NewError(err, md.ErrorMessage_TOPIC_UPDATE))
	}
	n.Host.Close()
}
