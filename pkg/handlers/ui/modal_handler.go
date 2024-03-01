package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/pkg/middleware/shared"
	templates "github.com/sonrhq/sonr/pkg/components/modals"
)

var Modals = modals{}

type modals struct{}

func (p modals) Deposit(c echo.Context) error {
	return shared.Render(c, templates.DepositModal(false))
}

func (p modals) Swap(c echo.Context) error {
	return shared.Render(c, templates.SwapModal(false))
}

func (p modals) Share(c echo.Context) error {
	return shared.Render(c, templates.ShareModal(false))
}

func (p modals) Settings(c echo.Context) error {
	return shared.Render(c, templates.SettingsModal(false))
}

func (p modals) Alert(c echo.Context) error {
	return shared.Render(c, templates.AlertModal(false))
}
