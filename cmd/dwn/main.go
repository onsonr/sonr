//go:build js && wasm
// +build js,wasm

package main

import (
	"errors"

	"github.com/labstack/echo/v4"
	wasmhttp "github.com/nlepage/go-wasm-http-server"

	"github.com/onsonr/sonr/internal/db"
)

var edb *db.DB

func main() {
	// Initialize the database
	initDB()

	e := echo.New()
	e.POST("/api/insert/:account/:address", itemsHandler)
	wasmhttp.Serve(e)
}

func initDB() {
	var err error
	edb, err = db.Open(db.New())
	if err != nil {
		panic(err.Error())
	}
}

func itemsHandler(e echo.Context) error {
	switch e.Request().Method {
	case "GET":
	case "POST":
		edb.AddAccount(e.Param("account"), e.Param("address"))
	default:
		e.Error(errors.New("Method not allowed"))
	}
	return e.JSON(200, "OK")
}
