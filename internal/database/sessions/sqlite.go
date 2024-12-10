package sessions

import (
	"os"
	"path/filepath"

	"github.com/onsonr/sonr/internal/gateway/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewGormDB initializes and returns a configured database connection
func NewGormDB(env config.Env) (*gorm.DB, error) {
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

func formatDBPath(path string) string {
	home := os.Getenv("HOME")
	if home == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		home = "."
	}

	configDir := filepath.Join(home, ".config", "hway")
	if err := os.MkdirAll(configDir, 0o755); err != nil {
		// If we can't create the directory, fall back to current directory
		return path
	}

	return filepath.Join(configDir, path)
}
