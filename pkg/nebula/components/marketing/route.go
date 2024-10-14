package marketing

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/internal/ctx"
)

func HomeRoute(c echo.Context) error {
	s, err := ctx.GetHWAYContext(c)
	if err != nil {
		return err
	}
	log.Printf("Session ID: %s", s.ID())
	return ctx.RenderTempl(c, View())
}
