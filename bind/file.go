package sonr

import (
	"bytes"
	"math"
	"os"

	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"

	"github.com/h2non/filetype"
	"github.com/nfnt/resize"
	"github.com/sonr-io/core/pkg/user"
)

const DEFAULT_MAX_WIDTH float64 = 320
const DEFAULT_MAX_HEIGHT float64 = 240

type Metadata struct {
	id        int
	name      string
	owner     string // Profile JSON String
	size      int64
	thumbnail []byte
	kind      string
	path      string
	received  string // DateTime as string
}

// ^ newMetadata generates file metadata and creates thumbnail if necessary ^ //
func newMetadata(ownr user.Info, filePath string) (*Metadata, error) {
	// Initialize
	meta := new(Metadata)
	meta.path = filePath
	file, _ := os.Open(filePath)

	// Get Info
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// Set Info
	meta.size = info.Size()
	meta.owner = ownr.String()

	// Get File Type
	head := make([]byte, 261)
	file.Read(head)
	kind, _ := filetype.Match(head)
	meta.kind = kind.Extension

	// Check for Image
	if filetype.IsImage(head) {
		fmt.Println("File is an image, Creating Thumbnail")

		// Generate Thumbnail
		thum, err := createThumbnail(filePath)
		if err != nil {
			return nil, err
		}

		// Set Thumbnail
		meta.thumbnail = thum

	}
	return meta, nil
}

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

	// Create Buffer
	buf := new(bytes.Buffer)

	// Encode to JPEG
	err = jpeg.Encode(buf, m, nil)
	if err != nil {
		return nil, err
	}

	// Convert Image to Byte List
	thumbnail := buf.Bytes()
	return thumbnail, nil
}

// Calculate the size of the image after scaling
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}
