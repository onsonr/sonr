package file

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"math"
	"os"

	"github.com/nfnt/resize"
)

// ** Resize Constants ** //
const MAX_WIDTH float64 = 320
const MAX_HEIGHT float64 = 240

// ^ Method to generate thumbnail for Image at File Path ^ //
func newThumbnail(path string) ([]byte, error) {
	// New File for ThumbNail
	thumbBuffer := new(bytes.Buffer)
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Convert to Image Object
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Find Image Bounds
	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	// Calculate Fit
	newWidth, newHeight := calculateRatioFit(width, height)
	scaledImage := resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos3)

	// Encode as Jpeg into buffer
	err = jpeg.Encode(thumbBuffer, scaledImage, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Close File
	file.Close()
	log.Println("Thumbnail created")
	return thumbBuffer.Bytes(), nil
}

// ^ Method to generate thumbnail for Image at File Path ^ //
func newVideoThumbnail(path string) ([]byte, error) {
	// New File for ThumbNail
	thumbBuffer := new(bytes.Buffer)
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Convert to Image Object
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Find Image Bounds
	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	// Calculate Fit
	newWidth, newHeight := calculateRatioFit(width, height)
	scaledImage := resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos3)

	// Encode as Jpeg into buffer
	err = jpeg.Encode(thumbBuffer, scaledImage, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Close File
	file.Close()
	log.Println("Thumbnail created")
	return thumbBuffer.Bytes(), nil
}

// @ Calculate the size of the image after scaling
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(MAX_WIDTH/float64(srcWidth), MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}
