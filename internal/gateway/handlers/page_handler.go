package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/context"
	"github.com/onsonr/sonr/internal/gateway/views"
	"net/http"
		"github.com/onsonr/sonr/internal/gateway/models"

	"github.com/onsonr/sonr/pkg/common/response"
)

func RenderIndex(c echo.Context) error {
	return response.TemplEcho(c, views.InitialView(context.IsUnavailableDevice(c)))
}

func RenderProfileCreate(c echo.Context) error {
	return response.TemplEcho(c, views.CreateProfileForm(context.GetCreateProfileData(c)))
}

func RenderPasskeyCreate(c echo.Context) error {
	return response.TemplEcho(c, views.CreatePasskeyForm(context.GetPasskeyCreateData(c)))
}

func RenderVaultLoading(c echo.Context) error {
	credentialJSON := c.FormValue("credential")
	if credentialJSON == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing credential data")
	}
	_, err := models.ExtractCredentialDescriptor(credentialJSON)
	if err != nil {
		return err
	}
	return response.TemplEcho(c, views.LoadingVaultView())
}
