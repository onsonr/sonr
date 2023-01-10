package common

import (
	"io/ioutil"

	"github.com/go-webauthn/webauthn/webauthn"
	crypto "github.com/libp2p/go-libp2p/core/crypto"
	tm_crypto "github.com/tendermint/tendermint/crypto"
	tm_json "github.com/tendermint/tendermint/libs/json"
)

// Extensions are discussed in §9. WebAuthn Extensions (https://www.w3.org/TR/webauthn/#extensions).

// For a list of commonly supported extensions, see §10. Defined Extensions
// (https://www.w3.org/TR/webauthn/#sctn-defined-extensions).

type AuthenticationExtensionsClientOutputs map[string]interface{}

const (
	ExtensionAppID        = "appid"
	ExtensionAppIDExclude = "appidExclude"
)

// Loads a private key from a JSON file and returns a `crypto.PrivKey` interface
func LoadPrivKeyFromJsonPath(path string) (crypto.PrivKey, error) {
	// Load the key from the given path.
	key, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Create new private key interface
	var vnPk tm_crypto.PrivKey

	// Unmarshal the key into the interface.
	err = tm_json.Unmarshal(key, &vnPk)
	if err != nil {
		return nil, err
	}
	priv, err := crypto.UnmarshalPrivateKey(vnPk.Bytes())
	if err != nil {
		return nil, err
	}
	return priv, nil
}

// It converts a `WebauthnCredential` to a `webauthn.Credential`
func ConvertToWebauthnCredential(credential *WebauthnCredential) webauthn.Credential {
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

// It converts a Package Struct to a WebauthnCredential Struct
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

func ConvertBoolToString(v bool) string {
	if v {
		return "TRUE"
	} else {
		return "FALSE"
	}
}

func ConvertStringToBool(v string) bool {
	if v == "TRUE" {
		return true
	}
	return false
}
