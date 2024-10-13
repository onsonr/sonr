package ctx

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dwngen "github.com/onsonr/sonr/internal/dwn/gen"
)

type DWNContext struct {
	echo.Context

	// Defaults
	id     string         // Generated ksuid http cookie; Initialized on first request
	dwnCfg *dwngen.Config // Provided by DWN frontend
}

func (s *DWNContext) ID() string {
	return s.id
}

func (s *DWNContext) Address() string {
	return s.dwnCfg.Motr.Address
}

func (s *DWNContext) ChainID() string {
	return s.dwnCfg.Sonr.ChainId
}

func GetDWNContext(c echo.Context) (*DWNContext, error) {
	ctx, ok := c.(*DWNContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "DWN Context not found")
	}
	return ctx, nil
}

// HighwaySessionMiddleware establishes a Session Cookie.
func DWNSessionMiddleware(config *dwngen.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sessionID := getSessionIDFromCookie(c)
			cc := &DWNContext{
				Context: c,
				id:      sessionID,
				dwnCfg:  config,
			}
			return next(cc)
		}
	}
}
