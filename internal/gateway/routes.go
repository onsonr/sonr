// Package gateway provides the default routes for the Sonr hway.
package gateway

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/context"
	"github.com/onsonr/sonr/internal/gateway/handlers"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/pkg/common/response"
	config "github.com/onsonr/sonr/pkg/config/hway"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, env config.Hway, db *gorm.DB) error {
	// Custom error handler for gateway
	e.HTTPErrorHandler = response.RedirectOnError("http://localhost:3000")

	// Inject session middleware with database connection
	e.Use(context.Middleware(db, env))

	// Register View Handlers
	e.GET("/", handlers.RenderIndex)
	e.GET("/register", handlers.RenderProfileCreate)
	e.POST("/register/passkey", handlers.RenderPasskeyCreate)
	e.POST("/register/loading", handlers.RenderVaultLoading)

	// Register Validation Handlers
	e.PUT("/register/profile/submit", handlers.ValidateProfileSubmit)
	e.PUT("/register/passkey/submit", handlers.ValidateCredentialSubmit)
	return nil
}

// NewGormDB initializes and returns a configured database connection
func NewDB(env config.Hway) (*gorm.DB, error) {
	// Try PostgreSQL first if DSN is provided
	if dsn := env.GetPsqlDSN(); dsn != "" && !strings.Contains(dsn, "password= ") {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			// Test the connection
			sqlDB, err := db.DB()
			if err == nil {
				if err = sqlDB.Ping(); err == nil {
					// Successfully connected to PostgreSQL
					db.AutoMigrate(&models.Credential{})
					db.AutoMigrate(&models.Session{})
					db.AutoMigrate(&models.User{})
					return db, nil
				}
			}
		}
	}

	// Fall back to SQLite
	path := formatDBPath(env.GetSqliteFile())
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&models.Credential{})
	db.AutoMigrate(&models.Session{})
	db.AutoMigrate(&models.User{})
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
