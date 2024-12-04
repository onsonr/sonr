package session

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/common/cookie"
	"github.com/onsonr/sonr/pkg/common/header"
	"github.com/onsonr/sonr/pkg/gateway/config"
	"github.com/onsonr/sonr/pkg/vault/types"
)

// HwayMiddleware establishes a Session Cookie.
func HwayMiddleware(env config.Env) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := injectSession(c, common.RoleHway)
			return next(cc)
		}
	}
}

// MotrMiddleware establishes a Session Cookie.
func MotrMiddleware(config *types.Config) echo.MiddlewareFunc {
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
	header.Write(c, header.SonrAPIURL, config.SonrApiUrl)
	header.Write(c, header.SonrRPCURL, config.SonrRpcUrl)

	cookie.Write(c, cookie.SonrAddress, config.MotrAddress)
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

// HasAuthorization checks if the request has an authorization header
func HasAuthorization(c echo.Context) bool {
	return header.Exists(c, header.Authorization)
}

// HasUserHandle checks if the request has a user handle cookie
func HasUserHandle(c echo.Context) bool {
	return cookie.Exists(c, cookie.UserHandle)
}

// HasVaultAddress checks if the request has a vault address cookie
func HasVaultAddress(c echo.Context) bool {
	return cookie.Exists(c, cookie.SonrAddress)
}
