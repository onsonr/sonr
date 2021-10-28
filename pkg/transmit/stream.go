package transmit

import (
	"bytes"
	"sync"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
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
	}
	s.Close()

	// unmarshal it
	req := &InviteRequest{}
	err = proto.Unmarshal(buf, req)
	if err != nil {
		logger.Errorf("%s - Failed to Unmarshal Invite REQUEST buffer.", err)
	}

	// generate response message
	err = p.sessionQueue.AddIncoming(remotePeer, req)
	if err != nil {
		logger.Errorf("%s - Failed to add incoming session to queue.", err)
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
	}

	// Create New Reader
	err = entry.ReadFrom(stream, p.node)
	if err != nil {
		logger.Errorf("%s - Failed to Read From Stream", err)
		stream.Close()
	}
}

// onInviteResponse response handler
func (p *TransmitProtocol) onInviteResponse(s network.Stream) {
	logger.Debug("Received Invite Response")
	// Initialize Stream Data
	remotePeer := s.Conn().RemotePeer()
	r := msgio.NewReader(s)

	// Read the request
	buf, err := r.ReadMsg()
	if err != nil {
		logger.Errorf("%s - Failed to Read Invite RESPONSE buffer.", err)
	}
	s.Close()

	// Unmarshal response
	resp := &InviteResponse{}
	err = proto.Unmarshal(buf, resp)
	if err != nil {
		logger.Errorf("%s - Failed to Unmarshal Invite RESPONSE buffer.", err)
	}

	// Locate request data and remove it if found
	entry, err := p.sessionQueue.Validate(resp)
	if err != nil {
		logger.Errorf("%s - Failed to Validate Invite RESPONSE buffer.", err)
	}

	// Check for Decision and Start Outgoing Transfer
	if resp.GetDecision() {
		// Create a new stream
		stream, err := p.host.NewStream(p.ctx, remotePeer, SessionPID)
		if err != nil {
			logger.Errorf("%s - Failed to create new stream.", err)
		}

		// Call Outgoing Transfer
		p.onOutgoingTransfer(entry, stream)
	}
	p.node.OnDecision(resp.ToEvent())
}

// onOutgoingTransfer is called by onInviteResponse if Validated
func (p *TransmitProtocol) onOutgoingTransfer(entry Session, stream network.Stream) {
	logger.Debug("Received Accept Decision, Starting Outgoing Transfer")
	// Create New Writer
	err := entry.WriteTo(stream, p.node)
	if err != nil {
		logger.Errorf("%s - Failed to Write To Stream", err)
		stream.Close()
		return
	}
}

// itemReader is a Reader for a FileItem
type itemReader struct {
	item         *common.FileItem
	buffer       bytes.Buffer
	index        int
	count        int
	size         int64
	node         api.NodeImpl
	written      int
	progressChan chan int
	doneChan     chan bool
}

// getProgressEvent returns a ProgressEvent for the current ItemReader
func (p *itemReader) getProgressEvent() *api.ProgressEvent {
	if (p.written % ITEM_INTERVAL) == 0 {
		// Create Progress Event
		return &api.ProgressEvent{
			Direction: common.Direction_INCOMING,
			Progress:  (float64(p.written) / float64(p.size)),
			Current:   int32(p.index),
			Total:     int32(p.count),
		}
	}
	return nil
}

// handleChannels handles the channels for the ItemReader
func (p *itemReader) handleChannels(wg sync.WaitGroup, compChan chan FileItemStreamResult) {
	for {
		select {
		case n := <-p.progressChan:
			p.written += n
			if ev := p.getProgressEvent(); ev != nil {
				p.node.OnProgress(ev)
			}
		case r := <-p.doneChan:
			if r {
				// Write Buffer to File
				err := p.item.WriteFile(p.buffer.Bytes())
				if err != nil {
					logger.Errorf("%s - Failed to Close item on Read Stream", err)
					return
				}
			}
			compChan <- p.toResult(r)
			wg.Done()
			return
		}
	}
}

// isItemComplete returns true if the item has been completely read
func (ir *itemReader) isItemComplete() bool {
	return ir.written >= int(ir.size)
}

// toResult returns a FileItemStreamResult for the current ItemReader
func (ir *itemReader) toResult(success bool) FileItemStreamResult {
	return FileItemStreamResult{
		index:     ir.index,
		item:      ir.item.ToTransferItem(),
		direction: common.Direction_INCOMING,
	}
}

// itemWriter is a Writer for FileItems
type itemWriter struct {
	item         *common.FileItem
	index        int
	count        int
	size         int64
	node         api.NodeImpl
	written      int
	progressChan chan int
	doneChan     chan bool
}

// getProgressEvent returns a ProgressEvent for the current ItemReader
func (p *itemWriter) getProgressEvent() *api.ProgressEvent {
	if (p.written % ITEM_INTERVAL) == 0 {
		// Create Progress Event
		return &api.ProgressEvent{
			Direction: common.Direction_OUTGOING,
			Progress:  (float64(p.written) / float64(p.size)),
			Current:   int32(p.index),
			Total:     int32(p.count),
		}
	}
	return nil
}

// handleChannels handles the channels for the ItemReader
func (p *itemWriter) handleChannels(wg sync.WaitGroup, compChan chan FileItemStreamResult) {
	for {
		select {
		case n := <-p.progressChan:
			p.written += n
			if ev := p.getProgressEvent(); ev != nil {
				p.node.OnProgress(ev)
			}
		case r := <-p.doneChan:
			compChan <- p.toResult(r)
			wg.Done()
			return
		}
	}
}

// isItemComplete returns true if the item has been completely written
func (ir *itemWriter) isItemComplete() bool {
	return ir.written >= int(ir.size)
}

// toResult returns a FileItemStreamResult for the current ItemReader
func (ir *itemWriter) toResult(success bool) FileItemStreamResult {
	return FileItemStreamResult{
		index:     ir.index,
		item:      ir.item.ToTransferItem(),
		direction: common.Direction_OUTGOING,
		success:   success,
	}
}
