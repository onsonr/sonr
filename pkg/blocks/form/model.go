package form

import "github.com/go-webauthn/webauthn/protocol"

type (
	CredentialDescriptor     = protocol.CredentialDescriptor
	AuthenticationExtensions = protocol.AuthenticationExtensions
	LoginOptions             = protocol.PublicKeyCredentialRequestOptions
	RegisterOptions          = protocol.PublicKeyCredentialCreationOptions
)
