package transmit

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
	"github.com/sonr-io/core/pkg/common"
	"google.golang.org/protobuf/proto"
)

// onInviteRequest peer requests handler
func (p *TransmitProtocol) onInviteRequest(s network.Stream) {
	logger.Debug("Received Invite Request")
	// Initialize Stream Data
	remotePeer := s.Conn().RemotePeer()
	r := msgio.NewReader(s)

	// Read the request
	buf, err := r.ReadMsg()
	if err != nil {
		s.Reset()
		logger.Errorf("%s - Failed to Read Invite Request buffer.", err)
		return
	}
	s.Close()

	// unmarshal it
	req := &InviteRequest{}
	err = proto.Unmarshal(buf, req)
	if err != nil {
		logger.Errorf("%s - Failed to Unmarshal Invite REQUEST buffer.", err)
		return
	}

	// generate response message
	err = p.sessionQueue.AddIncoming(remotePeer, req)
	if err != nil {
		logger.Errorf("%s - Failed to add incoming session to queue.", err)
		return
	}

	// store request data into Context
	p.node.OnInvite(req.ToEvent())
}

// onIncomingTransfer incoming transfer handler
func (p *TransmitProtocol) onIncomingTransfer(stream network.Stream) {
	logger.Debug("Received Incoming Transfer")
	// Find Entry in Queue
	entry, err := p.sessionQueue.Next()
	if err != nil {
		logger.Errorf("%s - Failed to find transfer request", err)
		stream.Close()
		return
	}

	// Create New Reader
	event, err := entry.ReadFrom(stream, p.node)
	if err != nil {
		logger.Errorf("%s - Failed to Read From Stream", err)
		stream.Reset()
		return
	}
	p.node.OnComplete(event)
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
func NewItemReader(index int, count int, item *common.Payload_Item, node api.NodeImpl) (*itemReader, error) {
	// generate path
	path, err := fs.Downloads.GenPath(item.GetFile().GetPath())
	if err != nil {
		return nil, err
	}
	f := item.GetFile()
	f.Path = path

	// Create New Item Reader
	return &itemReader{
		item:   f,
		size:   f.GetSize(),
		index:  index,
		count:  count,
		path:   path,
		buffer: bytes.Buffer{},
		node:   node,
	}, nil
}

// Returns Progress of File, Given the written number of bytes
func (p *itemReader) Progress(i int) {
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
	for i := 0; i < int(ir.size); {
		// Read Next Message
		buf, err := reader.ReadMsg()
		if err == io.EOF {
			break
		} else if err != nil {
			logger.Errorf("%s - Failed to Read Next Message on Read Stream", err)
			return
		} else {
			// Write Chunk to File
			n, err := ir.buffer.Write(buf)
			if err != nil {
				logger.Errorf("%s - Failed to Write Buffer to File on Read Stream", err)
				return
			}
			i += n
			ir.Progress(i)
		}
	}

	// Write File Buffer to File
	err := ioutil.WriteFile(ir.path, ir.buffer.Bytes(), 0644)
	if err != nil {
		logger.Errorf("%s - Failed to Close item on Read Stream", err)
		return
	}
	logger.Debug("Completed writing to file: " + ir.path)
	return
}
