package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RedirectOnError(target string) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			// Log the error if needed
			c.Logger().Errorf("Error: %v", he.Message)
		}
		// Redirect to main site
		c.Redirect(http.StatusFound, target)
	}
}
