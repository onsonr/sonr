//go:build js && wasm
// +build js,wasm

package main

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/vfs/wasm"

	"github.com/donseba/go-htmx"
	"github.com/onsonr/sonr/internal/db"
)

var dwn *DWN

type DWN struct {
	*echo.Echo
	DB   *db.DB
	Htmx *htmx.HTMX
}

func main() {
	dwn = initRails()

	dwn.GET("/", htmxHandler)
	dwn.POST("/accounts", accountsHandler)

	wasm.Serve(dwn.Echo)
}

// initRails initializes the Rails application
func initRails() *DWN {
	// Open the database
	cnfg := db.New()
	db, err := cnfg.Open()
	if err != nil {
		panic(err.Error())
	}

	// Initialize the htmx handler
	return &DWN{
		Echo: echo.New(),
		DB:   db,
		Htmx: htmx.New(),
	}
}

func accountsHandler(e echo.Context) error {
	switch e.Request().Method {
	case "GET":
	case "PUT":
		// dwn.DB.AddAccount(e.Param("account"), e.Param("address"))
	default:
		e.Error(errors.New("Method not allowed"))
	}
	return e.JSON(200, "OK")
}

func htmxHandler(e echo.Context) error {
	// initiate a new htmx handler
	h := dwn.Htmx.NewHandler(e.Response(), e.Request())

	// check if the request is a htmx request
	if h.IsHxRequest() {
		// do something
	}

	// check if the request is boosted
	if h.IsHxBoosted() {
		// do something
	}

	// check if the request is a history restore request
	if h.IsHxHistoryRestoreRequest() {
		// do something
	}

	// check if the request is a prompt request
	if h.RenderPartial() {
		// do something
	}

	// set the headers for the response, see docs for more options
	h.PushURL("http://push.url")
	h.ReTarget("#ReTarged")

	// write the output like you normally do.
	// check the inspector tool in the browser to see that the headers are set.
	_, err := h.Write([]byte("OK"))
	return err
}

//
// func credentialsHandler(e echo.Context) error {
// 	switch e.Request().Method {
// 	case "GET":
// 		dwn.DB.GetCredentials(e.Param("account"))
// 	case "PUT":
// 		dwn.DB.AddCredentials(e.Param("account"), e.Param("address"))
// 	default:
// 		e.Error(errors.New("Method not allowed"))
// 	}
// 	return e.JSON(200, "OK")
// }
//
// func keysharesHandler(e echo.Context) error {
// 	switch e.Request().Method {
// 	case "GET":
// 		dwn.DB.GetKeyshares(e.Param("account"))
// 	case "PUT":
// 		dwn.DB.AddKeyshares(e.Param("account"), e.Param("address"))
// 	default:
// 		e.Error(errors.New("Method not allowed"))
// 	}
// 	return e.JSON(200, "OK")
// }
