package transfer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	sync "sync"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/tools/config"
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
func (p *TransferProtocol) onIncomingTransfer(s network.Stream) {
	// Find Entry in Queue
	e, err := p.sessionQueue.Next()
	if err != nil {
		logger.Error("Failed to find transfer request", err)
		s.Reset()
		return
	}

	// Initialize Params
	logger.Info("Started Incoming Transfer...")
	reader := msgio.NewReader(s)
	wg := sync.WaitGroup{}
	wg.Add(1)
	// Handle incoming stream
	go func(entry *Session, rs msgio.ReadCloser) {

		// Write All Files
		for i, v := range entry.Items() {
			// Create Reader
			r, err := NewReader(i, entry.Count(), v, p.emitter)
			if err != nil {
				logger.Error("Failed to create reader", err)
				rs.Close()
				return
			}

			// Write to File
			if err := r.ReadFrom(rs); err != nil {
				logger.Error("Failed to write to file", err)
				rs.Close()
				return
			}

			// Update Progress
			logger.Info(fmt.Sprintf("Finished RECEIVING File (%v/%v)", i+1, entry.Count()))
		}

		wg.Done()
	}(e, reader)
	wg.Wait()

	// Complete the transfer
	event, err := p.sessionQueue.Done()
	if err != nil {
		logger.Error("Failed to Complete Incoming Transfer", err)
		return
	}

	// Emit Event
	p.emitter.Emit(Event_COMPLETED, event)

	// Await WaitGroup
	err = reader.Close()
	if err != nil {
		logger.Error("Failed to close stream for incoming transfer", err)
		return
	}
}

// itemReader is a Reader for a FileItem
type itemReader struct {
	emitter *state.Emitter
	mutex   sync.Mutex
	logger  *golog.Logger
	item    *common.FileItem
	path    string
	index   int
	count   int
	size    int64
}

// NewReader Returns a new Reader for the given FileItem
func NewReader(index int, count int, item *common.Payload_Item, em *state.Emitter) (*itemReader, error) {
	// Determine Path for File
	path, err := device.NewDownloadsPath(item.GetFile().GetPath())
	if err != nil {
		logger.Error("Failed to determine downloads path", err)
		return nil, err
	}

	// Return Reader
	return &itemReader{
		item:    item.GetFile(),
		size:    item.GetSize(),
		logger:  logger,
		emitter: em,
		index:   index,
		count:   count,
		path:    path,
	}, nil
}

// Returns Progress of File, Given the written number of bytes
func (p *itemReader) Progress(i int) {
	// Create Progress Event
	event := &api.ProgressEvent{
		Progress: (float64(i) / float64(p.size)),
		Current:  int32(p.index),
		Total:    int32(p.count),
	}

	// Push ProgressEvent to Emitter
	p.emitter.Emit(Event_PROGRESS, event)
}

// ReadFrom Reads from the given Reader and writes to the given File
func (ir *itemReader) ReadFrom(reader msgio.ReadCloser) error {
	// Create File Buffer
	fileBuffer := bytes.Buffer{}

	// Route Data from Stream
	for i := 0; i < int(ir.size); {
		// Read Length Fixed Bytes
		buffer, err := reader.ReadMsg()
		if err != nil {
			ir.logger.Error("Failed to Read Next Message on Read Stream", err)
			return err
		}

		// Decode Chunk
		buf, err := decodeChunk(buffer)
		if err != nil {
			ir.logger.Error("Failed to Decode Chunk on Read Stream", err)
			return err
		}

		// Write to File, and Update Progress
		ir.mutex.Lock()

		n, err := fileBuffer.Write(buf.Data)
		if err != nil {
			ir.logger.Error("Failed to Write Buffer to File on Read Stream", err)
			return err
		}
		i += n
		ir.mutex.Unlock()

		// Emit Progress
		if (i % ITEM_INTERVAL) == 0 {
			ir.Progress(i)
		}
	}

	// Write File Buffer to File
	err := ioutil.WriteFile(ir.path, fileBuffer.Bytes(), 0644)
	if err != nil {
		ir.logger.Error("Failed to Close item on Read Stream", err)
		return err
	}
	ir.logger.Info("Completed writing to file: " + ir.path)
	return nil
}

// decodeChunk Decodes a Chunk from a Message
func decodeChunk(buf []byte) (config.Chunk, error) {
	// Decode Chunk
	chunk := &Chunk{}
	err := proto.Unmarshal(buf, chunk)
	if err != nil {
		logger.Error("Failed to Unmarshal Chunk.", err)
		return config.Chunk{}, err
	}

	// Convert to Chunk
	return config.Chunk{
		Offset:      int(chunk.Offset),
		Length:      int(chunk.Length),
		Data:        chunk.Data,
		Fingerprint: uint64(chunk.Fingerprint),
	}, nil
}
