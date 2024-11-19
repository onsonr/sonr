package ctx

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ╭───────────────────────────────────────────────────────────╮
// │                  HwayContext struct methods               │
// ╰───────────────────────────────────────────────────────────╯

// HwayContext is the context for Highway endpoints.
type HwayContext struct {
	echo.Context

	// Defaults
	id string // Generated ksuid http cookie; Initialized on first request
}

// ID returns the ksuid http cookie
func (s *HwayContext) ID() string {
	return s.id
}

// GetHwayContext returns the HwayContext from the echo context.
func GetHWAYContext(c echo.Context) (*HwayContext, error) {
	ctx, ok := c.(*HwayContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Highway Context not found")
	}
	return ctx, nil
}

// HighwaySessionMiddleware establishes a Session Cookie.
func HighwaySessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID := GetSessionID(c)
		cc := &HwayContext{
			Context: c,
			id:      sessionID,
		}
		return next(cc)
	}
}
