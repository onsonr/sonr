package main

import (
	_ "embed"

	"github.com/labstack/echo/v4"
)

//go:embed styles.css
var cssData string

func staticCSS() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Content-Type", "text/css")
			return c.String(200, cssData)
		}
	}
}
