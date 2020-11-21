package file

import (
	"bytes"
	"image"
	"os"

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

// ^ GetMetadata generates file metadata ^ //
func GetMetadata(filePath string) pb.Metadata {

	// Initialize
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
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
	file.Close()

	// Return Protobuf
	return pb.Metadata{
		FileId: uuid.New().String(),
		Name:   fileName(filePath),
		Path:   filePath,
		Size:   info.Size(),
		Kind:   kind.MIME.Type,
	}
}

// ^ SetMetadataThumbnail creates thumbnail for given metadata ^ //
func SetMetadataThumbnail(meta *pb.Metadata) error {
	fmt.Println("Metadata type: " + meta.GetKind())
	// Check for Image
	if meta.GetKind() == "image" {
		fmt.Println("File is an image, Create a thumbnail")

		// Open File Path
		file, _ := os.Open(meta.Path)
		defer file.Close()

		// Convert to Image Object
		img, _, err := image.Decode(file)
		if err != nil {
			return err
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

		// Set Thumbnail
		meta.Thumbnail = buf.Bytes()

		// Log And Wait for Group
		fmt.Println("Thumbnail created")
		return nil
	}
	return nil
}
