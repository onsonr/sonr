package routes

import (
	"bytes"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/nebula/components/auth"
	"github.com/onsonr/sonr/pkg/nebula/components/home"
)

// ╭───────────────────────────────────────────────────────────╮
// │                  Marketing Pages                          │
// ╰───────────────────────────────────────────────────────────╯

func Home(c echo.Context) error {
	return render(c, home.View())
}

// ╭───────────────────────────────────────────────────────────╮
// │                  Authentication Views                     │
// ╰───────────────────────────────────────────────────────────╯

func AuthorizeStart(c echo.Context) error {
	return render(c, auth.Modal(c))
}

func AuthorizeFinish(c echo.Context) error {
	return render(c, auth.Modal(c))
}

func LoginDevice(c echo.Context) error {
	return render(c, auth.Modal(c))
}

func LoginStart(c echo.Context) error {
	return render(c, auth.Modal(c))
}

func LoginFinish(c echo.Context) error {
	return render(c, auth.Modal(c))
}

func RegisterStart(c echo.Context) error {
	return render(c, auth.Modal(c))
}

func RegisterFinish(c echo.Context) error {
	return render(c, auth.Modal(c))
}

// ╭───────────────────────────────────────────────────────────╮
// │                       Helper Methods                      │
// ╰───────────────────────────────────────────────────────────╯

func render(c echo.Context, cmp templ.Component) error {
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
