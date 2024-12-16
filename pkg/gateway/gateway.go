// Package gateway provides the default routes for the Sonr hway.
package gateway

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	config "github.com/onsonr/sonr/internal/config/hway"
	"github.com/onsonr/sonr/internal/database"
	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/gateway/middleware"
	"github.com/onsonr/sonr/pkg/gateway/routes"
)

type Gateway = *echo.Echo

func New(env config.Hway, ipc common.IPFS) (Gateway, error) {
	db, err := database.NewDB(env)
	if err != nil {
		return nil, err
	}
	e := initServer(env, ipc, db)
	routes.RegisterPages(e)
	return e, nil
}

// initServer sets up the server
func initServer(env config.Hway, ipc common.IPFS, db *sql.DB) *echo.Echo {
	e := middleware.UseBase()
	e.Use(middleware.UseSessions(db))
	e.Use(middleware.UseCredentials(db))
	e.Use(middleware.UseResolvers(env))
	e.Use(middleware.UseProfiles(db))
	e.Use(middleware.UseVaultProvider(ipc))
	e.Use(middleware.UseRender(env))
	return e
}
