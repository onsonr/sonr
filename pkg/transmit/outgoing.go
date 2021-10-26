package transmit

import (
	"io"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
	"github.com/sonr-io/core/pkg/common"

	"google.golang.org/protobuf/proto"
)

// onInviteResponse response handler
func (p *TransmitProtocol) onInviteResponse(s network.Stream) {
	logger.Debug("Received Invite Response")
	// Initialize Stream Data
	remotePeer := s.Conn().RemotePeer()
	r := msgio.NewReader(s)

	// Read the request
	buf, err := r.ReadMsg()
	if err != nil {
		s.Reset()
		logger.Errorf("%s - Failed to Read Invite RESPONSE buffer.", err)
		return
	}
	s.Close()

	// Unmarshal response
	resp := &InviteResponse{}
	err = proto.Unmarshal(buf, resp)
	if err != nil {
		logger.Errorf("%s - Failed to Unmarshal Invite RESPONSE buffer.", err)
		return
	}

	// Locate request data and remove it if found
	entry, err := p.sessionQueue.Validate(resp)
	if err != nil {
		logger.Errorf("%s - Failed to Validate Invite RESPONSE buffer.", err)
		return
	}

	// Check for Decision and Start Outgoing Transfer
	if resp.GetDecision() {
		// Create a new stream
		stream, err := p.host.NewStream(p.ctx, remotePeer, SessionPID)
		if err != nil {
			logger.Errorf("%s - Failed to create new stream.", err)
			return
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
	event, err := entry.WriteTo(stream, p.node)
	if err != nil {
		logger.Errorf("%s - Failed to Write To Stream", err)
		stream.Reset()
		return
	}
	p.node.OnComplete(event)
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

// handleItemWrite handles the writing of a FileItem to a Stream
func handleItemWrite(config itemConfig) error {
	// Create New Writer
	iw := &itemWriter{}
	config.ApplyWriter(iw)
	defer config.wg.Done()

	// Define Chunker Opts
	chunker, err := fs.NewFileChunker(config.Path(), config.Size())
	if err != nil {
		logger.Errorf("%s - Failed to create new chunker.", err)
		return err
	}

	go iw.handleChannels()

	// Loop through File
	for {
		c, err := chunker.Next()
		if err != nil {
			// Handle EOF
			if err == io.EOF {
				iw.progressChan <- int(iw.size)
				iw.doneChan <- true
				break
			}

			// Unexpected Error
			iw.doneChan <- false
			logger.Errorf("%s - Error reading chunk.", err)
			return err
		}

		// Write Message Bytes to Stream
		err = config.writer.WriteMsg(c.Data)
		if err != nil {
			logger.Errorf("%s - Error Writing data to msgio.Writer", err)
			return err
		}
		iw.progressChan <- c.Length
	}
	return nil
}

// getProgressEvent returns a ProgressEvent for the current ItemReader
func (p *itemWriter) getProgressEvent(isComplete bool) *api.ProgressEvent {
	// Create Completed Progress Event
	if isComplete {
		return &api.ProgressEvent{
			Progress: float64(1.0),
			Current:  int32(p.index),
			Total:    int32(p.count),
		}
	}

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
func (p *itemWriter) handleChannels() {
	for {
		select {
		case n := <-p.progressChan:
			finalProgress := n == int(p.size)
			if !finalProgress {
				p.written += n
			}
			if ev := p.getProgressEvent(finalProgress); ev != nil {
				p.node.OnProgress(ev)
			}
		case <-p.doneChan:
			return
		}
	}
}
