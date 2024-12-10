package context

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/internal/vault/types"
	"github.com/onsonr/sonr/pkg/common"
)

type SessionCtx interface {
	ID() string
	BrowserName() string
	BrowserVersion() string
}

type contextKey string

// Context keys
const (
	DataContextKey contextKey = "http_session_data"
)

// Get returns the session.Context from the echo context.
func Get(c echo.Context) (SessionCtx, error) {
	ctx, ok := c.(*HTTPContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Session Context not found")
	}
	return ctx, nil
}

// WebNodeMiddleware establishes a Session Cookie.
func Middleware(config *types.Config) echo.MiddlewareFunc {
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
