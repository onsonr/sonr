package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/gateway/context"
)

// AI! Fix this error handler to only return on 500 errors
func ErrorHandler(err error, c echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		// Log the error if needed
		c.Logger().Errorf("Error: %v", he.Message)
		context.RenderError(c, he)
	}
}
