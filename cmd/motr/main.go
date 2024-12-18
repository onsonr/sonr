//go:build js && wasm
// +build js,wasm

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"syscall/js"

	"github.com/labstack/echo/v4"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/onsonr/sonr/cmd/motr/wasm"
	"github.com/onsonr/sonr/internal/config/motr"
	sink "github.com/onsonr/sonr/internal/models/sink/sqlite"
	vault "github.com/onsonr/sonr/pkg/vault/routes"
)

var (
	env    *motr.Environment
	config *motr.Config
	err    error
)

func broadcastTx(this js.Value, args []js.Value) interface{} {
	return nil
}

func simulateTx(this js.Value, args []js.Value) interface{} {
	return nil
}

func syncData(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return nil
	}

	configString := args[0].String()
	if err := json.Unmarshal([]byte(configString), &config); err != nil {
		println("Error parsing config:", err.Error())
		return nil
	}
	return nil
}

func main() {
	// Load dwn config
	js.Global().Set("broadcastTx", js.FuncOf(broadcastTx))
	js.Global().Set("simulateTx", js.FuncOf(simulateTx))
	js.Global().Set("syncData", js.FuncOf(syncData))

	e := echo.New()
	e.Use(wasm.ContextMiddleware)
	// e.Use(controller.Middleware(nil))
	vault.RegisterRoutes(e, config)
	wasm.ServeFetch(e)
}

// NewDB initializes and returns a configured database connection
func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	// create tables
	if _, err := db.ExecContext(context.Background(), sink.SchemaMotrSQL); err != nil {
		return nil, err
	}
	return db, nil
}
