package file

import (
	"log"

	md "github.com/sonr-io/core/internal/models"
	dt "github.com/sonr-io/core/pkg/data"
)

// ^ Adds File Transfer from Invite Request ^ //
func (fs *FileSystem) AddFromRequest(req *md.InviteRequest, p *md.Peer) error {

	// Add Single File Transfer
	safeFile := dt.NewProcessedFile(req, p, fs.Call)

	// Add an item to the queue
	err := fs.Queue.Enqueue(safeFile)
	if err != nil {
		return err
	}
	return nil
}

// ^ CurrentFile returns last file in Processed Files ^ //
func (fs *FileSystem) CurrentFile() *dt.ProcessedFile {
	// Initialize
	var iface interface{}
	var err error

	// Dequeue the next item in the queue and block until one is available
	if iface, err = fs.Queue.DequeueBlock(); err != nil {
		log.Fatal("Error dequeuing item ", err)
	}

	// Assert type of the response to an Item pointer so we can work with it
	item, ok := iface.(*dt.ProcessedFile)
	if !ok {
		log.Fatal("Dequeued object is not an ProcessedFile pointer")
	}
	return item
}

// ^ Reset Current Queued File Metadata ^ //
func (fs *FileSystem) Close() {
	fs.Queue.Close()
}
