package register

import (
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/pkg/common/response"
	"golang.org/x/exp/rand"
)

func HandleCreateProfile(c echo.Context) error {
	return response.TemplEcho(c, ProfileFormView(randomCreateProfileData()))
}

func HandlePasskeyStart(c echo.Context) error {
	challenge, _ := protocol.CreateChallenge()
	handle := c.FormValue("handle")
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")

	ks, err := mpc.NewKeyset()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	dat := RegisterPasskeyData{
		Address:       ks.Address(),
		Handle:        handle,
		Name:          fmt.Sprintf("%s %s", firstName, lastName),
		Challenge:     challenge.String(),
		CreationBlock: "00001",
	}
	return response.TemplEcho(c, LinkCredentialView(dat))
}

func HandlePasskeyFinish(c echo.Context) error {
	// Get the raw credential JSON string
	credentialJSON := c.FormValue("credential")
	if credentialJSON == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing credential data")
	}
	_, err := extractCredentialDescriptor(credentialJSON)
	if err != nil {
		return err
	}
	return response.TemplEcho(c, LoadingVaultView())
}

func randomCreateProfileData() CreateProfileData {
	return CreateProfileData{
		FirstNumber: rand.Intn(5) + 1,
		LastNumber:  rand.Intn(4) + 1,
	}
}
