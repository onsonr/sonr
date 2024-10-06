package router

import (
	"bytes"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/nebula/components/grant"
	"github.com/onsonr/sonr/pkg/nebula/components/home"
	"github.com/onsonr/sonr/pkg/nebula/components/login"
	"github.com/onsonr/sonr/pkg/nebula/components/profile"
	"github.com/onsonr/sonr/pkg/nebula/components/register"
	"github.com/onsonr/sonr/pkg/nebula/models"
)

func Authorize(c echo.Context) error {
	return echoResponse(c, grant.View(c))
}

func Home(c echo.Context) error {
	mdls, err := models.GetModels()
	if err != nil {
		return err
	}
	return echoResponse(c, home.View(mdls.Home))
}

func Login(c echo.Context) error {
	return echoResponse(c, login.Modal(c))
}

func Profile(c echo.Context) error {
	return echoResponse(c, profile.View(c))
}

func Register(c echo.Context) error {
	return echoResponse(c, register.Modal(c))
}

// ╭───────────────────────────────────────────────────────────╮
// │                       Helper Methods                      │
// ╰───────────────────────────────────────────────────────────╯

func echoResponse(c echo.Context, cmp templ.Component) error {
	// Create a buffer to store the rendered HTML
	buf := &bytes.Buffer{}
	// Render the component to the buffer
	err := cmp.Render(c.Request().Context(), buf)
	if err != nil {
		return err
	}

	// Set the content type
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	// Write the buffered content to the response
	_, err = c.Response().Write(buf.Bytes())
	return err
}
