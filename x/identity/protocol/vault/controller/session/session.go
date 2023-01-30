package session

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/sonrhq/core/x/identity/types"
)

var (
	// Default Origins
	defaultRpOrigins = []string{
		"https://auth.sonr.io",
		"https://sonr.id",
		"https://sandbox.sonr.network",
		"http://localhost:3000",
	}

	// Default Icon to display
	defaultRpIcon = "https://raw.githubusercontent.com/sonrhq/core/master/docs/static/favicon.png"

	// Default name to display
	defaultRpName = "Sonr"

	// defaultAttestionPreference
	defaultAttestationPreference = protocol.PreferDirectAttestation

	// defaultAuthSelect
	defaultAuthSelect = protocol.AuthenticatorSelection{
		AuthenticatorAttachment: protocol.AuthenticatorAttachment("platform"),
	}

	// defaultTimeout
	defaultTimeout = 60000
)

// NewSession creates a new session with challenge to be used to register a new account
func NewSession(rpId string, aka string) (*Session, error) {
	s := defaultSession(rpId, aka)
	err := s.Apply()
	if err != nil {
		return nil, fmt.Errorf("failed to apply options to Webauthn config: %w", err)
	}
	return s, nil
}

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
	// Session ID
	rpid string
	AKA  string

	// Relying Party ID
	webauthn *webauthn.WebAuthn

	// User Data
	didDoc     *types.DidDocument
	data       *webauthn.SessionData
	isExisting bool
}

// Option is a function that configures a session
type Option func(*webauthn.Config)

// WithRPIcon sets the RPIcon
func WithRPIcon(icon string) Option {
	return func(s *webauthn.Config) {
		s.RPIcon = icon
	}
}

// WithRPOrigins sets the RPOrigins
func WithRPOrigins(origins []string) Option {
	return func(s *webauthn.Config) {
		s.RPOrigins = origins
	}
}

// WithTimeout sets the Timeout
func WithTimeout(timeout int) Option {
	return func(s *webauthn.Config) {
		s.Timeout = timeout
	}
}

// WithAttestionPreference sets the AttestionPreference
func WithAttestionPreference(pref protocol.ConveyancePreference) Option {
	return func(s *webauthn.Config) {
		s.AttestationPreference = pref
	}
}

// WithAuthenticatorSelect sets the AuthenticatorSelect
func WithAuthenticatorSelect(selectAuth protocol.AuthenticatorSelection) Option {
	return func(s *webauthn.Config) {
		s.AuthenticatorSelection = selectAuth
	}
}

// Apply applies the options to the session
func (s *Session) Apply(opts ...Option) error {
	c := &webauthn.Config{
		RPID:                   s.rpid,
		RPDisplayName:          s.AKA,
		RPIcon:                 defaultRpIcon,
		RPOrigins:              defaultRpOrigins,
		Timeout:                defaultTimeout,
		AttestationPreference:  defaultAttestationPreference,
		AuthenticatorSelection: defaultAuthSelect,
	}
	for _, opt := range opts {
		opt(c)
	}
	wauth, err := webauthn.New(c)
	if err != nil {
		return err
	}
	s.webauthn = wauth
	return nil
}

// defaultSession returns a default session
func defaultSession(rpid string, aka string) *Session {
	return &Session{
		isExisting: false,
		// didDoc:     types.NewBaseDocument(aka),
		rpid: rpid,
		AKA:  aka,
	}
}
