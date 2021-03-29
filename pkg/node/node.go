package node

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/crypto"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
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
	// client brprot.Service

	// Networking Properties
	Host    *net.HostNode
	router  *net.ProtocolRouter
	session *se.Session
}

// ^ NewNode Initializes Node with Router ^
func NewNode(ctx context.Context, cr *md.ConnectionRequest, call md.NodeCallback) *Node {
	return &Node{
		ctx:    ctx,
		call:   call,
		router: net.NewProtocolRouter(cr),
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
func (n *Node) Bootstrap() error {
	return n.Host.Bootstrap()
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Node) JoinLobby(name string) (*tpc.TopicManager, error) {
	if t, err := tpc.NewTopic(n.ctx, n.Host, n.router.Topic(name), false, n); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Node) JoinLocal() (*tpc.TopicManager, error) {
	if t, err := tpc.NewTopic(n.ctx, n.Host, n.router.LocalTopic(), true, n); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) InviteLink(req *md.InviteRequest, t *tpc.TopicManager, p *md.Peer) error {
	// @ 3. Send Invite to Peer
	if t.HasPeer(req.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := t.FindPeerInTopic(req.To.Id.Peer)
		if err != nil {
			return err
		}

		// Create Invite
		invite := md.GetAuthInviteWithURL(req, p)

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
func (n *Node) InviteContact(req *md.InviteRequest, t *tpc.TopicManager, p *md.Peer, c *md.Contact) error {
	// @ 3. Send Invite to Peer
	if t.HasPeer(req.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := t.FindPeerInTopic(req.To.Id.Peer)
		if err != nil {
			return err
		}

		// Build Invite Message
		invite := md.GetAuthInviteWithContact(req, p, c)

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
func (n *Node) InviteFile(req *md.InviteRequest, t *tpc.TopicManager, p *md.Peer, fs *sf.FileSystem) error {
	// Start New Session
	session := se.NewOutSession(p, req, fs, n.call)
	card := session.OutgoingCard()

	// Create Invite Message
	invite := md.AuthInvite{
		From:    p,
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
func (n *Node) Respond(decision bool, fs *sf.FileSystem, p *md.Peer, t *tpc.TopicManager, c *md.Contact) {
	t.RespondToInvite(decision, fs, p, c)
}

// ^ Send Direct Message to Peer in Lobby ^ //
func (n *Node) Message(t *tpc.TopicManager, msg string, to string, p *md.Peer) error {
	if t.HasPeer(to) {
		// Inform Lobby
		if err := t.Send(&md.LobbyEvent{
			Event:   md.LobbyEvent_MESSAGE,
			From:    p,
			Id:      p.Id.Peer,
			Message: msg,
			To:      to,
		}); err != nil {
			return err
		}
	}
	return nil
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (n *Node) Update(t *tpc.TopicManager, p *md.Peer) error {
	// Inform Lobby
	if err := t.Send(&md.LobbyEvent{
		Event: md.LobbyEvent_UPDATE,
		From:  p,
		Id:    p.Id.Peer,
	}); err != nil {
		return err
	}
	return nil
}

// ^ Close Ends All Network Communication ^
func (n *Node) Close() {
	n.Host.Host.Close()
}
