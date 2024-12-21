package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/gateway/context"
	"github.com/onsonr/sonr/pkg/gateway/views"
)

func RenderProfileCreate(c echo.Context) error {
	// numF, numL := middleware.GetHumanVerificationNumbers(c)
	params := context.CreateProfileParams{
		FirstNumber: int(context.StatusBlock(c)),
		LastNumber:  int(context.StatusBlock(c)),
	}
	return context.Render(c, views.RegisterProfileView(params.FirstNumber, params.LastNumber))
}

func RenderPasskeyCreate(c echo.Context) error {
	return context.Render(c, views.RegisterPasskeyView("", "", "", "", ""))
}

func RenderVaultLoading(c echo.Context) error {
	return context.Render(c, views.LoadingView())
}
