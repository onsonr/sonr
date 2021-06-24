package service

import (
	"context"

	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	net "github.com/sonr-io/core/internal/host"

	crypto "github.com/libp2p/go-libp2p-crypto"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/api/common"
)

type ServiceHandler interface {
	OnInvite([]byte)
	OnReply(id peer.ID, data []byte)
	OnConfirmed(inv *md.InviteRequest)
}

type ServiceClient interface {
	Invite(id peer.ID, inv *md.InviteRequest) error
	Respond(rep *md.InviteResponse)
	Close()
}

type serviceClient struct {
	ServiceClient

	// Common
	ctx         context.Context
	apiKeys     *md.APIKeys
	handler     ServiceHandler
	host        net.HostNode
	textileOpts *md.ConnectionRequest_TextileOptions
	user        *md.User

	// Services
	Auth    *AuthService
	Device  *DeviceService
	Textile *TextileService
}

func NewService(ctx context.Context, h net.HostNode, u *md.User, req *md.ConnectionRequest, sh ServiceHandler) (ServiceClient, *md.SonrError) {
	// Create Client
	client := &serviceClient{
		ctx:         ctx,
		apiKeys:     req.GetApiKeys(),
		handler:     sh,
		host:        h,
		textileOpts: req.GetTextileOptions(),
		user:        u,
	}

	// Begin Auth Service
	err := client.StartAuth()
	if err != nil {
		return nil, err
	}

	// Begin Textile Service

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
