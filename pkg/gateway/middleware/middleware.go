package middleware

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/medama-io/go-useragent"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/internal/config/hway"
	"github.com/onsonr/sonr/internal/database/repository"
	"github.com/onsonr/sonr/pkg/common"
)

type GatewayContext struct {
	echo.Context
	agent          useragent.UserAgent
	id             string
	dbq            *repository.Queries
	ipfsClient     common.IPFS
	tokenStore     common.IPFSTokenStore
	stagedEnclaves map[string]mpc.Enclave
	grpcAddr       string
}

func UseGateway(env hway.Hway, ipc common.IPFS, db *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ua := useragent.NewParser()
			ctx := &GatewayContext{
				agent: ua.Parse(c.Request().UserAgent()),
				Context:    c,
				dbq:        repository.New(db),
				ipfsClient: ipc,
				grpcAddr:   env.GetSonrGrpcUrl(),
				tokenStore: common.NewUCANStore(ipc),
			}
			return next(ctx)
		}
	}
}
