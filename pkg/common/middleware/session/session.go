package session

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/common/middleware/cookie"
	"github.com/onsonr/sonr/pkg/common/types"
)

// HTTPContext is the context for DWN endpoints.
type HTTPContext struct {
	echo.Context
	role common.PeerRole
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
	// Create HTTPContext with all the extracted data
	WithData(c.Request().Context(), injectSessionData(c))
	cc := &HTTPContext{
		Context: c,
		role:    common.PeerRole(cookie.ReadUnsafe(c, cookie.SessionRole)),
	}
	// Set the HTTPContext in the context
	return cc
}

func (s *HTTPContext) ID() string {
	return s.GetData().Id
}

func (s *HTTPContext) LoginOptions(credentials []common.CredDescriptor) *common.LoginOptions {
	ch, _ := common.Base64Decode(s.GetData().Challenge)
	return &common.LoginOptions{
		Challenge:          ch,
		Timeout:            10000,
		AllowedCredentials: credentials,
	}
}

func (s *HTTPContext) RegisterOptions(subject string) *common.RegisterOptions {
	ch, _ := common.Base64Decode(s.GetData().Challenge)
	opts := baseRegisterOptions()
	opts.Challenge = ch
	opts.User = buildUserEntity(subject)
	return opts
}

func (s *HTTPContext) GetData() *types.Session {
	return GetData(s.Context.Request().Context())
}
