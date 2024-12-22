package handlers

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/app/gateway/context"
	"github.com/onsonr/sonr/app/gateway/islands"
	"github.com/onsonr/sonr/app/gateway/views"
	hwayorm "github.com/onsonr/sonr/internal/database/hwayorm"
	"github.com/onsonr/sonr/pkg/common"
)

func RegisterHandler(g *echo.Group) {
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
	cc, err := context.GetGateway(c)
	if err != nil {
		return err
	}
	handle := c.FormValue("handle")
	origin := c.FormValue("origin")
	name := c.FormValue("name")
	cc.InsertProfile(context.BG(), hwayorm.InsertProfileParams{
		Handle: handle,
		Origin: origin,
		Name:   name,
	})

	params, err := cc.Spawn(handle, origin)
	if err != nil {
		return context.RenderError(c, err)
	}
	return context.Render(c, views.RegisterPasskeyView(params.Address, params.Handle, params.Name, params.Challenge, params.CreationBlock))
}

func renderVaultStatus(c echo.Context) error {
	return context.Render(c, views.LoadingView())
}

// ╭─────────────────────────────────────────────────────────╮
// │                  Validation Components                  │
// ╰─────────────────────────────────────────────────────────╯

func validateProfileForm(c echo.Context) error {
	cc, err := context.GetGateway(c)
	if err != nil {
		return context.RenderError(c, err)
	}
	handle := c.FormValue("handle")
	if handle == "" {
		return context.Render(c, islands.InputHandleError(handle, "Please enter a 4-16 character handle"))
	}
	notok, err := cc.CheckHandleExists(context.BG(), handle)
	if err != nil {
		return err
	}
	if notok {
		return context.Render(c, islands.InputHandleError(handle, "Handle is already taken"))
	}
	cc.WriteCookie(common.UserHandle, handle)
	return context.Render(c, islands.InputHandleSuccess(handle))
}

func validatePasskeyForm(c echo.Context) error {
	cc, err := context.GetGateway(c)
	if err != nil {
		return context.RenderError(c, err)
	}
	handle := context.GetProfileHandle(c)
	origin := c.Request().Host
	credentialJSON := c.FormValue("credential")
	cred := &context.CredentialDescriptor{}

	// Unmarshal the credential JSON
	err = json.Unmarshal([]byte(credentialJSON), cred)
	if err != nil {
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
	return context.Render(c, views.LoadingView())
}
