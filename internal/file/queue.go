package file

import (
	"errors"

	md "github.com/sonr-io/core/internal/models"
)

const K_QUEUE_SIZE = 16

// ^ Adds File Transfer from Invite Request ^ //
func (fs *FileSystem) EnqueueFromRequest(req *md.InviteRequest, p *md.Peer) error {
	// Add Single File Transfer
	safeFile := NewOutgoingFileItem(req, p, fs.Call)

	// Validate Files not Null
	if safeFile == nil {
		return errors.New("Request or Profile not Provided")
	}

	fs.Files = append(fs.Files, safeFile)
	fs.CurrentCount = 1
	return nil
}

// ^ CurrentFile returns last file in Processed Files ^ //
func (fs *FileSystem) CurrentFile() (*FileItem, error) {
	if len(fs.Files) > 0 {
		return fs.Files[len(fs.Files)-1], nil
	} else {
		return nil, errors.New("File does not exist")
	}
}

// ^ Removes Last File ^ //
func (fs *FileSystem) CompleteLast() {
	if len(fs.Files) > 0 {
		fs.Files = fs.Files[:len(fs.Files)-1]
	}
	fs.CurrentCount = 0
}

// ^ Reset Current Queued File Metadata ^ //
func (fs *FileSystem) Reset() {
	fs.Files = nil
	fs.Files = make([]*FileItem, K_QUEUE_SIZE)
	fs.CurrentCount = 0
}
