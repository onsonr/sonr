package session

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/x/identity/types"
)

type SessionEntry struct {
	ID                 string
	RPID               string
	WebauthnCredential common.WebauthnCredential
	DidDoc             types.DidDocument
	Data               webauthn.SessionData
	Label              string
	AlsoKnownAs        string
}

// NewEntry creates a new session with challenge to be used to register a new account
func NewEntry(rpId string, aka string, label string) (*SessionEntry, error) {
	sessionID := uuid.New().String()[:8]
	doc := types.NewBaseDocument(aka, sessionID)
	// Create Entry
	return &SessionEntry{
		ID:          sessionID,
		RPID:        rpId,
		DidDoc:      *doc,
		Label:       label,
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
	return &SessionEntry{
		ID:                 sessionID,
		RPID:               rpId,
		WebauthnCredential: *wb,
	}, nil
}

// BeginRegistration starts the registration process for the underlying Webauthn instance
func (s *SessionEntry) BeginRegistration() (string, error) {
	wauth, err := s.WebAuthn()
	if err != nil {
		return "", err
	}

	opts, sessionData, err := wauth.BeginRegistration(&s.DidDoc, webauthn.WithAuthenticatorSelection(defaultAuthSelect))
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
		return nil, err
	}

	err = s.DidDoc.AddWebauthnCredential(common.ConvertFromWebauthnCredential(cred), s.Label)
	if err != nil {
		return nil, err
	}
	return &s.DidDoc, nil
}

// BeginLogin creates a new AssertionChallenge for client to verify
func (s *SessionEntry) BeginLogin() (string, error) {
	wauth, err := s.WebAuthn()
	if err != nil {
		return "", err
	}

	allowList := make([]protocol.CredentialDescriptor, 1)
	allowList[0] = protocol.CredentialDescriptor{
		CredentialID: s.WebauthnCredential.Id,
		Type:         protocol.CredentialType("public-key"),
	}
	opts, session, err := wauth.BeginLogin(&s.DidDoc, webauthn.WithAllowedCredentials(allowList))
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
	wauth, err := s.WebAuthn()
	if err != nil {
		return false, err
	}

	pca, err := getParsedCredentialRequestData(credentialRequestData)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Failed to get parsed creation data: %s", err))
	}
	cred, err := wauth.ValidateLogin(&s.DidDoc, s.Data, pca)
	if err != nil {
		return false, err
	}
	if err := s.WebauthnCredential.Validate(cred); err != nil {
		return false, err
	}
	return true, nil
}

func GetEntry(id string, cache *gocache.Cache) (*SessionEntry, error) {
	val, ok := cache.Get(id)
	if !ok {
		return nil, errors.New("Failed to find entry for ID")
	}
	e, ok := val.(*SessionEntry)
	if !ok {
		return nil, errors.New("Invalid type for session entry")
	}
	return e, nil
}

func PutEntry(entry *SessionEntry, cache *gocache.Cache) error {
	if entry == nil || cache == nil {
		return errors.New("Entry or Cache cannot be nil to put Entry")
	}
	return cache.Add(entry.ID, entry, -1)
}
