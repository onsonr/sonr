package matrix

import (
	"github.com/labstack/echo/v4"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/appservice"
	"maunium.net/go/mautrix/id"
)

type Matrix struct {
	HomeServer   string
	ChatService  appservice.AppService
	EventService appservice.AppService
}

func NewMatrix() *Matrix {
	return &Matrix{}
}

func UseMatrixClient(c echo.Context) *mautrix.Client {
	m := NewMatrix()
	handle := c.Param("handle")
	return m.ChatService.NewMautrixClient(id.NewUserID(handle, m.HomeServer))
}
