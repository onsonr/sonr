package common

import (
	"os"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

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

// IsMedia Checks if Mime is any media
func (m *MIME) IsMedia() bool {
	return m.Type == MIME_AUDIO || m.Type == MIME_IMAGE || m.Type == MIME_VIDEO
}

// IsImage Checks if Mime is Image
func (m *MIME) IsImage() bool {
	return m.Type == MIME_IMAGE
}

// IsVideo Checks if Mime is Video
func (m *MIME) IsVideo() bool {
	return m.Type == MIME_VIDEO
}

// ToTransferItem Returns Transfer for FileItem
func (f *FileItem) ToTransferItem() *Payload_Item {
	return &Payload_Item{
		Data: &Payload_Item_File{
			File: f,
		},
		Type: Payload_Item_FILE,
	}
}

// NewTransferFileItem creates a new transfer file item
func NewTransferFileItem(path string) (*Payload_Item, error) {
	// Extracts File Infrom from path
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// Get MIME Type and Set Proto Enum
	mtype, err := mimetype.DetectFile(path)
	if err != nil {
		return nil, err
	}

	// Create File Item
	fileItem := &FileItem{
		Path: path,
		Size: int32(fi.Size()),
		Name: fi.Name(),
		Mime: &MIME{
			Type:    MIME_Type(MIME_Type_value[strings.ToUpper(mtype.Parent().String())]),
			Value:   mtype.String(),
			Subtype: mtype.Extension(),
		},
	}

	// // // Check if File is Image
	// if fileItem.Mime.IsImage() {
	// 	// Create Thumbnail
	// 	name := filepath.Base(path)
	// 	dir := filepath.Dir(path)
	// 	outPath := filepath.Join(dir, name+"_thumb")
	// 	err := thumbnail.NewImage(path, outPath)
	// 	if err != nil {
	// 		return nil, err
	// 	} else {
	// 		fileItem.Thumb = &Thumbnail{
	// 			Path: outPath,
	// 		}
	// 	}
	// }

	// Returns transfer item
	return &Payload_Item{
		Type: Payload_Item_FILE,
		Data: &Payload_Item_File{
			File: fileItem,
		},
	}, nil
}
