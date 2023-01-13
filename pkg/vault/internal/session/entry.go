package session

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/x/identity/types"
)

// `SessionEntry` is a struct that contains a `string` (`ID`), a `string` (`RPID`), a
// `common.WebauthnCredential` (`WebauthnCredential`), a `types.DidDocument` (`DidDoc`), a
// `webauthn.SessionData` (`Data`), and a `string` (`AlsoKnownAs`).
// @property {string} ID - The session ID.
// @property {string} RPID - The Relying Party ID. This is the domain of the relying party.
// @property WebauthnCredential - This is the credential that was created by the user.
// @property DidDoc - The DID Document of the user.
// @property Data - This is the data that is returned from the webauthn.Create() function.
// @property {string} AlsoKnownAs - The user's username.
type SessionEntry struct {
	ID                 string
	RPID               string
	WebauthnCredential common.WebauthnCredential
	DidDoc             *types.DidDocument
	Data               webauthn.SessionData
	Webauthn           *webauthn.WebAuthn
	AlsoKnownAs        string
}

// NewEntry creates a new session with challenge to be used to register a new account
func NewEntry(rpId string, aka string) (*SessionEntry, error) {
	sessionID := uuid.New().String()[:8]
	doc := types.NewBaseDocument(aka, sessionID)
	// Create Entry
	return &SessionEntry{
		ID:          sessionID,
		RPID:        rpId,
		DidDoc:      doc,
		AlsoKnownAs: aka,
	}, nil
}

// LoadEntry starts a new webauthn session with a given VerificationMethod
func LoadEntry(rpId string, vm *types.VerificationMethod) (*SessionEntry, error) {
	wb, err := vm.WebAuthnCredential()
	if err != nil {
		return nil, err
	}
	sessionID := uuid.New().String()[:8]
	// Create the Webauthn Instance
	wauth, err := webauthn.New(&webauthn.Config{
		RPDisplayName:          defaultRpName,
		RPID:                   rpId,
		RPIcon:                 defaultRpIcon,
		RPOrigins:              defaultRpOrigins,
		Timeout:                60000,
		AttestationPreference:  protocol.PreferDirectAttestation,
		AuthenticatorSelection: defaultAuthSelect,
		Debug:                  true,
	})

	return &SessionEntry{
		ID:                 sessionID,
		RPID:               rpId,
		WebauthnCredential: *wb,
		Webauthn:           wauth,
	}, nil
}

// BeginRegistration starts the registration process for the underlying Webauthn instance
func (s *SessionEntry) BeginRegistration() (string, error) {

	opts, sessionData, err := s.Webauthn.BeginRegistration(s.DidDoc, webauthn.WithAuthenticatorSelection(defaultAuthSelect))
	if err != nil {
		return "", err
	}
	s.Data = *sessionData
	opts = s.SetRPID(opts)

	bz, err := json.Marshal(opts)
	if err != nil {
		return "", err
	}
	return string(bz), nil
}

// FinishRegistration creates a credential which can be stored to use with User Authentication
func (s *SessionEntry) FinishRegistration(credentialCreationData string) (*types.DidDocument, error) {
	// Parse Client Credential Data
	pcc, err := getParsedCredentialCreationData(credentialCreationData)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get parsed creation data: %s", err))
	}
	// err = pcc.Verify(s.Data.Challenge, false, s.RPID, defaultRpOrigins)
	// if err != nil {
	// 	return nil, err
	// }
	cred, err := webauthn.MakeNewCredential(pcc)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to make new credential: %s", err))
	}
	keyIdx := s.DidDoc.Authentication.Count() + 1
	err = s.DidDoc.AddWebauthnCredential(common.ConvertFromWebauthnCredential(cred), fmt.Sprintf("key-%v", keyIdx))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to add webauthn credential: %s", err))
	}
	return s.DidDoc, nil
}

// BeginLogin creates a new AssertionChallenge for client to verify
func (s *SessionEntry) BeginLogin() (string, error) {

	allowList := make([]protocol.CredentialDescriptor, 1)
	allowList[0] = protocol.CredentialDescriptor{
		CredentialID: s.WebauthnCredential.Id,
		Type:         protocol.CredentialType("public-key"),
	}
	opts, session, err := s.Webauthn.BeginLogin(s.DidDoc, webauthn.WithAllowedCredentials(allowList))
	if err != nil {
		return "", err
	}
	s.Data = *session
	bz, err := json.Marshal(opts)
	if err != nil {
		return "", err
	}
	return string(bz), nil
}

// FinishLogin authenticates from the signature provided to the client
func (s *SessionEntry) FinishLogin(credentialRequestData string) (bool, error) {

	pca, err := getParsedCredentialRequestData(credentialRequestData)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Failed to get parsed creation data: %s", err))
	}
	cred, err := s.Webauthn.ValidateLogin(s.DidDoc, s.Data, pca)
	if err != nil {
		return false, err
	}
	if err := s.WebauthnCredential.Validate(cred); err != nil {
		return false, err
	}
	return true, nil
}
