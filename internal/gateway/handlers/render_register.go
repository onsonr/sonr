package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/context"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/internal/gateway/views"
	"github.com/onsonr/sonr/pkg/common/response"
)

func RenderProfileCreate(c echo.Context) error {
	d, err := context.GetCreateProfileData(c)
	if err != nil {
		return err
	}
	return response.TemplEcho(c, views.CreateProfileForm(d))
}

func RenderPasskeyCreate(c echo.Context) error {
	dat, err := context.GetPasskeyCreateData(c)
	if err != nil {
		return err
	}
	return response.TemplEcho(c, views.CreatePasskeyForm(dat))
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
