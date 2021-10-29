package transmit

import (
	"io"
	"sync"
	"time"

	"github.com/kataras/golog"
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
	success     map[int32]bool
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
func (s Session) Event() *api.CompleteEvent {
	return &api.CompleteEvent{
		From:       s.from,
		To:         s.to,
		Direction:  s.direction,
		Payload:    s.payload,
		CreatedAt:  s.payload.GetCreatedAt(),
		ReceivedAt: int64(time.Now().Unix()),
		Success:    s.success,
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
func (s Session) HandleComplete(n api.NodeImpl, wg *sync.WaitGroup) {
	for {
		select {
		case r := <-s.compChan:
			// Update Success
			s.success[int32(r.index)] = r.success

			// Complete Wait Group
			logger.Debug("Received Item Result", golog.Fields{"success": r.success})
			wg.Done()

			// Replace Incoming Item
			if r.IsIncoming() {
				s.payload.Items[r.index] = r.item
				s.lastUpdated = int64(time.Now().Unix())
			}

			// Check if Complete
			if r.index == s.Count()-1 {
				return
			}
		}
	}
}

// -----------------------------------------------------------------------------
// Transmit Reader Interface
// -----------------------------------------------------------------------------

// ReadFrom reads the next Session from the given stream.
func (s Session) ReadFrom(stream network.Stream, n api.NodeImpl) {
	// Initialize Params
	logger.Debug("Beginning INCOMING Transmit Stream")

	// Handle incoming stream
	rs := msgio.NewReader(stream)
	var wg sync.WaitGroup

	// Write All Files
	for i, v := range s.Items() {
		// Write File to Stream
		wg.Add(1)
		go s.HandleComplete(n, &wg)

		// Configure Reader
		config := itemConfig{
			index:  i,
			count:  s.Count(),
			item:   v,
			node:   n,
			reader: rs,
		}

		// Create Reader
		logger.Debugf("Start: Reading Item - %v", i)
		handleItemRead(config, s.compChan)
		wg.Wait()
		logger.Debugf("Done: Reading Item - %v", i)
	}

	// Wait for all items to be written
	stream.Close()
	n.OnComplete(s.Event())
}

// handleItemRead Returns a new Reader for the given FileItem
func handleItemRead(config itemConfig, compChan chan itemResult) {
	// Create New Writer
	ir := &itemReader{}
	config.ApplyReader(ir)

	// Define Finish Function and Start Channels
	go ir.handleProgress()
	callFinishFunc := func(r bool) {
		ir.doneChan <- r
		compChan <- ir.toResult(r)
	}

	// Route Data from Stream
	for int64(ir.written) < ir.size {
		// Read Next Message
		buf, err := config.reader.ReadMsg()
		if err != nil {
			logger.Errorf("%s - Failed to Read Next Message on Read Stream", err)
			callFinishFunc(false)
			return
		} else {
			// Write Chunk to File
			if err := ir.WriteChunk(buf); err != nil {
				logger.Errorf("%s - Failed to Write Buffer to File on Read Stream", err)
				callFinishFunc(false)
				return
			}
		}
	}

	// Flush Buffer to File
	if err := ir.FlushBuffer(); err != nil {
		logger.Errorf("%s - Failed to Sync File on Read Stream", err)
		callFinishFunc(false)
	}

	// Complete Writing to File
	callFinishFunc(true)
	return
}

// FlushBuffer writes the current buffer to the file.
func (p *itemReader) FlushBuffer() error {
	// Stop Channels
	logger.Debug("Item Read is Complete")

	// Write Buffer to File
	err := p.item.WriteFile(p.buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// handleProgress handles the channels for the ItemReader
func (p *itemReader) handleProgress() {
	for {
		select {
		case n := <-p.progressChan:
			p.written += n
			if ev := p.getProgressEvent(); ev != nil {
				p.node.OnProgress(ev)
			}
		case <-p.doneChan:
			return
		}
	}
}

// -----------------------------------------------------------------------------
// Transmit Writer Interface
// -----------------------------------------------------------------------------

// WriteTo writes the Session to the given stream.
func (s Session) WriteTo(stream network.Stream, n api.NodeImpl) {
	// Initialize Params
	logger.Debug("Beginning OUTGOING Transmit Stream")
	wc := msgio.NewWriter(stream)
	var wg sync.WaitGroup

	// Create New Writer
	for i, v := range s.Items() {
		// Write File to Stream
		wg.Add(1)
		go s.HandleComplete(n, &wg)

		// Configure Writer
		config := itemConfig{
			index:  i,
			count:  s.Count(),
			item:   v,
			node:   n,
			writer: wc,
		}

		// Create New Writer
		logger.Debugf("Start: Reading Item - %v", i)
		handleItemWrite(config, s.compChan)
		wg.Wait()
		logger.Debugf("Done: Reading Item - %v", i)
	}
	n.OnComplete(s.Event())
}

// handleItemWrite handles the writing of a FileItem to a Stream
func handleItemWrite(config itemConfig, compChan chan itemResult) {
	// Create New Writer
	iw := &itemWriter{}
	config.ApplyWriter(iw)

	// Define Finish Function and Start Channels
	go iw.handleProgress()
	callFinishFunc := func(r bool) {
		iw.doneChan <- r
		compChan <- iw.toResult(r)
	}

	// Define Chunker Opts
	chunker, err := fs.NewFileChunker(config.Path())
	if err != nil {
		logger.Errorf("%s - Failed to create new chunker.", err)
		iw.doneChan <- false
	}

	// Loop through File
	for {
		c, err := chunker.Next()
		if err != nil {
			// Handle EOF
			if err == io.EOF {
				logger.Debug("Chunker has reached end of file.")
				callFinishFunc(true)
				return
			}

			// Unexpected Error
			logger.Errorf("%s - Error reading chunk.", err)
			callFinishFunc(false)
			return
		}

		// Write Chunk to Stream
		if err := iw.WriteChunk(c.Data, c.Length); err != nil {
			logger.Errorf("%s - Error Writing data to msgio.Writer", err)
			callFinishFunc(false)
			return
		}

		// Check if Item is Complete
		if iw.isItemComplete() {
			logger.Debug("Item Write is Complete")
			callFinishFunc(true)
			return
		}
	}
}

// handleProgress handles the channels for the ItemReader
func (p *itemWriter) handleProgress() {
	for {
		select {
		case n := <-p.progressChan:
			p.written += n
			if ev := p.getProgressEvent(); ev != nil {
				p.node.OnProgress(ev)
			}
		case <-p.doneChan:
			return
		}
	}
}
