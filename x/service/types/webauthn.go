package types

import (
	"encoding/base64"
	"errors"
	fmt "fmt"
	"strconv"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/shengdoushi/base58"
	"github.com/sonrhq/core/internal/crypto"
)

// PubKeyFromWebAuthn takes a webauthncose.Key and returns a PubKey
func PubKeyFromWebAuthn(cred *WebauthnCredential) (*crypto.PubKey, error) {
	if cred == nil {
		return nil, errors.New("credential is nil")
	}
	pub, err := webauthncose.ParsePublicKey(cred.PublicKey)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case webauthncose.EC2PublicKeyData:
		return crypto.NewSecp256k1PubKey(pub.XCoord), nil
	case webauthncose.OKPPublicKeyData:
		return crypto.NewEd25519PubKey(pub.XCoord), nil
	default:
		return nil, fmt.Errorf("unsupported public key type: %T", pub)
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

// ConvertStdCredential creates a common.WebauthnCredential from a webauthn.Credential from the go-webauthn package
func ConvertStdCredential(wa *webauthn.Credential) *WebauthnCredential {
	transportsStr := []string{}
	for _, t := range wa.Transport {
		transportsStr = append(transportsStr, string(t))
	}
	return &WebauthnCredential{
		Id:              wa.ID,
		PublicKey:       wa.PublicKey,
		AttestationType: wa.AttestationType,
		Transport:       transportsStr,
		Authenticator: &WebauthnAuthenticator{
			Aaguid:       wa.Authenticator.AAGUID,
			SignCount:    wa.Authenticator.SignCount,
			CloneWarning: wa.Authenticator.CloneWarning,
		},
	}
}

// ToStdCredential converts a common WebauthnCredential to one that can be used for the go-webauthn package
func (c *WebauthnCredential) ToStdCredential() *webauthn.Credential {
	transports := []protocol.AuthenticatorTransport{}
	for _, t := range c.Transport {
		transports = append(transports, protocol.AuthenticatorTransport(t))
	}
	return &webauthn.Credential{
		ID:              c.Id,
		PublicKey:       c.PublicKey,
		AttestationType: c.AttestationType,
		Transport:       transports,
		Authenticator: webauthn.Authenticator{
			AAGUID:       c.Authenticator.Aaguid,
			SignCount:    c.Authenticator.SignCount,
			CloneWarning: c.Authenticator.CloneWarning,
		},
	}
}

// Did returns the DID for a WebauthnCredential
func (c *WebauthnCredential) DID() string {
	return fmt.Sprintf("did:key:%s#%s", c.PubKey().Multibase(), base58.Encode([]byte(c.Id), base58.BitcoinAlphabet))
}

// PublicKeyMultibase returns the public key in multibase format
func (c *WebauthnCredential) PubKey() *crypto.PubKey {
	return crypto.NewEd25519PubKey(c.PublicKey)
}

// CredentialFromDIDString converts a DID string into a WebauthnCredential
func CredentialFromDIDString(did string) (*WebauthnCredential, error) {
	parts := strings.Split(did, "#")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid DID string format")
	}

	multibaseKey := parts[0][8:]
	credIdBz, err := base58.Decode(parts[1], base58.BitcoinAlphabet)
	if err != nil {
		return nil, fmt.Errorf("failed to decode device label: %v", err)
	}

	if !strings.HasPrefix(multibaseKey, "z") {
		return nil, fmt.Errorf("invalid multibase prefix")
	}

	pubKeyBytes, err := base64.StdEncoding.DecodeString(multibaseKey[1:])
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key: %v", err)
	}
	return &WebauthnCredential{PublicKey: pubKeyBytes, Id: credIdBz}, nil
}

// FromMetadata converts a map[string]string into a common WebauthnCredential
func (c *WebauthnCredential) FromMetadata(m map[string]string) error {
	if m["webauthn"] != ConvertBoolToString(true) {
		return errors.New("not a webauthn credential")
	}
	signCount, err := strconv.ParseUint(m["authenticator.sign_count"], 10, 32)
	if err != nil {
		return err
	}
	c.Id, _ = base64.StdEncoding.DecodeString(m["credential_id"])
	c.Authenticator.Aaguid, _ = base64.StdEncoding.DecodeString(m["authenticator.aaguid"])
	c.Authenticator.CloneWarning = ConvertStringToBool(m["authenticator.clone_warning"])
	c.Authenticator.SignCount = uint32(signCount)
	c.Transport = strings.Split(m["transport"], ",")
	c.AttestationType = m["attestation_type"]
	return nil
}

// ToMetadata converts a common WebauthnCredential into a map[string]string
func (c *WebauthnCredential) ToMetadata() map[string]string {
	return map[string]string{
		"credential_id":               base64.StdEncoding.EncodeToString(c.Id),
		"authenticator.aaguid":        base64.StdEncoding.EncodeToString(c.Authenticator.Aaguid),
		"authenticator.clone_warning": ConvertBoolToString(c.Authenticator.CloneWarning),
		"authenticator.sign_count":    strconv.FormatUint(uint64(c.Authenticator.SignCount), 10),
		"transport":                   strings.Join(c.Transport, ","),
		"attestion_type":              c.AttestationType,
		"webauthn":                    ConvertBoolToString(true),
	}
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
