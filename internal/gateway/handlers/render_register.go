package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/internal/gateway/views"
	"github.com/onsonr/sonr/pkg/common/response"
	"golang.org/x/exp/rand"
)

func RenderProfileCreate(c echo.Context) error {
	d := models.CreateProfileData{
		FirstNumber: rand.Intn(5) + 1,
		LastNumber:  rand.Intn(4) + 1,
	}
	return response.TemplEcho(c, views.CreateProfileForm(d))
}

func RenderPasskeyCreate(c echo.Context) error {
	challenge, _ := protocol.CreateChallenge()
	handle := c.FormValue("handle")
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")

	ks, err := mpc.GenEnclave()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	dat := models.CreatePasskeyData{
		Address:       ks.Address(),
		Handle:        handle,
		Name:          fmt.Sprintf("%s %s", firstName, lastName),
		Challenge:     challenge.String(),
		CreationBlock: "00001",
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
