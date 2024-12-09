package passkeys

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/pkg/common"
)

func Create(c echo.Context, handle string, ks mpc.Keyset) protocol.PublicKeyCredentialCreationOptions {
	origin := c.Request().Host
	svcName := c.Request().Host
	addr := ks.Address()
	return buildRegisterOptions(addr, handle, ks, origin, svcName)
}

func buildRegisterOptions(addr string, handle string, ks mpc.Keyset, origin string, svcName string) protocol.PublicKeyCredentialCreationOptions {
	return protocol.PublicKeyCredentialCreationOptions{
		Attestation:            protocol.PreferDirectAttestation,
		AttestationFormats:     defaultPrimaryAttestationFormats(),
		AuthenticatorSelection: defaultAuthenticatorSelection(),
		RelyingParty:           buildServiceEntity(origin, svcName),
		Extensions:             buildExtensions(ks),
		Parameters:             buildCredentialParameters(),
		Timeout:                10000,
		User:                   buildUserEntity(addr, handle),
	}
}

func buildExtensions(ks mpc.Keyset) protocol.AuthenticationExtensions {
	return protocol.AuthenticationExtensions{
		"largeBlob": common.LargeBlob{
			Support: "required",
			Write:   ks.UserJSON(),
		},
		"payment": common.Payment{
			IsPayment: true,
		},
	}
}

func buildServiceEntity(name string, host string) protocol.RelyingPartyEntity {
	return protocol.RelyingPartyEntity{
		CredentialEntity: protocol.CredentialEntity{
			Name: name,
		},
		ID: host,
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
