package home

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/internal/ctx"
)

func Route(c echo.Context) error {
	s := ctx.GetSession(c)
	log.Printf("Session ID: %s", s.ID)
	log.Printf("Session Origin: %s", s.Origin)
	log.Printf("Session Address: %s", s.Address)
	log.Printf("Session ChainID: %s", s.ChainID)
	return ctx.RenderTempl(c, View())
}
