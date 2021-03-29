package node

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/crypto"
	us "github.com/sonr-io/core/internal/user"
	md "github.com/sonr-io/core/internal/models"
	pn "github.com/sonr-io/core/pkg/peer"
	tpc "github.com/sonr-io/core/pkg/topic"

	// Local
	// brprot "berty.tech/berty/v2/go/pkg/bertyprotocol"
	net "github.com/sonr-io/core/internal/network"
	se "github.com/sonr-io/core/internal/session"
)

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
	// Properties
	ctx  context.Context
	call md.NodeCallback
	peer *pn.PeerNode
	// client brprot.Service

	// Networking Properties
	Host    *net.HostNode
	router  *ProtocolRouter
	session *se.Session
}

// ^ NewNode Initializes Node with Router ^
func NewNode(ctx context.Context, cr *md.ConnectionRequest, call md.NodeCallback) *Node {
	return &Node{
		ctx:    ctx,
		call:   call,
		router: NewProtocolRouter(cr),
	}
}

// ^ Connects Host Node from Private Key ^
func (n *Node) Connect(key crypto.PrivKey) error {
	hn, err := net.NewHost(n.ctx, n.router.Rendevouz(), key)
	if err != nil {
		return err
	}

	n.Host = hn
	return nil
}

// ^ Begins Bootstrapping HostNode ^
func (n *Node) Bootstrap(peer *pn.PeerNode) error {
	n.peer = peer
	return n.Host.Bootstrap()
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Node) JoinLobby(name string) (*tpc.TopicManager, error) {
	if t, err := tpc.NewTopic(n.ctx, n.Host, n.peer, n.router.Topic(name), false, n); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Node) JoinLocal() (*tpc.TopicManager, error) {
	if t, err := tpc.NewTopic(n.ctx, n.Host, n.peer, n.router.LocalTopic(), true, n); err != nil {
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
		invite := md.GetAuthInviteWithURL(req, n.peer.Get())

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
		invite := md.GetAuthInviteWithContact(req, n.peer.Get(), c)

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
	session := se.NewOutSession(n.peer.Get(), req, fs, n.call)
	card := session.OutgoingCard()

	// Create Invite Message
	invite := md.AuthInvite{
		From:    n.peer.Get(),
		Payload: card.Payload,
		Card:    card,
	}

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
func (n *Node) Respond(decision bool, fs *us.FileSystem, t *tpc.TopicManager, c *md.Contact) {
	t.RespondToInvite(decision, fs, c)
}

// ^ Send Direct Message to Peer in Lobby ^ //
func (n *Node) Message(t *tpc.TopicManager, msg string, to string) error {
	if t.HasPeer(to) {
		// Inform Lobby
		if err := t.Send(n.peer.SignMessage(msg, to)); err != nil {
			return err
		}
	}
	return nil
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (n *Node) Update(t *tpc.TopicManager) error {
	// Inform Lobby
	if err := t.Send(n.peer.SignUpdate()); err != nil {
		return err
	}
	return nil
}

// ^ Close Ends All Network Communication ^
func (n *Node) Close() {
	n.Host.Host.Close()
}
