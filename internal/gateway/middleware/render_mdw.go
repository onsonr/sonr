package middleware

import (
	"bytes"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type RenderContext struct {
	echo.Context
}

func UseRender() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &RenderContext{Context: c}
			return next(ctx)
		}
	}
}

func Render(c echo.Context, cmp templ.Component) error {
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
