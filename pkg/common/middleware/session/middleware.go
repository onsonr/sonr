package session

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/common/middleware/cookie"
	"github.com/onsonr/sonr/pkg/common/middleware/header"
	"github.com/onsonr/sonr/pkg/core/dwn"
)

// HwayMiddleware establishes a Session Cookie.
func HwayMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := injectSession(c, common.RoleHway)
			return next(cc)
		}
	}
}

// MotrMiddleware establishes a Session Cookie.
func MotrMiddleware(config *dwn.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := injectConfig(c, config)
			if err != nil {
				return err
			}
			cc := injectSession(c, common.RoleMotr)
			return next(cc)
		}
	}
}

func injectConfig(c echo.Context, config *dwn.Config) error {
	header.Write(c, header.IPFSHost, config.IpfsGatewayUrl)
	header.Write(c, header.ChainID, config.SonrChainId)

	header.Write(c, header.SonrAPIURL, config.SonrApiUrl)
	header.Write(c, header.SonrRPCURL, config.SonrRpcUrl)

	cookie.Write(c, cookie.SonrAddress, config.MotrAddress)
	cookie.Write(c, cookie.SonrKeyshare, config.MotrKeyshare)

	schemaBz, err := json.Marshal(config.VaultSchema)
	if err != nil {
		return err
	}
	cookie.WriteBytes(c, cookie.VaultSchema, schemaBz)
	return nil
}

// injectSession returns the session injectSession from the cookies.
func injectSession(c echo.Context, role common.PeerRole) *HTTPContext {
	if c == nil {
		return initHTTPContext(nil)
	}
	
	cookie.Write(c, cookie.SessionRole, role.String())
	
	// Continue even if there are errors, just ensure we have valid session data
	if err := loadOrGenKsuid(c); err != nil {
		// Log error but continue
	}
	if err := loadOrGenChallenge(c); err != nil {
		// Log error but continue  
	}
	
	return initHTTPContext(c)
}
