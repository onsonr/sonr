package file

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func EncodeImage(buffer bytes.Buffer, path string) {
	imgByte := buffer.Bytes()
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}

	out, _ := os.Create(path)
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 1

	err = jpeg.Encode(out, img, &opts)
	//jpeg.Encode(out, img, nil)
	if err != nil {
		log.Println(err)
	}
}
