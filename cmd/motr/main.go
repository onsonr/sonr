//go:build js && wasm
// +build js,wasm

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"syscall/js"

	"github.com/labstack/echo/v4"
	promise "github.com/nlepage/go-js-promise"

	"github.com/onsonr/sonr/internal/ctx"
	"github.com/onsonr/sonr/pkg/nebula/routes"
	"github.com/onsonr/sonr/pkg/nebula/worker"
)

func main() {
	e := echo.New()
	e.Use(ctx.UseSession)
	registerViews(e)
	registerState(e)
	Serve(e)
}

func registerState(e *echo.Echo) {
	g := e.Group("state")
	g.POST("/login/:identifier", worker.HandleCredentialAssertion)
	g.GET("/jwks", worker.GetJWKS)
	g.GET("/token", worker.GetToken)
	g.POST("/:origin/grant/:subject", worker.GrantAuthorization)
	g.POST("/register/:subject", worker.HandleCredentialCreation)
	g.POST("/register/:subject/check", worker.CheckSubjectIsValid)
}

func registerViews(e *echo.Echo) {
	e.GET("/home", routes.Home)
	e.GET("/login", routes.LoginStart)
	e.GET("/register", routes.RegisterStart)
}

// Serve serves HTTP requests using handler or http.DefaultServeMux if handler is nil.
func Serve(handler http.Handler) func() {
	h := handler
	if h == nil {
		h = http.DefaultServeMux
	}

	prefix := js.Global().Get("wasmhttp").Get("path").String()
	for strings.HasSuffix(prefix, "/") {
		prefix = strings.TrimSuffix(prefix, "/")
	}

	if prefix != "" {
		mux := http.NewServeMux()
		mux.Handle(prefix+"/", http.StripPrefix(prefix, h))
		h = mux
	}

	cb := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		resPromise, resolve, reject := promise.New()

		go func() {
			defer func() {
				if r := recover(); r != nil {
					if err, ok := r.(error); ok {
						reject(fmt.Sprintf("wasmhttp: panic: %+v\n", err))
					} else {
						reject(fmt.Sprintf("wasmhttp: panic: %v\n", r))
					}
				}
			}()

			res := NewResponseRecorder()

			h.ServeHTTP(res, Request(args[0]))

			resolve(res.JSResponse())
		}()

		return resPromise
	})

	js.Global().Get("wasmhttp").Call("setHandler", cb)

	return cb.Release
}

// Request builds and returns the equivalent http.Request
func Request(r js.Value) *http.Request {
	jsBody := js.Global().Get("Uint8Array").New(promise.Await(r.Call("arrayBuffer")))
	body := make([]byte, jsBody.Get("length").Int())
	js.CopyBytesToGo(body, jsBody)

	req := httptest.NewRequest(
		r.Get("method").String(),
		r.Get("url").String(),
		bytes.NewBuffer(body),
	)

	headersIt := r.Get("headers").Call("entries")
	for {
		e := headersIt.Call("next")
		if e.Get("done").Bool() {
			break
		}
		v := e.Get("value")
		req.Header.Set(v.Index(0).String(), v.Index(1).String())
	}

	return req
}

// ResponseRecorder uses httptest.ResponseRecorder to build a JS Response
type ResponseRecorder struct {
	*httptest.ResponseRecorder
}

// NewResponseRecorder returns a new ResponseRecorder
func NewResponseRecorder() ResponseRecorder {
	return ResponseRecorder{httptest.NewRecorder()}
}

// JSResponse builds and returns the equivalent JS Response
func (rr ResponseRecorder) JSResponse() js.Value {
	res := rr.Result()

	body := js.Undefined()
	if res.ContentLength != 0 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		body = js.Global().Get("Uint8Array").New(len(b))
		js.CopyBytesToJS(body, b)
	}

	init := make(map[string]interface{}, 2)

	if res.StatusCode != 0 {
		init["status"] = res.StatusCode
	}

	if len(res.Header) != 0 {
		headers := make(map[string]interface{}, len(res.Header))
		for k := range res.Header {
			headers[k] = res.Header.Get(k)
		}
		init["headers"] = headers
	}

	return js.Global().Get("Response").New(body, init)
}
