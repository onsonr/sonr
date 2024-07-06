package orm

import "github.com/go-webauthn/webauthn/protocol"

// Authenticator contains all needed information about an authenticator for storage.
type Authenticator struct {
	Attachment   protocol.AuthenticatorAttachment `json:"attachment"`
	AAGUID       []byte                           `json:"AAGUID"`
	SignCount    uint32                           `json:"signCount"`
	CloneWarning bool                             `json:"cloneWarning"`
}

// SelectAuthenticator allow for easy marshaling of authenticator options that are provided to the user.
func SelectAuthenticator(att string, rrk *bool, uv string) protocol.AuthenticatorSelection {
	return protocol.AuthenticatorSelection{
		AuthenticatorAttachment: protocol.AuthenticatorAttachment(att),
		RequireResidentKey:      rrk,
		UserVerification:        protocol.UserVerificationRequirement(uv),
	}
}

// UpdateCounter updates the authenticator and either sets the clone warning value or the sign count.
//
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
func (a *Authenticator) UpdateCounter(authDataCount uint32) {
	if authDataCount <= a.SignCount && (authDataCount != 0 || a.SignCount != 0) {
		a.CloneWarning = true

		return
	}

	a.SignCount = authDataCount
}
