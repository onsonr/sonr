package session

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/gateway/config"
)

// Middleware establishes a Session Cookie.
func Middleware(env config.Env) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := injectSession(c, common.RoleHway)
			return next(cc)
		}
	}
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
