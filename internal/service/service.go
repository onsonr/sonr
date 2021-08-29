package service

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	net "github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/data"
	"google.golang.org/protobuf/encoding/protojson"
)

type ServiceHandler interface {
	OnConnected(r *data.ConnectionResponse)
	OnLink(success bool, incoming bool, id peer.ID, resp []byte)
	OnInvite([]byte)
	OnReply(id peer.ID, buf []byte)
	OnConfirmed(inv *data.InviteRequest)
	OnMail(e *data.MailEvent)
	OnError(err *data.SonrError)
}

type ServiceClient interface {
	HandleLinking(req *data.LinkRequest)
	Link(id peer.ID, inv *data.LinkRequest) error
	Invite(id peer.ID, inv *data.InviteRequest) error
	Respond(rep *data.InviteResponse)
	SendMail(e *data.InviteRequest) *data.SonrError
	HandleMailbox(req *data.MailboxRequest) (*data.MailboxResponse, *data.SonrError)
	PushSingle(*data.PushMessage) *data.SonrError
	PushMultiple(*data.PushMessage, []*data.Peer) *data.SonrError
	Close()
}

type serviceClient struct {
	ServiceClient

	// Common
	ctx       context.Context
	apiKeys   *data.APIKeys
	handler   ServiceHandler
	host      net.HostNode
	pushToken string
	request   *data.ConnectionRequest
	device    *data.Device

	// Services
	Auth    *AuthService
	Push    *PushService
	Textile *TextileService
}

// Creates New Service Interface
func NewService(ctx context.Context, h net.HostNode, u *data.Device, req *data.ConnectionRequest, sh ServiceHandler) (ServiceClient, *data.SonrError) {
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
func (sc *serviceClient) SendMail(inv *data.InviteRequest) *data.SonrError {
	// Check Mail Enabled
	if sc.Textile.options.GetMailbox() {
		// Fetch Peer Thread Key
		pubKey, serr := inv.GetTo().GetActive().ThreadKey()
		if serr != nil {
			return serr
		}

		// Marshal Data
		buf, err := protojson.Marshal(inv)
		if err != nil {
			return data.NewMarshalError(err)
		}

		// Send to Mailbox
		serr = sc.Textile.sendMail(pubKey, buf)
		if serr != nil {
			return serr
		}
		data.LogSuccess("Sending Mail")
		return nil
	} else {
		data.LogInfo("Mail is not Ready")
	}

	// Send Push Message
	serr := sc.Push.push(inv.ToPushMessage())
	if serr != nil {
		return serr
	}
	return nil
}

// Method Handles a given Mailbox Request for a Message
func (sc *serviceClient) HandleMailbox(req *data.MailboxRequest) (*data.MailboxResponse, *data.SonrError) {
	if req.Action == data.MailboxRequest_READ {
		// Set Mailbox Message as Read
		err := sc.Textile.readMessage(req.ID)
		if err != nil {
			return &data.MailboxResponse{
				Success: false,
				Action:  data.MailboxResponse_Action(req.Action),
			}, err
		}

		// Return Success
		return &data.MailboxResponse{
			Success: true,
			Action:  data.MailboxResponse_Action(req.Action),
		}, nil
	} else if req.Action == data.MailboxRequest_DELETE {
		// Delete Mailbox Message
		err := sc.Textile.deleteMessage(req.ID)
		if err != nil {
			return &data.MailboxResponse{
				Success: false,
				Action:  data.MailboxResponse_Action(req.Action),
			}, err
		}
		return &data.MailboxResponse{
			Success: true,
			Action:  data.MailboxResponse_Action(req.Action),
		}, nil
	} else {
		return &data.MailboxResponse{
			Success: false,
			Action:  data.MailboxResponse_Action(req.Action),
		}, data.NewErrorWithType(data.ErrorEvent_MAILBOX_ACTION_INVALID)
	}
}

// Method Sends Push Notification to Peer
func (sc *serviceClient) PushSingle(msg *data.PushMessage) *data.SonrError {
	if isPushEnabled {
		return sc.Push.push(msg)
	}
	return nil
}

// Method Send Multiple Push Notifications to Peers
func (sc *serviceClient) PushMultiple(msg *data.PushMessage, peers []*data.Peer) *data.SonrError {
	if isPushEnabled {
		return sc.Push.pushMulti(msg, peers)
	}
	return nil
}
