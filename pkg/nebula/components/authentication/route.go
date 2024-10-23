package authentication

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/ctx"
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
	return ctx.RenderTempl(c, CurrentView(c))
}

// ╭───────────────────────────────────────────────────────────╮
// │              Hway Routes - Authentication                 │
// ╰───────────────────────────────────────────────────────────╯

// AuthorizeModalRoute returns the Authorize Modal route.
func AuthorizeModalRoute(c echo.Context) error {
	return ctx.RenderTempl(c, AuthorizeModal(c))
}

// LoginModalRoute returns the Login Modal route.
func LoginModalRoute(c echo.Context) error {
	return ctx.RenderTempl(c, LoginModal(c))
}

// RegisterModalRoute returns the Register Modal route.
func RegisterModalRoute(c echo.Context) error {
	return ctx.RenderTempl(c, RegisterModal(c))
}
