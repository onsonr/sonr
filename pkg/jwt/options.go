package jwt

import (
	"context"

	"github.com/di-dao/sonr/internal/local"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
)

func GetRegisterOptions(ctx context.Context, challenge protocol.URLEncodedBase64) (protocol.PublicKeyCredentialCreationOptions, error) {
	return protocol.PublicKeyCredentialCreationOptions{
		Challenge:              challenge,
		AuthenticatorSelection: defaultAuthenticationSelection(),
		RelyingParty:           GetRelayingPartyEntity(ctx),
		User:                   GetUserEntity(ctx),
		Parameters:             defaultRegistrationCredentialParameters(),
	}, nil
}

func GetUserEntity(ctx context.Context) protocol.UserEntity {
	snrctx := local.UnwrapContext(ctx)
	return protocol.UserEntity{
		ID:          snrctx.UserAddress,
		DisplayName: snrctx.UserAddress,
		CredentialEntity: protocol.CredentialEntity{
			Name: snrctx.UserAddress,
		},
	}
}

func GetRelayingPartyEntity(ctx context.Context) protocol.RelyingPartyEntity {
	snrctx := local.UnwrapContext(ctx)
	return protocol.RelyingPartyEntity{
		ID: snrctx.ServiceOrigin,
		CredentialEntity: protocol.CredentialEntity{
			Name: snrctx.ServiceOrigin,
		},
	}
}

func defaultAuthenticationSelection() protocol.AuthenticatorSelection {
	return protocol.AuthenticatorSelection{
		AuthenticatorAttachment: protocol.CrossPlatform,
	}
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
