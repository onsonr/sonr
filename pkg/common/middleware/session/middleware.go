package session

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common/middleware/cookie"
	"github.com/onsonr/sonr/pkg/common/middleware/header"
	commonv1 "github.com/onsonr/sonr/pkg/common/types"
	"github.com/onsonr/sonr/pkg/motr/config"
)

// HwayMiddleware establishes a Session Cookie.
func HwayMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := injectSession(c, commonv1.RoleHway)
			return next(cc)
		}
	}
}

// MotrMiddleware establishes a Session Cookie.
func MotrMiddleware(config *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := injectConfig(c, config)
			if err != nil {
				return err
			}
			cc := injectSession(c, commonv1.RoleMotr)
			return next(cc)
		}
	}
}

func injectConfig(c echo.Context, config *config.Config) error {
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
func injectSession(c echo.Context, role commonv1.PeerRole) *HTTPContext {
	cookie.Write(c, cookie.SessionRole, role.String())
	err := loadOrGenKsuid(c)
	if err != nil {
		return nil
	}
	err = loadOrGenChallenge(c)
	if err != nil {
		return nil
	}
	return initHTTPContext(c)
}
