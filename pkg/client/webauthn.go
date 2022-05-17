package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/duo-labs/webauthn.io/session"
	"github.com/duo-labs/webauthn/protocol"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/patrickmn/go-cache"
	"github.com/sonr-io/sonr/pkg/config"
	rtv1 "github.com/sonr-io/sonr/x/registry/types"
)

const (
	RegistrationSessionKey   = "registration_session"
	AuthenticationSessionKey = "authentication_session"
)

// WebAuthn manages the WebAuthn interface
type WebAuthn struct {
	instance *webauthn.WebAuthn

	ctx          context.Context
	cache        *cache.Cache
	config       *config.Config
	sessionStore *session.Store
}

// NewWebauthn creates a new WebAuthn instance with the given configuration
func NewWebauthn(ctx context.Context, config *config.Config) (*WebAuthn, error) {
	// Create a WebAuthn instance
	web, err := webauthn.New(config.WebauthnConfig())
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
		instance:     web,
		ctx:          ctx,
		cache:        cache.New(5*time.Minute, 10*time.Minute),
		config:       config,
		sessionStore: sessionStore,
	}, nil
}

// FinishAuthenticationSession returns the registration session for the given user
func (w *WebAuthn) FinishAuthenticationSession(r *http.Request, username string) (*webauthn.Credential, error) {
	// get user
	x, found := w.cache.Get(AuthenticationCacheKey(username))
	if !found {
		return nil, errors.New("user not found")
	}
	whois := x.(*rtv1.WhoIs)

	sessionData, err := w.sessionStore.GetWebauthnSession(AuthenticationSessionKey, r)
	if err != nil {
		return nil, err
	}

	credential, err := w.instance.FinishLogin(whois, sessionData, r)
	if err != nil {
		return nil, err
	}
	return credential, nil
}

// FinishRegistrationSession returns the registration session for the given user
func (w *WebAuthn) FinishRegistrationSession(r *http.Request, username string) (*webauthn.Credential, error) {
	// get user
	x, found := w.cache.Get(RegisterCacheKey(username))
	if !found {
		return nil, errors.New("user not found")
	}
	whois := x.(*rtv1.WhoIs)

	sessionData, err := w.sessionStore.GetWebauthnSession(RegistrationSessionKey, r)
	if err != nil {
		return nil, err
	}

	credential, err := w.instance.FinishRegistration(whois, sessionData, r)
	if err != nil {
		return nil, err
	}
	return credential, nil
}

// SaveAuthenticationSession saves the login session for the given user
func (w *WebAuthn) SaveAuthenticationSession(
	r *http.Request,
	rw http.ResponseWriter,
	whoIs *rtv1.WhoIs,
) (*protocol.CredentialAssertion, error) {
	// generate PublicKeyCredentialRequestOptions, session data
	w.cache.Set(AuthenticationCacheKey(whoIs.Name), whoIs, cache.DefaultExpiration)

	options, sessionData, err := w.instance.BeginLogin(whoIs)
	if err != nil {
		return nil, err
	}

	// store session data as marshaled JSON
	err = w.sessionStore.SaveWebauthnSession("authentication", sessionData, r, rw)
	if err != nil {
		return nil, err
	}

	return options, nil
}

// SaveRegistrationSession saves the registration session for the given user
func (w *WebAuthn) SaveRegistrationSession(
	r *http.Request,
	rw http.ResponseWriter,
	username string,
	creator string,
) (*protocol.CredentialCreation, error) {
	whoIs := NewBlankWhoIs(username, creator)
	w.cache.Set(RegisterCacheKey(username), whoIs, cache.DefaultExpiration)

	// generate PublicKeyCredentialCreationOptions session data
	options, sessionData, err := w.instance.BeginRegistration(
		whoIs,
		registerOptions(whoIs),
		webauthn.WithAuthenticatorSelection(authSelect()),
		webauthn.WithConveyancePreference(conveyancePreference()),
	)
	if err != nil {
		return nil, err
	}

	// store session data as marshaled JSON
	err = w.sessionStore.SaveWebauthnSession("registration", sessionData, r, rw)
	if err != nil {
		return nil, err
	}

	return options, nil
}

// ----------------
// HELPER FUNCTIONS
// ----------------

// AuthenticationCacheKey is a helper function to create a cache key for the given user
func AuthenticationCacheKey(username string) string {
	return fmt.Sprintf("%s_authentication", username)
}

// authSelect is a helper function to create a AuthenticatorSelection
func authSelect() protocol.AuthenticatorSelection {
	return protocol.AuthenticatorSelection{
		AuthenticatorAttachment: protocol.AuthenticatorAttachment("platform"),
		RequireResidentKey:      protocol.ResidentKeyUnrequired(),
		UserVerification:        protocol.VerificationRequired,
	}
}

// NewBlankWhoIs is a helper function to create a blank WhoIs
func NewBlankWhoIs(username string, creator string) *rtv1.WhoIs {
	return &rtv1.WhoIs{
		Name:        username,
		Did:         "",
		Document:    nil,
		Creator:     creator,
		Credentials: make([]*rtv1.Credential, 0),
	}
}

// conveyancePreference is a helper function to create a ConveyancePreference
func conveyancePreference() protocol.ConveyancePreference {
	return protocol.ConveyancePreference(protocol.PreferNoAttestation)
}

// RegisterCacheKey is a helper function to create a cache key for a registration session
func RegisterCacheKey(username string) string {
	return fmt.Sprintf("%s_registration", username)
}

// registerOptions is a helper function to create a PublicKeyCredentialCreationOptions
func registerOptions(whois *rtv1.WhoIs) func(credCreationOpts *protocol.PublicKeyCredentialCreationOptions) {
	return func(credCreationOpts *protocol.PublicKeyCredentialCreationOptions) {
		credCreationOpts.CredentialExcludeList = whois.CredentialExcludeList()
	}
}
