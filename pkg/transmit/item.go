package transmit

import (
	"bytes"
	"io"

	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
	"github.com/sonr-io/core/pkg/common"
)

// itemReader is a Reader for a FileItem
type itemReader struct {
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

// startRead reads from the given Reader and writes to the given Buffer.
func (ir *itemReader) startRead(buffer bytes.Buffer, reader msgio.ReadCloser) {
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
			if err := ir.WriteChunk(buf, buffer); err != nil {
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

// WriteChunk writes a chunk to the buffer
func (ir *itemReader) WriteChunk(b []byte, buffer bytes.Buffer) error {
	n, err := buffer.Write(b)
	if err != nil {
		return err
	}
	ir.progressChan <- n
	return nil
}

// getProgressEvent returns a ProgressEvent for the current ItemReader
func (p *itemReader) getProgressEvent() *api.ProgressEvent {
	if p.written%p.interval == 0 {
		return &api.ProgressEvent{
			Direction: common.Direction_INCOMING,
			Progress:  (float64(p.written) / float64(p.size)),
			Index:     int32(p.index),
			Count:     int32(p.count),
		}
	}
	return nil
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

// startWrite writes the chunks of the given file to the stream
func (iw *itemWriter) startWrite(chunker *fs.Chunker, writer msgio.WriteCloser) {
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
		if err := iw.WriteChunk(c.Data, writer); err != nil {
			logger.Errorf("%s - Error Writing data to msgio.Writer", err)
			iw.doneChan <- false
			return
		}
	}
	iw.doneChan <- true
}

// WriteChunk writes a chunk to the Stream
func (ir *itemWriter) WriteChunk(b []byte, writer msgio.WriteCloser) error {
	err := writer.WriteMsg(b)
	if err != nil {
		return err
	}
	ir.progressChan <- len(b)
	return nil
}

// getProgressEvent returns a ProgressEvent for the current ItemReader
func (p *itemWriter) getProgressEvent() *api.ProgressEvent {
	if p.written%p.interval == 0 {
		return &api.ProgressEvent{
			Direction: common.Direction_OUTGOING,
			Progress:  (float64(p.written) / float64(p.size)),
			Index:     int32(p.index),
			Count:     int32(p.count),
		}
	}
	return nil
}

// toResult returns a FileItemStreamResult for the current ItemReader
func (ir *itemWriter) toResult(success bool) itemResult {
	return itemResult{
		index:     ir.index,
		item:      ir.item.ToTransferItem(),
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
