package pages

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/nebula/components/profile"
)

func Profile(c echo.Context) error {
	return echoResponse(c, profile.View(c))
}
