package data

import (
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/pkg/models"
)

// @ Maximum Files in Node Cache
const maxFileBufferSize = 64

// ^ Adds File Transfer from Invite Request ^ //
func (fq *SonrFS) AddFromRequest(req *md.InviteRequest) {
	if req.Type == md.InviteRequest_File {
		// Add Single File Transfer
		safeFile := sf.NewProcessedFile(req, fq.profile, fq.call)
		fq.files = append(fq.files, safeFile)
		fq.currentCount = 1
	} else if req.Type == md.InviteRequest_MultiFiles {
		// Add Batch File Transfer
		safeFiles := sf.NewBatchProcessFiles(req, fq.profile, fq.call)
		fq.files = append(fq.files, safeFiles...)
		fq.currentCount = len(safeFiles)
	}
}

// ^ CurrentFile returns last file in Processed Files ^ //
func (fq *SonrFS) CurrentFile() *sf.ProcessedFile {
	return fq.files[len(fq.files)-1]
}

// ^ Removes Last File ^ //
func (fq *SonrFS) CompleteLast() {
	if len(fq.files) > 0 {
		fq.files = fq.files[:len(fq.files)-1]
	}
	fq.currentCount = 0
}

// ^ Reset Current Queued File Metadata ^ //
func (fq *SonrFS) Reset() {
	fq.files = nil
	fq.files = make([]*sf.ProcessedFile, maxFileBufferSize)
	fq.currentCount = 0
}
