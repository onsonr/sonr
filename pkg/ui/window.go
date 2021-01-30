package ui

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"log"

	"fyne.io/fyne/v2/canvas"
	"github.com/skip2/go-qrcode"
)

// ^ Opens a Window For QR Code ^ //
func (sm *AppInterface) OpenQRWindow(json string) {
	// Encode to QR
	print(json)
	qrData, err := qrcode.Encode(json, qrcode.Medium, 256)
	if err != nil {
		log.Panicln(err)
	}

	// Create Byte Reader
	reader := bytes.NewReader(qrData)

	// Initialize New App
	w := sm.fyApp.NewWindow("Sonr Device QR Code")

	image := canvas.NewImageFromReader(reader, UUID())
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)
	w.ShowAndRun()
}

// ^ Generates UUID ^ //
func UUID() string {
	// Generate UUID
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Panicln(err)
	}

	// Format String
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
