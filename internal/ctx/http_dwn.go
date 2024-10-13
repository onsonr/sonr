package ctx

import (
	"github.com/labstack/echo/v4"
)

type DWNContext struct {
	echo.Context

	// Defaults
	id string // Generated ksuid http cookie; Initialized on first request

	// Initialization
	address string // Webauthn mapping to User ID; Supplied by DWN frontend
	chainID string // Macaroon mapping to location; Supplied by DWN frontend
}

func (s *DWNContext) ID() string {
	return s.id
}

func (s *DWNContext) Address() string {
	return s.address
}

func (s *DWNContext) ChainID() string {
	return s.chainID
}

func GetDWNContext(c echo.Context) *DWNContext {
	return c.(*DWNContext)
}

// HighwaySessionMiddleware establishes a Session Cookie.
func DWNSessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID := getSessionIDFromCookie(c)
		cc := &DWNContext{
			Context: c,
			id:      sessionID,
			address: c.Request().Header.Get("X-Sonr-Address"),
			chainID: "",
		}
		return next(cc)
	}
}
