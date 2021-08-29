package client

import (
	"context"
	"errors"
	"time"

	"github.com/sonr-io/core/internal/emitter"
	net "github.com/sonr-io/core/internal/host"
	room "github.com/sonr-io/core/internal/room"
	srv "github.com/sonr-io/core/internal/service"
	ac "github.com/sonr-io/core/pkg/account"
	data "github.com/sonr-io/core/pkg/data"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// Interface: Main Client handles Networking/Identity/Streams
type Client interface {
	// Client Methods
	Connect(cr *data.ConnectionRequest) (*data.Peer, *data.SonrError)
	Bootstrap(cr *data.ConnectionRequest) (*room.RoomManager, *data.SonrError)
	Mail(req *data.MailboxRequest) (*data.MailboxResponse, *data.SonrError)
	Link(invite *data.LinkRequest, t *room.RoomManager) (*data.LinkResponse, *data.SonrError)
	Invite(invite *data.InviteRequest, t *room.RoomManager) *data.SonrError
	Respond(r *data.InviteResponse)
	Update(t *room.RoomManager) *data.SonrError
	Lifecycle(state data.Lifecycle, t *room.RoomManager)

	// Properties
	GetHost() net.HostNode
}

// Struct: Main Client handles Networking/Identity/Streams
type client struct {
	Client

	// Properties
	ctx      context.Context
	call     data.Callback
	isLinker bool
	account  ac.Account
	session  *data.Session
	request  *data.ConnectionRequest
	emitter  *emitter.Emitter

	// References
	Host    net.HostNode
	Service srv.ServiceClient
}

// NewClient Initializes Node with Router ^
func NewClient(ctx context.Context, a ac.Account, call data.Callback) Client {
	c := &client{
		ctx:     ctx,
		call:    call,
		account: a,
	}

	c.initEmitter()
	return c
}

// Connects Host Node from Private Key
func (c *client) Connect(cr *data.ConnectionRequest) (*data.Peer, *data.SonrError) {
	// Set Request
	c.request = cr
	c.isLinker = cr.GetIsLinker()

	// Set Host
	hn, err := net.NewHost(c.ctx, cr, c.account.AccountKeys(), c.emitter)
	if err != nil {
		return nil, err
	}

	// Get MultiAddrs
	maddr, err := hn.MultiAddr()
	if err != nil {
		return nil, err
	}

	// Set Peer
	peer, _ := c.account.CurrentDevice().SetPeer(hn.ID(), maddr, cr.GetIsLinker())

	// Set Host
	c.Host = hn
	return peer, nil
}

// Begins Bootstrapping HostNode
func (c *client) Bootstrap(cr *data.ConnectionRequest) (*room.RoomManager, *data.SonrError) {
	// Bootstrap Host
	err := c.Host.Bootstrap(c.account.CurrentDevice().GetId())
	if err != nil {
		return nil, err
	}

	// Start Services
	s, err := srv.NewService(c.ctx, c.Host, c.account.CurrentDevice(), c.request, c.emitter)
	if err != nil {
		return nil, err
	}
	c.Service = s

	// Join Local
	RoomName := c.account.CurrentDevice().NewLocalRoom(cr.GetServiceOptions())
	if t, err := room.JoinRoom(c.ctx, c.Host, c.account, RoomName, c.emitter); err != nil {
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
func (c *client) Mail(req *data.MailboxRequest) (*data.MailboxResponse, *data.SonrError) {
	return c.Service.HandleMailbox(req)
}

// Link handles a LinkRequest
func (c *client) Link(req *data.LinkRequest, t *room.RoomManager) (*data.LinkResponse, *data.SonrError) {
	// Check Request Type
	if req.IsSend() {
		// Find Peer
		if t.HasPeer(req.To.Id.Peer) && t.HasLinker(req.To.Id.Peer) {
			// Get PeerID and Check error
			id, err := t.FindPeer(req.To.Id.Peer)
			if err != nil {
				return nil, data.NewPeerFoundError(err, req.GetTo().GetId().GetPeer())
			}

			// Send Default Invite
			err = c.Service.Link(id, req)
			if err != nil {
				return nil, data.NewError(err, data.ErrorEvent_ROOM_RPC)
			}
			return &data.LinkResponse{
				Success: false,
				Type:    data.LinkResponse_Type(req.GetType()),
			}, nil
		}
		return nil, data.NewErrorWithType(data.ErrorEvent_PEER_NOT_FOUND_INVITE)
	} else {
		c.Service.HandleLinking(req)
		return &data.LinkResponse{
			Success: false,
			Type:    data.LinkResponse_Type(req.GetType()),
		}, nil
	}
}

// Invite Processes Data and Sends Invite to Peer
func (c *client) Invite(invite *data.InviteRequest, t *room.RoomManager) *data.SonrError {
	// Check for Peer
	if invite.GetType() == data.InviteRequest_REMOTE {
		err := c.Service.SendMail(invite)
		if err != nil {
			return err
		}
	} else {
		// Get PeerID and Check error
		id, err := t.FindPeer(invite.To.GetActive().Id.Peer)
		if err != nil {
			c.newExitEvent(invite)
			data.LogError(errors.New("Failed to find Peer ID"))
			return data.NewPeerFoundError(err, invite.GetTo().GetActive().GetId().GetPeer())
		}

		// Initialize Session if transfer
		if invite.IsPayloadTransfer() {
			// Start New Session
			invite.SetProtocol(data.SonrProtocol_LocalTransfer, id)
			c.session = data.NewOutSession(c.account.CurrentDevice(), invite, c.emitter)
		}

		// Run Routine
		go func(inv *data.InviteRequest) {
			// Send Default Invite
			err = c.Service.Invite(id, inv)
			if err != nil {
				c.call.OnError(data.NewError(err, data.ErrorEvent_ROOM_RPC))
				return
			}
		}(invite)
	}
	return nil
}

// Respond Sends a Response to Service
func (c *client) Respond(r *data.InviteResponse) {
	c.Service.Respond(r)
}

// Update proximity/direction and Notify Lobby
func (c *client) Update(t *room.RoomManager) *data.SonrError {
	// Create Event
	ev := c.account.NewUpdateEvent(t.Room(), c.Host.ID())

	// Inform Lobby
	if err := t.Publish(ev); err != nil {
		return data.NewError(err, data.ErrorEvent_ROOM_UPDATE)
	}
	return nil
}

// Handle Network Communication from Lifecycle State Network Communication
func (c *client) Lifecycle(state data.Lifecycle, t *room.RoomManager) {
	if state == data.Lifecycle_ACTIVE {
		// Inform Lobby
		ev := c.account.NewUpdateEvent(t.Room(), c.Host.ID())
		if err := t.Publish(ev); err != nil {
			data.NewError(err, data.ErrorEvent_ROOM_UPDATE)
		}
	} else if state == data.Lifecycle_PAUSED {
		// Inform Lobby
		ev := c.account.NewExitEvent(t.Room(), c.Host.ID())
		if err := t.Publish(ev); err != nil {
			data.NewError(err, data.ErrorEvent_ROOM_UPDATE)
		}
	} else if state == data.Lifecycle_STOPPED {
		// Inform Lobby
		ev := c.account.NewExitEvent(t.Room(), c.Host.ID())
		if err := t.Publish(ev); err != nil {
			data.NewError(err, data.ErrorEvent_ROOM_UPDATE)
		}
		c.Host.Close()
	}
	return
}

// Helper: Creates new Exit Event
func (c *client) newExitEvent(inv *data.InviteRequest) {
	// Create Exit Event
	event := data.RoomEvent{
		Id:      inv.To.GetActive().Id.Peer,
		Subject: data.RoomEvent_EXIT,
	}

	// Marshal Data
	buf, err := proto.Marshal(&event)
	if err != nil {
		return
	}

	// Callback Event and Return Peer Error
	c.call.OnEvent(buf)
	return
}

// Helper: Background Process to continuously ping nearby peers
func (c *client) sendPeriodicRoomEvents(t *room.RoomManager) {
	for {
		// Create Event
		ev := c.account.NewDefaultUpdateEvent(t.Room(), c.Host.ID())

		// Send Update
		if err := t.Publish(ev); err != nil {
			c.call.OnError(data.NewError(err, data.ErrorEvent_ROOM_UPDATE))
			continue
		}
		time.Sleep(util.AUTOUPDATE_INTERVAL)
		data.GetState()
	}
}

// Respond Sends a Response to Service
func (c *client) GetHost() net.HostNode {
	return c.Host
}
