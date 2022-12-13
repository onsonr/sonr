package types

import "github.com/go-webauthn/webauthn/webauthn"

func (wvm *VerificationMethod) WebAuthnID() []byte {
	return []byte(wvm.ID)
}

func (wvm *VerificationMethod) WebAuthnName() string {
	return wvm.ID
}

func (wvm *VerificationMethod) WebAuthnDisplayName() string {
	return wvm.ID
}

func (wvm *VerificationMethod) WebAuthnIcon() string {
	return ""
}

func (wvm *VerificationMethod) WebAuthnCredentials() []webauthn.Credential {
	return []webauthn.Credential{
		convertToWebauthnCredential(wvm.WebauthnCredential),
	}
}

func convertToWebauthnCredential(credential *WebauthnCredential) webauthn.Credential {
	return webauthn.Credential{
		ID:        credential.Id,
		PublicKey: credential.PublicKey,
		Authenticator: webauthn.Authenticator{
			AAGUID:       credential.Authenticator.Aaguid,
			SignCount:    credential.Authenticator.SignCount,
			CloneWarning: credential.Authenticator.CloneWarning,
		},
	}
}

func ConvertFromWebauthnCredential(credential *webauthn.Credential) *WebauthnCredential {
	return &WebauthnCredential{
		Id:        credential.ID,
		PublicKey: credential.PublicKey,
		Authenticator: &WebauthnAuthenticator{
			Aaguid:       credential.Authenticator.AAGUID,
			SignCount:    credential.Authenticator.SignCount,
			CloneWarning: credential.Authenticator.CloneWarning,
		},
	}
}

// VerifyCounter
// Step 17 of §7.2. about verifying attestation. If the signature counter value authData.signCount
// is nonzero or the value stored in conjunction with credential’s id attribute is nonzero, then
// run the following sub-step:
//
//	If the signature counter value authData.signCount is
//
//	→ Greater than the signature counter value stored in conjunction with credential’s id attribute.
//	Update the stored signature counter value, associated with credential’s id attribute, to be the value of
//	authData.signCount.
//
//	→ Less than or equal to the signature counter value stored in conjunction with credential’s id attribute.
//	This is a signal that the authenticator may be cloned, see CloneWarning above for more information.
func (a *WebauthnAuthenticator) UpdateCounter(authDataCount uint32) {
	if authDataCount <= a.SignCount && (authDataCount != 0 || a.SignCount != 0) {
		a.CloneWarning = true
		return
	}
	a.SignCount = authDataCount
}
