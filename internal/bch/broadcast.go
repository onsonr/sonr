//go:build js && wasm

package channel

import (
	"syscall/js"

	"github.com/labstack/echo/v4"
)

type BroadcastContext struct {
	echo.Context
	Channel js.Value
}

func (c *BroadcastContext) BroadcastMessage(message string) {
	c.Channel.Call("postMessage", message)
}
