package file

import (
	"os"
	"sync"

	"fmt"
	_ "image/gif"
	_ "image/jpeg"

	"github.com/google/uuid"
	"github.com/h2non/filetype"
	pb "github.com/sonr-io/core/pkg/models"
)

const DEFAULT_MAX_WIDTH float64 = 320
const DEFAULT_MAX_HEIGHT float64 = 240

// ^ GetMetadata generates file metadata ^ //
func GetMetadata(filePath string) *pb.Metadata {
	// Initialize
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	// Get Info
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Get File Type
	head := make([]byte, 261)
	file.Read(head)
	kind, _ := filetype.Match(head)
	file.Close()
	return &pb.Metadata{
		FileId: uuid.New().String(),
		Name:   fileName(filePath),
		Path:   filePath,
		Size:   info.Size(),
		Kind:   kind.MIME.Type,
	}
}

// ^ GetThumbnail creates thumbnail if necessary ^ //
func GetThumbnail(wg *sync.WaitGroup, meta *pb.Metadata) []byte {
	fmt.Println("Metadata type: " + meta.GetKind())
	// Check for Image
	if meta.GetKind() == "image" {
		fmt.Println("File is an image, Create a thumbnail")

		// Generate Thumbnail
		thumb, err := createThumbnail(meta.GetPath())
		if err != nil {
			fmt.Println("Error Creating Thumbnail")
			return nil
		}

		// Log And Wait for Group
		fmt.Println("Thumbnail created")
		return thumb
	}
	fmt.Println("Not an image returning nil")
	return nil
}
