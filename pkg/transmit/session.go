package transmit

import (
	"io"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
	"github.com/sonr-io/core/pkg/common"
)

// Session is a single entry in the Transmit queue.
type Session struct {
	direction   common.Direction
	from        *common.Peer
	to          *common.Peer
	payload     *common.Payload
	compChan    chan itemResult
	lastUpdated int64
}

// MapItems performs PayloadItemFunc on each item in the Payload.
func (s Session) Items() []*common.Payload_Item {
	return s.payload.GetItems()
}

// Count returns the number of items in the payload.
func (s Session) Count() int {
	return len(s.Items())
}

// Event returns the CompleteEvent for the given session.
func (s Session) Event(success map[int32]bool) *api.CompleteEvent {
	return &api.CompleteEvent{
		From:       s.from,
		To:         s.to,
		Direction:  s.direction,
		Payload:    s.payload,
		CreatedAt:  s.payload.GetCreatedAt(),
		ReceivedAt: int64(time.Now().Unix()),
		Success:    success,
	}
}

// IsIncoming returns true if the session is incoming.
func (s Session) IsIncoming() bool {
	return s.direction == common.Direction_INCOMING
}

// IsOutgoing returns true if the session is outgoing.
func (s Session) IsOutgoing() bool {
	return s.direction == common.Direction_OUTGOING
}

// HandleComplete handles the completion of a session item.
func (s Session) HandleComplete(wg sync.WaitGroup, n api.NodeImpl) {
	success := make(map[int32]bool)
	for {
		select {
		case result := <-s.compChan:
			wg.Done()
			// Update Success
			success[int32(result.index)] = result.success

			// Replace Incoming Item
			if result.IsIncoming() {
				s.payload.Items[result.index] = result.item
				s.lastUpdated = int64(time.Now().Unix())
			}

			// Check if all items have been received
			if result.IsAllCompleted(s.Count()) {
				logger.Debug("All items completed")
				n.OnComplete(s.Event(success))
				return
			}
		}
	}
}

// -----------------------------------------------------------------------------
// Transmit Reader Interface
// -----------------------------------------------------------------------------

// ReadFrom reads the next Session from the given stream.
func (s Session) ReadFrom(stream network.Stream, n api.NodeImpl) error {
	// Initialize Params
	logger.Debug("Beginning INCOMING Transmit Stream")

	// Handle incoming stream
	rs := msgio.NewReader(stream)
	var wg sync.WaitGroup
	go s.HandleComplete(wg, n)

	// Write All Files
	for i, v := range s.Items() {
		// Write File to Stream
		wg.Add(1)

		// Configure Reader
		config := itemConfig{
			index:    i,
			count:    s.Count(),
			item:     v,
			node:     n,
			reader:   rs,
		}

		// Create Reader
		err := handleItemRead(config, s.compChan)
		if err != nil {
			logger.Errorf("%s - Failed to create new reader.", err)
			return err
		}
	}

	// Wait for all items to be written
	wg.Wait()
	stream.Close()
	return nil
}

// handleItemRead Returns a new Reader for the given FileItem
func handleItemRead(config itemConfig, compChan chan itemResult) error {
	// Create New Writer
	ir := &itemReader{}
	config.ApplyReader(ir)

	// Start Channels
	go ir.handleChannels(compChan)

	// Route Data from Stream
	for {
		// Read Next Message
		buf, err := config.reader.ReadMsg()
		if err != nil {
			logger.Errorf("%s - Failed to Read Next Message on Read Stream", err)
			ir.doneChan <- false
			return err
		} else {
			// Write Chunk to File
			if err := ir.WriteChunk(buf); err != nil {
				logger.Errorf("%s - Failed to Write Buffer to File on Read Stream", err)
				ir.doneChan <- false
				return err
			}
		}

		// Check if Item is Complete
		if ir.isItemComplete() {
			logger.Debug("Item Read is Complete")
			ir.doneChan <- true
			return nil
		}
	}
}

// handleChannels handles the channels for the ItemReader
func (p *itemReader) handleChannels(compChan chan itemResult) {
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
			return
		}
	}
}

// -----------------------------------------------------------------------------
// Transmit Writer Interface
// -----------------------------------------------------------------------------

// WriteTo writes the Session to the given stream.
func (s Session) WriteTo(stream network.Stream, n api.NodeImpl) error {
	// Initialize Params
	logger.Debug("Beginning OUTGOING Transmit Stream")
	wc := msgio.NewWriter(stream)
	var wg sync.WaitGroup
	go s.HandleComplete(wg, n)

	// Create New Writer
	for i, v := range s.Items() {
		// Write File to Stream
		wg.Add(1)

		// Configure Writer
		config := itemConfig{
			index:    i,
			count:    s.Count(),
			item:     v,
			node:     n,
			writer:   wc,
		}

		// Create New Writer
		err := handleItemWrite(config, s.compChan)
		if err != nil {
			logger.Errorf("%s - Failed to create new writer.", err)
			return err
		}
	}

	// Wait for all writes to finish
	wg.Wait()
	return nil
}

// handleItemWrite handles the writing of a FileItem to a Stream
func handleItemWrite(config itemConfig, compChan chan itemResult) error {
	// Create New Writer
	iw := &itemWriter{}
	config.ApplyWriter(iw)

	// Define Chunker Opts
	chunker, err := fs.NewFileChunker(config.Path())
	if err != nil {
		logger.Errorf("%s - Failed to create new chunker.", err)
		return err
	}

	// Start Channels
	go iw.handleChannels(compChan)

	// Loop through File
	for {
		c, err := chunker.Next()
		if err != nil {
			// Handle EOF
			if err == io.EOF {
				iw.doneChan <- true
				return nil
			}

			// Unexpected Error
			logger.Errorf("%s - Error reading chunk.", err)
			iw.doneChan <- false
			return err
		}

		// Write Chunk to Stream
		if err := iw.WriteChunk(c.Data); err != nil {
			logger.Errorf("%s - Error Writing data to msgio.Writer", err)
			iw.doneChan <- false
			return err
		}

		// Check if Item is Complete
		if iw.isItemComplete() {
			logger.Debug("Item Write is Complete")
			iw.doneChan <- true
			return nil
		}
	}
}

// handleChannels handles the channels for the ItemReader
func (p *itemWriter) handleChannels(compChan chan itemResult) {
	for {
		select {
		case n := <-p.progressChan:
			p.written += n
			if ev := p.getProgressEvent(); ev != nil {
				p.node.OnProgress(ev)
			}
		case r := <-p.doneChan:
			compChan <- p.toResult(r)
			return
		}
	}
}
