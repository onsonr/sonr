package file

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func EncodeImage(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Opening File: ", err)
	}
	defer file.Close()

	// Convert to Image Object
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Decoding: ", err)
	}

	err = jpeg.Encode(file, img, nil)
	if err != nil {
		fmt.Println("Encoding: ", err)
	}
}
