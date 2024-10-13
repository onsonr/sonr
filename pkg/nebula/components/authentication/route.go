package authentication

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/ctx"
)

func AuthorizeRoute(c echo.Context) error {
	s := ctx.GetDWNContext(c)
	log.Printf("Session ID: %s", s.ID())
	return ctx.RenderTempl(c, AuthorizeModal(c))
}

func CurrentRoute(c echo.Context) error {
	s := ctx.GetDWNContext(c)
	log.Printf("Session ID: %s", s.ID())
	return ctx.RenderTempl(c, CurrentView(c))
}

func LoginRoute(c echo.Context) error {
	s := ctx.GetDWNContext(c)
	log.Printf("Session ID: %s", s.ID())
	return ctx.RenderTempl(c, LoginModal(c))
}

func RegisterRoute(c echo.Context) error {
	s := ctx.GetDWNContext(c)
	log.Printf("Session ID: %s", s.ID())
	return ctx.RenderTempl(c, RegisterModal(c))
}
