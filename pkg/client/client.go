package client

import (
	"context"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
	tpc "github.com/sonr-io/core/internal/topic"
	md "github.com/sonr-io/core/pkg/models"

	// Local
	net "github.com/sonr-io/core/internal/host"
)

// ^ Struct: Main Client handles Networking/Identity/Streams ^
type Client struct {
	tpc.ClientHandler

	// Properties
	isLinker bool
	ctx      context.Context
	call     md.NodeCallback
	global   net.GlobalTopic
	user     *md.User
	session  *md.Session

	// References
	Host *net.HostNode
}

// ^ NewClient Initializes Node with Router ^
func NewClient(ctx context.Context, u *md.User, call md.NodeCallback) *Client {
	// Returns Storj Enabled Client
	return &Client{
		ctx:  ctx,
		call: call,
		user: u,
	}
}

// ^ Connects Host Node from Private Key ^
func (c *Client) Connect(pk crypto.PrivKey) *md.SonrError {
	// Set Host
	hn, err := net.NewHost(c.ctx, c.user.GetRouter().Rendevouz, pk)
	if err != nil {
		return err
	}

	// Get MultiAddrs
	maddr, err := hn.MultiAddr()
	if err != nil {
		return err
	}

	// Set Peer
	err = c.user.NewPeer(hn.ID, maddr)
	if err != nil {
		return err
	}

	// Set Host
	c.Host = hn
	return nil
}

// ^ Begins Bootstrapping HostNode ^
func (c *Client) Bootstrap() (*tpc.TopicManager, *md.SonrError) {
	// Bootstrap Host
	err := c.Host.Bootstrap()
	if err != nil {
		return nil, err
	}

	// Join Global
	global, err := c.Host.StartGlobal(c.user.SName())
	if err != nil {
		return nil, err
	}

	// Set Client Global Ref
	c.global = global

	// Join Local Topic
	if t, err := tpc.NewLocal(c.ctx, c.Host, c.user, c.user.GetRouter().LocalTopic, c); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// ^ Creates Remote from Lobby Data ^
func (n *Client) CreateRemote(r *md.RemoteCreateRequest) (*tpc.TopicManager, *md.RemoteCreateResponse, *md.SonrError) {
	if t, resp, err := tpc.NewRemote(n.ctx, n.Host, n.user, r, n); err != nil {
		return nil, nil, err
	} else {
		return t, resp, nil
	}
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Client) JoinRemote(r *md.RemoteJoinRequest) (*tpc.TopicManager, *md.RemoteJoinResponse, *md.SonrError) {
	// @ Returns error if Lobby doesnt Exist
	if t, resp, err := tpc.JoinRemote(n.ctx, n.Host, n.user, r, n); err != nil {
		return nil, nil, err
	} else {
		return t, resp, nil
	}
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Client) LeaveLobby(lob *tpc.TopicManager) *md.SonrError {
	if err := lob.LeaveTopic(); err != nil {
		return md.NewError(err, md.ErrorMessage_TOPIC_LEAVE)
	}
	return nil
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
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
				n.call.Error(md.NewError(err, md.ErrorMessage_TOPIC_RPC))
			}
		}(invite)
	} else {
		return md.NewErrorWithType(md.ErrorMessage_PEER_NOT_FOUND_INVITE)
	}
	return nil
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
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
					n.call.Error(md.NewError(err, md.ErrorMessage_TOPIC_RPC))
				}
			} else {
				// Request Invite for Non Flat
				err = t.Invite(id, inv)
				if err != nil {
					n.call.Error(md.NewError(err, md.ErrorMessage_TOPIC_RPC))
				}
			}
		}(invite)
	} else {
		return md.NewErrorWithType(md.ErrorMessage_PEER_NOT_FOUND_INVITE)
	}
	return nil
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
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
			n.call.Error(md.NewError(err, md.ErrorMessage_TOPIC_RPC))
		}
	}(invite)
	return nil
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (n *Client) Update(t *tpc.TopicManager) *md.SonrError {
	// Inform Lobby
	if err := t.SendLocal(n.user.Peer.SignUpdate()); err != nil {
		return md.NewError(err, md.ErrorMessage_TOPIC_UPDATE)
	}
	return nil
}

// ^ Close Ends All Network Communication ^
func (n *Client) Close() {
	n.Host.Host.Close()
}
