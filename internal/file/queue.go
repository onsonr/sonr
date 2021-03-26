package file

import (
	"errors"
	"log"
	"path/filepath"

	"github.com/libp2p/go-libp2p-core/network"
	msg "github.com/libp2p/go-msgio"
	md "github.com/sonr-io/core/internal/models"
	dt "github.com/sonr-io/core/pkg/data"
)

const K_QUEUE_SIZE = 16

type FileQueue struct {
	incoming []*FileItem
	outgoing []*FileItem
}

// @ Adds Item to File Queue
func (fs *FileSystem) Enqueue(element *FileItem) {
	fs.Queue.outgoing = append(fs.Queue.outgoing, element) // Simply append to enqueue.
}

// ^ Adds File Transfer from Invite Request ^ //
func (fs *FileSystem) EnqueueFromRequest(req *md.InviteRequest, p *md.Peer, done chan bool) error {
	// Add Single File Transfer
	safeFile, err := NewOutgoingFileItem(req, p, done)
	if err != nil {
		return err
	}

	// Add to Queue
	fs.Queue.outgoing = append(fs.Queue.outgoing, safeFile) // Simply append to enqueue.
	return nil
}

// ^ Adds File Transfer from Invite Request ^ //
func (fs *FileSystem) EnqueueFromInvite(inv *md.AuthInvite, p *md.Peer, done chan bool) error {
	// Get File Path
	var path string

	// Check for Desktop
	if fs.IsDesktop {
		path = filepath.Join(fs.Downloads, inv.Card.Properties.Name)
	} else {
		// Check Load
		if inv.Payload == md.Payload_MEDIA {
			path = filepath.Join(fs.Temporary, inv.Card.Properties.Name)
		} else {
			path = filepath.Join(fs.Main, inv.Card.Properties.Name)
		}
	}

	// Add Single File Transfer
	safeFile, err := NewIncomingFileItem(inv, path)
	if err != nil {
		return err
	}

	// Add to Queue
	fs.Queue.incoming = append(fs.Queue.incoming, safeFile) // Simply append to enqueue.
	return nil
}

// @ Returns Queue Length
func (fs *FileSystem) CountIn() int {
	return len(fs.Queue.incoming)
}

// @ Returns Queue Length
func (fs *FileSystem) CountOut() int {
	return len(fs.Queue.outgoing)
}

// ^ DequeueOut returns last file in Processed Files ^ //
func (fs *FileSystem) DequeueIn() (*FileItem, error) {
	// Return FileItem
	if fs.HasIn() {
		file := fs.Queue.incoming[0]              // The first element is the one to be dequeued.
		fs.Queue.incoming = fs.Queue.incoming[1:] // Slice off the element once it is dequeued.
		return file, nil
	}
	return nil, errors.New("No File in Queue")
}

// ^ DequeueOut returns last file in Processed Files ^ //
func (fs *FileSystem) DequeueOut() (*FileItem, error) {
	// Return FileItem
	if fs.HasOut() {
		file := fs.Queue.outgoing[0]              // The first element is the one to be dequeued.
		fs.Queue.outgoing = fs.Queue.outgoing[1:] // Slice off the element once it is dequeued.
		return file, nil
	}
	return nil, errors.New("No File in Queue")
}

// ^ Checks if Queue has any elements ^
func (fs *FileSystem) HasIn() bool {
	return fs.CountIn() > 0
}

// ^ Checks if Queue has any elements ^
func (fs *FileSystem) HasOut() bool {
	return fs.CountOut() > 0
}

// ^ HandleIncomingStream Writes to Current Incoming File ^ //
func (fs *FileSystem) HandleIncomingStream(stream network.Stream) {
	// Get current incoming file
	inFile, err := fs.DequeueIn()
	if err != nil {
		log.Println(err)
		return
	}

	// Process Stream Events
	go func(reader msg.ReadCloser, f *FileItem) {
		for i := 0; ; i++ {
			// @ Read Length Fixed Bytes
			buffer, err := reader.ReadMsg()
			if err != nil {
				fs.Call.Error(err, "HandleIncoming:ReadMsg")
				break
			}

			// @ Unmarshal Bytes into Proto
			res, err := f.WriteFromStream(i, buffer)
			if err != nil {
				fs.Call.Error(err, "HandleIncoming:AddBuffer")
				break
			}

			// @ Callback with Progress
			if res.MetInterval {
				fs.Call.Progressed(res.Progress)
			}

			// @ Check if All Buffer Received to Save
			if res.HasCompleted {
				// Save File
				if err := fs.SaveFile(f); err != nil {
					fs.Call.Error(err, "HandleIncoming:Save")
				}
				break
			}
			dt.GetState().NeedsWait()
		}
	}(msg.NewReader(stream), inFile)
}
