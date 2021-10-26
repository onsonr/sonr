package transmit

import (
	"bytes"
	"errors"
	"io"
	"os"

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
	item         *common.FileItem
	buffer       bytes.Buffer
	path         string
	index        int
	count        int
	size         int64
	node         api.NodeImpl
	written      int
	progressChan chan int
	completeChan chan bool // completeChan is closed when the stream finishes reading
	doneChan     chan bool // doneChan is closed when the file is written to disk
}

// handleItemRead Returns a new Reader for the given FileItem
func handleItemRead(config itemConfig) (*common.Payload_Item, error) {
	defer config.wg.Done()
	// generate path
	path, err := fs.Downloads.GenPath(config.item.GetFile().GetPath())
	if err != nil {
		logger.Errorf("%s - Failed to create new ItemReader", err)
		return nil, err
	}

	// Get File Item
	f := config.item.GetFile()
	err = f.SetPath(path)
	if err != nil {
		logger.Errorf("%s - Failed to create new ItemReader", err)
		return nil, err
	}

	// Create New Item Reader
	ir := &itemReader{
		item:         f,
		size:         f.GetSize(),
		index:        config.index,
		count:        config.count,
		path:         path,
		buffer:       bytes.Buffer{},
		node:         config.node,
		written:      0,
		progressChan: make(chan int),
		completeChan: make(chan bool),
		doneChan:     make(chan bool),
	}

	// Start Channels
	go ir.handleChannels()

	// Route Data from Stream
	for {
		// Read Next Message
		buf, err := config.reader.ReadMsg()
		if err == io.EOF {
			ir.completeChan <- true
			break
		} else if err != nil {
			logger.Errorf("%s - Failed to Read Next Message on Read Stream", err)
			ir.completeChan <- false
			return nil, err
		} else {
			// Write Chunk to File
			n, err := ir.buffer.Write(buf)
			if err != nil {
				logger.Errorf("%s - Failed to Write Buffer to File on Read Stream", err)
				ir.completeChan <- false
				return nil, err
			}
			ir.progressChan <- n
		}
	}
	return ir.toPayloadItem()
}

// getProgressEvent returns a ProgressEvent for the current ItemReader
func (p *itemReader) getProgressEvent() *api.ProgressEvent {
	if (p.written % ITEM_INTERVAL) == 0 {
		// Create Progress Event
		return &api.ProgressEvent{
			Progress: (float64(p.written) / float64(p.size)),
			Current:  int32(p.index),
			Total:    int32(p.count),
		}
	}
	return nil
}

// handleChannels handles the channels for the ItemReader
func (p *itemReader) handleChannels() {
	for {
		select {
		case n := <-p.progressChan:
			p.written += n
			if ev := p.getProgressEvent(); ev != nil {
				p.node.OnProgress(ev)
			}
		case r := <-p.completeChan:
			if r {
				// Write Buffer to File
				err := os.WriteFile(p.path, p.buffer.Bytes(), 0644)
				if err != nil {
					logger.Errorf("%s - Failed to Close item on Read Stream", err)
					p.doneChan <- false
					return
				}
			} else {
				// Delete File
				err := os.Remove(p.path)
				if err != nil {
					logger.Errorf("%s - Failed to Close item on Read Stream", err)
					p.doneChan <- false
					return
				}
			}
			p.doneChan <- true
			return
		}
	}
}

func (ir *itemReader) waitForDone() {
	<-ir.completeChan
}

// toPayloadItem Waits for the File to be completed
func (ir *itemReader) toPayloadItem() (*common.Payload_Item, error) {
	r := <-ir.doneChan
	if !r {
		return nil, errors.New("Failed to Write File to Disk")
	}

	// Create Payload Item
	item := ir.item.ToTransferItem()
	return item, nil
}
