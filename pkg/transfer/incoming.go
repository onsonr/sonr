package transfer

import (
	"bytes"
	"io/ioutil"

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
	emitter *state.Emitter
	item    *common.FileItem
	buffer  bytes.Buffer
	path    string
	index   int
	count   int
	size    int64
}

// NewReader Returns a new Reader for the given FileItem
func NewReader(index int, count int, item *common.Payload_Item) (*itemReader, error) {
	// Return Reader
	return &itemReader{
		item:    item.GetFile(),
		size:    item.GetSize(),
		emitter: state.NewEmitter(2048),
		index:   index,
		count:   count,
		path:    device.NewDownloadsPath(item.GetFile().GetPath()),
		buffer:  bytes.Buffer{},
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
func (ir *itemReader) ReadFrom(reader msgio.ReadCloser) {
	// Defer Closing of Reader and WaitGroup
	defer reader.Close()

	// Route Data from Stream
	for i := 0; i < int(ir.size); {
		// Read Length Fixed Bytes
		buffer, err := reader.ReadMsg()
		if err != nil {
			logger.Error("Failed to Read Next Message on Read Stream", err)
			return
		}

		// Decode Chunk
		buf, err := decodeChunk(buffer)
		if err != nil {
			logger.Error("Failed to Decode Chunk on Read Stream", err)
			return
		}

		n, err := ir.buffer.Write(buf.Data)
		if err != nil {
			logger.Error("Failed to Write Buffer to File on Read Stream", err)
			return
		}
		i += n

		// Emit Progress
		if (i % ITEM_INTERVAL) == 0 {
			ir.Progress(i)
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
