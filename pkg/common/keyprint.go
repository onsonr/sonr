package common

import (
	"encoding/base64"
	"errors"
	fmt "fmt"
	"strconv"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/shengdoushi/base58"
)

// ConvertProtocolCredential creates a common.WebauthnCredential from a webauthn.Credential from the go-webauthn package
func ConvertProtocolCredential(wa *webauthn.Credential) *WebauthnCredential {
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

// ToProtocolCredential converts a common WebauthnCredential to one that can be used for the go-webauthn package
func (c *WebauthnCredential) ToProtocolCredential() *webauthn.Credential {
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

func (c *WebauthnCredential) Did() string {
	return fmt.Sprintf("did:webauth:%s", base58.Encode(c.Id, base58.BitcoinAlphabet))
}

func (c *WebauthnCredential) PublicKeyMultibase() string {
	return "z" + base64.StdEncoding.EncodeToString(c.PublicKey)
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
	}
}

// NewWebAuthnCredential creates a new WebauthnCredential from a ParsedCredentialCreationData and contains all needed information about a WebAuthn credential for storage.
// This is then used to create a VerificationMethod for the DID Document.
func NewWebAuthnCredential(c *protocol.ParsedCredentialCreationData) *WebauthnCredential {
	transportsStr := []string{}
	for _, t := range c.Transports {
		transportsStr = append(transportsStr, string(t))
	}
	return &WebauthnCredential{
		Id:              c.Response.AttestationObject.AuthData.AttData.CredentialID,
		PublicKey:       c.Response.AttestationObject.AuthData.AttData.CredentialPublicKey,
		AttestationType: c.Response.AttestationObject.Format,
		Transport:       transportsStr,
		Authenticator: &WebauthnAuthenticator{
			Aaguid:    c.Response.AttestationObject.AuthData.AttData.AAGUID,
			SignCount: c.Response.AttestationObject.AuthData.Counter,
		},
	}
}

// Validate verifies that this WebauthnCredential is identical to the go-webauthn package credential
func (c *WebauthnCredential) Validate(pc *webauthn.Credential) error {
	if len(c.PublicKey) != len(pc.PublicKey) {
		return errors.New("Credential Public Keys do not match")
	}
	return nil
}
