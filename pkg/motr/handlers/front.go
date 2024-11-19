package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/app/nebula/views"
	"github.com/onsonr/sonr/app/nebula/views/auth"
	"github.com/onsonr/sonr/pkg/common/middleware/render"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
)

// ╭───────────────────────────────────────────────────────────╮
// │               DWN Routes - Authentication                 │
// ╰───────────────────────────────────────────────────────────╯

// CurrentViewRoute returns the current view route.
func CurrentViewRoute(c echo.Context) error {
	cc, err := session.Get(c)
	if err != nil {
		return err
	}
	return render.Templ(c, views.CurrentView(cc))
}

// ╭───────────────────────────────────────────────────────────╮
// │              Hway Routes - Authentication                 │
// ╰───────────────────────────────────────────────────────────╯

// AuthorizeModalRoute returns the Authorize Modal route.
func AuthorizeModalRoute(c echo.Context) error {
	cc, err := session.Get(c)
	if err != nil {
		return err
	}
	return render.Templ(c, auth.AuthorizeModal(cc))
}

// LoginModalRoute returns the Login Modal route.
func LoginModalRoute(c echo.Context) error {
	cc, err := session.Get(c)
	if err != nil {
		return err
	}
	return render.Templ(c, auth.LoginModal(cc))
}

// RegisterModalRoute returns the Register Modal route.
func RegisterModalRoute(c echo.Context) error {
	cc, err := session.Get(c)
	if err != nil {
		return err
	}
	return render.Templ(c, auth.RegisterModal(cc))
}
