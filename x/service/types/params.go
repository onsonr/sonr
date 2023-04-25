package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
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
func (p Params) NewWebauthnCreationOptions(s *ServiceRecord, alias string, challenge protocol.URLEncodedBase64, isMobile bool) (protocol.CredentialCreation, error) {
	entityUser := protocol.UserEntity{
		ID:          alias,
		DisplayName: alias,
		CredentialEntity: protocol.CredentialEntity{
			Name: alias,
		},
	}
	return protocol.CredentialCreation{
		Response: protocol.PublicKeyCredentialCreationOptions{
			Challenge:              challenge,
			Timeout:                int(60000),
			User:                   entityUser,
			Parameters:             defaultRegistrationCredentialParameters(),
			RelyingParty:           s.RelyingPartyEntity(),
			AuthenticatorSelection: getUserAuthenticationSelectionForDevice(isMobile),
			Attestation:            protocol.PreferDirectAttestation,
		},
	}, nil
}

// NewWebauthnAssertionOptions returns the webauthn assertion options.
func (p Params) NewWebauthnAssertionOptions(s *ServiceRecord, challenge protocol.URLEncodedBase64, allowedCredentials []protocol.CredentialDescriptor, isMobile bool) (protocol.CredentialAssertion, error) {

	// Build the credential assertion options.
	opts := protocol.PublicKeyCredentialRequestOptions{
		// Generated Challenge.
		Challenge:        challenge,
		RelyingPartyID:   s.Origin,
		UserVerification: getUserVerificationForDevice(isMobile),

		// Preconfigured parameters.
		Timeout:            int(60000),
		AllowedCredentials: allowedCredentials,
	}
	return protocol.CredentialAssertion{Response: opts}, nil
}

func getUserAuthenticationSelectionForDevice(isMobile bool) protocol.AuthenticatorSelection {
	if isMobile {
		return protocol.AuthenticatorSelection{
			ResidentKey:             protocol.ResidentKeyRequirementRequired,
			UserVerification:        protocol.VerificationPreferred,
			AuthenticatorAttachment: protocol.Platform,
		}
	}
	return protocol.AuthenticatorSelection{
		ResidentKey:             protocol.ResidentKeyRequirementPreferred,
		UserVerification:        protocol.VerificationRequired,
		AuthenticatorAttachment: protocol.CrossPlatform,
	}
}

func getUserVerificationForDevice(isMobile bool) protocol.UserVerificationRequirement {
	if isMobile {
		return protocol.VerificationPreferred
	}
	return protocol.VerificationRequired
}

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
