package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common/middleware/render"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/pkg/design/pages/auth"
	"github.com/onsonr/sonr/pkg/design/pages/home"
)

// ╭───────────────────────────────────────────────────────────╮
// │                  Home Routes - Marketing                  │
// ╰───────────────────────────────────────────────────────────╯

func RenderHome(c echo.Context) error {
	return render.Templ(c, home.View())
}

// RenderLogin returns the Login Modal route.
func RenderLogin(c echo.Context) error {
	cc, err := session.Get(c)
	if err != nil {
		return err
	}
	return render.Templ(c, auth.LoginModal(cc))
}

// RenderRegister returns the Register Modal route.
func RenderRegister(c echo.Context) error {
	cc, err := session.Get(c)
	if err != nil {
		return err
	}
	return render.Templ(c, auth.RegisterModal(cc))
}
