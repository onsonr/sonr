package transfer

import (
	"bytes"
	"context"
	"io/ioutil"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/tools/state"
	"google.golang.org/protobuf/proto"
)

// onInviteRequest peer requests handler
func (p *TransferProtocol) onInviteRequest(s network.Stream) {
	// Initialize Stream Data
	remotePeer := s.Conn().RemotePeer()
	r := msgio.NewReader(s)

	// Read the request
	buf, err := r.ReadMsg()
	if err != nil {
		s.Reset()
		logger.Error("Failed to Read Invite Request buffer.", err)
		return
	}
	s.Close()

	// unmarshal it
	req := &InviteRequest{}
	err = proto.Unmarshal(buf, req)
	if err != nil {
		logger.Error("Failed to Unmarshal Invite REQUEST buffer.", err)
		return
	}
	// store request data into Context
	p.emitter.Emit(Event_INVITED, req.ToEvent())

	// generate response message
	p.sessionQueue.AddIncoming(remotePeer, req)
}

// onIncomingTransfer incoming transfer handler
func (p *TransferProtocol) onIncomingTransfer(stream network.Stream) {
	// Find Entry in Queue
	entry, err := p.sessionQueue.Next()
	if err != nil {
		logger.Error("Failed to find transfer request", err)
		stream.Reset()
		return
	}

	// Create New Writer
	if event := entry.ReadFrom(stream); event != nil {
		p.emitter.Emit(Event_COMPLETED, event)
	}
}

// itemReader is a Reader for a FileItem
type itemReader struct {
	controller state.HandController
	emitter    *state.Emitter
	item       *common.FileItem
	buffer     bytes.Buffer
	path       string
	index      int
	count      int
	size       int64
}

// NewReader Returns a new Reader for the given FileItem
func NewReader(index int, count int, item *common.Payload_Item) *itemReader {
	return &itemReader{
		item:       item.GetFile(),
		size:       item.GetSize(),
		emitter:    state.NewEmitter(2048),
		index:      index,
		count:      count,
		path:       device.NewDownloadsPath(item.GetFile().GetPath()),
		buffer:     bytes.Buffer{},
		controller: state.NewHands(),
	}
}

// Returns Progress of File, Given the written number of bytes
func (p *itemReader) Progress(i int, n int) {
	i += n
	if (i % ITEM_INTERVAL) == 0 {
		// Create Progress Event
		event := &api.ProgressEvent{
			Progress: (float64(i) / float64(p.size)),
			Current:  int32(p.index),
			Total:    int32(p.count),
		}

		// Push ProgressEvent to Emitter
		p.emitter.Emit(Event_PROGRESS, event)
	}
}

// ReadFrom Reads from the given Reader and writes to the given File
func (ir *itemReader) ReadFrom(reader msgio.ReadCloser) {
	// Defer Closing of Reader and WaitGroup
	defer reader.Close()

	// Route Data from Stream
	for i := 0; i < int(ir.size); {
		var buf []byte
		// Function to Parse Chunk
		ir.controller.Do(func(ctx context.Context) error {
			// Read Length Fixed Bytes
			var err error
			buf, err = reader.ReadMsg()
			if err != nil {
				logger.Error("Failed to Read Next Message on Read Stream", err)
				return err
			}
			return nil
		}, state.P(1))

		// Function to Parse Chunk
		ir.controller.Do(func(ctx context.Context) error {
			// Write Chunk to File
			n, err := ir.buffer.Write(buf)
			if err != nil {
				logger.Error("Failed to Write Buffer to File on Read Stream", err)
				return err
			}
			ir.Progress(i, n)
			return nil
		}, state.P(2))

		// Run Handlers
		if err := ir.controller.Run(); err != nil {
			logger.Error("Failed to Run Handlers on Read Stream", err)
			return
		}
	}

	// Write File Buffer to File
	err := ioutil.WriteFile(ir.path, ir.buffer.Bytes(), 0644)
	if err != nil {
		logger.Error("Failed to Close item on Read Stream", err)
		return
	}
	logger.Info("Completed writing to file: " + ir.path)
	return
}
