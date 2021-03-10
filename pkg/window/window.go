package window

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"

	"github.com/progrium/macdriver/bridge"
)

func NewMacWindow() {
	h, err := bridge.NewHost(os.Stderr)
	if err != nil {
		log.Fatal(err)
	}
	go h.Run()

	window := bridge.Window{
		Title:       "Hello 1",
		Size:        bridge.Size{W: 480, H: 240},
		Position:    bridge.Point{X: 200, Y: 200},
		Closable:    true,
		Minimizable: false,
		Resizable:   false,
		Borderless:  false,
		Image:       base64.StdEncoding.EncodeToString(data),
		Background:  &bridge.Color{R: 0, G: 0, B: 1, A: 0.5},
	}
	h.Sync(&window)
}
