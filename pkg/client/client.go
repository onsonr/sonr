package client

import (
	"context"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
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
func (c *Client) Connect(pk crypto.PrivKey) *md.SonrError {
	// Set Host
	hn, err := net.NewHost(c.ctx, c.router.Rendevouz(), pk)
	if err != nil {
		return err
	}

	// Set Peer
	c.Peer, err = md.NewPeer(c.req, hn.ID)
	if err != nil {
		return err
	}

	c.Host = hn
	return nil
}

// ^ Begins Bootstrapping HostNode ^
func (c *Client) Bootstrap() *md.SonrError {
	return c.Host.Bootstrap()
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Client) JoinLobby(name string, isCreated bool) (*tpc.TopicManager, *md.SonrError) {
	// @ Check for Topic being Created
	if isCreated {
		if t, err := tpc.NewTopic(n.ctx, n.Host, n.Peer, n.router.Topic(name), false, n); err != nil {
			return nil, err
		} else {
			return t, nil
		}
	} else {
		// @ Returns error if Lobby doesnt Exist
		if t, err := tpc.JoinTopic(n.ctx, n.Host, n.Peer, n.router.Topic(name), false, n); err != nil {
			return nil, err
		} else {
			return t, nil
		}
	}
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Client) JoinLocal() (*tpc.TopicManager, *md.SonrError) {
	if t, err := tpc.NewTopic(n.ctx, n.Host, n.Peer, n.router.LocalTopic(), true, n); err != nil {
		return nil, err
	} else {
		return t, nil
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
func (n *Client) InviteLink(req *md.InviteRequest, t *tpc.TopicManager) *md.SonrError {
	// @ 3. Send Invite to Peer
	if t.HasPeer(req.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := t.FindPeerInTopic(req.To.Id.Peer)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_PEER_NOT_FOUND_INVITE)
		}

		// Create Invite
		invite := n.Peer.SignInviteWithLink(req)

		// Run Routine
		go func(inv *md.AuthInvite) {
			err = t.Invite(id, inv, nil)
			if err != nil {
				n.call.Error(md.NewError(err, md.ErrorMessage_TOPIC_RPC))
			}
		}(&invite)
	} else {
		return md.NewErrorWithType(md.ErrorMessage_PEER_NOT_FOUND_INVITE)
	}
	return nil
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Client) InviteContact(req *md.InviteRequest, t *tpc.TopicManager, c *md.Contact) *md.SonrError {
	// @ 3. Send Invite to Peer
	if t.HasPeer(req.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := t.FindPeerInTopic(req.To.Id.Peer)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_PEER_NOT_FOUND_INVITE)
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
					n.call.Error(md.NewError(err, md.ErrorMessage_TOPIC_RPC))
				}
			} else {
				// Request Invite for Non Flat
				err = t.Invite(id, inv, nil)
				if err != nil {
					n.call.Error(md.NewError(err, md.ErrorMessage_TOPIC_RPC))
				}
			}
		}(&invite)
	} else {
		return md.NewErrorWithType(md.ErrorMessage_PEER_NOT_FOUND_INVITE)
	}
	return nil
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Client) InviteFile(req *md.InviteRequest, t *tpc.TopicManager, fs *us.FileSystem) *md.SonrError {
	// Start New Session
	session := se.NewOutSession(n.Peer, req, fs, n.call)
	card := session.OutgoingCard()

	// Create Invite Message
	invite := n.Peer.SignInviteWithFile(card, req)

	// Get PeerID
	id, _, err := t.FindPeerInTopic(req.To.Id.Peer)
	if err != nil {
		return md.NewError(err, md.ErrorMessage_PEER_NOT_FOUND_INVITE)
	}

	// Run Routine
	go func(inv *md.AuthInvite) {
		err = t.Invite(id, inv, session)
		if err != nil {
			n.call.Error(md.NewError(err, md.ErrorMessage_TOPIC_RPC))
		}
	}(&invite)
	return nil
}

// ^ Respond to an Invitation ^ //
func (n *Client) Respond(req *md.RespondRequest, t *tpc.TopicManager, fs *us.FileSystem, c *md.Contact) {
	t.RespondToInvite(req, fs, n.Peer, c)
}

// ^ Send Direct Message to Peer in Lobby ^ //
func (n *Client) Message(t *tpc.TopicManager, msg string, to *md.Peer) *md.SonrError {
	if t.HasPeer(to.PeerID()) {
		// Inform Lobby
		if err := t.Send(n.Peer.SignMessage(msg, to)); err != nil {
			return md.NewError(err, md.ErrorMessage_TOPIC_MESSAGE)
		}
	}
	return nil
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (n *Client) Update(t *tpc.TopicManager) *md.SonrError {

	// Inform Lobby
	if err := t.Send(n.Peer.SignUpdate()); err != nil {
		return md.NewError(err, md.ErrorMessage_TOPIC_UPDATE)
	}
	return nil
}

// ^ Close Ends All Network Communication ^
func (n *Client) Close() {
	n.Host.Host.Close()
}
