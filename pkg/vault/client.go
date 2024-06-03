package vault

import (
	"context"

	"github.com/go-webauthn/webauthn/protocol"
)

type Client interface {
	Assign(ctx context.Context, resp *protocol.CredentialCreationResponse) error
}
