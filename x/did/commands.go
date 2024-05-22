package module

import (
	"context"

	"github.com/labstack/echo/v4"
)

// RunAsyncLocalAuthServer runs the local auth server asynchronously to authenticate via the command line.
func runAsyncLocalAuthServer(_ context.Context) error {
	e := echo.New()
	SetRouterLocal(e)
	return e.Start("sonr.local")
}
