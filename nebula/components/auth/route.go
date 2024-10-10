package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/ctx"
)

func AuthorizeRoute(c echo.Context) error {
	return ctx.RenderTempl(c, Modal(c))
}

func LoginRoute(c echo.Context) error {
	return ctx.RenderTempl(c, Modal(c))
}

func RegisterRoute(c echo.Context) error {
	return ctx.RenderTempl(c, Modal(c))
}
