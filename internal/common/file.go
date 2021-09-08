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
func (f *FileItem) ToTransferItem() *Transfer_Item {
	return &Transfer_Item{
		Data: &Transfer_Item_File{
			File: f,
		},
		Type: Transfer_Item_FILE,
	}
}

// // NewThumbnail creates a new thumbnail with the given path and size
// func NewThumbnail(width, height int, path string) (*Thumbnail, error) {
// 	// decoder wants []byte, so read the whole file into a buffer
// 	inputBuf, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	decoder, err := lilliput.NewDecoder(inputBuf)
// 	// this error reflects very basic checks,
// 	// mostly just for the magic bytes of the file to match known image formats
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer decoder.Close()

// 	header, err := decoder.Header()
// 	// this error is much more comprehensive and reflects
// 	// format errors
// 	if err != nil {
// 		return nil, err
// 	}

// 	// print some basic info about the image
// 	fmt.Printf("file type: %s\n", decoder.Description())
// 	fmt.Printf("%dpx x %dpx\n", header.Width(), header.Height())

// 	ops := lilliput.NewImageOps(8192)
// 	defer ops.Close()

// 	// create a buffer to store the output image, 50MB in this case
// 	outputImg := make([]byte, 50*1024*1024)
// 	opts := &lilliput.ImageOptions{
// 		Width:                width,
// 		Height:               height,
// 		NormalizeOrientation: true,
// 	}

// 	// resize and transcode image
// 	buf, err := ops.Transform(decoder, opts, outputImg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Thumbnail{
// 		Buffer: buf,
// 	}, nil
// }

// NewTransferFileItem creates a new transfer file item
func NewTransferFileItem(path string) (*Transfer_Item, error) {
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

	// // Check if File is Image
	// if fileItem.Mime.IsImage() {
	// 	// Create Thumbnail
	// 	thumbnail, err := NewThumbnail(100, 100, path)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if thumbnail != nil {
	// 		fileItem.Thumb = thumbnail
	// 	}
	// }

	// Returns transfer item
	return &Transfer_Item{
		Type: Transfer_Item_FILE,
		Data: &Transfer_Item_File{
			File: fileItem,
		},
	}, nil
}
