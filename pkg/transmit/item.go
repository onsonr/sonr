package transmit

import (
	"bytes"
	"io"
	"os"

	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
	"github.com/sonr-io/core/pkg/common"
)

// itemReader is a Reader for a FileItem
type itemReader struct {
	item         *common.FileItem
	buffer       bytes.Buffer
	interval     int
	index        int
	count        int
	size         int64
	node         api.NodeImpl
	path         string
	written      int
	progressChan chan int
	doneChan     chan bool
}

// handleRead reads from the given Reader and writes to the given Buffer.
func (ir *itemReader) handleRead(reader msgio.ReadCloser) {
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
		}

		// Write Chunk to Buffer
		n, err := ir.buffer.Write(buf)
		if err != nil {
			logger.Errorf("%s - Failed to Write Chunk to Buffer Read Stream", err)
			ir.doneChan <- false
			return
		}
		i += n
		ir.progressChan <- n
	}

	// Flush Buffer to File
	data := ir.buffer.Bytes()
	err := os.WriteFile(ir.path, data, 0644)
	if err != nil {
		logger.Errorf("%s - Failed to Flush Buffer to File on Read Stream", err)
		ir.doneChan <- false
		return
	}

	// Complete Writing to File
	ir.doneChan <- true
}

// getProgressEvent returns a ProgressEvent for the current ItemReader
func (ir *itemReader) getProgressEvent() *api.ProgressEvent {
	return &api.ProgressEvent{
		Direction: common.Direction_INCOMING,
		Progress:  (float64(ir.written) / float64(ir.size)),
		Index:     int32(ir.index),
		Count:     int32(ir.count),
	}
}

// toResult returns a FileItemStreamResult for the current ItemReader
func (ir *itemReader) toResult(success bool) itemResult {
	return itemResult{
		index:     ir.index,
		item:      ir.item.ToTransferItem(),
		direction: common.Direction_INCOMING,
		success:   success,
	}
}

// itemWriter is a Writer for FileItems
type itemWriter struct {
	chunker      *fs.Chunker
	item         *common.FileItem
	interval     int
	index        int
	count        int
	size         int64
	node         api.NodeImpl
	written      int
	progressChan chan int
	doneChan     chan bool
}

// handleWrite writes the chunks of the given file to the stream
func (iw *itemWriter) handleWrite(writer msgio.WriteCloser) {
	// Loop through File
	for {
		c, err := iw.chunker.Next()
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
		err = writer.WriteMsg(c.Data)
		if err != nil {
			logger.Errorf("%s - Error Writing data to msgio.Writer", err)
			iw.doneChan <- false
			return
		}
		iw.progressChan <- c.Length
	}
	iw.doneChan <- true
}

// getProgressEvent returns a ProgressEvent for the current ItemReader
func (iw *itemWriter) getProgressEvent() *api.ProgressEvent {
	return &api.ProgressEvent{
		Direction: common.Direction_OUTGOING,
		Progress:  (float64(iw.written) / float64(iw.size)),
		Index:     int32(iw.index),
		Count:     int32(iw.count),
	}
}

// toResult returns a FileItemStreamResult for the current ItemReader
func (iw *itemWriter) toResult(success bool) itemResult {
	return itemResult{
		index:     iw.index,
		item:      iw.item.ToTransferItem(),
		direction: common.Direction_OUTGOING,
		success:   success,
	}
}

// itemResult is the result of a FileItemStream
type itemResult struct {
	index     int
	direction common.Direction
	item      *common.Payload_Item
	success   bool
}

// IsIncoming returns true if the item is incoming
func (r itemResult) IsIncoming() bool {
	return r.direction == common.Direction_INCOMING && r.success
}

// IsOutgoing returns true if the item is outgoing
func (r itemResult) IsOutgoing() bool {
	return r.direction == common.Direction_OUTGOING && r.success
}
