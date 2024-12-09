package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/pkg/blocks/forms"
	"github.com/onsonr/sonr/pkg/common/response"
	"github.com/onsonr/sonr/pkg/gateway/config"
	"github.com/onsonr/sonr/pkg/gateway/internal/pages/register"
)

func HandleRegisterView(env config.Env) echo.HandlerFunc {
	return func(c echo.Context) error {
		return response.TemplEcho(c, register.ProfileFormView(env.GetTurnstileSiteKey()))
	}
}

func HandleRegisterStart(c echo.Context) error {
	challenge, _ := protocol.CreateChallenge()
	handle := c.FormValue("handle")
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")

	ks, err := mpc.NewKeyset()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	dat := forms.RegisterPasskeyData{
		Address:       ks.Address(),
		Handle:        handle,
		Name:          fmt.Sprintf("%s %s", firstName, lastName),
		Challenge:     challenge.String(),
		CreationBlock: "00001",
	}
	return response.TemplEcho(c, register.LinkCredentialView(dat))
}

type Credential struct {
	ID       string `json:"id"`
	RawID    []int  `json:"rawId"`
	Type     string `json:"type"`
	Response struct {
		AttestationObject []int `json:"attestationObject"`
		ClientDataJSON    []int `json:"clientDataJSON"`
	} `json:"response"`
}

func HandleRegisterFinish(c echo.Context) error {
	credB64 := c.FormValue("credential")
	
	// Decode base64 credential
	credJSON, err := base64.StdEncoding.DecodeString(credB64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid credential encoding")
	}

	// Unmarshal credential
	var cred Credential
	if err := json.Unmarshal(credJSON, &cred); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid credential format")
	}

	// Log credential details
	fmt.Printf("Credential ID: %s\n", cred.ID)
	fmt.Printf("Credential Type: %s\n", cred.Type)
	fmt.Printf("Raw ID Length: %d\n", len(cred.RawID))
	fmt.Printf("Attestation Object Length: %d\n", len(cred.Response.AttestationObject))
	fmt.Printf("Client Data Length: %d\n", len(cred.Response.ClientDataJSON))

	return response.TemplEcho(c, register.LoadingVaultView())
}
