package transmit

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"sync"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
	"github.com/sonr-io/core/pkg/common"
)

// Transfer Protocol ID's
const (
	IncomingPID   protocol.ID = "/transmit/incoming/0.0.1"
	OutgoingPID   protocol.ID = "/transmit/outgoing/0.0.1"
	ITEM_INTERVAL             = 25
)

// Error Definitions
var (
	logger        = golog.Default.Child("protocols/transmit")
	ErrNoSession  = errors.New("Failed to get current session, set to nil")
	ErrFailedAuth = errors.New("Failed to Authenticate message")
)

// calculateInterval calculates the interval for the progress callback
func calculateInterval(size int64) int {
	// Calculate Interval
	interval := size / 100
	if interval < 1 {
		interval = 1
	}
	logger.Debugf("Calculated Item progress interval: %v", interval)
	return int(interval)
}

// pushProgress pushes a progress event to the node
func pushProgress(n api.NodeImpl, dir common.Direction, written int, size int64, index, count int) {
	// Create Progress Event
	if (written % ITEM_INTERVAL) == 0 {
		event := &api.ProgressEvent{
			Direction: dir,
			Progress:  (float64(written) / float64(size)),
			Index:     int32(index),
			Count:     int32(count),
		}

		// Push ProgressEvent to Emitter
		n.OnProgress(event)
	}
}

// CurrentSession returns the current session
func (p *TransmitProtocol) CurrentSession() (*Session, error) {
	if p.current != nil {
		return p.current, nil
	}
	return nil, ErrNoSession
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

// NewItemWriter Returns a new Reader for the given FileItem
func WriteItem(index int, count int, pi *common.Payload_Item, wg *sync.WaitGroup, node api.NodeImpl, writer msgio.WriteCloser) {
	// Properties
	defer wg.Done()
	size := pi.GetSize()
	item := pi.GetFile()

	// Create New Chunker
	chunker, err := fs.NewFileChunker(item.Path)
	if err != nil {
		logger.Errorf("%s - Failed to create new chunker.", err)
		return
	}

	// Loop through File
	for i := 0; i < int(size); {
		c, err := chunker.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Errorf("%s - Error reading chunk.", err)
			return
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(c.Data)
		if err != nil {
			logger.Errorf("%s - Error Writing data to msgio.Writer", err)
			return
		}

		// Unexpected Error
		if err != nil && err != io.EOF {
			logger.Errorf("%s - Unexpected Error occurred on Write Stream", err)
			return
		}
		// Update Progress
		i += c.Length

		// Update Progress
		pushProgress(node, common.Direction_OUTGOING, i, size, index, count)
	}
	return
}
