package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/context"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/internal/gateway/views"
	"golang.org/x/exp/rand"

	"github.com/onsonr/sonr/pkg/common/response"
)

func RenderIndex(c echo.Context) error {
	return response.TemplEcho(c, views.InitialView(context.ForbiddenDevice(c)))
}

func RenderProfileCreate(c echo.Context) error {
	return response.TemplEcho(c, views.CreateProfileForm(getCreateProfileData()))
}

func RenderPasskeyCreate(c echo.Context) error {
	return response.TemplEcho(c, views.CreatePasskeyForm(models.CreatePasskeyParams{}))
}

func RenderVaultLoading(c echo.Context) error {
	return response.TemplEcho(c, views.LoadingVaultView())
}

func getCreateProfileData() models.CreateProfileParams {
	return models.CreateProfileParams{
		FirstNumber: rand.Intn(5) + 1,
		LastNumber:  rand.Intn(4) + 1,
	}
}
