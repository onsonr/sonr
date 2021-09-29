package common

import (
	//"errors"
	//"image/jpeg"
	"os"
	//"github.com/bakape/thumbnailer/v2"
	//"github.com/sonr-io/core/internal/device"
	//"github.com/sonr-io/core/tools/logger"
	//"go.uber.org/zap"
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

	// // Check if File is Image
	// if fileItem.Mime.PermitsThumbnail() {
	// 	// Create Thumbnail
	// 	thumb, err := fileItem.NewThumbnail()
	// 	if err != nil {
	// 		logger.Error("Failed to create Thumbnail", zap.Error(err))
	// 		return nil, err
	// 	}

	// 	// Set Thumbnail
	// 	fileItem.Thumbnail = thumb

	// 	// Returns transfer item
	// 	return &Payload_Item{
	// 		Size: fi.Size(),
	// 		Mime: mime,
	// 		Data: &Payload_Item_File{
	// 			File: fileItem,
	// 		},
	// 		Preview: &Payload_Item_Thumbnail{
	// 			Thumbnail: fileItem.GetThumbnail(),
	// 		},
	// 	}, nil
	// }

	// Returns transfer item
	return &Payload_Item{
		Size: fi.Size(),
		Mime: mime,
		Data: &Payload_Item_File{
			File: fileItem,
		},
	}, nil
}

// // NewThumbnail creates a thumbnail from source path
// func (fi *FileItem) NewThumbnail() (*Thumbnail, error) {
// 	// Check if path is valid and a File
// 	if fi.Path != "" && IsFile(fi.Path) {
// 		// Open File
// 		file, err := os.Open(fi.GetPath())
// 		if err != nil {
// 			logger.Error("Failed to open file", zap.Error(err))
// 			return nil, err
// 		}
// 		defer file.Close()

// 		// Create FFmpeg Reader
// 		ffctx, err := thumbnailer.NewFFContext(file)
// 		if err != nil {
// 			logger.Error("Failed to create FFContext", zap.Error(err))
// 			return nil, err
// 		}

// 		// Create Thumbnail
// 		thumbImg, err := ffctx.Thumbnail(thumbnailer.Dims{
// 			Width:  MIN_THUMBNAIL_BOUNDS,
// 			Height: MIN_THUMBNAIL_BOUNDS,
// 		})
// 		if err != nil {
// 			logger.Error("Failed to create thumbnail", zap.Error(err))
// 			return nil, err
// 		}

// 		// Init Thumbnail Path
// 		thumbPath, err := device.NewTempPath(fi.Path, device.WithSuffix("tmb"))
// 		if err != nil {
// 			logger.Error("Failed to retreive Temporary Path", zap.Error(err))
// 			return nil, err
// 		}

// 		// Create Thumbnail at path
// 		thumbFile, err := os.Create(thumbPath)
// 		if err != nil {
// 			logger.Error("Failed to create new image", zap.Error(err))
// 			return nil, err
// 		}
// 		defer thumbFile.Close()

// 		// Encode Thumbnail
// 		err = jpeg.Encode(thumbFile, thumbImg, &jpeg.Options{Quality: 100})
// 		if err != nil {
// 			logger.Error("Failed to encode image", zap.Error(err))
// 			return nil, err
// 		}

// 		// Read all bytes from thumbnail
// 		thumbBytes, err := os.ReadFile(thumbPath)
// 		if err != nil {
// 			logger.Error("Failed to read thumbnail bytes after processing.", zap.Error(err))
// 			return nil, err
// 		}

// 		// Create Thumbnail
// 		return &Thumbnail{
// 			Size:   int64(len(thumbBytes)),
// 			Buffer: thumbBytes,
// 			Mime:   fi.GetMime(),
// 		}, nil
// 	} else {
// 		return nil, errors.New("Invalid File Path provided for item.")
// 	}
// }

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
