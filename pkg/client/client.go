package client

import (
	"context"

	tpc "github.com/sonr-io/core/internal/topic"
	md "github.com/sonr-io/core/pkg/models"

	// Local
	net "github.com/sonr-io/core/internal/host"
)

// Struct: Main Client handles Networking/Identity/Streams
type Client struct {
	tpc.ClientHandler

	// Properties
	isLinker bool
	ctx      context.Context
	call     md.Callback
	user     *md.User
	session  *md.Session

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

	// Join Global
	err = c.Host.StartGlobal(c.user.SName())
	if err != nil {
		return nil, err
	}

	// Join Local Topic
	if t, err := tpc.NewLocal(c.ctx, c.Host, c.user, c.user.GetRouter().LocalTopic, c); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// @ Join Lobby Adds Node to Named Topic
func (n *Client) LeaveLobby(lob *tpc.TopicManager) *md.SonrError {
	if err := lob.LeaveTopic(); err != nil {
		return md.NewError(err, md.ErrorMessage_TOPIC_LEAVE)
	}
	return nil
}

// @ Invite Processes Data and Sends Invite to Peer
func (n *Client) InviteLink(invite *md.InviteRequest, t *tpc.TopicManager) *md.SonrError {
	// @ 3. Send Invite to Peer
	if t.HasPeer(invite.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := t.FindPeerInTopic(invite.To.Id.Peer)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_PEER_NOT_FOUND_INVITE)
		}

		// Run Routine
		go func(inv *md.InviteRequest) {
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

// @ Invite Processes Data and Sends Invite to Peer
func (n *Client) InviteContact(invite *md.InviteRequest, t *tpc.TopicManager, c *md.Contact) *md.SonrError {
	// @ 3. Send Invite to Peer
	if t.HasPeer(invite.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := t.FindPeerInTopic(invite.To.Id.Peer)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_PEER_NOT_FOUND_INVITE)
		}

		// Run Routine
		go func(inv *md.InviteRequest) {
			// Direct Invite for Flat
			if inv.IsFlat() {
				err = t.Flat(id, inv)
				if err != nil {
					n.call.OnError(md.NewError(err, md.ErrorMessage_TOPIC_RPC))
				}
			} else {
				// Request Invite for Non Flat
				err = t.Invite(id, inv)
				if err != nil {
					n.call.OnError(md.NewError(err, md.ErrorMessage_TOPIC_RPC))
				}
			}
		}(invite)
	} else {
		return md.NewErrorWithType(md.ErrorMessage_PEER_NOT_FOUND_INVITE)
	}
	return nil
}

// @ Invite Processes Data and Sends Invite to Peer
func (n *Client) InviteFile(invite *md.InviteRequest, t *tpc.TopicManager) *md.SonrError {
	// Start New Session
	n.session = md.NewOutSession(n.user, invite, n.call)

	// Get PeerID
	id, _, err := t.FindPeerInTopic(invite.To.Id.Peer)
	if err != nil {
		return md.NewError(err, md.ErrorMessage_PEER_NOT_FOUND_INVITE)
	}

	// Run Routine
	go func(inv *md.InviteRequest) {
		err = t.Invite(id, inv)
		if err != nil {
			n.call.OnError(md.NewError(err, md.ErrorMessage_TOPIC_RPC))
		}
	}(invite)
	return nil
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
