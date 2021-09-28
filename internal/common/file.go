package common

import (
	"os"
)

// IsFile Checks if Path is Valid File
func IsFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

// ToTransferItem Returns Transfer for FileItem
func (f *FileItem) ToTransferItem() *Payload_Item {
	return &Payload_Item{
		Mime:     f.GetMime(),
		Data: &Payload_Item_File{
			File: f,
		},
	}
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

	// Returns transfer item
	return &Payload_Item{
		Size:     fi.Size(),
		Mime:     mime,
		Data: &Payload_Item_File{
			File: fileItem,
		},
	}, nil
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
