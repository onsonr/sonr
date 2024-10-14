package ctx

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HighwayContext struct {
	echo.Context

	// Defaults
	id string // Generated ksuid http cookie; Initialized on first request
}

func (s *HighwayContext) ID() string {
	return s.id
}

func GetHWAYContext(c echo.Context) (*HighwayContext, error) {
	ctx, ok := c.(*HighwayContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Highway Context not found")
	}
	return ctx, nil
}

// HWAYSessionMiddleware establishes a Session Cookie.
func HWAYSessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID := GetSessionID(c)
		cc := &HighwayContext{
			Context: c,
			id:      sessionID,
		}
		return next(cc)
	}
}
