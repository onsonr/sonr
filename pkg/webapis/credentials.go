package webapis

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"

	modulev1 "github.com/sonrhq/sonr/api/service/module/v1"
)

func GetPublicKeyCredentialCreationOptions(record *modulev1.ServiceRecord, entity protocol.UserEntity) protocol.PublicKeyCredentialCreationOptions {
	return protocol.PublicKeyCredentialCreationOptions{
		RelyingParty: protocol.RelyingPartyEntity{
			ID: record.Origin,
		},
		Challenge: GenerateChallenge(),
		AuthenticatorSelection: protocol.AuthenticatorSelection{
			AuthenticatorAttachment: protocol.CrossPlatform,
			UserVerification:        protocol.VerificationPreferred,
			ResidentKey:             protocol.ResidentKeyRequirementPreferred,
		},
		Attestation: protocol.PreferIndirectAttestation,
		User:        entity,
	}
}

func GetPublicKeyCredentialRequestOptions(record *modulev1.ServiceRecord, creds []protocol.CredentialDescriptor) protocol.PublicKeyCredentialRequestOptions {
	return protocol.PublicKeyCredentialRequestOptions{
		Challenge:          GenerateChallenge(),
		UserVerification:   protocol.VerificationPreferred,
		RelyingPartyID:     record.Origin,
		AllowedCredentials: creds,
	}
}

// GenerateChallenge generates a new challenge for the registration/authentication process.
func GenerateChallenge() protocol.URLEncodedBase64 {
	challenge, err := protocol.CreateChallenge()
	if err != nil {
		panic(err)
	}
	return challenge
}

// GetCOSEAlgorithmIdentifier returns the COSEAlgorithmIdentifier for the given numerical value.
func GetCOSEAlgorithmIdentifier(val int64) (webauthncose.COSEAlgorithmIdentifier, error) {
	if val == -7 {
		return webauthncose.AlgES256, nil
	}
	if val == -8 {
		return webauthncose.AlgEdDSA, nil
	}
	return webauthncose.COSEAlgorithmIdentifier(val), fmt.Errorf("unsupported COSE algorithm identifier: %d", val)
}

// GetAuthenticationSelection returns the authentication selection for the given textual value.
func GetAuthenticationSelection(val string) (protocol.AuthenticatorAttachment, error) {
	if val == "platform" {
		return protocol.Platform, nil
	}
	if val == "cross-platform" {
		return protocol.CrossPlatform, nil
	}
	return protocol.AuthenticatorAttachment(val), fmt.Errorf("unsupported authentication selection: %s", val)
}

// GetCredentialEntity returns the credential entity for the given textual value.
func GetCredentialEntity(val string) (protocol.CredentialEntity, error) {
	return protocol.CredentialEntity{}, nil
}
