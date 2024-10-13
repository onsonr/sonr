package ctx

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HighwayContext struct {
	echo.Context

	// Defaults
	id string // Generated ksuid http cookie; Initialized on first request

	// Authentication
	challenge WebBytes // Webauthn mapping to Challenge; Per session based on origin
}

func (s *HighwayContext) ID() string {
	return s.id
}

func GetHighwayContext(c echo.Context) (*HighwayContext, error) {
	ctx, ok := c.(*HighwayContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Highway Context not found")
	}
	return ctx, nil
}

// HighwaySessionMiddleware establishes a Session Cookie.
func HighwaySessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID := getSessionIDFromCookie(c)
		cc := &HighwayContext{
			Context: c,
			id:      sessionID,
		}
		return next(cc)
	}
}
