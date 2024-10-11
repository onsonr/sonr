package home

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/internal/ctx"
)

func Route(c echo.Context) error {
	s := ctx.GetSession(c)
	log.Println(s.ID())
	return ctx.RenderTempl(c, View())
}
