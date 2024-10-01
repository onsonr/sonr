package pages

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/nebula/components/grant"
)

func Authorize(c echo.Context) error {
	return echoResponse(c, grant.View(c))
}
