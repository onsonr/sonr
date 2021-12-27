package main

import (
	"github.com/sonr-io/core/app"
	"github.com/sonr-io/core/node/api"
)

func main() {
	req := api.DefaultInitializeRequest()

	// Initialize App
	app.Start(req, app.WithMode(api.StubMode_BIN),
		app.WithHost(":"),
		app.WithPort(26225),
		app.WithLogLevel(app.InfoLevel),
	)
}
