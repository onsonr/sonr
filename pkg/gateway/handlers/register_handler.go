package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/gateway/middleware"
	"github.com/onsonr/sonr/pkg/gateway/views"
)

func RenderProfileCreate(c echo.Context) error {
	// numF, numL := middleware.GetHumanVerificationNumbers(c)
	params := middleware.CreateProfileParams{
		FirstNumber: int(middleware.CurrentBlock(c)),
		LastNumber:  int(middleware.CurrentBlock(c)),
	}
	return middleware.Render(c, views.RegisterProfileView(params.FirstNumber, params.LastNumber))
}

func RenderPasskeyCreate(c echo.Context) error {
	return middleware.Render(c, views.RegisterPasskeyView("", "", "", "", ""))
}

func RenderVaultLoading(c echo.Context) error {
	return middleware.Render(c, views.LoadingView())
}
