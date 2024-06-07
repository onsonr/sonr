package vault

import (
	"context"

	"github.com/go-webauthn/webauthn/protocol"
)

type Client interface {
	Assign(ctx context.Context, c *protocol.ParsedCredentialCreationData) error
	GetLoginOptions(ctx context.Context, challenge protocol.URLEncodedBase64) (protocol.PublicKeyCredentialRequestOptions, error)
	Sign(ctx context.Context, msg []byte) ([]byte, error)
	Verify(ctx context.Context, msg, sig []byte) (bool, error)
}

type client struct {
	address string
}

func (c *client) Assign(ctx context.Context, cred *protocol.ParsedCredentialCreationData) error {
	return nil
}

func (c *client) GetLoginOptions(ctx context.Context, challenge protocol.URLEncodedBase64) (protocol.PublicKeyCredentialRequestOptions, error) {
	return protocol.PublicKeyCredentialRequestOptions{}, nil
}

func (c *client) Sign(ctx context.Context, msg []byte) ([]byte, error) {
	return nil, nil
}

func (c *client) Verify(ctx context.Context, msg, sig []byte) (bool, error) {
	return false, nil
}
