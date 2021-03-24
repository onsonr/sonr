package user

import (
	"os"
	"path/filepath"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
)

// @ Maximum Files in Node Cache
const maxFileBufferSize = 64

// ^ Adds File Transfer from Invite Request ^ //
func (fq *FileSystem) AddFromRequest(req *md.InviteRequest) {

	// Add Single File Transfer
	safeFile := sf.NewProcessedFile(req, fq.Profile, fq.Call)

	// Validate Files not Null
	if safeFile != nil {
		fq.Files = append(fq.Files, safeFile)
		fq.CurrentCount = 1
	} else {
		fq.Call.Error(errors.New("Request or Profile not Provided"), "NewProcessedFile:GetFileInfo")
	}
}

// ^ CurrentFile returns last file in Processed Files ^ //
func (fq *FileSystem) CurrentFile() *sf.ProcessedFile {
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
	fq.Files = make([]*sf.ProcessedFile, maxFileBufferSize)
	fq.CurrentCount = 0
}

// ^ Write User Data at Path ^
func (sfs *FileSystem) WriteIncomingFile(load md.Payload, props *md.TransferCard_Properties, data []byte) (string, string) {
	// Create File Name
	fileName := props.Name + "." + props.Mime.Subtype
	path := sfs.getIncomingFilePath(load, fileName)

	// Check for User File at Path
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		sentry.CaptureException(err)
	}

	// Defer Close
	defer file.Close()

	// Write User Data to File
	_, err = file.Write(data)
	if err != nil {
		sentry.CaptureException(err)
	}
	return fileName, path
}

// @ Helper: Finds Write Path for Incoming File
func (sfs *FileSystem) getIncomingFilePath(load md.Payload, fileName string) string {
	// Check for Desktop
	if sfs.IsDesktop {
		return filepath.Join(sfs.Downloads, fileName)
	} else {
		// Check Load
		if load == md.Payload_MEDIA {
			return filepath.Join(sfs.Temporary, fileName)
		} else {
			return filepath.Join(sfs.Root, fileName)
		}
	}
}
