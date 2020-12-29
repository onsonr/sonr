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
)

// ** Const UI Resource Path ** //
const path = "/Users/prad/Sonr/core/pkg/res/"

func (d Icon) String() string {
	return [...]string{"SystemTray", "Close", "User", "Peer"}[d]
}

// ^ Returns Buffer of Image by Icon Type
func GetIcon(i Icon) []byte {
	var image []byte
	switch i {
	case SystemTray:
		return readIcon(path + "icon.png")
	case Close:
		log.Println(" goes down.")
	case User:
		log.Println(" goes down.")
	case Peer:
		log.Println(" goes down.")
	default:
		log.Println(" stays put.")
	}
	return image
}

// ^ Reads Decodes/Encodes Image from Path ^ //
func readIcon(path string) []byte {
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
