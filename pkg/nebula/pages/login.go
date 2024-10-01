package pages

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/nebula/components/login"
)

func Login(c echo.Context) error {
	return echoResponse(c, login.Modal(c))
}
