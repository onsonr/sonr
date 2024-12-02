package credentials

import "github.com/labstack/echo/v4"

type CredentialsManager interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	List(c echo.Context) error
	HasWebAuthn(c echo.Context) error
}
