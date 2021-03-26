package file

import (
	"errors"
	"sync"

	md "github.com/sonr-io/core/internal/models"
)

const K_QUEUE_SIZE = 16

type FileQueue struct {
	mutex    sync.Mutex
	elements []*FileItem
}

// @ Adds Item to File Queue
func (fs *FileSystem) Enqueue(element *FileItem) {
	fs.Queue.mutex.Lock()
	fs.Queue.elements = append(fs.Queue.elements, element) // Simply append to enqueue.
	fs.Queue.mutex.Unlock()
}

// ^ Adds File Transfer from Invite Request ^ //
func (fs *FileSystem) EnqueueFromRequest(req *md.InviteRequest, p *md.Profile, done chan bool) error {
	// Handle Mutex
	fs.Queue.mutex.Lock()
	defer fs.Queue.mutex.Unlock()

	// Add Single File Transfer
	safeFile, err := NewOutgoingFileItem(req, p, done)
	if err != nil {
		return err
	}

	// Add to Queue
	fs.Queue.elements = append(fs.Queue.elements, safeFile) // Simply append to enqueue.
	return nil
}

// @ Returns Queue Length
func (fq *FileQueue) Count() int {
	return len(fq.elements)
}

// ^ Dequeue returns last file in Processed Files ^ //
func (fs *FileSystem) Dequeue() (*FileItem, error) {
	// Handle Mutex
	fs.Queue.mutex.Lock()
	defer fs.Queue.mutex.Unlock()

	// Return FileItem
	if fs.IsNotEmpty() {
		file := fs.Queue.elements[0]              // The first element is the one to be dequeued.
		fs.Queue.elements = fs.Queue.elements[1:] // Slice off the element once it is dequeued.
		return file, nil
	}
	return nil, errors.New("No File in Queue")
}

// @ Checks if Queue does not have any elements
func (fs *FileSystem) IsEmpty() bool {
	return len(fs.Queue.elements) == 0
}

// @ Checks if Queue has any elements
func (fs *FileSystem) IsNotEmpty() bool {
	return len(fs.Queue.elements) > 0
}
