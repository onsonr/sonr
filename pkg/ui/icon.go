package ui

import (
	"bytes"
	"image/png"
	"log"
	"os"
)

// ** Icon Type for Image
type Icon int

const (
	SystemTray Icon = iota
	Close
	User
	Peer
	Invite
)

// ** Const UI Resource Path ** //
const RES_PATH = "/Users/prad/Sonr/core/pkg/res/"

func (d Icon) File() string {
	return [...]string{"systray.png", "close.png", "user.png", "peer.png", "invite.png"}[d]
}

// ^ Returns Buffer of Image by Icon Type
func GetIcon(i Icon) []byte {
	// Get Path
	path := RES_PATH + i.File()

	// Initialize
	imgBuffer := new(bytes.Buffer)
	existingImageFile, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer existingImageFile.Close()

	// Decode PNG
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		log.Println(err)
		return nil
	}
	// Encode PNG into Memory
	err = png.Encode(imgBuffer, loadedImage)
	if err != nil {
		log.Println(err)
		return nil
	}
	return imgBuffer.Bytes()
}
