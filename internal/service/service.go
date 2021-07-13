package service

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	net "github.com/sonr-io/core/internal/host"
	"google.golang.org/protobuf/proto"

	md "github.com/sonr-io/core/pkg/models"
)

type ServiceHandler interface {
	OnConnected(r *md.ConnectionResponse)
	OnInvite([]byte)
	OnReply(id peer.ID, data []byte)
	OnConfirmed(inv *md.InviteRequest)
	OnMail([]byte)
}

type ServiceClient interface {
	Invite(id peer.ID, inv *md.InviteRequest) error
	Respond(rep *md.InviteResponse)
	SendMail(e *md.InviteRequest) *md.SonrError
	ReadMail() *md.SonrError
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
	Textile *TextileService
}

// @ Creates New Service Interface
func NewService(ctx context.Context, h net.HostNode, u *md.User, req *md.ConnectionRequest, call md.Callback, sh ServiceHandler) (ServiceClient, *md.SonrError) {
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

	// Begin Textile Service
	go func(c *serviceClient) {
		if err := c.StartTextile(); err != nil {
			return
		}
	}(client)

	// Return Instance
	return client, nil
}

// @ Method Reads Inbox and Returns List of Mail Entries
func (sc *serviceClient) ReadMail() *md.SonrError {
	// Check Mail Enabled
	if sc.Textile.options.GetMailbox() {
		// Fetch Mail Event
		event, serr := sc.Textile.readMail()
		if serr != nil {
			return serr
		}
		// Create Mail and Marshal Data
		buf, err := proto.Marshal(event)
		if err != nil {
			return md.NewMarshalError(err)
		}

		// Callback Event
		sc.handler.OnMail(buf)
		md.LogSuccess("Reading Mail")
		return nil
	}
	md.LogInfo("Mail is not Ready")
	return nil
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
		md.LogInfo("To Public Key: " + pubKey.String())

		// Marshal Data
		buf, err := proto.Marshal(inv)
		if err != nil {
			return md.NewMarshalError(err)
		}

		// Send to Mailbox
		resp, serr := sc.Textile.sendMail(pubKey, buf)
		if serr != nil {
			return serr
		}
		sc.handler.OnReply(peer.ID(""), resp)
		md.LogSuccess("Sending Mail")
		return nil
	} else {
		md.LogInfo("Mail is not Ready")
	}
	return nil
}
