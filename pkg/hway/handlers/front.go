package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common/middleware/render"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/pkg/nebula/marketing"
	"github.com/onsonr/sonr/pkg/nebula/modals"
)

// ╭───────────────────────────────────────────────────────────╮
// │                  Home Routes - Marketing                  │
// ╰───────────────────────────────────────────────────────────╯

func RenderHome(c echo.Context) error {
	return render.Templ(c, marketing.View())
}

// RenderLogin returns the Login Modal route.
func RenderLogin(c echo.Context) error {
	cc, err := session.Get(c)
	if err != nil {
		return err
	}
	return render.Templ(c, modals.LoginModal(cc))
}

// RenderRegister returns the Register Modal route.
func RenderRegister(c echo.Context) error {
	cc, err := session.Get(c)
	if err != nil {
		return err
	}
	return render.Templ(c, modals.RegisterModal(cc))
}
