package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/internal/components/views/auth/login"
	"github.com/sonrhq/sonr/internal/components/views/auth/register"
	"github.com/sonrhq/sonr/internal/components/views/dash/chats"
	"github.com/sonrhq/sonr/internal/components/views/landing/changelog"
	"github.com/sonrhq/sonr/internal/components/views/landing/home"
	"github.com/sonrhq/sonr/internal/components/views/utility"
	"github.com/sonrhq/sonr/pkg/middleware/common"
)

var Pages = pages{}

type pages struct{}

func (p pages) Root(c echo.Context) error {
	if common.JWT(c).HasController() {
		return common.Render(c, chats.Page(c))
	}
	if common.Cookies(c).HasHandle() {
		return common.Render(c, login.PageView(c))
	}
	if common.Requests(c).PathIs("/register") {
		return common.Render(c, register.PageView(c))
	}
	return common.Render(c, home.Page(c))
}

func (p pages) Login(c echo.Context) error {
	return common.Render(c, login.PageView(c))
}

func (p pages) Register(c echo.Context) error {
	return common.Render(c, register.PageView(c))
}

func (p pages) Home(c echo.Context) error {
	return common.Render(c, home.Page(c))
}

func (p pages) Chats(c echo.Context) error {
	return common.Render(c, chats.Page(c))
}

func (p pages) Changelog(c echo.Context) error {
	return common.Render(c, changelog.Page(c))
}

func (p pages) Error(c echo.Context) error {
	return common.Render(c, utility.ErrorNotFound(c))
}
