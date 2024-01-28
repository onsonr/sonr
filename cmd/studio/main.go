package main

import (
	"log"

	"github.com/wailsapp/wails/v2"

	desktop "github.com/sonrhq/sonr/config/studio"
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
