package pages

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/nebula/components/register"
)

func Register(c echo.Context) error {
	return echoResponse(c, register.Modal(c))
}
