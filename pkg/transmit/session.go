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
}

// CreateReadConfig creates a new itemConfig for the given index for reading.
func (s Session) CreateReadConfig(index int, n api.NodeImpl, reader msgio.ReadCloser) itemConfig {
	return itemConfig{
		index:  index,
		node:   n,
		reader: reader,
		item:   s.Items()[index],
		count:  len(s.Items()),
	}
}

// CreateWriteConfig creates a new itemConfig for the given index for writing.
func (s Session) CreateWriteConfig(index int, n api.NodeImpl, writer msgio.WriteCloser) itemConfig {
	return itemConfig{
		index:  index,
		node:   n,
		writer: writer,
		item:   s.Items()[index],
		count:  len(s.Items()),
	}
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
func (s Session) HandleComplete(ctx context.Context, n api.NodeImpl, wg *sync.WaitGroup, compChan chan itemResult) {
	defer wg.Done()
	r := <-compChan
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
	close(compChan)
	return
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
	for i := range s.Items() {
		// Initialize Sync Management
		compChan := make(chan itemResult, 1)
		wg.Add(1)
		go s.HandleComplete(ctx, n, &wg, compChan)

		// Create Reader
		logger.Debugf("Start: Reading Item - %v", i)
		handleItemRead(s.CreateReadConfig(i, n, rs), compChan)
		logger.Debugf("Done: Reading Item - %v", i)
	}

	// Wait for all items to be written
	wg.Wait()
	stream.Close()
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
	go ir.startRead(config.reader)

	// Route Data from Stream
	for {
		select {
		case n := <-ir.progressChan:
			ir.written += n
			if ev := ir.getProgressEvent(); ev != nil {
				ir.node.OnProgress(ev)
			}
		case r := <-ir.doneChan:
			// Check for Success
			if r {
				logger.Debug("Item Read has Completed, successfully")
				// Write Buffer to File
				if err := ir.item.WriteFile(ir.buffer.Bytes()); err != nil {
					logger.Errorf("%s - Failed to Sync File on Read Stream", err)
					compChan <- ir.toResult(false)
				} else {
					compChan <- ir.toResult(true)
				}
			} else {
				logger.Error("Item Read has Completed, unsuccessfully")
				compChan <- ir.toResult(false)
			}
			return
		}
	}
}

// startRead reads from the given Reader and writes to the given Buffer.
func (ir *itemReader) startRead(reader msgio.ReadCloser) {
	// Route Data from Stream
	for i := 0; i < int(ir.size); {
		// Read Next Message
		buf, err := reader.ReadMsg()
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
			i += len(buf)
		}
	}

	// Complete Writing to File
	ir.doneChan <- true
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
	for i := range s.Items() {
		// Initialize Sync Management
		compChan := make(chan itemResult, 1)
		wg.Add(1)
		go s.HandleComplete(ctx, n, &wg, compChan)

		// Create New Writer
		handleItemWrite(s.CreateWriteConfig(i, n, wc), compChan)
		logger.Debugf("Done: Writing Item - %v", i)
	}

	// Wait for all items to be written
	wg.Wait()
	n.OnComplete(s.Event())
}

// handleItemWrite handles the writing of a FileItem to a Stream
func handleItemWrite(config itemConfig, compChan chan itemResult) {
	logger.Debugf("Start: Writing Item - %v", config.index)
	// Create New Writer
	iw := &itemWriter{}
	config.ApplyWriter(iw)

	// Define Chunker Opts
	chunker, err := fs.NewFileChunker(config.Path())
	if err != nil {
		logger.Errorf("%s - Failed to create new chunker.", err)
		iw.doneChan <- false
		compChan <- iw.toResult(false)
		return
	}

	// Write Chunks to Stream
	go iw.startWrite(chunker)

	// Await Progress and Result
	for {
		select {
		case n := <-iw.progressChan:
			iw.written += n
			if ev := iw.getProgressEvent(); ev != nil {
				iw.node.OnProgress(ev)
			}
		case r := <-iw.doneChan:
			logger.Debug("Item Write is Complete")
			compChan <- iw.toResult(r)
			return
		}
	}
}

// startWrite writes the chunks of the given file to the stream
func (iw *itemWriter) startWrite(chunker *fs.Chunker) {
	// Loop through File
	for {
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
	iw.doneChan <- true
}
