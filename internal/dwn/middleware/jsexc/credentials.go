//go:build js && wasm
// +build js,wasm

package jsexc

import (
	"errors"
	"syscall/js"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/config/dwn"
)

type Navigator struct {
	echo.Context
	navigator      js.Value
	hasCredentials bool
}

func NewNavigator(c echo.Context, cnfg *dwn.Config) *Navigator {
	navigator := js.Global().Get("navigator")
	credentials := navigator.Get("credentials")
	hasCredentials := !credentials.IsUndefined()

	return &Navigator{
		Context:        c,
		navigator:      navigator,
		hasCredentials: hasCredentials,
	}
}

func (c *Navigator) CreateCredential(options js.Value) (js.Value, error) {
	if !c.hasCredentials {
		return js.Null(), errors.New("navigator.credentials is undefined")
	}
	promise := c.navigator.Get("credentials").Call("create", map[string]interface{}{"publicKey": options})
	result, err := awaitPromise(promise)
	return result, err
}

func (c *Navigator) GetCredential(options js.Value) (js.Value, error) {
	if !c.hasCredentials {
		return js.Null(), errors.New("navigator.credentials is undefined")
	}
	promise := c.navigator.Get("credentials").Call("get", map[string]interface{}{"publicKey": options})
	result, err := awaitPromise(promise)
	return result, err
}

func awaitPromise(promise js.Value) (js.Value, error) {
	done := make(chan struct{})
	var result js.Value
	var err error

	thenFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result = args[0]
		close(done)
		return nil
	})
	catchFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		err = errors.New(args[0].String())
		close(done)
		return nil
	})
	defer thenFunc.Release()
	defer catchFunc.Release()

	promise.Call("then", thenFunc).Call("catch", catchFunc)

	<-done

	return result, err
}
