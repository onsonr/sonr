package middleware

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

func (m *Matrix) RegisterUser(handle string) *MatrixClient {
	mc := m.ChatService.NewMautrixClient(id.NewUserID(handle, m.HomeServer))
	return &MatrixClient{Client: mc, serverName: m.HomeServer}
}

type MatrixClient struct {
	Client     *mautrix.Client
	serverName string
}

// SendMessage sends a message to a room
func (mc *MatrixClient) SendMessage(c echo.Context) error {
	sentEvent, err := mc.Client.SendText(c.Request().Context(), id.RoomID((c.Param("room"))), c.Param("message"))
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, sentEvent)
	return nil
}

// GetMessages returns all messages in a room
func (mc *MatrixClient) GetMessages(c echo.Context) error {
	events, err := mc.Client.Messages(c.Request().Context(), id.RoomID((c.Param("room"))), c.Param("from"), c.Param("to"), mautrix.DirectionForward, &mautrix.FilterPart{}, 15)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, events)
	return nil
}

// JoinedRooms returns all rooms the user is in
func (mc *MatrixClient) JoinedRooms(c echo.Context) error {
	rooms, err := mc.Client.JoinedRooms(c.Request().Context())
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, rooms)
	return nil
}

// JoinRoom joins a room
func (mc *MatrixClient) JoinRoom(c echo.Context) error {
	resp, err := mc.Client.JoinRoom(c.Request().Context(), c.Param("room"), mc.serverName, nil)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, resp)
	return nil
}

// LeaveRoom leaves a room
func (mc *MatrixClient) LeaveRoom(c echo.Context) error {
	resp, err := mc.Client.LeaveRoom(c.Request().Context(), id.RoomID(c.Param("room")))
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, resp)
	return nil
}

// GetRoomMembers returns all members in a room
func (mc *MatrixClient) GetRoomMembers(c echo.Context) error {
	members, err := mc.Client.Members(c.Request().Context(), id.RoomID(c.Param("room")))
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, members)
	return nil
}

// GetRoomState returns the state of a room
func (mc *MatrixClient) GetRoomState(c echo.Context) error {
	state, err := mc.Client.State(c.Request().Context(), id.RoomID(c.Param("room")))
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, state)
	return nil
}
