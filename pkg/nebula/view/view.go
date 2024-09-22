package view

import (
	"bytes"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/nebula/routes"
)

// render renders a templ.Component
func Render(r routes.Route) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Create a buffer to store the rendered HTML
		buf := bytes.NewBuffer(nil)

		// Render the component to the buffer
		err := r.Component(c).Render(c.Request().Context(), buf)
		if err != nil {
			return err
		}

		// Set the content type
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

		// Write the buffered content to the response
		_, err = c.Response().Write(buf.Bytes())
		return err
	}
}
