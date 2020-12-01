package file

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func GetBase64Image(errorCallback OnError, path string) (string, error) {
	imgBuffer := new(bytes.Buffer)
	// New File for ThumbNail
	file, err := os.Open(meta.Path)
	if err != nil {
		fmt.Println(err)
		errorCallback(err, "getBase64Image")
		return nil, err
	}
	defer file.Close()

	// Convert to Image Object
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		errorCallback(err, "getBase64Image")
		return nil, err
	}

	// Encode as Jpeg into buffer
	err = jpeg.Encode(imgBuffer, img, nil)
	if err != nil {
		fmt.Println(err)
		errorCallback(err, "getBase64Image")
		return nil, err
	}

	return base64.StdEncoding.EncodeToString(imgBuffer.Bytes()), nil
}
