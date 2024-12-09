package handlers

import (
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
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
	// firstName := c.FormValue("first_name")
	// lastName := c.FormValue("last_name")
	ks, err := mpc.NewKeyset()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return response.TemplEcho(c, register.LinkCredentialView(ks.Address(), handle, challenge.String()))
}

func HandleRegisterFinish(c echo.Context) error {
	// cred := c.FormValue("credential")
	return response.TemplEcho(c, register.LoadingVaultView())
}
