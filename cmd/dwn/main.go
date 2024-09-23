//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/vfs/wasm"

	"github.com/onsonr/sonr/internal/db"
)

var dwn *DWN

type DWN struct {
	*echo.Echo
	DB *db.DB
}

func main() {
	dwn = initRails()
	wasm.Serve(dwn.Echo)
}

// initRails initializes the Rails application
func initRails() *DWN {
	// Open the database
	e := echo.New()
	db, err := db.New()
	if err != nil {
		panic(err.Error())
	}
	db.ServeEcho(e.Group("/dwn"))

	// Initialize the htmx handler
	return &DWN{
		Echo: e,
		DB:   db,
	}
}
