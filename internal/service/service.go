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

// // Method Sends Mail Entry to Peer
// func (sc *serviceClient) SendMail(inv *data.InviteRequest) *data.SonrError {
// 	// Check Mail Enabled
// 	if sc.Textile.options.GetMailbox() {
// 		// Fetch Peer Thread Key
// 		pubKey, serr := inv.GetTo().GetActive().ThreadKey()
// 		if serr != nil {
// 			return serr
// 		}

// 		// Marshal Data
// 		buf, err := protojson.Marshal(inv)
// 		if err != nil {
// 			return data.NewMarshalError(err)
// 		}

// 		// Send to Mailbox
// 		serr = sc.Textile.sendMail(pubKey, buf)
// 		if serr != nil {
// 			return serr
// 		}
// 		logger.Info("Succesfully sent mail!")
// 		return nil
// 	}
// 	return nil
// }

// // Method Handles a given Mailbox Request for a Message
// func (sc *serviceClient) HandleMailbox(req *data.MailboxRequest) (*data.MailboxResponse, *data.SonrError) {
// 	if req.Action == data.MailboxRequest_READ {
// 		// Set Mailbox Message as Read
// 		err := sc.Textile.readMessage(req.ID)
// 		if err != nil {
// 			return &data.MailboxResponse{
// 				Success: false,
// 				Action:  data.MailboxResponse_Action(req.Action),
// 			}, err
// 		}

// 		// Return Success
// 		return &data.MailboxResponse{
// 			Success: true,
// 			Action:  data.MailboxResponse_Action(req.Action),
// 		}, nil
// 	} else if req.Action == data.MailboxRequest_DELETE {
// 		// Delete Mailbox Message
// 		err := sc.Textile.deleteMessage(req.ID)
// 		if err != nil {
// 			return &data.MailboxResponse{
// 				Success: false,
// 				Action:  data.MailboxResponse_Action(req.Action),
// 			}, err
// 		}
// 		return &data.MailboxResponse{
// 			Success: true,
// 			Action:  data.MailboxResponse_Action(req.Action),
// 		}, nil
// 	} else {
// 		return &data.MailboxResponse{
// 			Success: false,
// 			Action:  data.MailboxResponse_Action(req.Action),
// 		}, data.NewErrorWithType(data.ErrorEvent_MAILBOX_ACTION_INVALID)
// 	}
// }
