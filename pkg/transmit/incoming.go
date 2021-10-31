package transmit

import (
	"bytes"
	"io"
	"io/ioutil"
	"sync"

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

// NewItemReader Returns a new Reader for the given FileItem
func ReadItem(index int, count int, pi *common.Payload_Item, wg *sync.WaitGroup, node api.NodeImpl, reader msgio.ReadCloser) {
	defer wg.Done()
	// generate path
	item := pi.GetFile()
	size := item.GetSize()
	path, err := item.ResetPath(fs.Downloads)
	if err != nil {
		logger.Errorf("%s - Failed to generate path for file: %s", err, item.Name)
		return
	}

	buffer := bytes.Buffer{}

	// Route Data from Stream
	for i := 0; i < int(size); {
		// Read Next Message
		buf, err := reader.ReadMsg()
		if err == io.EOF {
			break
		} else if err != nil {
			logger.Errorf("%s - Failed to Read Next Message on Read Stream", err)
			return
		} else {
			// Write Chunk to File
			n, err := buffer.Write(buf)
			if err != nil {
				logger.Errorf("%s - Failed to Write Buffer to File on Read Stream", err)
				return
			}
			i += n

			// Update Progress
			pushProgress(node, common.Direction_INCOMING, i, size, index, count)
		}
	}

	// Write File Buffer to File
	err = ioutil.WriteFile(path, buffer.Bytes(), 0644)
	if err != nil {
		logger.Errorf("%s - Failed to Close item on Read Stream", err)
		return
	}
	logger.Debug("Completed writing to file: " + path)
	return
}
