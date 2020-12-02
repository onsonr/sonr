package file

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"math"
	"os"

	"github.com/nfnt/resize"
)

// ** Resize Constants ** //
const MAX_WIDTH float64 = 320
const MAX_HEIGHT float64 = 240

func NewThumbnail(path string) ([]byte, error) {
	// New File for ThumbNail
	thumbBuffer := new(bytes.Buffer)
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

	// Find Image Bounds
	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y
	fmt.Println("width = ", width, " height = ", height)

	// Calculate Fit
	newWidth, newHeight := calculateRatioFit(width, height)
	fmt.Println("w = ", newWidth, " h = ", newHeight)

	// Call the resize library for image scaling
	scaledImage := resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos3)

	// Encode as Jpeg into buffer
	err = jpeg.Encode(thumbBuffer, scaledImage, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Thumbnail created")
	return thumbBuffer.Bytes(), nil
}

// @ Calculate the size of the image after scaling
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(MAX_WIDTH/float64(srcWidth), MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}
