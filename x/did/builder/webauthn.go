package builder

import (
	"bytes"
	"encoding/base64"
	"reflect"

	"github.com/go-webauthn/webauthn/protocol"
)

// Credential contains all needed information about a WebAuthn credential for storage.
type Credential struct {
	Subject         string   `json:"handle"`
	AttestationType string   `json:"attestationType"`
	Origin          string   `json:"origin"`
	CredentialID    []byte   `json:"id"`
	PublicKey       []byte   `json:"publicKey"`
	Transport       []string `json:"transport"`
	SignCount       uint32   `json:"signCount"`
	UserPresent     bool     `json:"userPresent"`
	UserVerified    bool     `json:"userVerified"`
	BackupEligible  bool     `json:"backupEligible"`
	BackupState     bool     `json:"backupState"`
	CloneWarning    bool     `json:"cloneWarning"`
}

// NewCredential will return a credential pointer on successful validation of a registration response.
func NewCredential(c *protocol.ParsedCredentialCreationData, origin, handle string) *Credential {
	return &Credential{
		Subject:         handle,
		Origin:          origin,
		AttestationType: c.Response.AttestationObject.Format,
		CredentialID:    c.Response.AttestationObject.AuthData.AttData.CredentialID,
		PublicKey:       c.Response.AttestationObject.AuthData.AttData.CredentialPublicKey,
		Transport:       NormalizeTransports(c.Response.Transports),
		SignCount:       c.Response.AttestationObject.AuthData.Counter,
		UserPresent:     c.Response.AttestationObject.AuthData.Flags.HasUserPresent(),
		UserVerified:    c.Response.AttestationObject.AuthData.Flags.HasUserVerified(),
		BackupEligible:  c.Response.AttestationObject.AuthData.Flags.HasBackupEligible(),
		BackupState:     c.Response.AttestationObject.AuthData.Flags.HasAttestedCredentialData(),
	}
}

// Descriptor converts a Credential into a protocol.CredentialDescriptor.
func (c *Credential) Descriptor() protocol.CredentialDescriptor {
	return protocol.CredentialDescriptor{
		Type:            protocol.PublicKeyCredentialType,
		CredentialID:    c.CredentialID,
		Transport:       ModuleTransportsToProtocol(c.Transport),
		AttestationType: c.AttestationType,
	}
}

// This is a signal that the authenticator may be cloned, see CloneWarning above for more information.
func (a *Credential) UpdateCounter(authDataCount uint32) {
	if authDataCount <= a.SignCount && (authDataCount != 0 || a.SignCount != 0) {
		a.CloneWarning = true
		return
	}

	a.SignCount = authDataCount
}

type CredentialDescriptor struct {
	// The valid credential types.
	Type CredentialType `json:"type"`

	// CredentialID The ID of a credential to allow/disallow.
	CredentialID URLEncodedBase64 `json:"id"`

	// The authenticator transports that can be used.
	Transport []AuthenticatorTransport `json:"transports,omitempty"`

	// The AttestationType from the Credential. Used internally only.
	AttestationType string `json:"-"`
}

func NewCredentialDescriptor(credentialID string, transports []AuthenticatorTransport, attestationType string) *CredentialDescriptor {
	return &CredentialDescriptor{
		CredentialID:    URLEncodedBase64(credentialID),
		Transport:       transports,
		AttestationType: attestationType,
		Type:            CredentialTypePublicKeyCredential,
	}
}

type AuthenticatorSelection struct {
	// AuthenticatorAttachment If this member is present, eligible authenticators are filtered to only
	// authenticators attached with the specified AuthenticatorAttachment enum.
	AuthenticatorAttachment AuthenticatorAttachment `json:"authenticatorAttachment,omitempty"`

	// RequireResidentKey this member describes the Relying Party's requirements regarding resident
	// credentials. If the parameter is set to true, the authenticator MUST create a client-side-resident
	// public key credential source when creating a public key credential.
	RequireResidentKey *bool `json:"requireResidentKey,omitempty"`

	// ResidentKey this member describes the Relying Party's requirements regarding resident
	// credentials per Webauthn Level 2.
	ResidentKey ResidentKeyRequirement `json:"residentKey,omitempty"`

	// UserVerification This member describes the Relying Party's requirements regarding user verification for
	// the create() operation. Eligible authenticators are filtered to only those capable of satisfying this
	// requirement.
	UserVerification UserVerificationRequirement `json:"userVerification,omitempty"`
}

type AuthenticatorData struct {
	RPIDHash []byte                 `json:"rpid"`
	Flags    AuthenticatorFlags     `json:"flags"`
	Counter  uint32                 `json:"sign_count"`
	AttData  AttestedCredentialData `json:"att_data"`
	ExtData  []byte                 `json:"ext_data"`
}

type AttestationObject struct {
	// The authenticator data, including the newly created public key. See AuthenticatorData for more info
	AuthData AuthenticatorData

	// The byteform version of the authenticator data, used in part for signature validation
	RawAuthData []byte `json:"authData"`

	// The format of the Attestation data.
	Format string `json:"fmt"`

	// The attestation statement data sent back if attestation is requested.
	AttStatement map[string]any `json:"attStmt,omitempty"`
}

type URLEncodedBase64 []byte

func (e URLEncodedBase64) String() string {
	return base64.RawURLEncoding.EncodeToString(e)
}

// UnmarshalJSON base64 decodes a URL-encoded value, storing the result in the
// provided byte slice.
func (e *URLEncodedBase64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}

	// Trim the leading spaces.
	data = bytes.Trim(data, "\"")

	// Trim the trailing equal characters.
	data = bytes.TrimRight(data, "=")

	out := make([]byte, base64.RawURLEncoding.DecodedLen(len(data)))

	n, err := base64.RawURLEncoding.Decode(out, data)
	if err != nil {
		return err
	}

	v := reflect.ValueOf(e).Elem()
	v.SetBytes(out[:n])

	return nil
}

// MarshalJSON base64 encodes a non URL-encoded value, storing the result in the
// provided byte slice.
func (e URLEncodedBase64) MarshalJSON() ([]byte, error) {
	if e == nil {
		return []byte("null"), nil
	}

	return []byte(`"` + base64.RawURLEncoding.EncodeToString(e) + `"`), nil
}
