package passkeys

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"github.com/onsonr/sonr/pkg/common"

	"github.com/labstack/echo/v4"
)

type (
	CredDescriptor  = protocol.CredentialDescriptor
	LoginOptions    = protocol.PublicKeyCredentialRequestOptions
	RegisterOptions = protocol.PublicKeyCredentialCreationOptions
)

func GetLoginOptions(c echo.Context, credentials []common.CredDescriptor) *common.LoginOptions {
	return &common.LoginOptions{
		Timeout:            10000,
		AllowedCredentials: credentials,
	}
}

//
// func GetRegisterOptions(subject string) *common.RegisterOptions {
// 	opts.User = buildUserEntity(subject)
// 	return opts
// }

// returns the base options for registering a new user without challenge or user entity.
func baseRegisterOptions(user protocol.UserEntity, blob common.LargeBlob, service protocol.RelyingPartyEntity) *common.RegisterOptions {
	return &protocol.PublicKeyCredentialCreationOptions{
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

func buildUserEntity(userAddress string, userHandle string, vaultCID string) protocol.UserEntity {
	return protocol.UserEntity{
		ID:          vaultCID,
		DisplayName: userHandle,
		CredentialEntity: protocol.CredentialEntity{
			Name: userAddress,
		},
	}
}

func buildServiceEntity(serviceName string, serviceIcon string, serviceOrigin string) protocol.RelyingPartyEntity {
	return protocol.RelyingPartyEntity{
		CredentialEntity: protocol.CredentialEntity{
			Name: serviceName,
		},
		ID: serviceOrigin,
	}
}
