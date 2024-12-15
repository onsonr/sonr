// Package gateway provides the default routes for the Sonr hway.
package gateway

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	hwayctx "github.com/onsonr/sonr/internal/gateway/context"
	"github.com/onsonr/sonr/internal/gateway/handlers"
	"github.com/onsonr/sonr/pkg/common/response"
	config "github.com/onsonr/sonr/pkg/config/hway"
	"github.com/onsonr/sonr/pkg/ipfsapi"

	_ "github.com/mattn/go-sqlite3"
)

type Gateway = *echo.Echo

func New(env config.Hway, ipc ipfsapi.Client) (Gateway, error) {
}

func RegisterRoutes(e *echo.Echo, env config.Hway, ipc ipfsapi.Client) error {
	// Custom error handler for gateway
	e.HTTPErrorHandler = response.RedirectOnError("http://localhost:3000")

	// Inject session middleware with database connection
	e.Use(hwayctx.Middleware(env))

	// Register View Handlers
	e.GET("/", handlers.RenderIndex)
	e.GET("/register", handlers.RenderProfileCreate)
	e.POST("/register/passkey", handlers.RenderPasskeyCreate)
	e.POST("/register/finish", handlers.RenderVaultLoading)

	// Register Validation Handlers
	e.POST("/register/profile/handle", handlers.ValidateProfileHandle)
	e.POST("/register/profile/is_human", handlers.ValidateIsHumanSum)
	e.POST("/submit/profile/handle", handlers.SubmitProfileHandle)
	e.POST("/submit/credential", handlers.SubmitPublicKeyCredential)
	return nil
}

// createServer sets up the server
func createServer(env config.Hway, ipc ipfsapi.Client) *echo.Echo {
	e := echo.New()
	e.Use(echoprometheus.NewMiddleware("hway"))
	e.IPExtractor = echo.ExtractIPDirect()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	RegisterRoutes(e, env, ipc)
	return e
}
