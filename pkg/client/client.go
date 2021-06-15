package client

import (
	"context"
	"errors"

	tpc "github.com/sonr-io/core/internal/topic"
	md "github.com/sonr-io/core/pkg/models"

	// Local
	net "github.com/sonr-io/core/internal/host"
)

// Struct: Main Client handles Networking/Identity/Streams
type Client struct {
	tpc.ClientHandler

	// Properties
	ctx     context.Context
	call    md.Callback
	user    *md.User
	session *md.Session

	// References
	Host net.HostNode
}

// ^ NewClient Initializes Node with Router ^
func NewClient(ctx context.Context, u *md.User, call md.Callback) *Client {
	// Returns Storj Enabled Client
	return &Client{
		ctx:  ctx,
		call: call,
		user: u,
	}
}

// @ Connects Host Node from Private Key
func (c *Client) Connect(api *md.APIKeys, keys *md.KeyPair) *md.SonrError {
	// Set Host
	hn, err := net.NewHost(c.ctx, c.user.GetRouter().Rendevouz, api, keys)
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
	return nil
}

// @ Begins Bootstrapping HostNode
func (c *Client) Bootstrap() (*tpc.TopicManager, *md.SonrError) {
	// Bootstrap Host
	err := c.Host.Bootstrap()
	if err != nil {
		return nil, err
	}

	// Start Textile
	err = c.Host.StartTextile(c.user.GetDevice())
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
func (n *Client) Invite(invite *md.InviteRequest, t *tpc.TopicManager) *md.SonrError {
	// Check for Peer
	if t.HasPeer(invite.To.Id.Peer) {
		// Initialize Session if transfer
		if invite.IsPayloadFile() {
			// Start New Session
			n.session = md.NewOutSession(n.user, invite, n.call)
		}

		// Get PeerID and Check error
		id, _, err := t.FindPeerInTopic(invite.To.Id.Peer)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_PEER_NOT_FOUND_INVITE)
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
func (c *Client) Mail(mr *md.MailRequest) *md.SonrError {
	if mr.Method == md.MailRequest_READ {

	} else if mr.Method == md.MailRequest_SEND {

	}
	return md.NewError(errors.New("Invalid MailRequest Method"), md.ErrorMessage_HOST_TEXTILE)
}

// @ Update proximity/direction and Notify Lobby
func (n *Client) Update(t *tpc.TopicManager) *md.SonrError {
	// Inform Lobby
	if err := t.SendLocal(n.user.Peer.SignUpdate()); err != nil {
		return md.NewError(err, md.ErrorMessage_TOPIC_UPDATE)
	}
	return nil
}

// @ Close Ends All Network Communication
func (n *Client) Close() {
	n.Host.Close()
}
