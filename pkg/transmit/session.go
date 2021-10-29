package transmit

import (
	"context"
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
	ctx         context.Context
	direction   common.Direction
	from        *common.Peer
	to          *common.Peer
	payload     *common.Payload
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
func (s Session) HandleComplete(ctx context.Context, n api.NodeImpl, wg *sync.WaitGroup, compChan chan itemResult) {
	for {
		select {
		case r := <-compChan:
			// Complete Wait Group
			wg.Done()

			// Update Success
			logger.Debug("Received Item Result", golog.Fields{"success": r.success})
			s.success[int32(r.index)] = r.success

			// Replace Incoming Item
			if r.IsIncoming() {
				s.payload.Items[r.index] = r.item
				s.lastUpdated = int64(time.Now().Unix())
			}

			// Handle Completion
			if r.index == s.Count()-1 {
				close(compChan)
				return
			} else {
				continue
			}
		case <-ctx.Done():
			close(compChan)
			return
		}
	}
}

// -----------------------------------------------------------------------------
// Transmit Reader Interface
// -----------------------------------------------------------------------------

// ReadFrom reads the next Session from the given stream.
func (s Session) ReadFrom(stream network.Stream, n api.NodeImpl) {
	// Initialize Params
	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()
	logger.Debug("Beginning INCOMING Transmit Stream")

	// Handle incoming stream
	rs := msgio.NewReader(stream)
	compChan := make(chan itemResult)
	var wg sync.WaitGroup

	// Add to Wait Group
	wg.Add(len(s.Items()))
	go s.HandleComplete(ctx, n, &wg, compChan)

	// Write All Files
	for i, v := range s.Items() {
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
		handleItemRead(cancel, config, compChan)
		logger.Debugf("Done: Reading Item - %v", i)
	}

	// Wait for all items to be written
	wg.Wait()
	stream.Close()
	n.OnComplete(s.Event())
}

// handleItemRead Returns a new Reader for the given FileItem
func handleItemRead(cancel context.CancelFunc, config itemConfig, compChan chan itemResult) {
	// Create New Writer
	ir := &itemReader{}
	config.ApplyReader(ir)

	// Define Finish Function and Start Channels
	go ir.handleProgress()
	callFinishFunc := func(r bool) {
		ir.doneChan <- r
		compChan <- ir.toResult(r)

		if ir.index == ir.count-1 {
			cancel()
		}
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
	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()
	logger.Debug("Beginning OUTGOING Transmit Stream")
	wc := msgio.NewWriter(stream)
	compChan := make(chan itemResult)
	var wg sync.WaitGroup

	// Add to Wait Group
	wg.Add(len(s.Items()))
	go s.HandleComplete(ctx, n, &wg, compChan)

	// Create New Writer
	for i, v := range s.Items() {
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
		handleItemWrite(cancel, config, compChan)
		logger.Debugf("Done: Reading Item - %v", i)
	}

	// Wait for all items to be written
	wg.Wait()
	n.OnComplete(s.Event())
}

// handleItemWrite handles the writing of a FileItem to a Stream
func handleItemWrite(cancel context.CancelFunc, config itemConfig, compChan chan itemResult) {
	// Create New Writer
	iw := &itemWriter{}
	config.ApplyWriter(iw)

	// Define Finish Function and Start Channels
	go iw.handleProgress()
	callFinishFunc := func(r bool) {
		iw.doneChan <- r
		compChan <- iw.toResult(r)

		// Check if Complete
		if iw.index == iw.count-1 {
			cancel()
		}
	}

	// Define Chunker Opts
	chunker, err := fs.NewFileChunker(config.Path())
	if err != nil {
		logger.Errorf("%s - Failed to create new chunker.", err)
		iw.doneChan <- false
	}

	// Loop through File
	for iw.written < int(iw.size) {
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
	}

	// Flush Buffer to Stream
	logger.Debug("Item Write is Complete")
	callFinishFunc(true)
	return
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
