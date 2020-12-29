package main

import (
	"log"
	"time"

	"github.com/getlantern/systray"
	"github.com/sonr-io/core/pkg/ui"
	//	"github.com/sonr-io/core/bind"
)

func main() {
	onExit := func() {
		now := time.Now()
		log.Println(now)
	}

	systray.Run(onReady, onExit)
}

func onReady() {
	// node := sonr.NewNode(reqBytes []byte, call sonr.Callback)
	// We can manipulate the systray in other goroutines
	ui.StartTray()
}
