//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/ctx"
	"github.com/onsonr/sonr/internal/dwn"
	dwngen "github.com/onsonr/sonr/internal/dwn/gen"
	"github.com/onsonr/sonr/pkg/workers/routes"
)

var config *dwngen.Config

func main() {
	// Load dwn config
	if err := loadDwnConfig(); err != nil {
		panic(err)
	}

	// Setup HTTP server
	e := echo.New()
	e.Use(ctx.SessionMiddleware)
	routes.RegisterClientAPI(e)
	routes.RegisterClientViews(e)
	dwn.Serve(e)
}

func loadDwnConfig() error {
	// Read dwn.json config
	dwnBz, err := os.ReadFile("dwn.json")
	if err != nil {
		return err
	}
	dwnConfig := &dwngen.Config{}
	err = json.Unmarshal(dwnBz, dwnConfig)
	if err != nil {
		return err
	}
	config = dwnConfig
	return nil
}
