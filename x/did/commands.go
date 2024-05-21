package module

import (
	"context"

	"github.com/labstack/echo/v4"
)

// RunAsyncLocalAuthServer runs the local auth server asynchronously to authenticate via the command line.
func runAsyncLocalAuthServer(ctx context.Context) error {
	e := echo.New()
	return e.Start("sonr.local")
}
