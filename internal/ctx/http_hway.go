package ctx

import (
	"github.com/labstack/echo/v4"
)

type HighwayContext struct {
	echo.Context

	// Defaults
	id string // Generated ksuid http cookie; Initialized on first request

	// Initialization
	address string // Webauthn mapping to User ID; Supplied by DWN frontend
	chainID string // Macaroon mapping to location; Supplied by DWN frontend

	// Authentication
	challenge WebBytes // Webauthn mapping to Challenge; Per session based on origin
}

func (s *HighwayContext) ID() string {
	return s.id
}

func (s *HighwayContext) Address() string {
	return s.address
}

func (s *HighwayContext) ChainID() string {
	return s.chainID
}

func GetHighwayContext(c echo.Context) *HighwayContext {
	return c.(*HighwayContext)
}

// HighwaySessionMiddleware establishes a Session Cookie.
func HighwaySessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID := getSessionIDFromCookie(c)
		cc := &HighwayContext{
			Context: c,
			id:      sessionID,
			address: c.Request().Header.Get("X-Sonr-Address"),
			chainID: "",
		}
		return next(cc)
	}
}
