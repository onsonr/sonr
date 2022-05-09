package types

import (
	"crypto"

	"github.com/duo-labs/webauthn/webauthn"
	"github.com/fxamacker/cbor/v2"
	rt "go.buf.build/grpc/go/sonr-io/blockchain/registry"
)

func NewCredentialListFromBuf(crds []*rt.Credential) []*Credential {
	creds := make([]*Credential, len(crds))
	for i, crd := range crds {
		creds[i] = NewCredentialFromBuf(crd)
	}
	return creds
}

func NewCredentialFromBuf(crd *rt.Credential) *Credential {
	return &Credential{
		ID:              crd.GetID(),
		PublicKey:       crd.GetPublicKey(),
		AttestationType: crd.GetAttestationType(),
		Authenticator:   NewAuthenticatorFromBuf(crd.GetAuthenticator()),
	}
}

func NewCredentialToBuf(crd *Credential) *rt.Credential {
	return &rt.Credential{
		ID:              crd.GetID(),
		PublicKey:       crd.GetPublicKey(),
		AttestationType: crd.GetAttestationType(),
		Authenticator:   NewAuthenticatorToBuf(crd.GetAuthenticator()),
	}
}

func NewCredentialListToBuf(crds []*Credential) []*rt.Credential {
	creds := make([]*rt.Credential, len(crds))
	for i, crd := range crds {
		creds[i] = NewCredentialToBuf(crd)
	}
	return creds
}

// ConvertToSonrCredential converts a webauthn.Credential to a sonrio.sonr.registry.Credential
func ConvertToSonrCredential(cred webauthn.Credential) *Credential {
	// Create a new credential
	c := &Credential{
		ID:              cred.ID,
		PublicKey:       cred.PublicKey,
		AttestationType: cred.AttestationType,
		Authenticator:   ConvertToSonrAuthenticator(cred.Authenticator),
	}
	return c
}

// ToWebAuthn converts a sonrio.sonr.registry.Credential to a webauthn.Credential
func (cred *Credential) ToWebAuthn() webauthn.Credential {
	return webauthn.Credential{
		ID:              cred.ID,
		PublicKey:       cred.PublicKey,
		AttestationType: cred.AttestationType,
		Authenticator:   cred.Authenticator.ToWebAuthn(),
	}
}

func NewAuthenticatorFromBuf(ath *rt.Authenticator) *Authenticator {
	return &Authenticator{
		Aaguid:       ath.GetAaguid(),
		SignCount:    ath.GetSignCount(),
		CloneWarning: ath.GetCloneWarning(),
	}
}

func NewAuthenticatorToBuf(ath *Authenticator) *rt.Authenticator {
	return &rt.Authenticator{
		Aaguid:       ath.GetAaguid(),
		SignCount:    ath.GetSignCount(),
		CloneWarning: ath.GetCloneWarning(),
	}
}

// ConvertToSonrAuthenticator converts a webauthn.Authenticator to a sonrio.sonr.registry.Authenticator
func ConvertToSonrAuthenticator(auth webauthn.Authenticator) *Authenticator {
	return &Authenticator{
		Aaguid:       auth.AAGUID,
		SignCount:    auth.SignCount,
		CloneWarning: auth.CloneWarning,
	}
}

// ToWebAuthn converts a sonrio.sonr.registry.Authenticator to a webauthn.Authenticator
func (auth *Authenticator) ToWebAuthn() webauthn.Authenticator {
	return webauthn.Authenticator{
		AAGUID:       auth.Aaguid,
		SignCount:    auth.SignCount,
		CloneWarning: auth.CloneWarning,
	}
}

// DecodePublicKey converts a public key from a CBOR encoded byte array to a crypto.PublicKey
func (cred *Credential) DecodePublicKey() (crypto.PublicKey, error) {
	coseKey := COSEKey{}
	err := cbor.Unmarshal(cred.PublicKey, &coseKey)
	if err != nil {
		return nil, err
	}
	pubKey, err := DecodePublicKey(&coseKey)
	if err != nil {
		return nil, err
	}
	return pubKey, nil
}
