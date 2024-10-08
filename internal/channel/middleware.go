//go:build js && wasm

package channel

import (
	"syscall/js"

	"github.com/labstack/echo/v4"
)

type JSHandler func(this js.Value, args []js.Value) interface{}

func UseBroadcastChannel(channelName string, handler JSHandler) echo.MiddlewareFunc {
	var channel js.Value

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if channel.IsUndefined() {
				channel = js.Global().Get("BroadcastChannel").New(channelName)
				channel.Call("addEventListener", "message", handler)
			}

			cc := &BroadcastContext{
				Context: c,
				Channel: channel,
			}
			return next(cc)
		}
	}
}

func PostBroadcastMessage(c echo.Context, message string) {
	cc := c.(*BroadcastContext)
	cc.BroadcastMessage(message)
}
