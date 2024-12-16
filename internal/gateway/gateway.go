// Package gateway provides the default routes for the Sonr hway.
package gateway

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/sonr/internal/gateway/middleware"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/internal/gateway/routes"
	config "github.com/onsonr/sonr/pkg/config/hway"
	"github.com/onsonr/sonr/pkg/ipfsapi"

	_ "github.com/mattn/go-sqlite3"
)

type Gateway = *echo.Echo

func New(env config.Hway, ipc ipfsapi.Client) (Gateway, error) {
	db, err := models.NewDB(env)
	if err != nil {
		return nil, err
	}
	e := createServer(env, ipc, db)
	routes.RegisterPages(e)
	return e, nil
}

// createServer sets up the server
func createServer(env config.Hway, ipc ipfsapi.Client, db *sql.DB) *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = RedirectOnError("http://localhost:3000")
	e.Use(echoprometheus.NewMiddleware("hway"))
	e.IPExtractor = echo.ExtractIPDirect()
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())

	// Custom middleware
	e.Use(middleware.UseSessions(db))
	e.Use(middleware.UseCredentials(db))
	e.Use(middleware.UseResolvers(env))
	e.Use(middleware.UseProfiles(db))
	e.Use(middleware.UseVaults(ipc))
	e.Use(middleware.UseRender())
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
