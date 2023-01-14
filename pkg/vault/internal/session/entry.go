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

// `Session` is a struct that contains a `string` (`ID`), a `string` (`RPID`), a
// `common.WebauthnCredential` (`WebauthnCredential`), a `types.DidDocument` (`DidDoc`), a
// `webauthn.SessionData` (`Data`), and a `string` (`AlsoKnownAs`).
// @property {string} ID - The session ID.
// @property {string} RPID - The Relying Party ID. This is the domain of the relying party.
// @property WebauthnCredential - This is the credential that was created by the user.
// @property DidDoc - The DID Document of the user.
// @property Data - This is the data that is returned from the webauthn.Create() function.
// @property {string} AlsoKnownAs - The user's username.
type Session struct {
	ID          string
	RPID        string
	DidDoc      *types.DidDocument
	Data        webauthn.SessionData
	Webauthn    *webauthn.WebAuthn
	AlsoKnownAs string
}

// NewEntry creates a new session with challenge to be used to register a new account
func NewEntry(rpId string, aka string) (*Session, error) {
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
	})
	if err != nil {
		return nil, err
	}

	// Create Entry
	return &Session{
		ID:          sessionID,
		RPID:        rpId,
		DidDoc:      types.NewBaseDocument(aka, sessionID),
		AlsoKnownAs: aka,
		Webauthn:    wauth,
	}, nil
}

// LoadEntry starts a new webauthn session with a given VerificationMethod
func LoadEntry(rpId string, vm *types.VerificationMethod) (*Session, error) {
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
	})
	if err != nil {
		return nil, err
	}

	return &Session{
		ID:       sessionID,
		RPID:     rpId,
		Webauthn: wauth,
	}, nil
}

// BeginRegistration starts the registration process for the underlying Webauthn instance
func (s *Session) BeginRegistration() (string, error) {
	opts, sessionData, err := s.Webauthn.BeginRegistration(s.DidDoc, webauthn.WithAuthenticatorSelection(defaultAuthSelect))
	if err != nil {
		return "", err
	}
	s.Data = *sessionData
	bz, err := json.Marshal(opts)
	if err != nil {
		return "", err
	}
	return string(bz), nil
}

// FinishRegistration creates a credential which can be stored to use with User Authentication
func (s *Session) FinishRegistration(credentialCreationData string) (*types.DidDocument, error) {
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
func (s *Session) BeginLogin() (string, error) {
	allowList := make([]protocol.CredentialDescriptor, 0)
	creds := s.DidDoc.WebAuthnCredentials()
	for _, cred := range creds {
		allowList = append(allowList, cred.Descriptor())
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
func (s *Session) FinishLogin(credentialRequestData string) (bool, error) {

	pca, err := getParsedCredentialRequestData(credentialRequestData)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Failed to get parsed creation data: %s", err))
	}
	_, err = s.Webauthn.ValidateLogin(s.DidDoc, s.Data, pca)
	if err != nil {
		return false, err
	}
	return true, nil
}
