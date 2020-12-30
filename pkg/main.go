package main

import (
	"log"
	"time"

	"context"

	"github.com/getlantern/systray"
	"github.com/sonr-io/core/pkg/ui"
)

// Define Context
var ctx context.Context

func main() {
	// Exit Function
	onExit := func() {
		now := time.Now()
		log.Println(now)
		ctx.Done()
	}

	// Start Function
	ctx = context.Background()
	systray.Run(onReady, onExit)
}

func onReady() {
	// Starts New Node
	menu := ui.StartTray()

	// Creates New Client
	NewClient(ctx, menu)
}
