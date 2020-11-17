package file

import (
	"math"
	"os"
	"sync"

	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"

	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
	"github.com/nfnt/resize"
	"github.com/sonr-io/core/pkg/lobby"
)

const DEFAULT_MAX_WIDTH float64 = 320
const DEFAULT_MAX_HEIGHT float64 = 240

type Metadata struct {
	id        int
	name      string
	owner     string // Profile JSON String
	size      int64
	thumbPath string
	kind      string
	path      string
	received  string // DateTime as string
}

// ^ GetMetadata generates file metadata and creates thumbnail if necessary ^ //
func GetMetadata(ownr lobby.Peer, filePath string, cacheDir string) (*Metadata, error) {
	// Start WaitGroup
	var wg sync.WaitGroup

	// Initialize
	meta := new(Metadata)
	meta.path = filePath
	fmt.Println("FilePath: ", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// Get Info
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fmt.Println("FileInfo: ", info)

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
		// Get Save Path
		savePath := fileThumbnailPath(filePath, cacheDir)

		// Begin Wait Group
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Generate Thumbnail
			err := createThumbnail(filePath, savePath)
			// Check for error
			if err != nil {
				fmt.Println("Error Creating Thumbnail")
			}

			// Set Thumbnail
			meta.thumbPath = savePath
		}()
		wg.Wait()
		fmt.Println("Thumbnail created")
	}
	fmt.Println("Metadata: ", meta)
	return meta, nil
}

// ^ createThumbnail generates thumbnail for path and provided save path ^ //
func createThumbnail(imagePath string, savePath string) error {
	// Open File Path
	file, _ := os.Open(imagePath)
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

	// files that need to be saved
	imgfile, _ := os.Create(savePath)
	defer imgfile.Close()

	// Encode to JPEG
	err = jpeg.Encode(imgfile, m, nil)
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

// Get FileName without Extension
func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// Get Path for Thumbnail
func fileThumbnailPath(filePath string, cacheDir string) string {
	// Create Save Path
	fileBase := filepath.Base(filePath)
	fileExt := filepath.Ext(filePath)
	fileName := fileNameWithoutExtension(fileBase)
	fmt.Println("File Name: ", fileName)
	return cacheDir + "/" + fileName + "_thumb" + fileExt
}
