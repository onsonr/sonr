package file

import (
	"errors"

	md "github.com/sonr-io/core/internal/models"
)

// // ^ Adds File Transfer from Invite Request ^ //
// func (fs *FileSystem) AddFromRequest(req *md.InviteRequest, pr *md.Profile) error {

// 	// Add Single File Transfer
// 	safeFile := NewProcessedFile(req, pr, fs.Call)

// 	// Add an item to the queue
// 	err := fs.Queue.Enqueue(safeFile)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // ^ CurrentFile returns last file in Processed Files ^ //
// func (fs *FileSystem) CurrentFile() *ProcessedFile {
// 	// Initialize
// 	var iface interface{}
// 	var err error

// 	// Dequeue the next item in the queue and block until one is available
// 	if iface, err = fs.Queue.DequeueBlock(); err != nil {
// 		log.Fatal("Error dequeuing item ", err)
// 	}

// 	// Assert type of the response to an Item pointer so we can work with it
// 	item, ok := iface.(*ProcessedFile)
// 	if !ok {
// 		log.Fatal("Dequeued object is not an ProcessedFile pointer")
// 	}
// 	return item
// }

// ^ Adds File Transfer from Invite Request ^ //
func (fq *FileSystem) AddFromRequest(req *md.InviteRequest, p *md.Profile) error {
	// Add Single File Transfer
	safeFile := NewProcessedFile(req, p, fq.Call)

	// Validate Files not Null
	if safeFile == nil {
		return errors.New("Request or Profile not Provided")
	}

	fq.Files = append(fq.Files, safeFile)
	fq.CurrentCount = 1
	return nil
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
	fq.Files = make([]*ProcessedFile, K_QUEUE_SIZE)
	fq.CurrentCount = 0
}
