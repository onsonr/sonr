package file

import (
	"errors"

	md "github.com/sonr-io/core/internal/models"
)

const K_BUF_CHUNK = 32000
const K_B64_CHUNK = 31998 // Adjusted for Base64 -- has to be divisible by 3

const K_QUEUE_SIZE = 16

type FileQueue struct {
	outgoing []*FileItem
}

// ^ Adds File Transfer from Invite Request ^ //
func (fq *FileSystem) AddFromRequest(req *md.InviteRequest, p *md.Profile, done chan bool) error {
	// Add Single File Transfer
	safeFile, err := NewOutgoingFileItem(req, p, done)
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

// @ Adds Item to File Queue
func (fq *FileQueue) Enqueue(element *FileItem) {
	fq.outgoing = append(fq.outgoing, element) // Simply append to enqueue.
}

// @ Pops Item from File Queue
func (fq *FileQueue) Dequeue() *FileItem {
	file := fq.outgoing[0]        // The first element is the one to be dequeued.
	fq.outgoing = fq.outgoing[1:] // Slice off the element once it is dequeued.
	return file
}

// @ Returns Queue Length
func (fq *FileQueue) Count() int {
	return len(fq.outgoing)
}

// @ Checks if Queue does not have any elements
func (fq *FileQueue) IsEmpty() bool {
	return len(fq.outgoing) == 0
}

// @ Checks if Queue has any elements
func (fq *FileQueue) IsNotEmpty() bool {
	return len(fq.outgoing) > 0
}
