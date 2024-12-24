package context

import (
	gocontext "context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/medama-io/go-useragent"
	"github.com/onsonr/sonr/internal/crypto/mpc"
	"github.com/onsonr/sonr/internal/common"
	"github.com/onsonr/sonr/internal/config/hway"
	hwayorm "github.com/onsonr/sonr/internal/database/hwayorm"
)

type GatewayContext struct {
	echo.Context
	*hwayorm.Queries
	agent            useragent.UserAgent
	id               string
	ipfsClient       common.IPFS
	tokenStore       common.IPFSTokenStore
	stagedEnclaves   map[string]mpc.Enclave
	grpcAddr         string
	turnstileSiteKey string
}

func GetGateway(c echo.Context) (*GatewayContext, error) {
	cc, ok := c.(*GatewayContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Gateway Context not found")
	}
	return cc, nil
}

func UseGateway(env hway.Hway, ipc common.IPFS, db *hwayorm.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ua := useragent.NewParser()
			ctx := &GatewayContext{
				Context:          c,
				turnstileSiteKey: env.GetTurnstileSiteKey(),
				agent:            ua.Parse(c.Request().UserAgent()),
				Queries:          db,
				ipfsClient:       ipc,
				grpcAddr:         env.GetSonrGrpcUrl(),
				tokenStore:       common.NewUCANStore(ipc),
			}
			return next(ctx)
		}
	}
}

func BG() gocontext.Context {
	ctx := gocontext.Background()
	return ctx
}

func (cc *GatewayContext) ReadCookie(k common.CookieKey) string {
	return common.ReadCookieUnsafe(cc.Context, k)
}

func (cc *GatewayContext) WriteCookie(k common.CookieKey, v string) {
	common.WriteCookie(cc.Context, k, v)
}
