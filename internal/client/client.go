package client

import (
	"context"

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
	Bootstrap() (*net.TopicManager, *md.SonrError)
	Invite(invite *md.InviteRequest, t *net.TopicManager) *md.SonrError
	ReadMail() *md.SonrError
	Respond(r *md.InviteResponse)
	Update(t *net.TopicManager) *md.SonrError
	Lifecycle(state md.LifecycleState, t *net.TopicManager)
	Restart(ur *md.UpdateRequest, keys *md.KeyPair) (*net.TopicManager, *md.SonrError)

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
	err = c.user.NewPeer(hn.ID(), maddr)
	if err != nil {
		return err
	}

	// Set Host
	c.Host = hn
	return nil
}

// @ Begins Bootstrapping HostNode
func (c *client) Bootstrap() (*net.TopicManager, *md.SonrError) {
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
	if t, err := c.Host.JoinTopic(c.ctx, c.user, c.user.NewLocalTopic(), c); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

// @ Invite Processes Data and Sends Invite to Peer
func (c *client) Invite(invite *md.InviteRequest, t *net.TopicManager) *md.SonrError {
	if c.user.IsReady() {
		// Check for Peer
		if invite.GetType() == md.InviteRequest_Remote {
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
						c.call.OnError(md.NewError(err, md.ErrorMessage_TOPIC_RPC))
						return
					}
				}(invite)
			} else {
				c.newExitEvent(invite)
				return md.NewErrorWithType(md.ErrorMessage_PEER_NOT_FOUND_INVITE)
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

// @ Method Calls to Read Local Mailbox
func (c *client) ReadMail() *md.SonrError {
	//return c.Service.ReadMail()
	return nil
}

// @ Update proximity/direction and Notify Lobby
func (c *client) Update(t *net.TopicManager) *md.SonrError {
	if c.user.IsReady() {
		// Inform Lobby
		if err := t.Publish(c.user.Peer.NewUpdateEvent(t.Topic())); err != nil {
			return md.NewError(err, md.ErrorMessage_TOPIC_UPDATE)
		}
	}
	return nil
}

// @ Handle Network Communication from Lifecycle State Network Communication
func (c *client) Lifecycle(state md.LifecycleState, t *net.TopicManager) {
	if state == md.LifecycleState_Active {
		// Inform Lobby
		if c.user.IsReady() {
			if err := t.Publish(c.user.Peer.NewUpdateEvent(t.Topic())); err != nil {
				md.NewError(err, md.ErrorMessage_TOPIC_UPDATE)
			}
		}
	} else if state == md.LifecycleState_Paused {
		// Inform Lobby
		if c.user.IsReady() {
			if err := t.Publish(c.user.Peer.NewExitEvent(t.Topic())); err != nil {
				md.NewError(err, md.ErrorMessage_TOPIC_UPDATE)
			}
		}
	} else if state == md.LifecycleState_Stopped {
		// Inform Lobby
		if c.user.IsReady() {
			if err := t.Publish(c.user.Peer.NewExitEvent(t.Topic())); err != nil {
				md.NewError(err, md.ErrorMessage_TOPIC_UPDATE)
			}
		}
		c.Host.Close()
	}
}

// @ Restart HostNode on Network Change
func (c *client) Restart(ur *md.UpdateRequest, keys *md.KeyPair) (*net.TopicManager, *md.SonrError) {
	switch ur.Data.(type) {
	case *md.UpdateRequest_Connectivity:
		if c.request != nil {
			// Update Request
			newRequest := c.request
			newRequest.Type = ur.GetConnectivity()

			// Connect
			err := c.Connect(newRequest, keys)
			if err != nil {
				return nil, err
			}

			// Bootstrap
			tpc, err := c.Bootstrap()
			if err != nil {
				return nil, err
			}
			return tpc, nil
		}
	}
	return nil, nil
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
