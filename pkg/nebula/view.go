package nebula

import (
	"bytes"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(cmp templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
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
}
