package transmit

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"

	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
)

// Read reads the item from the stream
func (si *SessionItem) Read(doneChan chan bool, node api.NodeImpl, part *multipart.Part) {
	buffer := bytes.Buffer{}
	writer := bufio.NewWriter(&buffer)
	chunker, err := fs.NewChunkerWithAvgSize(part, si.Item.AvgChunkSize())
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return
	}

	for {
		// Read from stream
		chunk, err := chunker.Next()
		if err != nil {
			if err == io.EOF {
				logger.Debug("Completed reading from stream: " + si.GetPath())
				break
			}
			log.Println("Error reading from stream:", err)
			doneChan <- false
			return
		}

		// Write to buffer
		n, err := writer.Write(chunk.Data)
		if err != nil {
			log.Println("Error writing to buffer:", err)
			doneChan <- false
			return
		}

		// Update Progress
		if done := si.Progress(n, node); done {
			break
		}
	}

	// Write File Buffer to File
	err = ioutil.WriteFile(si.GetPath(), buffer.Bytes(), 0644)
	if err != nil {
		logger.Errorf("%s - Failed to Close item on Read Stream", err)
		doneChan <- false
		return
	}
	logger.Debug("Completed reading from stream: " + si.GetPath())
	doneChan <- true
	return
}

// Write writes the item to the stream
func (si *SessionItem) Write(doneChan chan bool, node api.NodeImpl, mwr *multipart.Writer) {
	writer, err := mwr.CreateFormFile(fmt.Sprint(si.GetIndex()), si.GetItem().GetName())
	if err != nil {
		logger.Errorf("%s - Failed to Create Form File on Write Stream", err)
		doneChan <- false
		return
	}

	// Create New Chunker
	chunker, err := fs.NewFileChunker(si.GetPath())
	if err != nil {
		logger.Errorf("%s - Failed to create new chunker.", err)
		doneChan <- false
		return
	}

	// Loop through File
	for {
		c, err := chunker.Next()
		if err != nil {
			logger.Warnf("%s - Failed to get next chunk.", err)
			break
		}

		// Write Message Bytes to Stream
		n, err := writer.Write(c.Data)
		if err != nil {
			logger.Errorf("%s - Error Writing data to msgio.Writer", err)
			doneChan <- false
			return
		}

		// Update Progress
		if done := si.Progress(n, node); done {
			break
		}
	}
	logger.Debug("Completed writing to stream: " + si.GetPath())
	doneChan <- true
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
