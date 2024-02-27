package nitro

import "github.com/labstack/echo/v4"

// RegisterSynapseUser calls the Synapse server to register a new user with the given user handle and a deterministic password.
// [POST]: /_synapse/admin/v1/register
func RegisterSynapseUser(c echo.Context, handle string) error {
	return nil
}

// CreateRoom calls the Synapse server to create a new room with the given name and topic.
func CreateRoom(c echo.Context, name, topic string) error {
	return nil
}
