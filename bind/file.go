package sonr

import (
	"math"
	"os"

	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"

	"github.com/nfnt/resize"
)

const DEFAULT_MAX_WIDTH float64 = 320
const DEFAULT_MAX_HEIGHT float64 = 240

type Metadata struct {
	id        int
	name      string
	owner     string // Profile JSON String
	size      int
	thumbnail []byte
	fileType  string
	path      string
	received  string // DateTime as string
}

func NewFile() {

}

func NewThumbnail(imagePath, savePath string) error {
	file, _ := os.Open(imagePath)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	w, h := calculateRatioFit(width, height)

	fmt.Println("width = ", width, " height = ", height)
	fmt.Println("w = ", w, " h = ", h)

	// Call the resize library for image scaling
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	// files that need to be saved
	imgfile, _ := os.Create(savePath)
	defer imgfile.Close()

	// save the file in PNG format
	err = png.Encode(imgfile, m)
	if err != nil {
		return err
	}

	return nil
}

// Calculate the size of the image after scaling
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}
