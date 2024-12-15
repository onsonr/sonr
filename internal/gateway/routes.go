// Package gateway provides the default routes for the Sonr hway.
package gateway

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	hwayctx "github.com/onsonr/sonr/internal/gateway/context"
	"github.com/onsonr/sonr/internal/gateway/context/sink"
	"github.com/onsonr/sonr/internal/gateway/handlers"
	"github.com/onsonr/sonr/pkg/common/response"
	config "github.com/onsonr/sonr/pkg/config/hway"
	"github.com/onsonr/sonr/pkg/ipfsapi"

	_ "github.com/mattn/go-sqlite3"
)

func RegisterRoutes(e *echo.Echo, env config.Hway, db *sql.DB, ipc ipfsapi.Client) error {
	// Custom error handler for gateway
	e.HTTPErrorHandler = response.RedirectOnError("http://localhost:3000")

	// Inject session middleware with database connection
	e.Use(hwayctx.Middleware(db, env, ipc))

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

// NewDB initializes and returns a configured database connection
func NewDB(env config.Hway) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	// create tables
	if _, err := db.ExecContext(context.Background(), sink.SchemaSQL); err != nil {
		return nil, err
	}
	return db, nil
}

func formatDBPath(fileName string) string {
	configDir := filepath.Join(os.Getenv("XDG_CONFIG_HOME"), "hway")
	if err := os.MkdirAll(configDir, 0o755); err != nil {
		// If we can't create the directory, fall back to current directory
		return configDir
	}

	return filepath.Join(configDir, fileName)
}
