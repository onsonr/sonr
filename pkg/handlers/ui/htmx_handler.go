package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/internal/components/views/auth/login"
	"github.com/sonrhq/sonr/internal/components/views/auth/register"
	"github.com/sonrhq/sonr/internal/components/views/dash"
	"github.com/sonrhq/sonr/internal/components/views/landing/home"
	"github.com/sonrhq/sonr/internal/components/views/utility"
	"github.com/sonrhq/sonr/pkg/middleware/common"
)

var Pages = pages{}

type pages struct{}

func (p pages) Index(c echo.Context) error {
	if common.JWT(c).HasController() {
		return common.Render(c, dash.IndexView())
	}
	if common.Cookies(c).HasHandle() {
		return common.Render(c, login.IndexView())
	}
	if common.Requests(c).PathIs("/register") {
		return common.Render(c, register.IndexView())
	}
	return common.Render(c, home.Page(c))
}

func (p pages) Error(c echo.Context) error {
	return common.Render(c, utility.ErrorNotFound(c))
}
