package transmit

import (
	"bufio"
	"errors"
	"io"
	"os"

	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/sonr/pkg/motor/config"
	v1 "go.buf.build/grpc/go/sonr-io/core/host/transmit/v1"
	motor "go.buf.build/grpc/go/sonr-io/core/motor/v1"
	types "go.buf.build/grpc/go/sonr-io/core/types/v1"
)

// ReadFromStream reads the item from the stream
func ReadItemFromStream(si *v1.SessionItem, node config.CallbackImpl, reader msgio.ReadCloser) error {
	// Create New File
	dst, err := os.Create(si.GetPath())
	defer dst.Close()
	if err != nil {
		return err
	}

	// Route Data from Stream
	for {
		// Read Next Message
		buf, err := reader.ReadMsg()
		if buf == nil {
			logger.Debug("Completed reading from stream: " + si.GetPath())
			return nil
		}

		if err != nil {
			if err == io.EOF {
				logger.Debug("Completed reading from stream: " + si.GetPath())
				return nil
			}
			return err
		}

		// Write Chunk to File
		n, err := dst.Write(buf)
		if err != nil {
			logger.Errorf("%s - Failed to Write Buffer to File on Read Stream", err)
			return err
		}

		// Update Progress
		if done := ProgressItem(si, n, node); done {
			return nil
		}
	}
}

// WriteToStream writes the item to the stream
func WriteItemToStream(si *v1.SessionItem, node config.CallbackImpl, writer msgio.WriteCloser) error {
	// Create New Chunker
	f, err := os.Open(si.GetPath())
	defer f.Close()
	if err != nil {
		return err
	}

	// Create New Reader
	r := bufio.NewReader(f)
	buf := make([]byte, 0, 4*1024)

	// Loop through File
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				logger.Debug("Completed writing from stream: " + si.GetPath())
				return nil
			}
			return err
		}

		// process buf
		if err != nil && err != io.EOF {
			return err
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(buf)
		if err != nil {
			logger.Errorf("%s - Error Writing data to msgio.Writer", err)
			return err
		}

		// Update Progress
		ProgressItem(si, len(buf), node)
	}
}

// Progress pushes a progress event to the node. Returns true if the item is done.
func ProgressItem(si *v1.SessionItem, wrt int, n config.CallbackImpl) bool {
	// Update Progress
	si.Written += int64(wrt)

	// Create Progress Event
	if (si.GetWritten() % ITEM_INTERVAL) == 0 {
		event := &motor.OnTransmitProgressResponse{
			Direction: si.GetDirection(),
			Progress:  (float64(si.GetWritten()) / float64(si.GetTotalSize())),
			Current:   int32(si.GetIndex()) + 1,
			Total:     int32(si.GetCount()),
		}

		// Push ProgressEvent to Emitter
		go n.OnProgress(event)
	}

	// Return if Done
	return si.GetWritten() >= si.GetSize()
}

// ** ───────────────────────────────────────────────────────
// ** ─── Payload Management ────────────────────────────────
// ** ───────────────────────────────────────────────────────
// PayloadItemFunc is the Map function for PayloadItem
type PayloadItemFunc func(item *types.Payload_Item, index int, total int) error

// IsSingle returns true if the transfer is a single transfer. Error returned
// if No Items present in Payload
func IsSingle(p *types.Payload) (bool, error) {
	if len(p.GetItems()) == 0 {
		return false, errors.New("No Items present in Payload")
	}
	if len(p.GetItems()) > 1 {
		return false, nil
	}
	return true, nil
}

// IsMultiple returns true if the transfer is a multiple transfer. Error returned
// if No Items present in Payload
func IsMultiple(p *types.Payload) (bool, error) {
	if len(p.GetItems()) == 0 {
		return false, errors.New("No Items present in Payload")
	}
	if len(p.GetItems()) > 1 {
		return true, nil
	}
	return false, nil
}

// FileCount returns the number of files in the Payload
func FileCount(p *types.Payload) int {
	// Initialize
	count := 0

	// Iterate over Items
	for _, item := range p.GetItems() {
		// Check if Item is File
		if item.GetMime().Type != types.MIME_TYPE_URL {
			// Increase Count
			count++
		}
	}

	// Return Count
	return count
}

// URLCount returns the number of URLs in the Payload
func URLCount(p *types.Payload) int {
	// Initialize
	count := 0

	// Iterate over Items
	for _, item := range p.GetItems() {
		// Check if Item is File
		if item.GetMime().Type == types.MIME_TYPE_URL {
			// Increase Count
			count++
		}
	}

	// Return Count
	return count
}
