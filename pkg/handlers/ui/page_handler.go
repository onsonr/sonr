package handlers

import (
	"github.com/labstack/echo/v4"

	templates "github.com/sonrhq/sonr/pkg/components/pages"
	"github.com/sonrhq/sonr/pkg/middleware/common"
)

var Pages = pages{}

type pages struct{}

func (p pages) Index(c echo.Context) error {
	if common.JWT(c).HasController() {
		return common.Render(c, templates.Chat(c))
	}
	if common.Cookies(c).HasHandle() {
		return common.Render(c, templates.Login(c))
	}
	return common.Render(c, templates.Register(c))
}

func (p pages) Error(c echo.Context) error {
	return common.Render(c, templates.Error404View())
}
