package islands

import (
	"bytes"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func echoFormResponse(c echo.Context, cmp templ.Component, state FormState) error {
	// Create a buffer to store the rendered HTML
	buf := &bytes.Buffer{}
	// Render the component to the buffer
	err := cmp.Render(c.Request().Context(), buf)
	if err != nil {
		return err
	}

	// Set the content type
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	c.Response().Header().Set("X-Status", string(state))

	// Write the buffered content to the response
	_, err = c.Response().Write(buf.Bytes())
	return err
}

type FormState string

const (
	InitialForm FormState = "initial"
	ErrorForm   FormState = "error"
	SuccessForm FormState = "success"
	WarningForm FormState = "warning"
)

type Form struct {
	State       FormState
	Title       string
	Description string
}
