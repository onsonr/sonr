package ui

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/url"

	"github.com/zserge/lorca"
)

// ^ Opens a Window For QR Code ^ //
func (sm *AppInterface) OpenQRWindow(json string) {
	// Encode to QR
	print(json)
	// qrData, err := qrcode.Encode(json, qrcode.Medium, 256)
	// if err != nil {
	// 	log.Panicln(err)
	// }

	// Create Byte Reader
	//reader := bytes.NewReader(qrData)
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
	<html>
		<head><title>Hello</title></head>
		<body><h1>Hello, world!</h1></body>
	</html>
	`), "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
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
