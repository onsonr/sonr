package transmit

import (
	"io"
	"sync"

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
		stream, err := p.host.NewStream(p.ctx, remotePeer, IncomingPID)
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

// NewItemWriter Returns a new Reader for the given FileItem
func WriteItem(index int, count int, pi *common.Payload_Item, wg *sync.WaitGroup, node api.NodeImpl, writer msgio.WriteCloser) {
	// Properties
	defer wg.Done()
	size := pi.GetSize()
	item := pi.GetFile()

	// Create New Chunker
	chunker, err := fs.NewFileChunker(item.Path)
	if err != nil {
		logger.Errorf("%s - Failed to create new chunker.", err)
		return
	}

	// Loop through File
	for i := 0; i < int(size); {
		c, err := chunker.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Errorf("%s - Error reading chunk.", err)
			return
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(c.Data)
		if err != nil {
			logger.Errorf("%s - Error Writing data to msgio.Writer", err)
			return
		}

		// Unexpected Error
		if err != nil && err != io.EOF {
			logger.Errorf("%s - Unexpected Error occurred on Write Stream", err)
			return
		}
		// Update Progress
		i += c.Length

		// Update Progress
		pushProgress(node, common.Direction_OUTGOING, i, size, index, count)
	}
	return
}
