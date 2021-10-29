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
	results     map[int32]bool
	compChan    chan itemResult
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
		Results:    s.results,
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
func (s Session) HandleComplete(ctx context.Context, n api.NodeImpl, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case r := <-s.compChan:
			if !r.success {
				logger.Errorf("Failed to Complete File: %s", r.item.GetFile().GetName())
				return
			}
			// Update Success
			logger.Debug("Received Item Result", golog.Fields{"success": r.success})
			s.results[int32(r.index)] = r.success

			// Replace Incoming Item
			if r.IsIncoming() {
				s.payload.Items[r.index] = r.item
				s.lastUpdated = int64(time.Now().Unix())
			}
		case <-ctx.Done():
			return
		}
	}
}

// -----------------------------------------------------------------------------
// Transmit Reader Interface
// -----------------------------------------------------------------------------

// ReadFrom reads the next Session from the given stream.
func (s Session) ReadFrom(stream network.Stream, n api.NodeImpl) {
	logger.Debug("Beginning INCOMING Transmit Stream")
	// Initialize Params
	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()
	var wg sync.WaitGroup

	// Create Reader
	rs := msgio.NewReader(stream)
	for i, v := range s.Items() {
		wg.Add(1)
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
		go s.HandleComplete(ctx, n, &wg)
		logger.Debugf("Done: Reading Item - %v", i)
	}

	// Wait for all items to be written
	wg.Wait()
	stream.Close()
	close(s.compChan)
	n.OnComplete(s.Event())
}

// handleItemRead Returns a new Reader for the given FileItem
func handleItemRead(config itemConfig, compChan chan itemResult) {
	logger.Debug("Handling Item Read...")
	// Create New Reader
	ir := &itemReader{}
	if err := config.ApplyReader(ir); err != nil {
		logger.Errorf("Failed to Apply Reader: %s", err)
		return
	}
	go ir.handleProgress(compChan)

	// Route Data from Stream
	for int64(ir.written) < ir.size {
		// Read Next Message
		buf, err := config.reader.ReadMsg()
		if err != nil {
			// Handle EOF
			if err == io.EOF {
				logger.Debug("Reader has reached end of stream.")
				break
			}

			// Unexpected Error
			logger.Errorf("%s - Failed to Read Next Message on Read Stream", err)
			ir.doneChan <- false
			return
		} else {
			// Write Chunk to File
			if err := ir.WriteChunk(buf); err != nil {
				logger.Errorf("%s - Failed to Write Buffer to File on Read Stream", err)
				ir.doneChan <- false
				return
			}
		}
	}

	// Complete Writing to File
	ir.doneChan <- true
}

// handleProgress handles the channels for the ItemReader
func (p *itemReader) handleProgress(compChan chan itemResult) {
	for {
		select {
		case n := <-p.progressChan:
			p.written += n
			if ev := p.getProgressEvent(); ev != nil {
				p.node.OnProgress(ev)
			}
		case r := <-p.doneChan:
			if r {
				logger.Debug("Item Read has Completed, successfully")
				// Write Buffer to File
				if err := p.item.WriteFile(p.buffer.Bytes()); err != nil {
					logger.Errorf("%s - Failed to Sync File on Read Stream", err)
					compChan <- p.toResult(false)
				} else {
					compChan <- p.toResult(true)
				}
			} else {
				logger.Error("Item Read has Completed, unsuccessfully")
				compChan <- p.toResult(false)
			}
			return
		}
	}
}

// -----------------------------------------------------------------------------
// Transmit Writer Interface
// -----------------------------------------------------------------------------

// WriteTo writes the Session to the given stream.
func (s Session) WriteTo(stream network.Stream, n api.NodeImpl) {
	logger.Debug("Beginning OUTGOING Transmit Stream")
	// Initialize Params
	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()
	var wg sync.WaitGroup

	// Create New Writer
	wc := msgio.NewWriter(stream)
	for i, v := range s.Items() {
		wg.Add(1)
		// Configure Writer
		config := itemConfig{
			index:  i,
			count:  s.Count(),
			item:   v,
			node:   n,
			writer: wc,
		}

		// Create New Writer
		logger.Debugf("Start: Writing Item - %v", i)
		handleItemWrite(config, s.compChan)
		go s.HandleComplete(ctx, n, &wg)
		logger.Debugf("Done: Writing Item - %v", i)
	}

	// Wait for all items to be written
	wg.Wait()
	n.OnComplete(s.Event())
	close(s.compChan)
}

// handleItemWrite handles the writing of a FileItem to a Stream
func handleItemWrite(config itemConfig, compChan chan itemResult) {
	logger.Debug("Handling Item Write...")
	// Create New Writer
	iw := &itemWriter{}
	config.ApplyWriter(iw)
	go iw.handleProgress(compChan)

	// Define Chunker Opts
	chunker, err := fs.NewFileChunker(config.Path())
	if err != nil {
		logger.Errorf("%s - Failed to create new chunker.", err)
		iw.doneChan <- false
		return
	}

	// Loop through File
	for iw.written < int(iw.size) {
		c, err := chunker.Next()
		if err != nil {
			// Handle EOF
			if err == io.EOF {
				logger.Debug("Chunker has reached end of file.")
				break
			}

			// Unexpected Error
			logger.Errorf("%s - Error reading chunk.", err)
			iw.doneChan <- false
			return
		}

		// Write Chunk to Stream
		if err := iw.WriteChunk(c.Data, c.Length); err != nil {
			logger.Errorf("%s - Error Writing data to msgio.Writer", err)
			iw.doneChan <- false
			return
		}
	}

	// Flush Buffer to Stream
	iw.doneChan <- true
}

// handleProgress handles the channels for the ItemReader
func (p *itemWriter) handleProgress(compChan chan itemResult) {
	for {
		select {
		case n := <-p.progressChan:
			p.written += n
			if ev := p.getProgressEvent(); ev != nil {
				p.node.OnProgress(ev)
			}
		case r := <-p.doneChan:
			logger.Debug("Item Write is Complete")
			compChan <- p.toResult(r)
			return
		}
	}
}
