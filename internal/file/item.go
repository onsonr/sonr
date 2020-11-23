package file

import (
	"errors"
	"time"

	pb "github.com/sonr-io/core/internal/models"
)

// Define Function Types
type OnQueued func(data []byte)
type OnProgress func(data []byte)
type OnCompleted func(data []byte)
type OnError func(err error, method string)

// Struct to Implement Node Callback Methods
type FileCallback struct {
	Queued    OnQueued
	Progress  OnProgress
	Completed OnCompleted
	Error     OnError
}

// Define Block Size for Chunking Bytes on Stream
const BlockSize = 16000

// ^ Combo Struct that contains Item and Transfer ^ //
type Item struct {
	IsReady  bool
	HasMeta  bool
	Path     string
	Call     FileCallback
	safeMeta *SafeMeta
	transfer *TransferFile
}

func (i *Item) New() error {
	// Check if Path Provided
	if i.Path == "" {
		err := errors.New("Path wasnt provided")
		return err
	}

	// Check if CallBack Provided
	callPointer := &i.Call
	if callPointer == nil {
		err := errors.New("Callback wasnt provided")
		return err
	}

	// Set Default Properties
	i.IsReady = false
	i.HasMeta = false

	go i.safeMeta.Generate(i)
	return nil
}

func (i *Item) completedMetadata(sf *SafeMeta) {
	// Call Queued
	i.HasMeta = true
	i.Call.Queued(i.safeMeta.Bytes())

	// Begin Transfer Generation
	go i.transfer.Generate(i)
}

func (i *Item) completedTransfer() {
	i.IsReady = true
}

func (i *Item) GetBlocks() []*pb.Block {
	time.Sleep(time.Millisecond * 500)
	return i.transfer.Blocks()
}

func (i *Item) GetMetadata() *pb.Metadata {
	time.Sleep(time.Millisecond * 500)
	return i.safeMeta.Metadata()
}
