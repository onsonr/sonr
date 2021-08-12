package client

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	net "github.com/sonr-io/core/internal/host"
	srv "github.com/sonr-io/core/internal/service"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// Interface: Main Client handles Networking/Identity/Streams
type Client interface {
	// Client Methods
	Connect(cr *md.ConnectionRequest, keys *md.KeyPair) *md.SonrError
	Bootstrap(cr *md.ConnectionRequest) (*net.RoomManager, *md.SonrError)
	Mail(req *md.MailboxRequest) (*md.MailboxResponse, *md.SonrError)
	Link(invite *md.LinkRequest, t *net.RoomManager) (*md.LinkResponse, *md.SonrError)
	Invite(invite *md.InviteRequest, t *net.RoomManager) *md.SonrError
	Respond(r *md.InviteResponse)
	Update(t *net.RoomManager) *md.SonrError
	Lifecycle(state md.Lifecycle, t *net.RoomManager)

	// Room Callbacks
	OnConnected(*md.ConnectionResponse)
	OnEvent(*md.RoomEvent)
	OnError(err *md.SonrError)
	OnInvite(buf []byte)
	OnMail(e *md.MailEvent)
	OnReply(id peer.ID, data []byte)
	OnResponded(inv *md.InviteRequest)
	OnProgress(buf []byte)
	OnCompleted(stream network.Stream, pid protocol.ID, completeEvent *md.CompleteEvent)
}

// Struct: Main Client handles Networking/Identity/Streams
type client struct {
	Client

	// Properties
	ctx      context.Context
	call     md.Callback
	isLinker bool
	user     *md.User
	session  *md.Session
	request  *md.ConnectionRequest

	// References
	Host    net.HostNode
	Service srv.ServiceClient
}

// NewClient Initializes Node with Router ^
func NewClient(ctx context.Context, u *md.User, call md.Callback) Client {
	return &client{
		ctx:  ctx,
		call: call,
		user: u,
	}
}

// Connects Host Node from Private Key
func (c *client) Connect(cr *md.ConnectionRequest, keys *md.KeyPair) *md.SonrError {
	// Set Request
	c.request = cr
	c.isLinker = cr.GetIsLinker()

	// Set Host
	hn, err := net.NewHost(c.ctx, cr, keys, c)
	if err != nil {
		return err
	}

	// Get MultiAddrs
	maddr, err := hn.MultiAddr()
	if err != nil {
		return err
	}

	// Set Peer
	err = c.user.SetPeer(hn.ID(), maddr, cr.GetIsLinker())
	if err != nil {
		return err
	}

	// Set Host
	c.Host = hn
	return nil
}

// Begins Bootstrapping HostNode
func (c *client) Bootstrap(cr *md.ConnectionRequest) (*net.RoomManager, *md.SonrError) {
	// Bootstrap Host
	err := c.Host.Bootstrap(c.user.GetDevice().GetId())
	if err != nil {
		return nil, err
	}

	// Start Services
	s, err := srv.NewService(c.ctx, c.Host, c.user, c.request, c)
	if err != nil {
		return nil, err
	}
	c.Service = s

	// Join Local
	RoomName := c.user.NewLocalRoom(cr.GetServiceOptions())
	if t, err := c.Host.JoinRoom(c.ctx, c.user, RoomName, c); err != nil {
		return nil, err
	} else {
		// Check if Auto Update Events
		if cr.GetServiceOptions().GetAutoUpdate() {
			go c.sendPeriodicRoomEvents(t)
		}
		return t, nil
	}
}

// Handle a Mailbox Request from Node
func (c *client) Mail(req *md.MailboxRequest) (*md.MailboxResponse, *md.SonrError) {
	return c.Service.HandleMailbox(req)
}

// Link handles a LinkRequest
func (c *client) Link(req *md.LinkRequest, t *net.RoomManager) (*md.LinkResponse, *md.SonrError) {
	// Check Request Type
	if req.IsSend() {
		// Find Peer
		if t.HasPeer(req.To.Id.Peer) && t.HasLinker(req.To.Id.Peer) {
			// Get PeerID and Check error
			id, err := t.FindPeer(req.To.Id.Peer)
			if err != nil {
				return nil, md.NewPeerFoundError(err, req.GetTo().GetId().GetPeer())
			}

			// Send Default Invite
			err = c.Service.Link(id, req)
			if err != nil {
				return nil, md.NewError(err, md.ErrorEvent_TOPIC_RPC)
			}
			return c.user.VerifyLinkReceive(req), nil
		}
		return nil, md.NewErrorWithType(md.ErrorEvent_PEER_NOT_FOUND_INVITE)
	} else {
		c.Service.HandleLinking(req)
		return c.user.VerifyLinkReceive(req), nil
	}
}

