package client

import (
	"context"
	"errors"

	crypto "github.com/libp2p/go-libp2p-crypto"
	tpc "github.com/sonr-io/core/internal/topic"
	md "github.com/sonr-io/core/pkg/models"
	us "github.com/sonr-io/core/pkg/user"

	// Local
	// brprot "berty.tech/berty/v2/go/pkg/bertyprotocol"
	net "github.com/sonr-io/core/internal/network"
	se "github.com/sonr-io/core/internal/session"
)

// ^ Struct: Main Client handles Networking/Identity/Streams ^
type Client struct {
	// Properties
	ctx     context.Context
	call    md.NodeCallback
	req     *md.ConnectionRequest
	router  *ProtocolRouter
	session *se.Session

	// client brprot.Service

	// References
	Host *net.HostNode
	Peer *md.Peer
}

// ^ NewClient Initializes Node with Router ^
func NewClient(ctx context.Context, cr *md.ConnectionRequest, call md.NodeCallback) *Client {
	return &Client{
		ctx:    ctx,
		call:   call,
		req:    cr,
		router: NewProtocolRouter(cr),
	}
}

// ^ Connects Host Node from Private Key ^
func (c *Client) Connect(pk crypto.PrivKey) error {
	// Set Host
	hn, err := net.NewHost(c.ctx, c.router.Rendevouz(), pk)
	if err != nil {
		return err
	}

	// Set Peer
	c.Peer, err = md.NewPeer(c.req, hn.ID())
	if err != nil {
		return err
	}

	c.Host = hn
	return nil
}

// ^ Begins Bootstrapping HostNode ^
func (c *Client) Bootstrap() error {
	return c.Host.Bootstrap()
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Client) JoinLobby(name string) (*tpc.TopicManager, error) {
	if t, err := tpc.NewTopic(n.ctx, n.Host, n.Peer, n.router.Topic(name), false, n); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Client) JoinLocal() (*tpc.TopicManager, error) {
	if t, err := tpc.NewTopic(n.ctx, n.Host, n.Peer, n.router.LocalTopic(), true, n); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Client) LeaveLobby(lob *tpc.TopicManager) error {
	if err := lob.LeaveTopic(); err != nil {
		return err
	}
	return nil
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Client) InviteLink(req *md.InviteRequest, t *tpc.TopicManager) error {
	// @ 3. Send Invite to Peer
	if t.HasPeer(req.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := t.FindPeerInTopic(req.To.Id.Peer)
		if err != nil {
			return err
		}

		// Create Invite
		invite := n.Peer.SignInviteWithLink(req)

		// Run Routine
		go func(inv *md.AuthInvite) {
			err = t.Invite(id, inv, nil)
			if err != nil {
				n.call.Error(err, "InviteLink")
			}
		}(&invite)
	} else {
		return errors.New("Invalid Peer")
	}
	return nil
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Client) InviteContact(req *md.InviteRequest, t *tpc.TopicManager, c *md.Contact) error {
	// @ 3. Send Invite to Peer
	if t.HasPeer(req.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := t.FindPeerInTopic(req.To.Id.Peer)
		if err != nil {
			return err
		}

		// Build Invite Message
		isFlat := req.Type == md.InviteRequest_FlatContact
		invite := n.Peer.SignInviteWithContact(c, isFlat, req)

		// Run Routine
		go func(inv *md.AuthInvite) {
			// Direct Invite for Flat
			if isFlat {
				err = t.Direct(id, inv)
				if err != nil {
					n.call.Error(err, "InviteContact:Flat")
				}
			} else {
				// Request Invite for Non Flat
				err = t.Invite(id, inv, nil)
				if err != nil {
					n.call.Error(err, "InviteContact")
				}
			}
		}(&invite)
	} else {
		return errors.New("Invalid Peer")
	}
	return nil
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Client) InviteFile(req *md.InviteRequest, t *tpc.TopicManager, fs *us.FileSystem) error {
	// Start New Session
	session := se.NewOutSession(n.Peer, req, fs, n.call)
	card := session.OutgoingCard()

	// Create Invite Message
	invite := n.Peer.SignInviteWithFile(card, req)

	// Get PeerID
	id, _, err := t.FindPeerInTopic(req.To.Id.Peer)
	if err != nil {
		return err
	}

	// Run Routine
	go func(inv *md.AuthInvite) {
		err = t.Invite(id, inv, session)
		if err != nil {
			n.call.Error(err, "InviteFile")
		}
	}(&invite)
	return nil
}

// ^ Respond to an Invitation ^ //
func (n *Client) Respond(req *md.RespondRequest, t *tpc.TopicManager, fs *us.FileSystem, c *md.Contact) {
	t.RespondToInvite(req, fs, n.Peer, c)
}

// ^ Send Direct Message to Peer in Lobby ^ //
func (n *Client) Message(t *tpc.TopicManager, msg string, to *md.Peer) error {
	if t.HasPeer(to.PeerID()) {
		// Inform Lobby
		if err := t.Send(n.Peer.SignMessage(msg, to)); err != nil {
			return err
		}
	}
	return nil
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (n *Client) Update(t *tpc.TopicManager) error {

	// Inform Lobby
	if err := t.Send(n.Peer.SignUpdate()); err != nil {
		return err
	}
	return nil
}

// ^ Close Ends All Network Communication ^
func (n *Client) Close() {
	n.Host.Host.Close()
}
