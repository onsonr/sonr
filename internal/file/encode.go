package file

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func EncodeImage(imgByte []byte, path string) (string, error) {
	// Read Buffer
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}

	// Create File at Path
	out, _ := os.Create(path)
	defer out.Close()

	// Encode to Jpeg
	var opts jpeg.Options
	opts.Quality = 1

	// Encode bytes into file
	err = jpeg.Encode(out, img, &opts)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return path, nil
}

func DecodeImage(path string) (*bytes.Buffer, error) {
	// Initialize Buffer
	imgBuffer := new(bytes.Buffer)

	// New File for ThumbNail
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	// Convert to Image Object
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Encode as Jpeg into buffer
	err = jpeg.Encode(imgBuffer, img, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return imgBuffer, nil
}
