package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/internal/gateway/views"
	"github.com/onsonr/sonr/internal/nebula/form"
	"github.com/onsonr/sonr/pkg/common/response"
	"github.com/onsonr/sonr/pkg/passkeys"
)

func RenderProfileRegister(c echo.Context) error {
	d := form.RandomCreateProfileData()
	return response.TemplEcho(c, views.CreateProfileForm(d))
}

func RenderPasskeyStart(c echo.Context) error {
	challenge, _ := protocol.CreateChallenge()
	handle := c.FormValue("handle")
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")

	ks, err := mpc.NewKeyset()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	dat := form.CreatePasskeyData{
		Address:       ks.Address(),
		Handle:        handle,
		Name:          fmt.Sprintf("%s %s", firstName, lastName),
		Challenge:     challenge.String(),
		CreationBlock: "00001",
	}
	return response.TemplEcho(c, views.CreatePasskeyForm(dat))
}

func RenderPasskeyFinish(c echo.Context) error {
	// Get the raw credential JSON string
	credentialJSON := c.FormValue("credential")
	if credentialJSON == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing credential data")
	}
	_, err := passkeys.ExtractCredential(credentialJSON)
	if err != nil {
		return err
	}
	return response.TemplEcho(c, views.LoadingVaultView())
}
