package routes

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common/ctx"
	"github.com/onsonr/sonr/pkg/nebula/modals"
	"github.com/onsonr/sonr/pkg/nebula/views"
)

// ╭───────────────────────────────────────────────────────────╮
// │               DWN Routes - Authentication                 │
// ╰───────────────────────────────────────────────────────────╯

// CurrentViewRoute returns the current view route.
func CurrentViewRoute(c echo.Context) error {
	s, err := ctx.GetDWNContext(c)
	if err != nil {
		return err
	}
	log.Printf("Session ID: %s", s.ID())
	return ctx.RenderTempl(c, views.CurrentView(c))
}

// ╭───────────────────────────────────────────────────────────╮
// │              Hway Routes - Authentication                 │
// ╰───────────────────────────────────────────────────────────╯

// AuthorizeModalRoute returns the Authorize Modal route.
func AuthorizeModalRoute(c echo.Context) error {
	return ctx.RenderTempl(c, modals.AuthorizeModal(c))
}

// LoginModalRoute returns the Login Modal route.
func LoginModalRoute(c echo.Context) error {
	return ctx.RenderTempl(c, modals.LoginModal(c))
}

// RegisterModalRoute returns the Register Modal route.
func RegisterModalRoute(c echo.Context) error {
	return ctx.RenderTempl(c, modals.RegisterModal(c))
}
