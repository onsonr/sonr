package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"github.com/sonr-io/core/pkg/crypto"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
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

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// NewWebauthnCreationOptions returns the webauthn creation options.
func (p Params) NewWebauthnCreationOptions(rp protocol.RelyingPartyEntity, alias string, challenge protocol.URLEncodedBase64) (protocol.CredentialCreation, error) {
	// entityUser is the user entity for which the credential is being created.
	entityUser := protocol.UserEntity{
		ID:          crypto.Base64Encode([]byte(alias)),
		DisplayName: alias,
		CredentialEntity: protocol.CredentialEntity{
			Name: alias,
		},
	}

	// If the user is using a mobile device, we don't want to require a resident key.
	return protocol.CredentialCreation{
		Response: protocol.PublicKeyCredentialCreationOptions{
			Challenge:    challenge,                                 // a random value generated by the server to prevent replay attacks.
			Timeout:      int(60000),                                // the time allowed for the operation.
			User:         entityUser,                                // the user entity for which the credential is being created.
			Parameters:   defaultRegistrationCredentialParameters(), // the public key credential parameters specifying the cryptographic parameters for the credential.
			RelyingParty: rp,                                        // the relying party entity which is the service record that the user is interacting with.

			// preferences about the authenticator to be used.
			AuthenticatorSelection: protocol.AuthenticatorSelection{
				ResidentKey:             protocol.ResidentKeyRequirementRequired,
				AuthenticatorAttachment: protocol.Platform,
			},
			Attestation: protocol.PreferDirectAttestation, // preference for attestation conveyance.
		},
	}, nil
}

// NewWebauthnAssertionOptions returns the webauthn assertion options.
func (p Params) NewWebauthnAssertionOptions(s *ServiceRecord, challenge protocol.URLEncodedBase64, allowedCredentials []protocol.CredentialDescriptor) (protocol.CredentialAssertion, error) {
	opts := protocol.PublicKeyCredentialRequestOptions{
		Challenge:          challenge,                     // a random value generated by the server to prevent replay attacks.
		UserVerification:   protocol.VerificationRequired, // requirement for user verification.
		RelyingPartyID:     s.Id,                          // identifier of the relying party.
		Timeout:            int(60000),                    // the time allowed for the operation.
		AllowedCredentials: allowedCredentials,            // list of credentials acceptable to the relying party.
	}
	return protocol.CredentialAssertion{Response: opts}, nil
}

// Supported WebAuthn registration credential algorithms.
func defaultRegistrationCredentialParameters() []protocol.CredentialParameter {
	return []protocol.CredentialParameter{
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgES256,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgES384,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgES512,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgRS256,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgRS384,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgRS512,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgPS256,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgPS384,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgPS512,
		},
		{
			Type:      protocol.PublicKeyCredentialType,
			Algorithm: webauthncose.AlgEdDSA,
		},
	}
}
