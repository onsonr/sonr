package transmit

import (
	"bytes"
	"io"
	"io/ioutil"
	"sync"

	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
)

// Read reads the item from the stream
func (si *SessionItem) Read(wg *sync.WaitGroup, node api.NodeImpl, reader msgio.ReadCloser) {
	defer wg.Done()
	// generate path
	item := si.GetItem()
	size := item.GetSize()
	path, err := item.ResetPath(fs.Downloads)
	if err != nil {
		logger.Errorf("%s - Failed to generate path for file: %s", err, item.Name)
		return
	}

	buffer := bytes.Buffer{}

	// Route Data from Stream
	for int(si.Written) < int(size) {
		// Read Next Message
		buf, err := reader.ReadMsg()
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Errorf("%s - Failed to Read Next Message on Read Stream", err)
			return
		}

		// Write Chunk to File
		n, err := buffer.Write(buf)
		if err != nil {
			logger.Errorf("%s - Failed to Write Buffer to File on Read Stream", err)
			return
		}

		// Update Progress
		if done := si.Progress(n, node); done {
			break
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

// Write writes the item to the stream
func (si *SessionItem) Write(wg *sync.WaitGroup, node api.NodeImpl, writer msgio.WriteCloser) {
	// Properties
	defer wg.Done()
	item := si.GetItem()

	// Create New Chunker
	chunker, err := fs.NewFileChunker(item.Path)
	if err != nil {
		logger.Errorf("%s - Failed to create new chunker.", err)
		return
	}

	// Loop through File
	for {
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
		if done := si.Progress(c.Length, node); done {
			break
		}
	}
	return
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
