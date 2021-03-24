package file

import (
	"log"

	md "github.com/sonr-io/core/internal/models"
)

// ^ Adds File Transfer from Invite Request ^ //
func (fs *FileSystem) AddFromRequest(req *md.InviteRequest, pr *md.Profile) error {

	// Add Single File Transfer
	safeFile := NewProcessedFile(req, pr, fs.Call)

	// Add an item to the queue
	err := fs.Queue.Enqueue(safeFile)
	if err != nil {
		return err
	}
	return nil
}

// ^ CurrentFile returns last file in Processed Files ^ //
func (fs *FileSystem) CurrentFile() *ProcessedFile {
	// Initialize
	var iface interface{}
	var err error

	// Dequeue the next item in the queue and block until one is available
	if iface, err = fs.Queue.DequeueBlock(); err != nil {
		log.Fatal("Error dequeuing item ", err)
	}

	// Assert type of the response to an Item pointer so we can work with it
	item, ok := iface.(*ProcessedFile)
	if !ok {
		log.Fatal("Dequeued object is not an ProcessedFile pointer")
	}
	return item
}


