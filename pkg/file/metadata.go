package file

import (
	"bytes"
	"errors"
	"image"
	"os"
	"sync"

	"fmt"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"

	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"github.com/nfnt/resize"
	pb "github.com/sonr-io/core/pkg/models"
)

const DEFAULT_MAX_WIDTH float64 = 320
const DEFAULT_MAX_HEIGHT float64 = 240

// ** File that concurrently sets metadata and thumbnail ** //
type SafeFile struct {
	Path     string
	mutex    sync.Mutex
	metadata pb.Metadata
}

// ^ Create generates file metadata ^ //
func (sf *SafeFile) Create() {
	// ** Lock ** //
	sf.mutex.Lock()

	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(sf.Path)
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening File", err)
	}

	// Get Info
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	// Get File Type
	head := make([]byte, 261)
	file.Read(head)
	kind, _ := filetype.Match(head)

	// @ 2. Create Thumbnail
	thumbBuffer := new(bytes.Buffer)
	if filetype.IsImage(head) {
		// Convert to Image Object
		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Println(err)
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
		}
		fmt.Println("Thumbnail created")
	}

	// @ 3. Set Metadata Protobuf Values
	sf.metadata = pb.Metadata{
		FileId: uuid.New().String(),
		Name:   fileName(sf.Path),
		Path:   sf.Path,
		Size:   info.Size(),
		Kind:   kind.MIME.Type,
	}

	// @ 3. Set Thumbnail if it exists
	if thumbBuffer.Len() > 0 {
		sf.metadata.Thumbnail = thumbBuffer.Bytes()
	}

	// ** Unlock ** //
	sf.mutex.Unlock()
}

// ^ SetMetadataThumbnail creates thumbnail for given metadata ^ //
func (sf *SafeFile) Metadata() (*pb.Metadata, error) {
	// ** Lock File wait for access ** //
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	// @ 1. Check for Metadata
	if sf.metadata.GetPath() == "" {
		errMsg := fmt.Sprintf("Metadata was not found in SafeFile for '%s'", sf.Path)
		return nil, errors.New(errMsg)
	}

	// @ 2. Return Value
	return &sf.metadata, nil
}
