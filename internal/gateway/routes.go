// Package gateway provides the default routes for the Sonr hway.
package gateway

import (
	"os"
	"path/filepath"

	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/context"
	"github.com/onsonr/sonr/internal/gateway/handlers"
	"github.com/onsonr/sonr/pkg/common/response"
	config "github.com/onsonr/sonr/pkg/config/hway"
	"github.com/onsonr/sonr/pkg/ipfsapi"
	_ "modernc.org/sqlite"
)

func RegisterRoutes(e *echo.Echo, env config.Hway, db *sql.DB, ipc ipfsapi.Client) error {
	// Custom error handler for gateway
	e.HTTPErrorHandler = response.RedirectOnError("http://localhost:3000")

	// Inject session middleware with database connection
	e.Use(context.Middleware(db, env, ipc))

	// Register View Handlers
	e.GET("/", handlers.RenderIndex)
	e.GET("/register", handlers.RenderProfileCreate)
	e.POST("/register/passkey", handlers.RenderPasskeyCreate)
	e.POST("/register/finish", handlers.RenderVaultLoading)

	// Register Validation Handlers
	e.POST("/register/profile/handle", handlers.ValidateProfileHandle)
	e.POST("/register/profile/is_human", handlers.ValidateIsHumanSum)
	e.POST("/register/submit/credential", handlers.SubmitPublicKeyCredential)
	return nil
}

// NewDB initializes and returns a configured database connection
func NewDB(env config.Hway) (*sql.DB, error) {
	path := formatDBPath(env.GetSqliteFile())
	return sql.Open("sqlite3", path)
}

func formatDBPath(fileName string) string {
	configDir := filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "hway")
	if err := os.MkdirAll(configDir, 0o755); err != nil {
		// If we can't create the directory, fall back to current directory
		return configDir
	}

	return filepath.Join(configDir, fileName)
}
