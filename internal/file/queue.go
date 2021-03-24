package file

import (
	md "github.com/sonr-io/core/internal/models"
)

// @ Maximum Files in Node Cache
const maxFileBufferSize = 64

// ^ Adds File Transfer from Invite Request ^ //
func (fq *FileSystem) AddFromRequest(req *md.InviteRequest) {

	// // Add Single File Transfer
	// safeFile := sf.NewProcessedFile(req, fq.Profile, fq.Call)

	// // Validate Files not Null
	// if safeFile != nil {
	// 	fq.Files = append(fq.Files, safeFile)
	// 	fq.CurrentCount = 1
	// } else {
	// 	fq.Call.Error(errors.New("Request or Profile not Provided"), "NewProcessedFile:GetFileInfo")
	// }
}

// ^ CurrentFile returns last file in Processed Files ^ //
func (fq *FileSystem) CurrentFile() *ProcessedFile {
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
	fq.Files = make([]*ProcessedFile, maxFileBufferSize)
	fq.CurrentCount = 0
}
