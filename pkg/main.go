package main

import (
	"log"
	"time"

	"context"

	"github.com/getlantern/systray"
	ui "github.com/sonr-io/core/pkg/interface"
	sv "github.com/sonr-io/core/pkg/service"

	md "github.com/sonr-io/core/internal/models"
)

type SysInfo struct {
	OLC       string
	Device    md.Device
	Directory md.Directories
}

// Define Context
var desk *sv.Client
var ctx context.Context
var app ui.Interface

func main() {
	ctx = context.Background()
	systray.Run(onReady, onExit)
}

func onReady() {
	// Starts Menu Bar
	app = ui.Start()

	// Creates New Client
	desk = sv.NewClient(ctx, app)
}

func onExit() {
	log.Println(time.Now())
	ctx.Done()
}
