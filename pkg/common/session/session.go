package session

import (
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common"
)

// HTTPContext is the context for HTTP endpoints.
type HTTPContext struct {
	echo.Context
	role common.PeerRole
	id   string
	chal string
	bn   string
	bv   string
}

// Ensure HTTPContext implements context.Context
func (s *HTTPContext) Deadline() (deadline time.Time, ok bool) {
	return s.Context.Request().Context().Deadline()
}

func (s *HTTPContext) Done() <-chan struct{} {
	return s.Context.Request().Context().Done()
}

func (s *HTTPContext) Err() error {
	return s.Context.Request().Context().Err()
}

func (s *HTTPContext) Value(key interface{}) interface{} {
	return s.Context.Request().Context().Value(key)
}

// initHTTPContext loads the headers from the request.
func initHTTPContext(c echo.Context) *HTTPContext {
	if c == nil {
		return &HTTPContext{}
	}

	id, chal := extractPeerInfo(c)
	bn, bv := extractBrowserInfo(c)

	cc := &HTTPContext{
		Context: c,
		role:    common.PeerRole(common.ReadCookieUnsafe(c, common.SessionRole)),
		id:      id,
		chal:    chal,
		bn:      bn,
		bv:      bv,
	}

	// Set the session data in both contexts
	return cc
}

func (s *HTTPContext) ID() string {
	return s.id
}

func (s *HTTPContext) Challenge() string {
	return s.chal
}

func (s *HTTPContext) LoginOptions(credentials []protocol.CredentialDescriptor) *protocol.PublicKeyCredentialRequestOptions {
	ch, _ := common.Base64Decode(s.chal)
	return &protocol.PublicKeyCredentialRequestOptions{
		Challenge:          ch,
		Timeout:            10000,
		AllowedCredentials: credentials,
	}
}

func (s *HTTPContext) RegisterOptions(subject string) *protocol.PublicKeyCredentialCreationOptions {
	ch, _ := common.Base64Decode(s.chal)
	opts := baseRegisterOptions()
	opts.Challenge = ch
	opts.User = buildUserEntity(subject)
	return opts
}

func (s *HTTPContext) BrowserName() string {
	return s.bn
}

func (s *HTTPContext) BrowserVersion() string {
	return s.bv
}
