package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cfssl/log"
	"github.com/duo-labs/webauthn.io/session"
	"github.com/duo-labs/webauthn/protocol"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/patrickmn/go-cache"
	"github.com/sonr-io/sonr/pkg/config"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	rtv1 "github.com/sonr-io/sonr/x/registry/types"
)

const (
	REGISTRATION_SESSION_KEY   = "registration_session"
	AUTHENTICATION_SESSION_KEY = "authentication_session"
)

// WebAuthn manages the WebAuthn interface
type WebAuthn struct {
	instance *webauthn.WebAuthn

	ctx      context.Context
	cache    *cache.Cache
	config   *config.Config
	sessions *session.Store
}

// NewWebauthn creates a new WebAuthn instance with the given configuration
func NewWebauthn(ctx context.Context, config *config.Config) (*WebAuthn, error) {
	// Create a WebAuthn instance
	web, err := webauthn.New(toWebauthnConfig(config))
	if err != nil {
		return nil, err
	}

	// Create a new Session Store
	sessionStore, err := session.NewStore()
	if err != nil {
		return nil, err
	}

	// return a new WebAuthn instance
	return &WebAuthn{
		instance: web,
		ctx:      ctx,
		cache:    cache.New(5*time.Minute, 10*time.Minute),
		config:   config,
		sessions: sessionStore,
	}, nil
}

// FinishAuthenticationSession returns the registration session for the given user
func (w *WebAuthn) FinishAuthenticationSession(r *http.Request, username string) (*webauthn.Credential, error) {
	// get user
	x, found := w.cache.Get(authenticationCacheKey(username))
	if !found {
		return nil, errors.New("user not found")
	}
	whois := x.(*rtv1.WhoIs)

	sessionData, err := w.sessions.GetWebauthnSession(AUTHENTICATION_SESSION_KEY, r)
	if err != nil {
		return nil, err
	}

	doc, err := whois.UnmarshalDidDocument()
	if err != nil {
		return nil, err
	}

	credential, err := w.instance.FinishLogin(doc, sessionData, r)
	if err != nil {
		return nil, err
	}
	return credential, nil
}

// FinishRegistrationSession returns the registration session for the given user
func (w *WebAuthn) FinishRegistrationSession(r *http.Request, username string) (*webauthn.Credential, error) {
	// get user
	x, found := w.cache.Get(registerCacheKey(username))
	if !found {
		return nil, errors.New("user not found")
	}
	whois := x.(*rtv1.WhoIs)

	sessionData, err := w.sessions.GetWebauthnSession(REGISTRATION_SESSION_KEY, r)
	if err != nil {
		log.Errorf("error finishing registration: %s", err)
		return nil, err
	}

	doc, err := whois.UnmarshalDidDocument()
	if err != nil {
		log.Errorf("error finishing registration: %s", err)
		return nil, err
	}

	credential, err := w.instance.FinishRegistration(doc, sessionData, r)
	if err != nil {
		log.Errorf("error finishing registration: %s", err)
		return nil, err
	}
	return credential, nil
}

// SaveAuthenticationSession saves the login session for the given user
func (wan *WebAuthn) SaveAuthenticationSession(r *http.Request, w http.ResponseWriter, whoIs *rtv1.WhoIs) (*protocol.CredentialAssertion, error) {
	// generate PublicKeyCredentialRequestOptions, session data
	wan.cache.Set(authenticationCacheKey(whoIs.Owner), whoIs, cache.DefaultExpiration)
	doc, err := whoIs.UnmarshalDidDocument()
	if err != nil {
		return nil, err
	}

	options, sessionData, err := wan.instance.BeginLogin(doc)
	if err != nil {
		return nil, err
	}

	// store session data as marshaled JSON
	err = wan.sessions.SaveWebauthnSession(AUTHENTICATION_SESSION_KEY, sessionData, r, w)
	if err != nil {
		return nil, err
	}
	return options, nil
}

// SaveRegistrationSession saves the registration session for the given user
func (wan *WebAuthn) SaveRegistrationSession(r *http.Request, w http.ResponseWriter, username string, creator string) (*protocol.CredentialCreation, error) {
	// Create Blank WhoIs
	whoIs := blankWhoIs(username, creator)
	wan.cache.Set(registerCacheKey(username), whoIs, cache.DefaultExpiration)

	doc, err := whoIs.UnmarshalDidDocument()
	if err != nil {
		log.Errorf("error finishing registration: %s", err)
		return nil, err
	}

	// generate PublicKeyCredentialCreationOptions, session data
	options, sessionData, err := wan.instance.BeginRegistration(
		doc,
		registerOptions(whoIs),
		webauthn.WithAuthenticatorSelection(authSelect()),
		webauthn.WithConveyancePreference(conveyancePreference()),
	)
	if err != nil {
		log.Errorf("error finishing registration: %s", err)
		return nil, err
	}

	// store session data as marshaled JSON
	err = wan.sessions.SaveWebauthnSession(REGISTRATION_SESSION_KEY, sessionData, r, w)
	if err != nil {
		log.Errorf("error finishing registration: %s", err)
		return nil, err
	}
	return options, nil
}

// ----------------
// HELPER FUNCTIONS
// ----------------

// authenticationCacheKey is a helper function to create a cache key for the given user
func authenticationCacheKey(username string) string {
	return fmt.Sprintf("%s_authentication", username)
}

// authSelect is a helper function to create a AuthenticatorSelection
func authSelect() protocol.AuthenticatorSelection {
	return protocol.AuthenticatorSelection{
		AuthenticatorAttachment: protocol.AuthenticatorAttachment("platform"),
		RequireResidentKey:      protocol.ResidentKeyRequired(),
		UserVerification:        protocol.VerificationRequired,
	}
}

// blankWhoIs is a helper function to create a blank WhoIs
func blankWhoIs(username, creator string) *rtv1.WhoIs {
	didUrl, err := did.ParseDID(fmt.Sprintf("did:snr:%s", creator))
	if err != nil {
		return nil
	}
	ctxUri, err := ssi.ParseURI("https://www.w3.org/ns/did/v1")
	if err != nil {
		return nil
	}

	doc := did.Document{
		ID:      *didUrl,
		Context: []ssi.URI{*ctxUri},
	}

	docBuf, err := doc.MarshalJSON()
	if err != nil {
		return nil
	}

	return &rtv1.WhoIs{
		Owner:       creator,
		DidDocument: docBuf,
	}
}

// conveyancePreference is a helper function to create a ConveyancePreference
func conveyancePreference() protocol.ConveyancePreference {
	return protocol.ConveyancePreference(protocol.PreferNoAttestation)
}

// registerCacheKey is a helper function to create a cache key for a registration session
func registerCacheKey(username string) string {
	return fmt.Sprintf("%s_registration", username)
}

// registerOptions is a helper function to create a PublicKeyCredentialCreationOptions
func registerOptions(whois *rtv1.WhoIs) func(credCreationOpts *protocol.PublicKeyCredentialCreationOptions) {
	return func(credCreationOpts *protocol.PublicKeyCredentialCreationOptions) {
		doc, err := whois.UnmarshalDidDocument()
		if err != nil {
			return
		}
		credCreationOpts.CredentialExcludeList = doc.WebAuthnCredentialExcludeList()
	}
}