// Invite Processes Data and Sends Invite to Peer
func (c *client) Invite(invite *md.InviteRequest, t *net.RoomManager) *md.SonrError {
	if c.user.IsReady() {
		// Check for Peer
		if invite.GetType() == md.InviteRequest_REMOTE {
			err := c.Service.SendMail(invite)
			if err != nil {
				return err
			}
		} else {
			if t.HasPeer(invite.To.Id.Peer) {
				// Get PeerID and Check error
				id, err := t.FindPeer(invite.To.Id.Peer)
				if err != nil {
					c.newExitEvent(invite)
					return md.NewPeerFoundError(err, invite.GetTo().GetId().GetPeer())
				}

				// Initialize Session if transfer
				if invite.IsPayloadTransfer() {
					// Update Status
					c.call.SetStatus(md.Status_PENDING)

					// Start New Session
					invite.SetProtocol(md.SonrProtocol_LocalTransfer, id)
					c.session = md.NewOutSession(c.user, invite, c)
				}

				// Run Routine
				go func(inv *md.InviteRequest) {
					// Send Default Invite
					err = c.Service.Invite(id, inv)
					if err != nil {
						c.call.OnError(md.NewError(err, md.ErrorEvent_TOPIC_RPC))
						return
					}
				}(invite)
			} else {
				// Send Mail to Offline Peer
				err := c.Service.SendMail(invite)
				if err != nil {
					return err
				}

				// Record Peer is Offline
				c.newExitEvent(invite)
				return md.NewErrorWithType(md.ErrorEvent_PEER_NOT_FOUND_INVITE)
			}
		}
		return nil
	}
	return nil
}

// Respond Sends a Response to Service
func (c *client) Respond(r *md.InviteResponse) {
	c.Service.Respond(r)
}

// Update proximity/direction and Notify Lobby
func (c *client) Update(t *net.RoomManager) *md.SonrError {
	if c.user.IsReady() {
		// Create Event
		ev := c.user.NewUpdateEvent(t.RoomData(), c.Host.ID())

		// Inform Lobby
		if err := t.Publish(ev); err != nil {
			return md.NewError(err, md.ErrorEvent_TOPIC_UPDATE)
		}
	}
	return nil
}

// Handle Network Communication from Lifecycle State Network Communication
func (c *client) Lifecycle(state md.Lifecycle, t *net.RoomManager) {
	if state == md.Lifecycle_ACTIVE {
		// Inform Lobby
		if c.user.IsReady() {
			ev := c.user.NewUpdateEvent(t.RoomData(), c.Host.ID())
			if err := t.Publish(ev); err != nil {
				md.NewError(err, md.ErrorEvent_TOPIC_UPDATE)
			}
		}
	} else if state == md.Lifecycle_PAUSED {
		// Inform Lobby
		if c.user.IsReady() {
			ev := c.user.NewExitEvent(t.RoomData(), c.Host.ID())
			if err := t.Publish(ev); err != nil {
				md.NewError(err, md.ErrorEvent_TOPIC_UPDATE)
			}
		}
	} else if state == md.Lifecycle_STOPPED {
		// Inform Lobby
		if c.user.IsReady() {
			ev := c.user.NewExitEvent(t.RoomData(), c.Host.ID())
			if err := t.Publish(ev); err != nil {
				md.NewError(err, md.ErrorEvent_TOPIC_UPDATE)
			}
		}
		c.Host.Close()
	}
	return
}

// Helper: Creates new Exit Event
func (c *client) newExitEvent(inv *md.InviteRequest) {
	// Create Exit Event
	event := md.RoomEvent{
		Id:      inv.To.Id.Peer,
		Subject: md.RoomEvent_EXIT,
		Peer:    inv.To,
	}

	// Marshal Data
	buf, err := proto.Marshal(&event)
	if err != nil {
		return
	}

	// Callback Event and Return Peer Error
	c.call.OnEvent(buf)
	c.call.SetStatus(md.Status_AVAILABLE)
	return
}

// Helper: Background Process to continuously ping nearby peers
func (c *client) sendPeriodicRoomEvents(t *net.RoomManager) {
	for {
		if c.user.IsReady() {
			// Create Event
			ev := c.user.NewDefaultUpdateEvent(t.RoomData(), c.Host.ID())

			// Send Update
			if err := t.Publish(ev); err != nil {
				c.call.OnError(md.NewError(err, md.ErrorEvent_TOPIC_UPDATE))
				continue
			}
		}
		time.Sleep(util.AUTOUPDATE_INTERVAL)
		md.GetState()
	}
}
