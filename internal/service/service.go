package service

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	net "github.com/sonr-io/core/internal/host"
	"google.golang.org/protobuf/encoding/protojson"

	md "github.com/sonr-io/core/pkg/data"
)

type ServiceHandler interface {
	OnConnected(r *md.ConnectionResponse)
	OnLink(success bool, incoming bool, id peer.ID, resp []byte)
	OnInvite([]byte)
	OnReply(id peer.ID, data []byte)
	OnConfirmed(inv *md.InviteRequest)
	OnMail(e *md.MailEvent)
	OnError(err *md.SonrError)
}

type ServiceClient interface {
	HandleLinking(req *md.LinkRequest)
	Link(id peer.ID, inv *md.LinkRequest) error
	Invite(id peer.ID, inv *md.InviteRequest) error
	Respond(rep *md.InviteResponse)
	SendMail(e *md.InviteRequest) *md.SonrError
	HandleMailbox(req *md.MailboxRequest) (*md.MailboxResponse, *md.SonrError)
	PushSingle(*md.PushMessage) *md.SonrError
	PushMultiple(*md.PushMessage, []*md.Peer) *md.SonrError
	Close()
}

type serviceClient struct {
	ServiceClient

	// Common
	ctx       context.Context
	apiKeys   *md.APIKeys
	handler   ServiceHandler
	host      net.HostNode
	pushToken string
	request   *md.ConnectionRequest
	device    *md.Device

	// Services
	Auth    *AuthService
	Device  *DeviceService
	Push    *PushService
	Textile *TextileService
}

// Creates New Service Interface
func NewService(ctx context.Context, h net.HostNode, u *md.Device, req *md.ConnectionRequest, sh ServiceHandler) (ServiceClient, *md.SonrError) {
	// Create Client
	client := &serviceClient{
		ctx:       ctx,
		apiKeys:   req.GetApiKeys(),
		handler:   sh,
		host:      h,
		pushToken: req.GetPushToken(),
		request:   req,
		device:    u,
	}

	// Begin Auth Service
	err := client.StartAuth()
	if err != nil {
		return nil, err
	}

	// Begin Push Service
	err = client.StartPush()
	if err != nil {
		return nil, err
	}

	// Begin Textile Service
	go func(c *serviceClient) {
		if err := c.StartTextile(); err != nil {
			return
		}
	}(client)

	// Return Instance
	return client, nil
}

// Method Sends Mail Entry to Peer
func (sc *serviceClient) SendMail(inv *md.InviteRequest) *md.SonrError {
	// Check Mail Enabled
	if sc.Textile.options.GetMailbox() {
		// Fetch Peer Thread Key
		pubKey, serr := inv.GetTo().ThreadKey()
		if serr != nil {
			return serr
		}

		// Marshal Data
		buf, err := protojson.Marshal(inv)
		if err != nil {
			return md.NewMarshalError(err)
		}

		// Send to Mailbox
		serr = sc.Textile.sendMail(pubKey, buf)
		if serr != nil {
			return serr
		}
		md.LogSuccess("Sending Mail")
		return nil
	} else {
		md.LogInfo("Mail is not Ready")
	}

	// Send Push Message
	serr := sc.Push.push(inv.ToPushMessage())
	if serr != nil {
		return serr
	}
	return nil
}

// Method Handles a given Mailbox Request for a Message
func (sc *serviceClient) HandleMailbox(req *md.MailboxRequest) (*md.MailboxResponse, *md.SonrError) {
	if req.Action == md.MailboxRequest_READ {
		// Set Mailbox Message as Read
		err := sc.Textile.readMessage(req.ID)
		if err != nil {
			return &md.MailboxResponse{
				Success: false,
				Action:  md.MailboxResponse_Action(req.Action),
			}, err
		}

		// Return Success
		return &md.MailboxResponse{
			Success: true,
			Action:  md.MailboxResponse_Action(req.Action),
		}, nil
	} else if req.Action == md.MailboxRequest_DELETE {
		// Delete Mailbox Message
		err := sc.Textile.deleteMessage(req.ID)
		if err != nil {
			return &md.MailboxResponse{
				Success: false,
				Action:  md.MailboxResponse_Action(req.Action),
			}, err
		}
		return &md.MailboxResponse{
			Success: true,
			Action:  md.MailboxResponse_Action(req.Action),
		}, nil
	} else {
		return &md.MailboxResponse{
			Success: false,
			Action:  md.MailboxResponse_Action(req.Action),
		}, md.NewErrorWithType(md.ErrorEvent_MAILBOX_ACTION_INVALID)
	}
}

// Method Sends Push Notification to Peer
func (sc *serviceClient) PushSingle(msg *md.PushMessage) *md.SonrError {
	if isPushEnabled {
		return sc.Push.push(msg)
	}
	return nil
}

// Method Send Multiple Push Notifications to Peers
func (sc *serviceClient) PushMultiple(msg *md.PushMessage, peers []*md.Peer) *md.SonrError {
	if isPushEnabled {
		return sc.Push.pushMulti(msg, peers)
	}
	return nil
}
