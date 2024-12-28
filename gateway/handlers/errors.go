package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/gateway/context"
)

func ErrorHandler(err error, c echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		// Log the error if needed
		c.Logger().Errorf("Error: %v", he.Message)
		context.RenderError(c, he)
	}
}
