package passkeys

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
)

func defaultPrimaryAttestationFormats() []protocol.AttestationFormat {
	return []protocol.AttestationFormat{
		protocol.AttestationFormatApple,
		protocol.AttestationFormatAndroidKey,
		protocol.AttestationFormatAndroidSafetyNet,
		protocol.AttestationFormatFIDOUniversalSecondFactor,
	}
}

func defaultSecondaryAttestationFormats() []protocol.AttestationFormat {
	return []protocol.AttestationFormat{
		protocol.AttestationFormatPacked,
		protocol.AttestationFormatTPM,
	}
}

func defaultAuthenticatorSelection() protocol.AuthenticatorSelection {
	return protocol.AuthenticatorSelection{
		AuthenticatorAttachment: protocol.Platform,
		ResidentKey:             protocol.ResidentKeyRequirementRequired,
		UserVerification:        protocol.VerificationRequired,
	}
}

func buildCredentialParameters() []protocol.CredentialParameter {
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
