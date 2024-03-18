package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/didao-org/sonr/pkg/components/views/auth/login"
	"github.com/didao-org/sonr/pkg/components/views/auth/register"
	"github.com/didao-org/sonr/pkg/components/views/dash/chats"
	"github.com/didao-org/sonr/pkg/components/views/landing/about"
	"github.com/didao-org/sonr/pkg/components/views/landing/changelog"
	"github.com/didao-org/sonr/pkg/components/views/landing/ecosystem"
	"github.com/didao-org/sonr/pkg/components/views/landing/home"
	"github.com/didao-org/sonr/pkg/components/views/landing/research"
	"github.com/didao-org/sonr/pkg/components/views/utility"
	"github.com/didao-org/sonr/pkg/middleware/common"
)

var Pages = pages{}

type pages struct{}

func (p pages) Root(c echo.Context) error {
	if common.JWT(c).HasController() {
		return common.Render(c, chats.Page(c))
	}
	if common.Cookies(c).HasHandle() {
		return common.Render(c, login.Page(c))
	}
	if common.Requests(c).PathIs("/register") {
		return common.Render(c, register.Page(c))
	}
	return common.Render(c, home.Page(c))
}

func (p pages) Login(c echo.Context) error {
	return common.Render(c, login.Page(c))
}

func (p pages) Register(c echo.Context) error {
	return common.Render(c, register.Page(c))
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

func (p pages) NotFound(c echo.Context) error {
	return common.Render(c, utility.ErrorNotFound(c))
}

func (p pages) Ecosystem(c echo.Context) error {
	return common.Render(c, ecosystem.Page(c))
}

func (p pages) About(c echo.Context) error {
	return common.Render(c, about.Page(c))
}

func (p pages) Research(c echo.Context) error {
	return common.Render(c, research.Page(c))
}
