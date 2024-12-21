package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/medama-io/go-useragent"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/internal/config/hway"
	hwayorm "github.com/onsonr/sonr/internal/database/hwayorm"
	"github.com/onsonr/sonr/pkg/common"
)

type GatewayContext struct {
	echo.Context
	agent            useragent.UserAgent
	id               string
	dbq              *hwayorm.Queries
	ipfsClient       common.IPFS
	tokenStore       common.IPFSTokenStore
	stagedEnclaves   map[string]mpc.Enclave
	grpcAddr         string
	turnstileSiteKey string
}

func UseGateway(env hway.Hway, ipc common.IPFS, db *hwayorm.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ua := useragent.NewParser()
			ctx := &GatewayContext{
				turnstileSiteKey: env.GetTurnstileSiteKey(),
				agent:            ua.Parse(c.Request().UserAgent()),
				Context:          c,
				dbq:              db,
				ipfsClient:       ipc,
				grpcAddr:         env.GetSonrGrpcUrl(),
				tokenStore:       common.NewUCANStore(ipc),
			}
			return next(ctx)
		}
	}
}
