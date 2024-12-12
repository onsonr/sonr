package sessions

import (
	"os"
	"path/filepath"
	"strings"

	config "github.com/onsonr/sonr/pkg/config/hway"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewGormDB initializes and returns a configured database connection
func NewGormDB(env config.Hway) (*gorm.DB, error) {
	// Try PostgreSQL first if DSN is provided
	if dsn := env.GetPsqlDSN(); dsn != "" && !strings.Contains(dsn, "password= ") {
		if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err == nil {
			// Successfully connected to PostgreSQL
			db.AutoMigrate(&Session{})
			db.AutoMigrate(&User{})
			return db, nil
		}
	}

	// Fall back to SQLite
	path := formatDBPath(env.GetSqliteFile())
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&Session{})
	db.AutoMigrate(&User{})
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
