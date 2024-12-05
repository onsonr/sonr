package session

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/gateway/config"
	"github.com/onsonr/sonr/pkg/vault/types"
)

// GatewayMiddleware establishes a Session Cookie.
func GatewayMiddleware(env config.Env) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := injectSession(c, common.RoleHway)
			return next(cc)
		}
	}
}

// WebNodeMiddleware establishes a Session Cookie.
func WebNodeMiddleware(config *types.Config) echo.MiddlewareFunc {
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

func injectConfig(c echo.Context, config *types.Config) error {
	common.HeaderWrite(c, common.SonrAPIURL, config.SonrApiUrl)
	common.HeaderWrite(c, common.SonrRPCURL, config.SonrRpcUrl)

	common.WriteCookie(c, common.SonrAddress, config.MotrAddress)
	schemaBz, err := json.Marshal(config.VaultSchema)
	if err != nil {
		return err
	}
	common.WriteCookieBytes(c, common.VaultSchema, schemaBz)
	return nil
}

// injectSession returns the session injectSession from the cookies.
func injectSession(c echo.Context, role common.PeerRole) *HTTPContext {
	if c == nil {
		return initHTTPContext(nil)
	}

	common.WriteCookie(c, common.SessionRole, role.String())

	// Continue even if there are errors, just ensure we have valid session data
	if err := loadOrGenKsuid(c); err != nil {
		// Log error but continue
	}
	if err := loadOrGenChallenge(c); err != nil {
		// Log error but continue
	}

	return initHTTPContext(c)
}

// HasAuthorization checks if the request has an authorization header
func HasAuthorization(c echo.Context) bool {
	return common.HeaderExists(c, common.Authorization)
}

// HasUserHandle checks if the request has a user handle cookie
func HasUserHandle(c echo.Context) bool {
	return common.CookieExists(c, common.UserHandle)
}

// HasVaultAddress checks if the request has a vault address cookie
func HasVaultAddress(c echo.Context) bool {
	return common.CookieExists(c, common.SonrAddress)
}
