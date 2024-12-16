// Package gateway provides the default routes for the Sonr hway.
package gateway

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/sonr/pkg/gateway/middleware"
	"github.com/onsonr/sonr/internal/models"
	config "github.com/onsonr/sonr/internal/config/hway"
	"github.com/onsonr/sonr/pkg/gateway/routes"
	"github.com/onsonr/sonr/pkg/common/ipfs"

	_ "github.com/mattn/go-sqlite3"
)

type Gateway = *echo.Echo

func New(env config.Hway, ipc ipfs.Client) (Gateway, error) {
	db, err := models.NewDB(env)
	if err != nil {
		return nil, err
	}
	e := initServer(env, ipc, db)
	routes.RegisterPages(e)
	return e, nil
}

// initServer sets up the server
func initServer(env config.Hway, ipc ipfs.Client, db *sql.DB) *echo.Echo {
	e := echo.New()
	// Overrides
	e.HTTPErrorHandler = RedirectOnError("http://localhost:3000")
	e.IPExtractor = echo.ExtractIPDirect()

	// Built-in middleware
	e.Use(echoprometheus.NewMiddleware("hway"))
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())

	// Custom middleware
	e.Use(middleware.UseSessions(db))
	e.Use(middleware.UseCredentials(db))
	e.Use(middleware.UseResolvers(env))
	e.Use(middleware.UseProfiles(db))
	e.Use(middleware.UseVaultProvider(ipc))
	e.Use(middleware.UseRender(env))
	return e
}

func RedirectOnError(target string) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			// Log the error if needed
			c.Logger().Errorf("Error: %v", he.Message)
		}
		// Redirect to main site
		c.Redirect(http.StatusFound, target)
	}
}
