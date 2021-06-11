package host

import (
	"context"
	"crypto/tls"
	"time"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/api/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// @ Initializes New Textile Instance
func (hn *hostNode) StartTextile() *md.SonrError {
	// Initialize
	var err error
	creds := credentials.NewTLS(&tls.Config{})
	auth := common.Credentials{}

	// Dial GRPC
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(auth)}
	hn.tileClient, err = client.NewClient(textileApiUrl, opts...)
	if err != nil {
		return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
	}

	// Get Identity
	hn.tileIdentity, err = getIdentity(hn.keyPair.PrivKey())
	if err != nil {
		return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
	}

	// Create Auth Context
	hn.ctxTileAuth, err = newUserAuthCtx(context.Background(), hn.apiKeys)
	if err != nil {
		return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
	}

	// Create Token Context
	hn.ctxTileToken, err = hn.newTokenCtx()
	if err != nil {
		return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
	}

	// Return Instance
	return nil
}

// # Helper: Gets Thread Identity from Private Key
func getIdentity(privateKey crypto.PrivKey) (thread.Identity, error) {
	myIdentity := thread.NewLibp2pIdentity(privateKey)
	return myIdentity, nil
}

// # Helper: Creates User Auth Context from API Keys
func newUserAuthCtx(ctx context.Context, keys *md.APIKeys) (context.Context, error) {
	// Add our user group key to the context
	ctx = common.NewAPIKeyContext(ctx, keys.TextileKey)

	// Add a signature using our user group secret
	return common.CreateAPISigContext(ctx, time.Now().Add(time.Minute), keys.TextileSecret)
}

// # Helper: Creates Auth Token Context from AuthContext, Client, Identity
func (hn *hostNode) newTokenCtx() (context.Context, error) {
	// Generate a new token for the user
	token, err := hn.tileClient.GetToken(hn.ctxTileAuth, hn.tileIdentity)
	if err != nil {
		return nil, err
	}
	return thread.NewTokenContext(hn.ctxTileAuth, token), nil
}
