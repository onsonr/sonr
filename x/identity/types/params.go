package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"

	"github.com/go-webauthn/webauthn/protocol"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		DidBaseContext:                  "https://www.w3.org/ns/did/v1",
		DidMethodContext:                "https://docs.sonr.io/identity/1.0",
		DidMethodName:                   "sonr",
		DidMethodVersion:                "1.0",
		DidNetwork:                      "devnet",
		IpfsGateway:                     "https://sonr.space/ipfs",
		IpfsApi:                         "https://api.sonr.space",
		WebauthnAttestionPreference:     "direct",
		WebauthnAuthenticatorAttachment: "platform",
		WebauthnTimeout:                 60000,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// WebauthnConveyancePreference returns the webauthn conveyance preference.
func (p Params) WebauthnConveyancePreference() protocol.ConveyancePreference {
	return protocol.ConveyancePreference(p.WebauthnAttestionPreference)
}

// WebauthnAuthenticatorSelection returns the authenticator selection for webauthn.
func (p Params) WebauthnAuthenticatorSelection() protocol.AuthenticatorSelection {
	return protocol.AuthenticatorSelection{
		AuthenticatorAttachment: protocol.AuthenticatorAttachment(p.WebauthnAuthenticatorAttachment),
	}
}

// We return ECDSA P-256 with SHA-256 as the default credential parameter.
func (p Params) WebauthnRegistrationCredentialParameters() []protocol.CredentialParameter {
	return []protocol.CredentialParameter{
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgES256,
		},
	}
}

// WebauthnTimeoutInteger returns the webauthn timeout as an integer.
func (p Params) WebauthnTimeoutInteger() int {
	return int(p.WebauthnTimeout)
}
