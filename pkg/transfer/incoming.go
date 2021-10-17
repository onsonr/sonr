package transfer

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
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
	p.node.OnInvite(req.ToEvent())

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
	if event := entry.ReadFrom(stream, p.node); event != nil {
		p.node.OnComplete(event)
	}
}

// itemReader is a Reader for a FileItem
type itemReader struct {
	item   *common.FileItem
	buffer bytes.Buffer
	path   string
	index  int
	count  int
	size   int64
	node   api.NodeImpl
}

// NewItemReader Returns a new Reader for the given FileItem
func NewItemReader(index int, count int, item *common.Payload_Item, node api.NodeImpl) *itemReader {
	path, _ := device.NewDownloadsPath(item.GetFile().GetPath())
	return &itemReader{
		item:   item.GetFile(),
		size:   item.GetSize(),
		index:  index,
		count:  count,
		path:   path,
		buffer: bytes.Buffer{},
		node:   node,
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
		p.node.OnProgress(event)
	}
}

// ReadFrom Reads from the given Reader and writes to the given File
func (ir *itemReader) ReadFrom(reader msgio.ReadCloser) {
	// Defer Closing of Reader and WaitGroup
	defer reader.Close()

	// Route Data from Stream
	for {
		// Read Next Message
		buf, err := reader.ReadMsg()
		if err == io.EOF {
			break
		} else if err != nil {
			logger.Error("Failed to Read Next Message on Read Stream", err)
			return
		} else {
			// Write Chunk to File
			_, err := ir.buffer.Write(buf)
			if err != nil {
				logger.Error("Failed to Write Buffer to File on Read Stream", err)
				return
			}
			// ir.Progress(i, n)
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
