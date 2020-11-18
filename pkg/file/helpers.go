package file

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

// ^ createThumbnail generates thumbnail for path and provided save path ^ //
func createThumbnail(imagePath string) ([]byte, error) {
	// Open File Path
	file, _ := os.Open(imagePath)
	defer file.Close()

	// Convert to Image Object
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	// Find Image Bounds
	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y
	fmt.Println("width = ", width, " height = ", height)

	// Calculate Fit
	w, h := calculateRatioFit(width, height)
	fmt.Println("w = ", w, " h = ", h)

	// Call the resize library for image scaling
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, m, nil)
	return buf.Bytes(), nil
}

// Calculate the size of the image after scaling
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// Get FileName without Extension
func fileName(path string) string {
	fileBase := filepath.Base(path)
	return strings.TrimSuffix(fileBase, filepath.Ext(fileBase))
}
