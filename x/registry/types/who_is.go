package types

import (
	"fmt"

	"github.com/duo-labs/webauthn/protocol"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/sonr-io/sonr/pkg/did"
	rt "go.buf.build/grpc/go/sonr-io/blockchain/registry"
)

func NewWhoIsFromBuf(doc *rt.WhoIs) *WhoIs {
	return &WhoIs{
		Type:        WhoIs_Type(doc.GetType()),
		Name:        doc.GetName(),
		Did:         doc.GetDid(),
		Document:    doc.GetDocument(),
		Metadata:    doc.GetMetadata(),
		Credentials: NewCredentialListFromBuf(doc.GetCredentials()),
	}
}

func NewWhoIsToBuf(doc *WhoIs) *rt.WhoIs {
	return &rt.WhoIs{
		Type:        rt.WhoIs_Type(doc.GetType()),
		Name:        doc.GetName(),
		Did:         doc.GetDid(),
		Document:    doc.GetDocument(),
		Metadata:    doc.GetMetadata(),
		Credentials: NewCredentialListToBuf(doc.GetCredentials()),
	}
}

// AddCredential adds a webauthn credential to the whois object on the registry
func (w *WhoIs) AddCredential(cred *Credential) {
	if w.Credentials == nil {
		w.Credentials = make([]*Credential, 0)
	}
	w.Credentials = append(w.Credentials, cred)
}

// ApplicationName returns the app name from the WhoIs metadata
func (w *WhoIs) ApplicationName() (string, error) {
	if w.IsUser() {
		return "", fmt.Errorf("not an application")
	}
	m := w.GetMetadata()
	if m == nil {
		return "", fmt.Errorf("no metadata found")
	}

	if an := m["app_name"]; an != "" {
		return an, nil
	}
	return "", fmt.Errorf("no app_name found")
}

// ApplicationDescription returns the description of the application
func (w *WhoIs) ApplicationDescription() (string, error) {
	if w.IsUser() {
		return "", fmt.Errorf("not an application")
	}
	m := w.GetMetadata()
	if m == nil {
		return "", fmt.Errorf("no metadata found")
	}

	if an := m["app_description"]; an != "" {
		return an, nil
	}
	return "", fmt.Errorf("no app_description found")
}

// ApplicationCategory returns the category of the application
func (w *WhoIs) ApplicationCategory() (string, error) {
	if w.IsUser() {
		return "", fmt.Errorf("not an application")
	}
	m := w.GetMetadata()
	if m == nil {
		return "", fmt.Errorf("no metadata found")
	}

	if an := m["app_category"]; an != "" {
		return an, nil
	}
	return "", fmt.Errorf("no app_category found")
}

// ApplicationUrl returns the website of the application
func (w *WhoIs) ApplicationUrl() (string, error) {
	if w.IsUser() {
		return "", fmt.Errorf("not an application")
	}
	m := w.GetMetadata()
	if m == nil {
		return "", fmt.Errorf("no metadata found")
	}

	if an := m["app_url"]; an != "" {
		return an, nil
	}
	return "", fmt.Errorf("no app_description found")
}

// CredentialExcludeList returns a CredentialDescriptor array filled
// with all the user's credentials
func (w *WhoIs) CredentialExcludeList() []protocol.CredentialDescriptor {
	credentialExcludeList := []protocol.CredentialDescriptor{}
	for _, cred := range w.GetCredentials() {
		descriptor := protocol.CredentialDescriptor{
			Type:         protocol.PublicKeyCredentialType,
			CredentialID: cred.ID,
		}
		credentialExcludeList = append(credentialExcludeList, descriptor)
	}

	return credentialExcludeList
}

// DIDDocument returns the DID document of the whois object
func (w *WhoIs) DIDDocument() (*did.Document, error) {
	// Get the DID document from the registry
	buf := w.GetDocument()
	if buf == nil {
		return nil, fmt.Errorf("no document found")
	}

	// Unmarshal DID document from JSON
	doc := &did.Document{}
	err := doc.UnmarshalJSON(buf)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// DIDUrl returns the did.DID string of the whois object
func (w *WhoIs) DIDUrl() (*did.DID, error) {
	return did.ParseDID(w.GetDid())
}

// IsApplication returns true if WhoIs type is WhoIs_Application
func (w *WhoIs) IsApplication() bool {
	return w.Type == WhoIs_Application
}

// IsUser returns true if WhoIs type is WhoIs_User
func (w *WhoIs) IsUser() bool {
	return w.Type == WhoIs_User
}

// WebAuthnID returns the ID of the user's authenticator
func (w *WhoIs) WebAuthnID() []byte {
	// Unmarshal DID document from JSON
	return []byte(w.GetName())
}

// WebAuthnDisplayName returns the display name of the user's authenticator
func (w *WhoIs) WebAuthnName() string {
	return w.GetName()
}

// WebAuthnDisplayName returns the display name of the user's authenticator
func (w *WhoIs) WebAuthnDisplayName() string {
	return fmt.Sprintf("%s.snr", w.GetName())
}

// WebAuthnIcon returns the icon of the user's authenticator
func (w *WhoIs) WebAuthnIcon() string {
	return ""
}

// WebAuthnCredentials returns credentials owned by the user
func (w *WhoIs) WebAuthnCredentials() []webauthn.Credential {
	credentials := []webauthn.Credential{}
	for _, cred := range w.Credentials {
		credentials = append(credentials, cred.ToWebAuthn())
	}
	return credentials
}

// NewSessionFromBuf returns a new Session object from a registry buffer
func NewSessionFromBuf(doc *rt.Session) *Session {
	return &Session{
		BaseDid:    doc.BaseDid,
		Whois:      NewWhoIsFromBuf(doc.Whois),
		Credential: NewCredentialFromBuf(doc.Credential),
	}
}

// BaseDID returns the DID string as a did.DID
func (s *Session) BaseDID() (*did.DID, error) {
	return did.ParseDID(s.GetBaseDid())
}

// DIDDocument returns the DID Document for the underlying WhoIs
func (s *Session) Creator() string {
	return s.GetWhois().GetCreator()
}

// DIDDocument returns the DID Document for the underlying WhoIs
func (s *Session) DIDDocument() (*did.Document, error) {
	return s.GetWhois().DIDDocument()
}

// GenerateDID takes the input string and generates a DID URL
func (s *Session) GenerateDID(opts ...did.Option) (*did.DID, error) {
	return did.NewDID(s.GetBaseDid(), opts...)
}

// WebAuthnCredential returns the Sonr Credential as the Webauthn type
func (s *Session) WebAuthnCredential() webauthn.Credential {
	return s.GetCredential().ToWebAuthn()
}
