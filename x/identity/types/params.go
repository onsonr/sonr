package types

import (
	fmt "fmt"

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
		DidMethodVersion:                "0.6.3",
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

// NewWebauthnCreationOptions returns the webauthn creation options.
func (p Params) NewWebauthnCreationOptions(s *Service, uuid string, challenge protocol.URLEncodedBase64) (protocol.CredentialCreation, error) {


	// Build the credential creation options.
	opts := protocol.PublicKeyCredentialCreationOptions{
		// Generated Challenge.
		Challenge: challenge,

		// Service resulting properties.
		User: s.GetUserEntity(uuid),

		// Preconfigured parameters.
		Parameters: []protocol.CredentialParameter{
			{
				Type:      protocol.PublicKeyCredentialType,
				Algorithm: webauthncose.AlgES256,
			},
		},
		RelyingParty: protocol.RelyingPartyEntity{
			CredentialEntity: protocol.CredentialEntity{
				Name: s.Name,
			},
			ID:   s.Origin,
		},
		Timeout: int(p.WebauthnTimeout),
		AuthenticatorSelection: protocol.AuthenticatorSelection{
			AuthenticatorAttachment: protocol.AuthenticatorAttachment(p.WebauthnAuthenticatorAttachment),
		},
		Attestation: protocol.ConveyancePreference(p.WebauthnAttestionPreference),
	}
	return protocol.CredentialCreation{Response: opts}, nil
}

// NewWebauthnAssertionOptions returns the webauthn assertion options.
func (p Params) NewWebauthnAssertionOptions(s *Service, uuid string, deviceLabel string) (protocol.CredentialAssertion, error) {
	// Issue the challenge.
	chal, err := s.IssueChallenge()
	if err != nil {
		return protocol.CredentialAssertion{}, fmt.Errorf("failed to issue challenge: %w", err)
	}

	// Build the credential assertion options.
	opts := protocol.PublicKeyCredentialRequestOptions{
		// Generated Challenge.
		Challenge: chal,

		// Preconfigured parameters.
		Timeout: int(p.WebauthnTimeout),
	}
	return protocol.CredentialAssertion{Response: opts}, nil
}
