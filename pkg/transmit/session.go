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
	compChan    chan FileItemStreamResult
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
func (s Session) HandleComplete(n api.NodeImpl) {
	success := make(map[int32]bool)
	for {
		select {
		case result := <-s.compChan:
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

// ReadFrom reads the next Session from the given stream.
func (s Session) ReadFrom(stream network.Stream, n api.NodeImpl) error {
	// Initialize Params
	logger.Debug("Beginning INCOMING Transmit Stream")

	// Handle incoming stream
	rs := msgio.NewReader(stream)
	var wg sync.WaitGroup
	go s.HandleComplete(n)

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
			wg:       wg,
			compChan: s.compChan,
		}

		// Create Reader
		err := handleItemRead(config)
		if err != nil {
			logger.Errorf("%s - Failed to create new reader.", err)
			rs.Close()
			return err
		}
	}
	wg.Wait()
	stream.Close()

	// Return Complete Event
	return nil
}

// handleItemRead Returns a new Reader for the given FileItem
func handleItemRead(config itemConfig) error {
	// Create New Writer
	ir := &itemReader{}
	config.ApplyReader(ir)
	defer config.wg.Done()

	// Start Channels
	go ir.handleChannels(config.wg, config.compChan)

	// Route Data from Stream
	for !ir.isItemComplete() {
		// Read Next Message
		buf, err := config.reader.ReadMsg()
		if err != nil {
			logger.Errorf("%s - Failed to Read Next Message on Read Stream", err)
			ir.doneChan <- false
			return err
		} else {
			// Write Chunk to File
			n, err := ir.buffer.Write(buf)
			if err != nil {
				logger.Errorf("%s - Failed to Write Buffer to File on Read Stream", err)
				ir.doneChan <- false
				return err
			}
			ir.progressChan <- n
		}
	}
	ir.doneChan <- true
	return nil
}

// WriteTo writes the Session to the given stream.
func (s Session) WriteTo(stream network.Stream, n api.NodeImpl) error {
	// Initialize Params
	logger.Debug("Beginning OUTGOING Transmit Stream")
	wc := msgio.NewWriter(stream)
	var wg sync.WaitGroup
	go s.HandleComplete(n)

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
			wg:       wg,
			compChan: s.compChan,
		}

		// Create New Writer
		err := handleItemWrite(config)
		if err != nil {
			logger.Errorf("%s - Failed to create new writer.", err)
			wc.Close()
			return err
		}
	}

	// Wait for all writes to finish
	wg.Wait()
	return nil
}

// handleItemWrite handles the writing of a FileItem to a Stream
func handleItemWrite(config itemConfig) error {
	// Create New Writer
	iw := &itemWriter{}
	config.ApplyWriter(iw)

	// Define Chunker Opts
	chunker, err := fs.NewFileChunker(config.Path())
	if err != nil {
		logger.Errorf("%s - Failed to create new chunker.", err)
		return err
	}

	go iw.handleChannels(config.wg, config.compChan)

	// Loop through File
	for !iw.isItemComplete() {
		c, err := chunker.Next()
		if err != nil {
			// Handle EOF
			if err == io.EOF {
				iw.doneChan <- true
				break
			}

			// Unexpected Error
			logger.Errorf("%s - Error reading chunk.", err)
			iw.doneChan <- false
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
	iw.doneChan <- true
	return nil
}
