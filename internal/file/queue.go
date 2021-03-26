package file

import (
	"errors"

	md "github.com/sonr-io/core/internal/models"
)

// ^ Adds File Transfer from Invite Request ^ //
func (fq *FileSystem) AddFromRequest(req *md.InviteRequest, p *md.Profile, done chan bool) error {
	// Add Single File Transfer
	safeFile, err := NewFileItem(req, p, done)
	if err != nil {
		return err
	}

	fq.Queue.Enqueue(safeFile)
	return nil
}

// ^ CurrentFile returns last file in Processed Files ^ //
func (fq *FileSystem) CurrentFile() (*FileItem, error) {
	if fq.Queue.IsNotEmpty() {
		return fq.Queue.Dequeue(), nil
	}
	return nil, errors.New("No File in Queue")
}
