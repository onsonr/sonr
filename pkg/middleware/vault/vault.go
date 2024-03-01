package vault

import "github.com/labstack/echo/v4"

type vault struct {
	echo.Context
}

func Vault(ctx echo.Context) *vault {
	return &vault{
		Context: ctx,
	}
}
