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
	"google.golang.org/protobuf/proto"
)

// Interface: Main Client handles Networking/Identity/Streams
type Client interface {
	// Client Methods
	Connect(cr *md.ConnectionRequest, keys *md.KeyPair) *md.SonrError
	Bootstrap(cr *md.ConnectionRequest) (*net.TopicManager, *md.SonrError)
	Mail(req *md.MailboxRequest) (*md.MailboxResponse, *md.SonrError)
	Invite(invite *md.InviteRequest, t *net.TopicManager) *md.SonrError
	Respond(r *md.InviteResponse)
	Update(t *net.TopicManager) *md.SonrError
	Lifecycle(state md.Lifecycle, t *net.TopicManager)

	// Topic Callbacks
	OnConnected(*md.ConnectionResponse)
	OnEvent(*md.TopicEvent)
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
	ctx     context.Context
	call    md.Callback
	user    *md.User
	session *md.Session
	request *md.ConnectionRequest

	// References
	Host    net.HostNode
	Service srv.ServiceClient
}

// ^ NewClient Initializes Node with Router ^
func NewClient(ctx context.Context, u *md.User, call md.Callback) Client {
	return &client{
		ctx:  ctx,
		call: call,
		user: u,
	}
}

// @ Connects Host Node from Private Key
func (c *client) Connect(cr *md.ConnectionRequest, keys *md.KeyPair) *md.SonrError {
	// Set Request
	c.request = cr

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
	err = c.user.SetPeer(hn.ID(), maddr)
	if err != nil {
		return err
	}

	// Set Host
	c.Host = hn
	return nil
}

// @ Begins Bootstrapping HostNode
func (c *client) Bootstrap(cr *md.ConnectionRequest) (*net.TopicManager, *md.SonrError) {
	// Bootstrap Host
	err := c.Host.Bootstrap()
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
	topicName := c.user.NewLocalTopic(cr.GetServiceOptions())
	if t, err := c.Host.JoinTopic(c.ctx, c.user, topicName, c); err != nil {
		return nil, err
	} else {
		// Check if Auto Update Events
		if cr.GetServiceOptions().GetAutoUpdate() {
			go c.sendPeriodicTopicEvents(t)
		}
		return t, nil
	}
}

// @ Handle a Mailbox Request from Node
func (c *client) Mail(req *md.MailboxRequest) (*md.MailboxResponse, *md.SonrError) {
	return c.Service.HandleMailbox(req)
}

// @ Invite Processes Data and Sends Invite to Peer
func (c *client) Invite(invite *md.InviteRequest, t *net.TopicManager) *md.SonrError {
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
				id, err := t.FindPeerInTopic(invite.To.Id.Peer)
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
					// Send Invite
					err = c.Service.Invite(id, inv)
					if err != nil {
						c.call.OnError(md.NewError(err, md.ErrorEvent_TOPIC_RPC))
						return
					}
				}(invite)
			} else {
				c.newExitEvent(invite)
				return md.NewErrorWithType(md.ErrorEvent_PEER_NOT_FOUND_INVITE)
			}
		}
		return nil
	}
	return nil
}

// @ Respond Sends a Response to Service
func (c *client) Respond(r *md.InviteResponse) {
	c.Service.Respond(r)
}

// @ Update proximity/direction and Notify Lobby
func (c *client) Update(t *net.TopicManager) *md.SonrError {
	if c.user.IsReady() {
		// Create Event
		ev := c.user.NewUpdateEvent(t.Topic(), c.Host.ID())

		// Inform Lobby
		if err := t.Publish(ev); err != nil {
			return md.NewError(err, md.ErrorEvent_TOPIC_UPDATE)
		}
	}
	return nil
}

// @ Handle Network Communication from Lifecycle State Network Communication
func (c *client) Lifecycle(state md.Lifecycle, t *net.TopicManager) {
	if state == md.Lifecycle_ACTIVE {
		// Inform Lobby
		if c.user.IsReady() {
			ev := c.user.NewUpdateEvent(t.Topic(), c.Host.ID())
			if err := t.Publish(ev); err != nil {
				md.NewError(err, md.ErrorEvent_TOPIC_UPDATE)
			}
		}
	} else if state == md.Lifecycle_PAUSED {
		// Inform Lobby
		if c.user.IsReady() {
			ev := c.user.NewExitEvent(t.Topic(), c.Host.ID())
			if err := t.Publish(ev); err != nil {
				md.NewError(err, md.ErrorEvent_TOPIC_UPDATE)
			}
		}
	} else if state == md.Lifecycle_STOPPED {
		// Inform Lobby
		if c.user.IsReady() {
			ev := c.user.NewExitEvent(t.Topic(), c.Host.ID())
			if err := t.Publish(ev); err != nil {
				md.NewError(err, md.ErrorEvent_TOPIC_UPDATE)
			}
		}
		c.Host.Close()
	}
	return
}

func (c *client) newExitEvent(inv *md.InviteRequest) {
	// Create Exit Event
	event := md.TopicEvent{
		Id:      inv.To.Id.Peer,
		Subject: md.TopicEvent_EXIT,
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

// # Helper: Background Process to continuously ping nearby peers
func (s *client) sendPeriodicTopicEvents(t *net.TopicManager) {
	for {
		if s.user.IsReady() {
			// Create Event
			ev := s.user.NewDefaultUpdateEvent(t.Topic(), s.Host.ID())

			// Send Update
			if err := t.Publish(ev); err != nil {
				s.call.OnError(md.NewError(err, md.ErrorEvent_TOPIC_UPDATE))
				continue
			}
		}
		time.Sleep(3 * time.Second)
		md.GetState()
	}
}
