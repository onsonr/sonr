package database

import "github.com/labstack/echo/v4"

type DatabaseContext struct {
	echo.Context
	sqlitePath string
}
