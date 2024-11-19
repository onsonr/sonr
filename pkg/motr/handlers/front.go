package handlers

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common/middleware/render"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/pkg/nebula/modals"
	"github.com/onsonr/sonr/pkg/nebula/views"
)

// ╭───────────────────────────────────────────────────────────╮
// │               DWN Routes - Authentication                 │
// ╰───────────────────────────────────────────────────────────╯

// CurrentViewRoute returns the current view route.
func CurrentViewRoute(c echo.Context) error {
	s, err := session.Get(c)
	if err != nil {
		return err
	}
	log.Printf("Session ID: %s", s.ID())
	return render.Templ(c, views.CurrentView(c))
}

// ╭───────────────────────────────────────────────────────────╮
// │              Hway Routes - Authentication                 │
// ╰───────────────────────────────────────────────────────────╯

// AuthorizeModalRoute returns the Authorize Modal route.
func AuthorizeModalRoute(c echo.Context) error {
	return render.Templ(c, modals.AuthorizeModal(c))
}

// LoginModalRoute returns the Login Modal route.
func LoginModalRoute(c echo.Context) error {
	return render.Templ(c, modals.LoginModal(c))
}

// RegisterModalRoute returns the Register Modal route.
func RegisterModalRoute(c echo.Context) error {
	return render.Templ(c, modals.RegisterModal(c))
}
