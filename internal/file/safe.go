package file

import (
	"bytes"
	"image"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"fmt"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"

	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"github.com/nfnt/resize"
	pb "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ File that safely sets metadata and thumbnail in routine ^ //
type SafeMeta struct {
	mutex sync.Mutex
	meta  pb.Metadata
}

// ^ Create generates file metadata ^ //
func (sf *SafeMeta) Generate(i *Item) {
	// ** Lock ** //
	sf.mutex.Lock()

	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(i.Path)
	if err != nil {
		fmt.Println("Error opening File", err)
		i.Call.Error(err, "AddFile")
	}

	// Get Info
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		i.Call.Error(err, "AddFile")
	}

	// Get File Type
	head := make([]byte, 261)
	file.Read(head)
	kind, err := filetype.Match(head)
	if err != nil {
		fmt.Println(err)
		i.Call.Error(err, "AddFile")
	}
	file.Close()

	// @ 2. Create Thumbnail
	thumbBuffer := new(bytes.Buffer)
	if filetype.IsImage(head) {
		// New File for ThumbNail
		file, err := os.Open(i.Path)
		if err != nil {
			fmt.Println(err)
			i.Call.Error(err, "AddFile")
		}
		defer file.Close()

		// Convert to Image Object
		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Println(err)
			i.Call.Error(err, "AddFile")
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
			i.Call.Error(err, "AddFile")
		}
		fmt.Println("Thumbnail created")
	}

	// @ 3. Set Metadata Protobuf Values
	sf.meta = pb.Metadata{
		FileId:    uuid.New().String(),
		Name:      fileName(i.Path),
		Path:      i.Path,
		Size:      info.Size(),
		Blocks:    info.Size() / BlockSize,
		Kind:      kind.MIME.Type,
		Thumbnail: thumbBuffer.Bytes(),
	}
	// ** Unlock ** //
	sf.mutex.Unlock()
	i.completedMetadata(sf)
}

// ^ Safely returns metadata depending on lock ^ //
func (sf *SafeMeta) Metadata() *pb.Metadata {
	// ** Lock File wait for access ** //
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	// @ 2. Return Value
	return &sf.meta
}

// ^ Safely returns metadata depending on lock ^ //
func (sf *SafeMeta) Bytes() []byte {
	// Get Metadata
	meta := sf.Metadata()

	// Convert to bytes
	data, err := proto.Marshal(meta)
	if err != nil {
		fmt.Println("Error Marshaling Metadata ", err)
		return nil
	}
	return data
}

// ** Resize Constants ** //
const MAX_WIDTH float64 = 320
const MAX_HEIGHT float64 = 240

// @ Calculate the size of the image after scaling
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(MAX_WIDTH/float64(srcWidth), MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// @ Get FileName without Extension
func fileName(path string) string {
	fileBase := filepath.Base(path)
	return strings.TrimSuffix(fileBase, filepath.Ext(fileBase))
}
