//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/ctx"
	"github.com/onsonr/sonr/internal/dwn/fetch"
	dwngen "github.com/onsonr/sonr/internal/dwn/gen"
	"github.com/onsonr/sonr/pkg/workers/routes"
)

const FileNameConfigJSON = "dwn.json"

var config *dwngen.Config

func main() {
	// Load dwn config
	if err := loadDwnConfig(); err != nil {
		panic(err)
	}

	// Setup HTTP server
	e := echo.New()
	e.Use(ctx.DWNSessionMiddleware(config))
	routes.RegisterWebNodeAPI(e)
	routes.RegisterWebNodeViews(e)
	fetch.Serve(e)
}

func loadDwnConfig() error {
	// Read dwn.json config
	dwnBz, err := os.ReadFile(FileNameConfigJSON)
	if err != nil {
		return err
	}
	dwnConfig := new(dwngen.Config)
	err = json.Unmarshal(dwnBz, dwnConfig)
	if err != nil {
		return err
	}
	config = dwnConfig
	return nil
}
