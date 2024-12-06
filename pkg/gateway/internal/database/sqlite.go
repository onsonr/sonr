package database

import (
	"os"
	"path/filepath"

	"github.com/onsonr/sonr/pkg/gateway/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDB initializes and returns a configured database connection
func InitDB(env config.Env) (*gorm.DB, error) {
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
	return filepath.Join(home, ".config", "hway", path)
}
