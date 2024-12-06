package session

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Middleware creates a new session middleware
func Middleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := NewHTTPContext(c, db)
			if err := cc.InitSession(); err != nil {
				return err
			}
			return next(cc)
		}
	}
}
