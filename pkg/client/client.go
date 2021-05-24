package client

import (
	"context"
	"log"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
	tpc "github.com/sonr-io/core/internal/topic"
	md "github.com/sonr-io/core/pkg/models"

	// Local
	net "github.com/sonr-io/core/internal/host"
)

// ^ Struct: Main Client handles Networking/Identity/Streams ^
type Client struct {
	tpc.ClientCallback

	// Properties
	isLinker bool
	ctx      context.Context
	call     md.NodeCallback
	global   net.GlobalTopic
	user     *md.User
	session  *md.Session

	// References
	Host *net.HostNode

	// Linker Properties
	linker *md.Linker
}

// ^ NewClient Initializes Node with Router ^
func NewClient(ctx context.Context, u *md.User, call md.NodeCallback) *Client {
	return &Client{
		ctx:      ctx,
		call:     call,
		user:     u,
		isLinker: false,
	}
}

func NewLinkClient(ctx context.Context, lr *md.LinkRequest) (*Client, *md.SonrError) {
	// Create Linker/Client
	linker := md.NewLinker(lr)
	c := &Client{
		ctx:      ctx,
		linker:   linker,
		isLinker: true,
	}

	// Connect Linker
	err := c.Connect(linker.PrivateKey())
	if err != nil {
		return nil, err
	}

	// Bootstrap Linker
	err = c.Bootstrap()
	if err != nil {
		return nil, err
	}

	// Set Linker Peer ID
	id, serr := c.global.FindPeerID(c.linker.Username)
	if serr != nil {
		return nil, md.NewError(serr, md.ErrorMessage_HOST_INFO)
	}
	c.linker.UserID = id.String()
	return c, nil
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
	c.user.NewPeer(hn.ID, maddr)

	// Set Host
	c.Host = hn
	return nil
}

// ^ Begins Bootstrapping HostNode ^
func (c *Client) Bootstrap() *md.SonrError {
	// Bootstrap Host
	err := c.Host.Bootstrap()
	if err != nil {
		return err
	}

	// Join Global
	global, err := c.Host.StartGlobal(c.user.SName())
	if err != nil {
		log.Println("FAILED: Start Global Topic")
		return err
	}
	log.Println("SUCCESS: Start Global Topic")

	// Set Client Global Ref
	c.global = global
	return nil
}

// ^ Creates Remote from Lobby Data ^
func (n *Client) CreateRemote(l *md.Lobby) (*tpc.TopicManager, *md.SonrError) {
	if t, err := tpc.NewRemote(n.ctx, n.Host, n.user, l, n); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Client) JoinRemote(r *md.RemoteResponse) (*tpc.TopicManager, *md.SonrError) {
	// @ Returns error if Lobby doesnt Exist
	if t, err := tpc.JoinRemote(n.ctx, n.Host, n.user, r, n); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Client) JoinLocal() (*tpc.TopicManager, *md.SonrError) {
	if t, err := tpc.NewLocal(n.ctx, n.Host, n.user, n.user.GetRouter().LocalIPTopic, n); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// ^ Join Lobby Adds Node to Named Topic ^
func (n *Client) JoinLinkLocal() (*tpc.TopicManager, *md.SonrError) {
	if t, err := tpc.NewLocalLink(n.ctx, n.Host, n.linker, n.linker.GetRouter().LocalIPTopic, n); err != nil {
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
func (n *Client) InviteUrl(invite *md.AuthInvite, t *tpc.TopicManager) *md.SonrError {
	// @ 3. Send Invite to Peer
	if t.HasPeer(invite.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := t.FindPeerInTopic(invite.To.Id.Peer)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_PEER_NOT_FOUND_INVITE)
		}

		// Run Routine
		go func(inv *md.AuthInvite) {
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
func (n *Client) InviteContact(invite *md.AuthInvite, t *tpc.TopicManager, c *md.Contact) *md.SonrError {
	// @ 3. Send Invite to Peer
	if t.HasPeer(invite.To.Id.Peer) {
		// Get PeerID and Check error
		id, _, err := t.FindPeerInTopic(invite.To.Id.Peer)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_PEER_NOT_FOUND_INVITE)
		}

		// Run Routine
		go func(inv *md.AuthInvite) {
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
func (n *Client) InviteFile(invite *md.AuthInvite, t *tpc.TopicManager) *md.SonrError {
	// Start New Session
	n.session = md.NewOutSession(n.user, invite, n.call)

	// Get PeerID
	id, _, err := t.FindPeerInTopic(invite.To.Id.Peer)
	if err != nil {
		return md.NewError(err, md.ErrorMessage_PEER_NOT_FOUND_INVITE)
	}

	// Run Routine
	go func(inv *md.AuthInvite) {
		err = t.Invite(id, inv)
		if err != nil {
			n.call.Error(md.NewError(err, md.ErrorMessage_TOPIC_RPC))
		}
	}(invite)
	return nil
}

// ^ Respond to an Invitation ^ //
func (n *Client) Respond(req *md.AuthReply, t *tpc.TopicManager) {
	t.RespondToInvite(req)
}

// ^ Update proximity/direction and Notify Lobby ^ //
func (n *Client) Update(t *tpc.TopicManager) *md.SonrError {
	// Inform Lobby
	if err := t.Send(n.user.Peer.SignUpdate()); err != nil {
		return md.NewError(err, md.ErrorMessage_TOPIC_UPDATE)
	}
	return nil
}

// ^ Close Ends All Network Communication ^
func (n *Client) Close() {
	n.Host.Host.Close()
}
