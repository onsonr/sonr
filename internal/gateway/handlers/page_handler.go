package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/middleware"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/internal/gateway/views"
	"golang.org/x/exp/rand"
)

func RenderIndex(c echo.Context) error {
	return middleware.Render(c, views.InitialView(middleware.ForbiddenDevice(c)))
}

func RenderProfileCreate(c echo.Context) error {
	return middleware.Render(c, views.CreateProfileForm(getCreateProfileData()))
}

func RenderPasskeyCreate(c echo.Context) error {
	return middleware.Render(c, views.CreatePasskeyForm(models.CreatePasskeyParams{}))
}

func RenderVaultLoading(c echo.Context) error {
	return middleware.Render(c, views.LoadingVaultView())
}

func getCreateProfileData() models.CreateProfileParams {
	return models.CreateProfileParams{
		FirstNumber: rand.Intn(5) + 1,
		LastNumber:  rand.Intn(4) + 1,
	}
}
