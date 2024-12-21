package handlers

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	hwayorm "github.com/onsonr/sonr/internal/database/hwayorm"
	"github.com/onsonr/sonr/internal/nebula/input"
	"github.com/onsonr/sonr/pkg/gateway/context"
	"github.com/onsonr/sonr/pkg/gateway/views"
)

func HandleRegistration(g *echo.Group) {
	g.GET("/", renderProfileForm)
	g.POST("/profile", validateProfileForm)
	g.GET("/passkey", renderPasskeyForm)
	g.POST("/passkey", validatePasskeyForm)
	g.GET("/vault", renderVaultStatus)
}

// ╭──────────────────────────────────────────────────────╮
// │                  Registration Views                  │
// ╰──────────────────────────────────────────────────────

func renderProfileForm(c echo.Context) error {
	params := context.CreateProfileParams{
		FirstNumber: 6,
		LastNumber:  3,
	}
	return context.Render(c, views.RegisterProfileView(params.FirstNumber, params.LastNumber))
}

func renderPasskeyForm(c echo.Context) error {
	return context.Render(c, views.RegisterPasskeyView("", "", "", "", ""))
}

func renderVaultStatus(c echo.Context) error {
	return context.Render(c, views.LoadingView())
}

// ╭─────────────────────────────────────────────────────────╮
// │                  Validation Components                  │
// ╰─────────────────────────────────────────────────────────╯

func validateProfileForm(c echo.Context) error {
	// value := c.FormValue("is_human")
	handle := c.FormValue("handle")
	if handle == "" {
		return context.Render(c, input.HandleError(handle, "Please enter a valid handle"))
	}
	return context.Render(c, input.HandleSuccess(handle))
}

func validatePasskeyForm(c echo.Context) error {
	cc, err := context.GetGateway(c)
	if err != nil {
		return err
	}
	handle := context.GetProfileHandle(c)
	origin := c.Request().Host
	credentialJSON := c.FormValue("credential")
	cred := &context.CredentialDescriptor{}
	// Unmarshal the credential JSON
	if err := json.Unmarshal([]byte(credentialJSON), cred); err != nil {
		return context.RenderError(c, err)
	}
	md := cred.ToModel(handle, origin)
	_, err = cc.InsertCredential(context.BG(), hwayorm.InsertCredentialParams{
		Handle:       md.Handle,
		CredentialID: md.CredentialID,
		Origin:       md.Origin,
		Type:         md.Type,
		Transports:   md.Transports,
	})
	if err != nil {
		return context.RenderError(c, err)
	}
	return nil
}
