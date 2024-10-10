package pay

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/ctx"
)

func Route(c echo.Context) error {
	return ctx.RenderTempl(c, nil)
}
