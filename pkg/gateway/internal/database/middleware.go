package database

import (
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/gateway/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseContext struct {
	echo.Context
	db *gorm.DB
}

func Middleware(env config.Env) echo.MiddlewareFunc {
	cc := initDB(env.GetSqliteFile())
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc.Context = c
			return next(cc)
		}
	}
}

func (c *DatabaseContext) HasDB() bool {
	return c.db != nil
}

func initDB(path string) *DatabaseContext {
	cc := new(DatabaseContext)
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		cc.db = nil
		return cc
	}
	// Migrate the schema
	db.AutoMigrate(&Session{})
	db.AutoMigrate(&User{})

	return &DatabaseContext{
		db: db,
	}
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
