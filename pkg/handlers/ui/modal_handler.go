package handlers

import (
	"github.com/labstack/echo/v4"

	templates "github.com/sonrhq/sonr/pkg/components/modals"
	"github.com/sonrhq/sonr/pkg/middleware/common"
)

var Modals = modals{}

type modals struct{}

func (p modals) Deposit(c echo.Context) error {
	return common.Render(c, templates.DepositModal(false))
}

func (p modals) Swap(c echo.Context) error {
	return common.Render(c, templates.SwapModal(false))
}

func (p modals) Share(c echo.Context) error {
	return common.Render(c, templates.ShareModal(false))
}

func (p modals) Settings(c echo.Context) error {
	return common.Render(c, templates.SettingsModal(false))
}

func (p modals) Alert(c echo.Context) error {
	return common.Render(c, templates.AlertModal(false))
}
