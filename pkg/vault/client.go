package vault

import (
	"context"
	"errors"

	"github.com/di-dao/sonr/internal/local"
	"github.com/di-dao/sonr/pkg/ipfs"
	"github.com/di-dao/sonr/pkg/jwt"
	types "github.com/di-dao/sonr/pkg/wallet"
	"github.com/go-webauthn/webauthn/protocol"
)

type Client interface {
	Assign(ctx context.Context, c *protocol.ParsedCredentialCreationData) error

	Sign(ctx context.Context, msg []byte) ([]byte, error)
	Verify(ctx context.Context, msg, sig []byte) (bool, error)
}

type client struct {
	vfs       ipfs.VFS
	address   string
	sessionId string
}

func (c *client) Assign(ctx context.Context, parsedCreationData *protocol.ParsedCredentialCreationData) error {
	sctx := local.UnwrapCtx(ctx)
	if sctx.UserAddress != c.address {
		return errors.New("user address mismatch")
	}

	if sctx.SessionID != c.sessionId {
		return errors.New("session id mismatch")
	}

	vault, ok := vaultCache.Get(cacheKey(sctx.SessionID))
	if !ok {
		return errors.New("vault not found")
	}

	cred, err := types.MakeNewCredential(parsedCreationData)
	if err != nil {
		return err
	}

	err = vault.Creds.LinkCredential(sctx.ServiceOrigin, cred)
	if err != nil {
		return err
	}
	fmap, err := vault.ToFileMap()
	if err != nil {
		return err
	}
	err = c.vfs.AddFileMap(fmap)
	if err != nil {
		return err
	}
	err = ipfs.PublishFileSystem(ctx, c.vfs)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) GetLoginOptions(ctx context.Context) (protocol.PublicKeyCredentialRequestOptions, error) {
	return protocol.PublicKeyCredentialRequestOptions{}, nil
}

func (c *client) GetRegisterOptions(ctx context.Context) (protocol.PublicKeyCredentialCreationOptions, error) {
	chall := jwt.GenerateChallenge(ctx)
	return jwt.GetRegisterOptions(ctx, chall)
}

func (c *client) Sign(ctx context.Context, msg []byte) ([]byte, error) {
	return nil, nil
}

func (c *client) Verify(ctx context.Context, msg, sig []byte) (bool, error) {
	return false, nil
}
