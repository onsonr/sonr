package host

import (
	"context"
	"crypto/rand"
	"time"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/api/common"
)

func GetRandomUser() (thread.Identity, error) {
	privateKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}
	myIdentity := thread.NewLibp2pIdentity(privateKey)
	return myIdentity, nil
}

func NewUserAuthCtx(ctx context.Context, userGroupKey string, userGroupSecret string) (context.Context, error) {
	// Add our user group key to the context
	ctx = common.NewAPIKeyContext(ctx, userGroupKey)

	// Add a signature using our user group secret
	return common.CreateAPISigContext(ctx, time.Now().Add(time.Minute), userGroupSecret)
}

func NewTokenCtx(ctx context.Context, cli *client.Client, user thread.Identity) (context.Context, error) {
	// Generate a new token for the user
	token, err := cli.GetToken(ctx, user)
	if err != nil {
		return nil, err
	}
	return thread.NewTokenContext(ctx, token), nil
}
