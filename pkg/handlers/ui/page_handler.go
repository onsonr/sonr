package handlers

import (
	"github.com/labstack/echo/v4"

	templates "github.com/sonrhq/sonr/pkg/components/pages"
	"github.com/sonrhq/sonr/pkg/middleware/shared"
)

var Pages = pages{}

type pages struct{}

func (p pages) Index(c echo.Context) error {
	if !shared.Cookies(c).HasHandle() {
		return shared.Render(c, templates.Register(c))
	}
	return shared.Render(c, templates.Chat(c))
}

func (p pages) Error(c echo.Context) error {
	return shared.Render(c, templates.Error404View())
}
