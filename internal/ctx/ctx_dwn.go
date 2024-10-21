package ctx

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	dwngen "github.com/onsonr/sonr/internal/dwn/gen"
)

// ╭───────────────────────────────────────────────────────────╮
// │                  DWNContext struct methods                │
// ╰───────────────────────────────────────────────────────────╯

// DWNContext is the context for DWN endpoints.
type DWNContext struct {
	echo.Context

	// Defaults
	id string // Generated ksuid http cookie; Initialized on first request
}

// HasAuthorization returns true if the request has an Authorization header.
func (s *DWNContext) HasAuthorization() bool {
	v := ReadHeader(s.Context, HeaderAuthorization)
	return v != ""
}

// ID returns the ksuid http cookie.
func (s *DWNContext) ID() string {
	return s.id
}

// Address returns the sonr address from the cookies.
func (s *DWNContext) Address() string {
	v, err := ReadCookie(s.Context, CookieKeySonrAddr)
	if err != nil {
		return ""
	}
	return v
}

// IPFSGatewayURL returns the IPFS gateway URL from the headers.
func (s *DWNContext) IPFSGatewayURL() string {
	return ReadHeader(s.Context, HeaderIPFSGatewayURL)
}

// ChainID returns the chain ID from the headers.
func (s *DWNContext) ChainID() string {
	return ReadHeader(s.Context, HeaderSonrChainID)
}

// Schema returns the vault schema from the cookies.
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

// GetDWNContext returns the DWNContext from the echo context.
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
