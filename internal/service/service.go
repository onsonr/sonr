package service

// import (
// 	"context"

// 	"github.com/libp2p/go-libp2p-core/peer"
// 	net "github.com/sonr-io/core/internal/host"
// 	"github.com/sonr-io/core/pkg/data"
// 	"github.com/sonr-io/core/tools/emitter"
// 	"github.com/sonr-io/core/tools/logger"
// 	"google.golang.org/protobuf/encoding/protojson"
// )

// type ServiceClient interface {
// 	HandleLinking(req *data.LinkRequest)
// 	Link(id peer.ID, inv *data.LinkRequest) error
// 	Invite(id peer.ID, inv *data.InviteRequest) error
// 	Respond(rep *data.InviteResponse)
// 	SendMail(e *data.InviteRequest) *data.SonrError
// 	HandleMailbox(req *data.MailboxRequest) (*data.MailboxResponse, *data.SonrError)
// 	Close()
// }

// type serviceClient struct {
// 	ServiceClient

// 	// Common
// 	ctx     context.Context
// 	apiKeys *data.APIKeys
// 	emitter *emitter.Emitter
// 	host    net.HostNode
// 	request *data.ConnectionRequest
// 	device  *data.Device

// 	// Services
// 	Auth    *AuthService
// 	Textile *TextileService
// }

// // Creates New Service Interface
// func NewService(ctx context.Context, h net.HostNode, u *data.Device, req *data.ConnectionRequest, em *emitter.Emitter) (ServiceClient, *data.SonrError) {
// 	// Create Client
// 	client := &serviceClient{
// 		ctx:     ctx,
// 		apiKeys: req.GetApiKeys(),
// 		emitter: em,
// 		host:    h,
// 		request: req,
// 		device:  u,
// 	}

// 	// Begin Auth Service
// 	err := client.StartAuth()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Begin Textile Service
// 	go func(c *serviceClient) {
// 		if err := c.StartTextile(); err != nil {
// 			return
// 		}
// 	}(client)

// 	// Return Instance
// 	return client, nil
// }
