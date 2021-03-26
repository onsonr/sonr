package node

import (
	"context"
	"errors"

	dht "github.com/libp2p/go-libp2p-kad-dht"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	tpc "github.com/sonr-io/core/pkg/topic"
	tr "github.com/sonr-io/core/pkg/transfer"

	// Imported
	"github.com/libp2p/go-libp2p-core/host"
	psub "github.com/libp2p/go-libp2p-pubsub"

	// Local
	net "github.com/sonr-io/core/internal/network"
	dt "github.com/sonr-io/core/pkg/data"
)

// ^ Struct: Main Node handles Networking/Identity/Streams ^
type Node struct {
	// Properties
	ctx  context.Context
	opts *net.HostOptions
	call dt.NodeCallback

	// Networking Properties
	Host   host.Host
	kdht   *dht.IpfsDHT
	pubsub *psub.PubSub
	router *net.ProtocolRouter

	// Data Handlers
	incoming *tr.IncomingFile
}

// ^ NewNode Initializes Node with a host and default properties ^
func NewNode(ctx context.Context, opts *net.HostOptions, call dt.NodeCallback) *Node {
	return &Node{
		ctx:    ctx,
		call:   call,
		opts:   opts,
		router: net.NewProtocolRouter(opts.ConnRequest),
	}
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Node) JoinLobby(name string) (*tpc.TopicManager, error) {
	if t, err := tpc.NewTopic(n.ctx, n.Host, n.pubsub, n.router.Topic(name), n.router, n); err != nil {
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
		err = t.Invite(id, &invite, p, nil)
		if err != nil {
			return err
		}
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
		invMsg := md.GetAuthInviteWithContact(req, p, c)

		// Run Routine
		err = t.Invite(id, &invMsg, p, nil)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Invalid Peer")
	}
	return nil
}

// ^ Invite Processes Data and Sends Invite to Peer ^ //
func (n *Node) InviteFile(req *md.InviteRequest, t *tpc.TopicManager, p *md.Peer, cf *sf.FileItem) error {
	// Create Invite
	info, err := cf.InfoOut()
	if err != nil {
		return err
	}
	invite := md.GetAuthInviteWithFile(req, p, info)

	// Get PeerID
	id, _, err := t.FindPeerInTopic(req.To.Id.Peer)
	if err != nil {
		return err
	}

	// Run Routine
	err = t.Invite(id, &invite, p, cf)
	if err != nil {
		return err
	}
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
