package common

import (
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"

	"github.com/liujiawm/graphics-go/graphics"
	"github.com/sonr-io/core/tools/logger"
	"go.uber.org/zap"
)

// MIN_THUMBNAIL_BOUNDS is the minimum size of the thumbnail
const MIN_THUMBNAIL_BOUNDS = 240

// IsFile Checks if Path is Valid File
func IsFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

// NewFileItem creates a new transfer file item
func NewFileItem(path string) (*Payload_Item, error) {
	// Extracts File Infrom from path
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// Create MIME
	mime, err := NewMime(path)
	if err != nil {
		return nil, err
	}

	// Create File Item
	fileItem := &FileItem{
		Mime:         mime,
		Path:         path,
		Size:         fi.Size(),
		Name:         fi.Name(),
		LastModified: fi.ModTime().Unix(),
	}

	// Check if File is Image
	if fileItem.Mime.IsImage() {
		// Init Thumbnail Path
		thumbTempPath, err := NewTempPath(fileItem.Path, WithSuffix("thumb"))
		if err != nil {
			logger.Error("Failed to retreive Temporary Path", zap.Error(err))
			return nil, err
		}

		// Create Thumbnail
		thumb, err := NewThumbnail(fileItem.Path, thumbTempPath)
		if err != nil {
			logger.Error("Failed to create Thumbnail", zap.Error(err))
			return nil, err
		}

		// Set Thumbnail
		fileItem.Thumbnail = thumb

		// Returns transfer item
		return &Payload_Item{
			Size: fi.Size(),
			Mime: mime,
			Data: &Payload_Item_File{
				File: fileItem,
			},
			Preview: &Payload_Item_Thumbnail{
				Thumbnail: fileItem.GetThumbnail(),
			},
		}, nil
	}

	// Returns transfer item
	return &Payload_Item{
		Size: fi.Size(),
		Mime: mime,
		Data: &Payload_Item_File{
			File: fileItem,
		},
	}, nil
}

// NewThumbnail creates a thumbnail from source path
func NewThumbnail(srcPath, destPath string) (*Thumbnail, error) {
	// Open File
	imagePath, err := os.Open(srcPath)
	if err != nil {
		logger.Error("Failed to open file", zap.Error(err))
		return nil, err
	}
	defer imagePath.Close()

	// Decode Image
	srcImage, _, err := image.Decode(imagePath)
	if err != nil {
		logger.Error("Failed to decode image", zap.Error(err))
		return nil, err
	}

	// Find Thumbnail Size
	srcWidth := srcImage.Bounds().Max.X
	srcHeight := srcImage.Bounds().Max.Y

	// Calculate Thumbnail Size
	w, h, ar := getThumbWidthHeight(srcWidth, srcHeight)

	// Dimension of new thumbnail 80 X 80
	dstImage := image.NewRGBA(image.Rect(0, 0, w, h))
	// Thumbnail function of Graphics
	err = graphics.Thumbnail(dstImage, srcImage)
	if err != nil {
		logger.Error("Failed to create Thumbnail", zap.Error(err))
		return nil, err
	}

	// Create Thumbnail
	newImage, err := os.Create(destPath)
	if err != nil {
		logger.Error("Failed to create new image", zap.Error(err))
		return nil, err
	}
	defer newImage.Close()

	// Encode Thumbnail
	err = jpeg.Encode(newImage, dstImage, &jpeg.Options{Quality: 100})
	if err != nil {
		logger.Error("Failed to encode image", zap.Error(err))
		return nil, err
	}

	// Read all bytes from thumbnail
	thumbBytes, err := ioutil.ReadFile(destPath)
	if err != nil {
		logger.Error("Failed to read thumbnail bytes after processing.", zap.Error(err))
		return nil, err
	}

	// Create Thumbnail
	return &Thumbnail{
		Size:        int64(len(thumbBytes)),
		Buffer:      thumbBytes,
		Width:       int32(w),
		Height:      int32(h),
		AspectRatio: float64(ar),
	}, nil
}

// ToTransferItem Returns Transfer for FileItem
func (f *FileItem) ToTransferItem() *Payload_Item {
	return &Payload_Item{
		Mime: f.GetMime(),
		Data: &Payload_Item_File{
			File: f,
		},
	}
}

// getThumbWidthHeight returns the width and height of the thumbnail by aspect ratio
func getThumbWidthHeight(srcWidth, srcHeight int) (int, int, float32) {
	// Set Min Width/Height
	width := srcWidth
	height := srcHeight
	aspectRatio := float32(srcWidth) / float32(srcHeight)

	// Calculate Bounds with larger width
	if width > height {
		width = MIN_THUMBNAIL_BOUNDS
		height = int(float32(width) / aspectRatio)
		return width, height, aspectRatio
	} else {
		height = MIN_THUMBNAIL_BOUNDS
		width = int(float32(height) * aspectRatio)
		return width, height, aspectRatio
	}
}
