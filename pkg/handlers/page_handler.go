package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/pkg/shared"
	"github.com/sonrhq/sonr/pkg/templates"
)

var Pages = pages{}

type pages struct{}

func (p pages) Index(c echo.Context) error {
	return shared.Render(c, templates.AuthView(false))
}

func (p pages) Error(c echo.Context) error {
	return shared.Render(c, templates.Error404View())
}

func (p pages) Home(c echo.Context) error {
	return shared.Render(c, templates.HomePanel(c))
}

func (p pages) Console(c echo.Context) error {
	return shared.Render(c, templates.ConsolePanel(c))
}

func (p pages) Chat(c echo.Context) error {
	return shared.Render(c, templates.ChatPanel(c))
}

func (p pages) Wallet(c echo.Context) error {
	return shared.Render(c, templates.WalletPanel(c))
}

func (p pages) Status(c echo.Context) error {
	return shared.Render(c, templates.StatusPanel(c))
}

func (p pages) Governance(c echo.Context) error {
	return shared.Render(c, templates.GovernancePanel(c))
}
