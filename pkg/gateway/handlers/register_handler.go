package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/gateway/middleware"
	"github.com/onsonr/sonr/pkg/gateway/types"
	"github.com/onsonr/sonr/pkg/gateway/views"
)

func RenderProfileCreate(c echo.Context) error {
	// numF, numL := middleware.GetHumanVerificationNumbers(c)
	params := types.CreateProfileParams{
		FirstNumber: int(middleware.CurrentBlock(c)),
		LastNumber:  int(middleware.CurrentBlock(c)),
	}
	return middleware.Render(c, views.RegisterProfileView(params))
}

func RenderPasskeyCreate(c echo.Context) error {
	return middleware.Render(c, views.RegisterPasskeyView(types.CreatePasskeyParams{}))
}

func RenderVaultLoading(c echo.Context) error {
	return middleware.Render(c, views.LoadingView())
}
