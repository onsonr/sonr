package routes

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/ctx"
	"github.com/onsonr/sonr/pkg/nebula/marketing"
)

// ╭───────────────────────────────────────────────────────────╮
// │                  Home Routes - Marketing                  │
// ╰───────────────────────────────────────────────────────────╯

func HomeRoute(c echo.Context) error {
	s, err := ctx.GetHWAYContext(c)
	if err != nil {
		return err
	}
	log.Printf("Session ID: %s", s.ID())
	return ctx.RenderTempl(c, marketing.View())
}
