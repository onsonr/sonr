package service

import (
	"context"
	"log"

	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	net "github.com/sonr-io/core/internal/host"

	crypto "github.com/libp2p/go-libp2p-crypto"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/api/common"
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

	// Status's
	isAuthReady    bool
	isDeviceReady  bool
	isMailReady    bool
	isThreadsReady bool
	isBucketsReady bool
	isHttpReady    bool
}

func NewService(ctx context.Context, h net.HostNode, u *md.User, req *md.ConnectionRequest, call md.Callback, sh ServiceHandler) (ServiceClient, *md.SonrError) {
	// Create Client
	client := &serviceClient{
		ctx:            ctx,
		apiKeys:        req.GetApiKeys(),
		handler:        sh,
		host:           h,
		request:        req,
		isAuthReady:    false,
		isDeviceReady:  false,
		isMailReady:    false,
		isThreadsReady: false,
		isBucketsReady: false,
		isHttpReady:    false,
		user:           u,
	}

	// Begin Auth Service
	err := client.StartAuth()
	if err != nil {
		return nil, err
	}

	// Begin Textile Service
	go func(c *serviceClient) {
		if err := c.StartTextile(); err != nil {
			log.Println(err)
			return
		}
	}(client)

	// Return Instance
	return client, nil
}

// # Helper: Gets Thread Identity from Private Key
func getIdentity(privateKey crypto.PrivKey) thread.Identity {
	myIdentity := thread.NewLibp2pIdentity(privateKey)
	return myIdentity
}

// # Helper: Creates User Auth Context from API Keys
func newUserAuthCtx(ctx context.Context, keys *md.APIKeys) (context.Context, error) {
	// Add our user group key to the context
	ctx = common.NewAPIKeyContext(ctx, keys.TextileKey)

	// Add a signature using our user group secret
	return common.CreateAPISigContext(ctx, time.Now().Add(time.Hour), keys.TextileSecret)
}

// # Helper: Creates Auth Token Context from AuthContext, Client, Identity
func (tn *TextileService) newTokenCtx() (context.Context, error) {
	// Generate a new token for the user
	token, err := tn.client.GetToken(tn.ctxAuth, tn.identity)
	if err != nil {
		return nil, err
	}
	return thread.NewTokenContext(tn.ctxAuth, token), nil
}

// @ Set Service Status for Auth
func (sc *serviceClient) SetAuthStatus(val bool) {
	sc.isAuthReady = val
}

// @ Set Service Status for Device
func (sc *serviceClient) SetDeviceStatus(val bool) {
	sc.isDeviceReady = val
}

// @ Set Service Status for Buckets
func (sc *serviceClient) SetBucketsStatus(val bool) {
	sc.isBucketsReady = val
}

// @ Set Service Status for HTTP
func (sc *serviceClient) SetHTTPStatus(val bool) {
	sc.isHttpReady = val
}

// @ Set Service Status for Mailbox
func (sc *serviceClient) SetMailboxStatus(val bool) {
	sc.isMailReady = val
}

// @ Set Service Status for Threads
func (sc *serviceClient) SetThreadsStatus(val bool) {
	sc.isThreadsReady = val
}
