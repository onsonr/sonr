package transmit

import (
	"bytes"
	"io/ioutil"

	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/pkg/common"
)

// Read reads the item from the stream
func (si *SessionItem) Read(node api.NodeImpl, reader msgio.ReadCloser) error {
	buffer := bytes.Buffer{}

	// Route Data from Stream
	for {
		// Read Next Message
		buf, err := reader.ReadMsg()
		if err != nil {
			logger.Warnf("%s - Failed to Read Next Message on Read Stream", err)
			break
		}

		// Write Chunk to File
		n, err := buffer.Write(buf)
		if err != nil {
			logger.Errorf("%s - Failed to Write Buffer to File on Read Stream", err)

			return err
		}

		// Update Progress
		if done := si.Progress(n, node); done {
			break
		}
	}

	// Write File Buffer to File
	err := ioutil.WriteFile(si.GetPath(), buffer.Bytes(), 0644)
	if err != nil {
		logger.Errorf("%s - Failed to Close item on Read Stream", err)
		return err
	}
	logger.Debug("Completed reading from stream: " + si.GetPath())
	return nil
}

// Write writes the item to the stream
func (si *SessionItem) Write(node api.NodeImpl, writer msgio.WriteCloser) error {
	// Create New Chunker
	chunker, err := common.NewFileChunker(si.GetPath())
	if err != nil {
		logger.Errorf("%s - Failed to create new chunker.", err)
		return err
	}

	// Loop through File
	for {
		c, err := chunker.Next()
		if err != nil {
			logger.Warnf("%s - Failed to get next chunk.", err)
			break
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(c.Data)
		if err != nil {
			logger.Errorf("%s - Error Writing data to msgio.Writer", err)
			return err
		}

		// Update Progress
		if done := si.Progress(c.Length, node); done {
			break
		}
	}
	logger.Debug("Completed writing to stream: " + si.GetPath())
	return nil
}

// Progress pushes a progress event to the node. Returns true if the item is done.
func (si *SessionItem) Progress(wrt int, n api.NodeImpl) bool {
	// Update Progress
	si.Written += int64(wrt)

	// Create Progress Event
	if (si.GetWritten() % ITEM_INTERVAL) == 0 {
		event := &api.ProgressEvent{
			Direction: si.GetDirection(),
			Progress:  (float64(si.GetWritten()) / float64(si.GetTotalSize())),
			Current:   int32(si.GetIndex()) + 1,
			Total:     int32(si.GetCount()),
		}

		// Push ProgressEvent to Emitter
		n.OnProgress(event)
	}
	return si.Written >= si.TotalSize
}
