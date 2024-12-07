package webauth

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common"
)

func buildRegisterOptions(user protocol.UserEntity, blob common.LargeBlob, service protocol.RelyingPartyEntity) protocol.PublicKeyCredentialCreationOptions {
	return protocol.PublicKeyCredentialCreationOptions{
		Timeout:     10000,
		Attestation: protocol.PreferDirectAttestation,
		AuthenticatorSelection: protocol.AuthenticatorSelection{
			AuthenticatorAttachment: "platform",
			ResidentKey:             protocol.ResidentKeyRequirementPreferred,
			UserVerification:        "preferred",
		},
		RelyingParty: service,
		User:         user,
		Extensions: protocol.AuthenticationExtensions{
			"largeBlob": blob,
		},
		Parameters: []protocol.CredentialParameter{
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
		},
	}
}

func buildLargeBlob(userKeyshareJSON string) common.LargeBlob {
	return common.LargeBlob{
		Support: "required",
		Write:   userKeyshareJSON,
	}
}

func buildUserEntity(userAddress string, userHandle string) protocol.UserEntity {
	return protocol.UserEntity{
		ID:          userAddress,
		DisplayName: userHandle,
		CredentialEntity: protocol.CredentialEntity{
			Name: userAddress,
		},
	}
}

func buildServiceEntity(c echo.Context) protocol.RelyingPartyEntity {
	return protocol.RelyingPartyEntity{
		CredentialEntity: protocol.CredentialEntity{
			Name: "Sonr.ID",
		},
		ID: c.Request().Host,
	}
}
