package ctx

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	dwngen "github.com/onsonr/sonr/internal/dwn/gen"
)

type DWNContext struct {
	echo.Context

	// Defaults
	id string // Generated ksuid http cookie; Initialized on first request
}

func (s *DWNContext) HasAuthorization() bool {
	v := ReadHeader(s.Context, HeaderAuthorization)
	return v != ""
}

func (s *DWNContext) ID() string {
	return s.id
}

func (s *DWNContext) Address() string {
	v, err := ReadCookie(s.Context, CookieKeySonrAddr)
	if err != nil {
		return ""
	}
	return v
}

func (s *DWNContext) IPFSGatewayURL() string {
	return ReadHeader(s.Context, HeaderIPFSGatewayURL)
}

func (s *DWNContext) ChainID() string {
	return ReadHeader(s.Context, HeaderSonrChainID)
}

func (s *DWNContext) Schema() *dwngen.Schema {
	v, err := ReadCookie(s.Context, CookieKeyVaultSchema)
	if err != nil {
		return nil
	}
	var schema dwngen.Schema
	err = json.Unmarshal([]byte(v), &schema)
	if err != nil {
		return nil
	}
	return &schema
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
			sessionID := GetSessionID(c)
			injectConfig(c, config)
			cc := &DWNContext{
				Context: c,
				id:      sessionID,
			}
			return next(cc)
		}
	}
}
