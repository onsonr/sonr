package transmit

import (
	"bytes"
	"container/list"
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/common"
)

// Session is a single entry in the Transmit queue.
type Session struct {
	ctx         context.Context
	direction   common.Direction
	from        *common.Peer
	to          *common.Peer
	payload     *common.Payload
	lastUpdated int64
	results     map[int32]bool
}

// Count returns the number of items in the payload.
func (s *Session) Count() int {
	return len(s.payload.GetItems())
}

// Event returns the CompleteEvent for the given session.
func (s *Session) Event() *api.CompleteEvent {
	return &api.CompleteEvent{
		From:       s.from,
		To:         s.to,
		Direction:  s.direction,
		Payload:    s.payload,
		CreatedAt:  s.payload.GetCreatedAt(),
		ReceivedAt: int64(time.Now().Unix()),
		Results:    s.results,
	}
}

// IndexAt returns the FileItem at the given index
func (s *Session) IndexAt(i int) *common.FileItem {
	return s.payload.GetItems()[i].GetFile()
}

// ReadItem Returns a new Reader for the given FileItem
func (s *Session) ReadItem(i int, n api.NodeImpl, str network.Stream, cchan chan itemResult) {
	// Initialize Properties
	reader := msgio.NewReader(str)
	fi := s.IndexAt(i)

	// Reset Item Path by OS FileSystem
	path, err := fi.ResetPath(fs.Downloads)
	if err != nil {
		logger.Errorf("Failed to Apply Reader: %s", err)
		return
	}

	// Create New Reader
	ir := &itemReader{
		item:         fi,
		index:        i,
		count:        s.Count(),
		size:         fi.GetSize(),
		node:         n,
		written:      0,
		progressChan: make(chan int),
		doneChan:     make(chan bool),
		interval:     calculateInterval(fi.GetSize()),
		buffer:       bytes.Buffer{},
		path:         path,
	}

	// Start Channels and Reader
	defer ir.Close()
	go handleRead(ir, reader)

	// Route Data from Stream
	for {
		select {
		case n := <-ir.progressChan:
			ir.written += n
			if ev := ir.getProgressEvent(); ev != nil {
				ir.node.OnProgress(ev)
			}
		case r := <-ir.doneChan:
			logger.Debug("Item Read is Complete")
			cchan <- ir.toResult(r)
			return
		}
	}
}

// WriteItem handles the writing of a FileItem to a Stream
func (s *Session) WriteItem(i int, n api.NodeImpl, str network.Stream, cchan chan itemResult) {
	// Initialize Properties
	writer := msgio.NewWriter(str)
	fi := s.IndexAt(i)

	// Create New File Chunker
	chunker, err := fs.NewFileChunker(fi.GetPath())
	if err != nil {
		logger.Errorf("%s - Failed to create new chunker.", err)
		return
	}

	// Create New Writer
	iw := &itemWriter{
		item:         fi,
		index:        i,
		count:        s.Count(),
		node:         n,
		size:         fi.GetSize(),
		written:      0,
		progressChan: make(chan int),
		doneChan:     make(chan bool),
		interval:     calculateInterval(fi.GetSize()),
		chunker:      chunker,
	}

	// Start Channels and Writer
	defer iw.Close()
	go handleWrite(iw, writer)

	// Await Progress and Result
	for {
		select {
		case n := <-iw.progressChan:
			iw.written += n
			if ev := iw.getProgressEvent(); ev != nil {
				iw.node.OnProgress(ev)
			}
		case r := <-iw.doneChan:
			logger.Debug("Item Write is Complete")
			cchan <- iw.toResult(r)
			return
		}
	}
}

// SessionQueue is a queue for incoming and outgoing requests.
type SessionQueue struct {
	ctx   context.Context
	host  *host.SNRHost
	queue *list.List
}

// AddIncoming adds Incoming Request to Transfer Queue
func (sq *SessionQueue) AddIncoming(from peer.ID, req *InviteRequest) error {
	// Authenticate Message
	valid := sq.host.AuthenticateMessage(req, req.Metadata)
	if !valid {
		return ErrFailedAuth
	}

	// Create New TransferEntry
	entry := Session{
		direction:   common.Direction_INCOMING,
		payload:     req.GetPayload(),
		from:        req.GetFrom(),
		to:          req.GetTo(),
		lastUpdated: int64(time.Now().Unix()),
		results:     make(map[int32]bool),
		ctx:         sq.ctx,
	}

	// Add to Requests
	sq.queue.PushBack(&entry)
	return nil
}

// AddOutgoing adds Outgoing Request to Transfer Queue
func (sq *SessionQueue) AddOutgoing(to peer.ID, req *InviteRequest) error {
	// Create New TransferEntry
	entry := Session{
		direction:   common.Direction_OUTGOING,
		payload:     req.GetPayload(),
		from:        req.GetFrom(),
		to:          req.GetTo(),
		lastUpdated: int64(time.Now().Unix()),
		results:     make(map[int32]bool),
		ctx:         sq.ctx,
	}

	// Add to Requests
	sq.queue.PushBack(&entry)
	return nil
}

// Next returns topmost entry in the queue.
func (sq *SessionQueue) Next() (*Session, error) {
	// Find Entry for Peer
	entry := sq.queue.Remove(sq.queue.Front()).(*Session)
	entry.lastUpdated = int64(time.Now().Unix())
	return entry, nil
}

// Validate takes list of Requests and returns true if Request exists in List and UUID is verified.
// Method also returns the InviteRequest that points to the Response.
func (sq *SessionQueue) Validate(resp *InviteResponse) (*Session, error) {
	// Authenticate Message
	valid := sq.host.AuthenticateMessage(resp, resp.Metadata)
	if !valid {
		return nil, ErrFailedAuth
	}

	// Check Decision
	if !resp.GetDecision() {
		return nil, nil
	}

	// Check if the request is valid
	if sq.queue.Len() == 0 {
		return nil, ErrEmptyRequests
	}

	// Get Next Entry
	entry, err := sq.Next()
	if err != nil {
		logger.Errorf("%s - Failed to get Transmit entry", err)
		return nil, err
	}

	entry.lastUpdated = int64(time.Now().Unix())
	return entry, nil
}
