package context

import (
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

func (s *HTTPContext) BrowserName() string {
	return s.bn
}

func (s *HTTPContext) BrowserVersion() string {
	return s.bv
}
