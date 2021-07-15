package service

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	net "github.com/sonr-io/core/internal/host"
	"google.golang.org/protobuf/encoding/protojson"

	md "github.com/sonr-io/core/pkg/models"
)

type ServiceHandler interface {
	OnConnected(r *md.ConnectionResponse)
	OnInvite([]byte)
	OnReply(id peer.ID, data []byte)
	OnConfirmed(inv *md.InviteRequest)
	OnMail(e *md.MailEvent)
	OnError(err *md.SonrError)
}

type ServiceClient interface {
	Invite(id peer.ID, inv *md.InviteRequest) error
	Respond(rep *md.InviteResponse)
	SendMail(e *md.InviteRequest) *md.SonrError
	ReadMail() *md.SonrError
	PushSingle(*md.PushMessage) *md.SonrError
	PushMultiple(*md.PushMessage, []*md.Peer) *md.SonrError
	Close()
}

type serviceClient struct {
	ServiceClient

	// Common
	ctx     context.Context
	apiKeys *md.APIKeys
	handler ServiceHandler
	host    net.HostNode
	request *md.ConnectionRequest
	user    *md.User

	// Services
	Auth    *AuthService
	Device  *DeviceService
	Push    *PushService
	Textile *TextileService
}

// @ Creates New Service Interface
func NewService(ctx context.Context, h net.HostNode, u *md.User, req *md.ConnectionRequest, sh ServiceHandler) (ServiceClient, *md.SonrError) {
	// Create Client
	client := &serviceClient{
		ctx:     ctx,
		apiKeys: req.GetApiKeys(),
		handler: sh,
		host:    h,
		request: req,
		user:    u,
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

// @ Method Sends Mail Entry to Peer
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
	return nil
}

// @ Method Sends Push Notification to Peer
func (sc *serviceClient) PushSingle(msg *md.PushMessage) *md.SonrError {
	if isPushEnabled {
		return sc.Push.push(msg)
	}
	return nil
}

// @ Method Send Multiple Push Notifications to Peers
func (sc *serviceClient) PushMultiple(msg *md.PushMessage, peers []*md.Peer) *md.SonrError {
	if isPushEnabled {
		return sc.Push.pushMulti(msg, peers)
	}
	return nil
}
