package file

import (
	"errors"

	md "github.com/sonr-io/core/internal/models"
)

type FileQueue struct {
	queue []*FileItem
}

func NewFileQueue(maxSize int) *FileQueue {
	return &FileQueue{
		queue: make([]*FileItem, maxSize),
	}
}

func (fq *FileQueue) Enqueue(element *FileItem) {
	fq.queue = append(fq.queue, element) // Simply append to enqueue.
}

func (fq *FileQueue) Dequeue() *FileItem {
	file := fq.queue[0]     // The first element is the one to be dequeued.
	fq.queue = fq.queue[1:] // Slice off the element once it is dequeued.
	return file
}

func (fq *FileQueue) Count() int {
	return len(fq.queue)
}

func (fq *FileQueue) IsEmpty() bool {
	return len(fq.queue) == 0
}

func (fq *FileQueue) IsNotEmpty() bool {
	return len(fq.queue) > 0
}

// ^ Adds File Transfer from Invite Request ^ //
func (fq *FileSystem) AddFromRequest(req *md.InviteRequest, p *md.Profile) error {
	// Add Single File Transfer
	safeFile, err := NewFileItem(req, p, fq.Call)
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
