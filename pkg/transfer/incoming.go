package transfer

import (
	"fmt"
	"os"
	sync "sync"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/tools/config"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"go.uber.org/zap"
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

	// generate response message
	p.queue.AddIncoming(remotePeer, req)

	// store request data into Context
	p.emitter.Emit(Event_INVITED, req.ToEvent())
}

// onIncomingTransfer incoming transfer handler
func (p *TransferProtocol) onIncomingTransfer(s network.Stream) {
	// Find Entry in Queue
	e, err := p.queue.Next()
	if err != nil {
		logger.Error("Failed to find transfer request", err)
		return
	}

	// Initialize Params
	logger.Info("Started Incoming Transfer...")
	waitGroup := sync.WaitGroup{}
	reader := msgio.NewReader(s)

	// Handle incoming stream
	go func(entry *Session, wg sync.WaitGroup, rs msgio.ReadCloser) {
		// Write All Files
		err = entry.MapItems(func(m *common.Payload_Item, i int, count int) error {
			// Add to WaitGroup
			logger.Info("Current Item: ", zap.String(fmt.Sprint(i), m.String()))
			wg.Add(1)

			// Create New Reader
			r, err := NewReader(m, i, count, p.emitter)
			if err != nil {
				return logger.Error("Error creating reader", err)
			}

			// Read From Stream
			err = r.ReadFrom(rs)
			if err != nil {
				return logger.Error("Failed to Read from Stream and Write to File.", err)
			}

			// Complete Writing
			logger.Info(fmt.Sprintf("Finished RECEIVING File (%v/%v)", i+1, count))
			wg.Done()
			return nil
		})
		if err != nil {
			logger.Error("Error writing stream", err)
			return
		}

		// Await WaitGroup
		waitGroup.Wait()
		reader.Close()

		// Complete the transfer
		event, err := p.queue.Done()
		if err != nil {
			logger.Error("Failed to Complete Incoming Transfer", err)
			return
		}

		// Emit Event
		p.emitter.Emit(Event_COMPLETED, event)
	}(e, waitGroup, reader)
}

// itemReader is a Reader for a FileItem
type itemReader struct {
	emitter *state.Emitter
	mutex   sync.Mutex
	item    *common.FileItem
	path    string
	index   int
	count   int
	size    int64
}

// NewReader Returns a new Reader for the given FileItem
func NewReader(pi *common.Payload_Item, index int, count int, em *state.Emitter) (*itemReader, error) {
	// Determine Path for File
	path, err := device.NewDownloadsPath(pi.GetFile().GetPath())
	if err != nil {
		return nil, logger.Error("Failed to determine downloads path", err)
	}

	// Return Reader
	return &itemReader{
		item:    pi.GetFile(),
		size:    pi.GetSize(),
		emitter: em,
		index:   index,
		count:   count,
		path:    path,
	}, nil
}

// Progress Returns Progress of File, Given the written number of bytes
func (p *itemReader) Progress(i int) {
	// Create Progress Event
	event := &common.ProgressEvent{
		Progress: (float64(i) / float64(p.size)),
		Current:  int32(p.index),
		Total:    int32(p.count),
	}

	// Push ProgressEvent to Emitter
	p.emitter.Emit(Event_PROGRESS, event)
}

// ReadFrom Reads from the given Reader and writes to the given File
func (ir *itemReader) ReadFrom(reader msgio.ReadCloser) error {
	// Return Created File
	f, err := os.Create(ir.path)
	if err != nil {
		return err
	}
	logger.Info("Created new file at path", zap.String("incoming.ItemPath", ir.path))
	defer f.Close()

	// Route Data from Stream
	for i := 0; i < int(ir.size); {
		// Read Length Fixed Bytes
		buffer, err := reader.ReadMsg()
		if err != nil {
			return logger.Error("Failed to Read Next Message on Read Stream", err)
		}

		// Decode Chunk
		buf, err := decodeChunk(buffer)
		if err != nil {
			return logger.Error("Failed to Decode Chunk on Read Stream", err)
		}

		// Write to File, and Update Progress
		ir.mutex.Lock()
		n, err := f.Write(buf.Data)
		if err != nil {
			return logger.Error("Failed to Write Buffer to File on Read Stream", err)
		}
		i += n
		ir.mutex.Unlock()

		// Emit Progress
		if (i % ITEM_INTERVAL) == 0 {
			ir.Progress(i)
		}
	}

	// Flush File Contents
	err = f.Sync()
	if err != nil {
		return logger.Error("Failed to Sync item on Read Stream", err)
	}

	// Close File
	err = f.Close()
	if err != nil {
		return logger.Error("Failed to Close item on Read Stream", err)
	}

	// Close Reader
	err = reader.Close()
	if err != nil {
		return logger.Error("Failed to Close Reader for Incoming Stream", err)
	}
	logger.Info("Completed writing to file.")
	return nil
}

// decodeChunk Decodes a Chunk from a Message
func decodeChunk(buf []byte) (config.Chunk, error) {
	// Decode Chunk
	chunk := &Chunk{}
	err := proto.Unmarshal(buf, chunk)
	if err != nil {
		return config.Chunk{}, logger.Error("Failed to Unmarshal Chunk.", err)
	}

	// Convert to Chunk
	return config.Chunk{
		Offset:      int(chunk.Offset),
		Length:      int(chunk.Length),
		Data:        chunk.Data,
		Fingerprint: uint64(chunk.Fingerprint),
	}, nil
}
