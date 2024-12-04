package database

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/gateway/database/internal"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseContext struct {
	echo.Context
	db *gorm.DB
}

func Middleware(sqlitePath string) echo.MiddlewareFunc {
	cc := initDB(sqlitePath)
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
	db.AutoMigrate(&internal.Session{})
	db.AutoMigrate(&internal.User{})

	return &DatabaseContext{
		db: db,
	}
}
