package orm

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
)

func NewCredentialCreationOptions(subject, address string) (*protocol.PublicKeyCredentialCreationOptions, error) {
	chl, err := protocol.CreateChallenge()
	if err != nil {
		return nil, err
	}
	return &protocol.PublicKeyCredentialCreationOptions{
		Challenge: chl,
		User: protocol.UserEntity{
			DisplayName: subject,
			ID:          address,
		},
		Attestation:            defaultAttestation(),
		AuthenticatorSelection: defaultAuthenticatorSelection(),
		Parameters:             defaultCredentialParameters(),
	}, nil
}

func buildUserEntity(userID string) protocol.UserEntity {
	return protocol.UserEntity{
		ID: userID,
	}
}

func defaultAttestation() protocol.ConveyancePreference {
	return protocol.PreferDirectAttestation
}

func defaultAuthenticatorSelection() protocol.AuthenticatorSelection {
	return protocol.AuthenticatorSelection{
		AuthenticatorAttachment: "platform",
		ResidentKey:             protocol.ResidentKeyRequirementPreferred,
		UserVerification:        "preferred",
	}
}

func defaultCredentialParameters() []protocol.CredentialParameter {
	return []protocol.CredentialParameter{
		{
			Type:      "public-key",
			Algorithm: webauthncose.AlgES256,
		},
		{
			Type:      "public-key",
			Algorithm: webauthncose.AlgES256K,
		},
		{
			Type:      "public-key",
			Algorithm: webauthncose.AlgEdDSA,
		},
	}
}
