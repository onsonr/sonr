package file

import (
	md "github.com/sonr-io/core/internal/models"
)

// ^ Adds File Transfer from Invite Request ^ //
func (fq *FileSystem) AddFromRequest(req *md.InviteRequest, p *md.Profile) error {
	// Add Single File Transfer
	safeFile, err := NewFileItem(req, p, fq.Call)
	if err != nil {
		return err
	}

	fq.Files = append(fq.Files, safeFile)
	fq.CurrentCount = 1
	return nil
}

// ^ CurrentFile returns last file in Processed Files ^ //
func (fq *FileSystem) CurrentFile() *FileItem {
	if len(fq.Files) > 0 {
		return fq.Files[len(fq.Files)-1]
	} else {
		return nil
	}
}

// ^ Removes Last File ^ //
func (fq *FileSystem) CompleteLast() {
	if len(fq.Files) > 0 {
		fq.Files = fq.Files[:len(fq.Files)-1]
	}
	fq.CurrentCount = 0
}

// ^ Reset Current Queued File Metadata ^ //
func (fq *FileSystem) Reset() {
	fq.Files = nil
	fq.Files = make([]*FileItem, K_QUEUE_SIZE)
	fq.CurrentCount = 0
}
