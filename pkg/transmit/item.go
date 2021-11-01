package transmit

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"math"
	"mime/multipart"
	"os"

	"github.com/sonr-io/core/internal/api"
)

// progressReader wraps an existing io.Reader.
//
// It simply forwards the Read() call, while displaying
// the results from individual calls to it.
type progressReader struct {
	io.Reader
	node  api.NodeImpl
	total int64 // Total # of bytes transferred
	item  *SessionItem
}

// initProgressReader creates a new ProgressReader
func (si *SessionItem) initProgressReader(r io.Reader, node api.NodeImpl) *progressReader {
	return &progressReader{r, node, 0, si}
}

// Read 'overrides' the underlying io.Reader's Read method.
// This is the one that will be called by io.Copy(). We simply
// use it to keep track of byte counts and then forward the call.
func (pr *progressReader) Read(p []byte) (int, error) {
	n, err := pr.Reader.Read(p)
	pr.total += int64(n)

	if err == nil {
		pr.item.Progress(n, pr.node)
	}
	return n, err
}

// Read reads the item from the stream
func (si *SessionItem) Read(doneChan chan bool, node api.NodeImpl, part *multipart.Part) {
	buffer := bytes.Buffer{}
	dst := bufio.NewWriter(&buffer)
	src := si.initProgressReader(part, node)

	// Copy from src to dst
	n, err := io.Copy(dst, src)
	if err != nil {
		doneChan <- false
		return
	}
	logger.Debug("Item Completed Read \n\tPath: %s \n\tSize: %v"+si.GetPath(), n)

	// Write to File
	err = ioutil.WriteFile(si.GetPath(), buffer.Bytes(), 0644)
	if err != nil {
		doneChan <- false
		return
	}

	// Update Progress
	doneChan <- true
	return
}

// ProgressReader wraps an existing io.Reader.
//
// It simply forwards the Read() call, while displaying
// the results from individual calls to it.
type progressWriter struct {
	io.Writer
	node  api.NodeImpl
	total int64 // Total # of bytes transferred
	item  *SessionItem
}

// initProgressWriter creates a new ProgressWriter
func (si *SessionItem) initProgressWriter(wr io.Writer, node api.NodeImpl) *progressWriter {
	return &progressWriter{wr, node, 0, si}
}

// Read 'overrides' the underlying io.Reader's Read method.
// This is the one that will be called by io.Copy(). We simply
// use it to keep track of byte counts and then forward the call.
func (pr *progressWriter) Write(p []byte) (int, error) {
	n, err := pr.Writer.Write(p)
	pr.total += int64(n)

	if err == nil {
		pr.item.Progress(n, pr.node)
	}

	return n, err
}

// Write writes the item to the stream
func (si *SessionItem) Write(doneChan chan bool, node api.NodeImpl, wr io.Writer) {
	// Create New Chunker
	data, err := os.ReadFile(si.GetPath())
	if err != nil {
		logger.Errorf("%s - Failed to read file data in Item writer")
		doneChan <- false
		return
	}

	// Create Source and Destination
	buffer := bytes.NewBuffer(data)
	src := bufio.NewReader(buffer)
	dst := si.initProgressWriter(wr, node)

	// Copy from src to dst
	n, err := io.Copy(dst, src)
	if err != nil {
		logger.Errorf("%s - Failed to copy from src to dst in Item writer")
		doneChan <- false
		return
	}
	logger.Debug("Item Completed Write \n\tPath: %s \n\tSize: %v"+si.GetPath(), n)

	// Update Progress
	doneChan <- true
	return
}

// Progress pushes a progress event to the node. Returns true if the item is done.
func (si *SessionItem) Progress(wrt int, n api.NodeImpl) {
	// Update Progress
	si.Written += int64(wrt)

	// Create Progress Event
	val := si.GetWritten() % ITEM_INTERVAL
	if math.Floor(float64(val)) == 0 {
		event := &api.ProgressEvent{
			Direction: si.GetDirection(),
			Progress:  (float64(si.GetWritten()) / float64(si.GetTotalSize())),
			Current:   int32(si.GetIndex()) + 1,
			Total:     int32(si.GetCount()),
		}

		// Push ProgressEvent to Emitter
		n.OnProgress(event)
	}
}
