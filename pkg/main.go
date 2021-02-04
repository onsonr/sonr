package main

import (
	"log"
	"time"

	"context"

	"github.com/getlantern/systray"
	"github.com/sonr-io/core/pkg/desktop"
	"github.com/sonr-io/core/pkg/ui"
)

// Define Context
var desk *desktop.Client
var ctx context.Context
var app ui.AppInterface

func main() {
	// Start Function
	ctx = context.Background()
	systray.Run(onReady, onExit)
}

func onReady() {
	// Starts Menu Bar
	app = ui.Start()

	// Creates New Client
	desk = desktop.NewClient(ctx, app)
}

func onExit() {
	log.Println(time.Now())
	ctx.Done()
}
