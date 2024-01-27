package main

import (
	"log"

	"github.com/wailsapp/wails/v2"

	"github.com/sonrhq/sonr/pkg/desktop"
)

func main() {
	// Create an instance of the app structure
	app := desktop.NewApp()

	// Create application with options
	err := wails.Run(app.WailsOptions())
	if err != nil {
		log.Fatal(err)
	}
}
