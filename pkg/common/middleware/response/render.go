package response

import (
	"bytes"
	"context"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Templ renders a component to the response
func Templ(cmp templ.Component) echo.HandlerFunc {
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
		if err != nil {
			return err
		}
		c.Response().WriteHeader(200)
		return nil
	}
}

// / TemplRawBytes renders a component to a byte slice
func TemplRawBytes(cmp templ.Component) ([]byte, error) {
	// Create a buffer to store the rendered HTML
	w := bytes.NewBuffer(nil)
	err := cmp.Render(context.Background(), w)
	if err != nil {
		return nil, err
	}
	dat := w.Bytes()
	return dat, nil
}
