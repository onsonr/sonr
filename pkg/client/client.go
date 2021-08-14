package client

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
		ac "github.com/sonr-io/core/pkg/account"
	net "github.com/sonr-io/core/internal/host"
	srv "github.com/sonr-io/core/internal/service"
	tp "github.com/sonr-io/core/internal/topic"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// Interface: Main Client handles Networking/Identity/Streams
type Client interface {
	// Client Methods
	Connect(cr *md.ConnectionRequest, a ac.Account) (*md.Peer, bool, *md.SonrError)
	Bootstrap(cr *md.ConnectionRequest) (*tp.RoomManager, *md.SonrError)
	Mail(req *md.MailboxRequest) (*md.MailboxResponse, *md.SonrError)
	Link(invite *md.LinkRequest, t *tp.RoomManager) (*md.LinkResponse, *md.SonrError)
	Invite(invite *md.InviteRequest, t *tp.RoomManager) *md.SonrError
	Respond(r *md.InviteResponse)
	Update(t *tp.RoomManager) *md.SonrError
	Lifecycle(state md.Lifecycle, t *tp.RoomManager)

	// Room Callbacks
	OnConnected(*md.ConnectionResponse)
	OnRoomEvent(*md.RoomEvent)
	OnSyncEvent(*md.SyncEvent)
	OnError(*md.SonrError)
	OnInvite([]byte)
	OnMail(*md.MailEvent)
	OnReply(peer.ID, []byte)
	OnResponded(*md.InviteRequest)
	OnProgress([]byte)
	OnCompleted(network.Stream, protocol.ID, *md.CompleteEvent)
}

// Struct: Main Client handles Networking/Identity/Streams
type client struct {
	Client

	// Properties
	ctx      context.Context
	call     md.Callback
	isLinker bool
	account  ac.Account
	device   *md.Device
	session  *md.Session
	request  *md.ConnectionRequest

	// References
	Host    net.HostNode
	Service srv.ServiceClient
}

// NewClient Initializes Node with Router ^
func NewClient(ctx context.Context, u *md.Device, call md.Callback) Client {
	return &client{
		ctx:    ctx,
		call:   call,
		device: u,
	}
}

// Connects Host Node from Private Key
func (c *client) Connect(cr *md.ConnectionRequest, a ac.Account) (*md.Peer, bool, *md.SonrError) {
	// Set Request
	c.request = cr
	c.isLinker = cr.GetIsLinker()
	c.account = a

	// Set Host
	hn, err := net.NewHost(c.ctx, cr, c.device.AccountKeys(), c)
	if err != nil {
		return nil, false, err
	}

	// Get MultiAddrs
	maddr, err := hn.MultiAddr()
	if err != nil {
		return nil, false, err
	}

	// Set Peer
	peer, isPrimary := c.device.SetPeer(hn.ID(), maddr, cr.GetIsLinker())

	// Set Host
	c.Host = hn
	return peer, isPrimary, nil
}

// Begins Bootstrapping HostNode
func (c *client) Bootstrap(cr *md.ConnectionRequest) (*tp.RoomManager, *md.SonrError) {
	// Bootstrap Host
	err := c.Host.Bootstrap(c.device.GetId())
	if err != nil {
		return nil, err
	}

	// Start Services
	s, err := srv.NewService(c.ctx, c.Host, c.device, c.request, c)
	if err != nil {
		return nil, err
	}
	c.Service = s

	// Join Local
	RoomName := c.device.NewLocalRoom(cr.GetServiceOptions())
	if t, err := tp.JoinRoom(c.ctx, c.Host, c.device, RoomName, c); err != nil {
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
func (c *client) Link(req *md.LinkRequest, t *tp.RoomManager) (*md.LinkResponse, *md.SonrError) {
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
				return nil, md.NewError(err, md.ErrorEvent_ROOM_RPC)
			}
			return &md.LinkResponse{
				Success: false,
				Type:    md.LinkResponse_Type(req.GetType()),
			}, nil
		}
		return nil, md.NewErrorWithType(md.ErrorEvent_PEER_NOT_FOUND_INVITE)
	} else {
		c.Service.HandleLinking(req)
		return &md.LinkResponse{
			Success: false,
			Type:    md.LinkResponse_Type(req.GetType()),
		}, nil
	}
}

// Invite Processes Data and Sends Invite to Peer
func (c *client) Invite(invite *md.InviteRequest, t *tp.RoomManager) *md.SonrError {
	if c.device.IsReady() {
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
					c.session = md.NewOutSession(c.device, invite, c)
				}

				// Run Routine
				go func(inv *md.InviteRequest) {
					// Send Default Invite
					err = c.Service.Invite(id, inv)
					if err != nil {
						c.call.OnError(md.NewError(err, md.ErrorEvent_ROOM_RPC))
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
func (c *client) Update(t *tp.RoomManager) *md.SonrError {
	if c.device.IsReady() {
		// Create Event
		ev := c.device.NewUpdateEvent(t.Room(), c.Host.ID())

		// Inform Lobby
		if err := t.Publish(ev); err != nil {
			return md.NewError(err, md.ErrorEvent_ROOM_UPDATE)
		}
	}
	return nil
}

// Handle Network Communication from Lifecycle State Network Communication
func (c *client) Lifecycle(state md.Lifecycle, t *tp.RoomManager) {
	if state == md.Lifecycle_ACTIVE {
		// Inform Lobby
		if c.device.IsReady() {
			ev := c.device.NewUpdateEvent(t.Room(), c.Host.ID())
			if err := t.Publish(ev); err != nil {
				md.NewError(err, md.ErrorEvent_ROOM_UPDATE)
			}
		}
	} else if state == md.Lifecycle_PAUSED {
		// Inform Lobby
		if c.device.IsReady() {
			ev := c.device.NewExitEvent(t.Room(), c.Host.ID())
			if err := t.Publish(ev); err != nil {
				md.NewError(err, md.ErrorEvent_ROOM_UPDATE)
			}
		}
	} else if state == md.Lifecycle_STOPPED {
		// Inform Lobby
		if c.device.IsReady() {
			ev := c.device.NewExitEvent(t.Room(), c.Host.ID())
			if err := t.Publish(ev); err != nil {
				md.NewError(err, md.ErrorEvent_ROOM_UPDATE)
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
func (c *client) sendPeriodicRoomEvents(t *tp.RoomManager) {
	for {
		if c.device.IsReady() {
			// Create Event
			ev := c.device.NewDefaultUpdateEvent(t.Room(), c.Host.ID())

			// Send Update
			if err := t.Publish(ev); err != nil {
				c.call.OnError(md.NewError(err, md.ErrorEvent_ROOM_UPDATE))
				continue
			}
		}
		time.Sleep(util.AUTOUPDATE_INTERVAL)
		md.GetState()
	}
}
