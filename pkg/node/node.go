package node

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/crypto"
	tpc "github.com/sonr-io/core/internal/topic"
	md "github.com/sonr-io/core/pkg/models"
	us "github.com/sonr-io/core/pkg/user"

	// Local
	// brprot "berty.tech/berty/v2/go/pkg/bertyprotocol"
	se "github.com/sonr-io/core/internal/session"
	net "github.com/sonr-io/core/pkg/network"
)

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
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

// ^ NewNode Initializes Node with Router ^
func NewNode(ctx context.Context, cr *md.ConnectionRequest, call md.NodeCallback) *Node {
	return &Node{
		ctx:    ctx,
		call:   call,
		req:    cr,
		router: NewProtocolRouter(cr),
	}
}

// ^ Connects Host Node from Private Key ^
func (n *Node) Connect(key crypto.PrivKey) error {
	// Set Host
	hn, err := net.NewHost(n.ctx, n.router.Rendevouz(), key)
	if err != nil {
		return err
	}

	// Set Peer
	n.Peer, err = md.NewPeer(n.req, hn.ID())
	if err != nil {
		return err
	}

	n.Host = hn
	return nil
}

// ^ Begins Bootstrapping HostNode ^
func (n *Node) Bootstrap() error {
	return n.Host.Bootstrap()
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Node) JoinLobby(name string) (*tpc.TopicManager, error) {
	if t, err := tpc.NewTopic(n.ctx, n.Host, n.Peer, n.router.Topic(name), false, n); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Node) JoinLocal() (*tpc.TopicManager, error) {
	if t, err := tpc.NewTopic(n.ctx, n.Host, n.Peer, n.router.LocalTopic(), true, n); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) InviteLink(req *md.InviteRequest, t *tpc.TopicManager) error {
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
func (n *Node) InviteContact(req *md.InviteRequest, t *tpc.TopicManager, c *md.Contact) error {
	// @ 3. Send Invite to Peer
	if t.HasPeer(req.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := t.FindPeerInTopic(req.To.Id.Peer)
		if err != nil {
			return err
		}

		// Build Invite Message
		invite := n.Peer.SignInviteWithContact(c)

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
func (n *Node) InviteFile(req *md.InviteRequest, t *tpc.TopicManager, fs *us.FileSystem) error {
	// Start New Session
	session := se.NewOutSession(n.Peer, req, fs, n.call)
	card := session.OutgoingCard()

	// Create Invite Message
	invite := n.Peer.SignInviteWithFile(card)

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
func (n *Node) Respond(decision bool, t *tpc.TopicManager, fs *us.FileSystem, c *md.Contact) {
	t.RespondToInvite(decision, fs, n.Peer, c)
}

// ^ Send Direct Message to Peer in Lobby ^ //
func (n *Node) Message(t *tpc.TopicManager, msg string, to string) error {
	if t.HasPeer(to) {
		// Inform Lobby
		if err := t.Send(n.Peer.SignMessage(msg, to)); err != nil {
			return err
		}
	}
	return nil
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (n *Node) Update(t *tpc.TopicManager, f float64, h float64) error {
	// Update Position
	n.Peer.SetPosition(f, h)

	// Inform Lobby
	if err := t.Send(n.Peer.SignUpdate()); err != nil {
		return err
	}
	return nil
}

// ^ Close Ends All Network Communication ^
func (n *Node) Close() {
	n.Host.Host.Close()
}
