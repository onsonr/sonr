package file

import (
	"bytes"
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
