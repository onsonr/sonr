package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/middleware"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/internal/gateway/views"
)

func RenderIndex(c echo.Context) error {
	isForbidden := middleware.ForbiddenDevice(c)
	return middleware.Render(c, views.InitialView(isForbidden))
}

func RenderProfileCreate(c echo.Context) error {
	numF, numL := middleware.GetHumanVerificationNumbers(c)
	params := models.CreateProfileParams{
		FirstNumber: int(numF),
		LastNumber:  int(numL),
	}
	return middleware.Render(c, views.CreateProfileForm(params))
}

func RenderPasskeyCreate(c echo.Context) error {
	return middleware.Render(c, views.CreatePasskeyForm(models.CreatePasskeyParams{}))
}

func RenderVaultLoading(c echo.Context) error {
	return middleware.Render(c, views.LoadingVaultView())
}
