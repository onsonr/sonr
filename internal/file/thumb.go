package file

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"math"
	"os"

	"github.com/nfnt/resize"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ** Resize Constants ** //
const MAX_WIDTH float64 = 320
const MAX_HEIGHT float64 = 240

// ^ Method to generate thumbnail for Image at File Path ^ //
func NewThumbnail(req *md.ProcessRequest, sm *SafePreview) {
	// Initialize
	thumbBuffer := new(bytes.Buffer)

	// @ 1. Check for External File Request
	if req.IsExternal {
		// Get Existing Thumbnail
		if req.ThumbnailPath != "" {
			file, err := os.Open(req.ThumbnailPath)
			checkError(err)

			// Convert to Image Object
			img, _, err := image.Decode(file)
			checkError(err)

			// Encode as Jpeg into buffer
			err = jpeg.Encode(thumbBuffer, img, nil)
			checkError(err)

			// Update Thumbnail Value
			sm.preview.Thumbnail = thumbBuffer.Bytes()
			file.Close()
		}

		// @ 2. Handle Created File Request
	} else {
		// Check File Type
		if sm.Type == md.MIME_image {
			// Open File
			file, err := os.Open(req.FilePath)
			checkError(err)

			// Convert to Image Object
			img, _, err := image.Decode(file)
			checkError(err)

			scaledImage := scaleImage(img)

			// Encode as Jpeg into buffer
			err = jpeg.Encode(thumbBuffer, scaledImage, nil)
			checkError(err)

			// Update Thumbnail Value
			sm.preview.Thumbnail = thumbBuffer.Bytes()
			file.Close()
		}
	}

	// ** Unlock ** //
	sm.mutex.Unlock()

	// Get Metadata
	preview := sm.GetPreview()

	// Convert to bytes
	data, err := proto.Marshal(preview)
	if err != nil {
		log.Println("Error Marshaling Metadata ", err)
	}

	// @ 3. Callback with Preview
	sm.OnQueued(data)
}

// @ Helper Method Checks for Error ^ //
func checkError(err error) {
	if err != nil {
		onError(err, "Thumbnail")
		log.Fatalln(err)
	}
}

// @ Calculate the size of the image after scaling
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(MAX_WIDTH/float64(srcWidth), MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

func scaleImage(img image.Image) image.Image {
	// Find Image Bounds
	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	// Calculate Fit
	newWidth, newHeight := calculateRatioFit(width, height)
	return resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos3)
}
