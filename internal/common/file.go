package common

import (
	//"errors"
	//"image/jpeg"
	"os"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/sonr-io/core/internal/device"

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

// ReplaceDir replaces the directory of the item path
func (fi *FileItem) ReplaceDir(dir string) error {
	// Get New Path
	path, err := device.NewPath(fi.GetPath(), dir)
	if err != nil {
		logger.Error("Failed to replace directory for FileItem", err)
		return err
	}

	// Set Path
	fi.Path = path
	return nil
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

// ** ───────────────────────────────────────────────────────
// ** ─── MIME Management ───────────────────────────────────
// ** ───────────────────────────────────────────────────────
// DefaultUrlMIME is the standard MIME type for URLs
func DefaultUrlMIME() *MIME {
	return &MIME{
		Type:    MIME_URL,
		Subtype: ".html",
		Value:   "url/html",
	}
}

// NewMime creates a new MIME type from Path
func NewMime(path string) (*MIME, error) {
	// Check if path is URL
	if IsUrl(path) {
		return DefaultUrlMIME(), nil
	}

	// Check if path to file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	// Get MIME Type and Set Proto Enum
	mtype, err := mimetype.DetectFile(path)
	if err != nil {
		return nil, err
	}

	// Return MIME Type
	return &MIME{
		Type:    MIME_Type(MIME_Type_value[strings.ToUpper(mtype.Parent().String())]),
		Value:   mtype.String(),
		Subtype: mtype.Extension(),
	}, nil
}

// Ext Method adjusts extension for JPEG
func (m *MIME) Ext() string {
	if m.Subtype == "jpg" || m.Subtype == "jpeg" {
		return "jpeg"
	}
	return m.Subtype
}

// IsAudio Checks if Mime is Audio
func (m *MIME) IsAudio() bool {
	return m.Type == MIME_AUDIO
}

// IsImage Checks if Mime is Image
func (m *MIME) IsImage() bool {
	return m.Type == MIME_IMAGE
}

// IsMedia Checks if Mime is any media
func (m *MIME) IsMedia() bool {
	return m.IsAudio() || m.IsImage() || m.IsVideo()
}

// IsPDF Checks if Mime is PDF
func (m *MIME) IsPDF() bool {
	return strings.Contains(m.GetSubtype(), "pdf")
}

// IsVideo Checks if Mime is Video
func (m *MIME) IsVideo() bool {
	return m.Type == MIME_VIDEO
}

// IsUrl Checks if Path is a URL
func (m *MIME) IsUrl() bool {
	return m.Type == MIME_URL
}

// PermitsThumbnail Checks if Mime Type Allows Thumbnail Generation.
// Image, Video, Audio, and PDF are allowed.
func (m *MIME) PermitsThumbnail() bool {
	return m.IsImage() || m.IsVideo() || m.IsAudio() || m.IsPDF()
}
