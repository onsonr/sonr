package transmit

import (
	"bufio"
	"errors"
	"io"
	"os"

	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/sonr/pkg/fs"
	"github.com/sonr-io/sonr/pkg/host"
	motor "go.buf.build/grpc/go/sonr-io/motor/common/v1"
	v1 "go.buf.build/grpc/go/sonr-io/motor/service/v1"
)

// ReadFromStream reads the item from the stream
func ReadItemFromStream(si *v1.SessionItem, node host.SonrHost, reader msgio.ReadCloser) error {
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
func WriteItemToStream(si *v1.SessionItem, h host.SonrHost, writer msgio.WriteCloser) error {
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
		ProgressItem(si, len(buf), h)
	}
}

// Progress pushes a progress event to the node. Returns true if the item is done.
func ProgressItem(si *v1.SessionItem, wrt int, h host.SonrHost) bool {
	// Update Progress
	si.Written += int64(wrt)

	// Create Progress Event
	if (si.GetWritten() % ITEM_INTERVAL) == 0 {
		// event := &motor.OnTransmitProgressResponse{
		// 	Direction: si.GetDirection(),
		// 	Progress:  (float64(si.GetWritten()) / float64(si.GetTotalSize())),
		// 	Current:   int32(si.GetIndex()) + 1,
		// 	Total:     int32(si.GetCount()),
		// }

		// Push ProgressEvent to Emitter
		// h.Events().Emit(t.ON_PROGRESS, event)
	}

	// Return if Done
	return si.GetWritten() >= si.GetSize()
}

// ** ───────────────────────────────────────────────────────
// ** ─── Payload Management ────────────────────────────────
// ** ───────────────────────────────────────────────────────
// PayloadItemFunc is the Map function for PayloadItem
type PayloadItemFunc func(item *v1.Payload_Item, index int, total int) error

// IsSingle returns true if the transfer is a single transfer. Error returned
// if No Items present in Payload
func IsSingle(p *v1.Payload) (bool, error) {
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
func IsMultiple(p *v1.Payload) (bool, error) {
	if len(p.GetItems()) == 0 {
		return false, errors.New("No Items present in Payload")
	}
	if len(p.GetItems()) > 1 {
		return true, nil
	}
	return false, nil
}

// FileCount returns the number of files in the Payload
func FileCount(p *v1.Payload) int {
	// Initialize
	count := 0

	// Iterate over Items
	for _, item := range p.GetItems() {
		// Check if Item is File
		if item.GetMime().Type != motor.MIME_TYPE_URL {
			// Increase Count
			count++
		}
	}

	// Return Count
	return count
}

// URLCount returns the number of URLs in the Payload
func URLCount(p *v1.Payload) int {
	// Initialize
	count := 0

	// Iterate over Items
	for _, item := range p.GetItems() {
		// Check if Item is File
		if item.GetMime().Type == motor.MIME_TYPE_URL {
			// Increase Count
			count++
		}
	}

	// Return Count
	return count
}

// SetPathFromFolder sets the path of the FileItem
func SetPathFromFolder(f *v1.FileItem, folder fs.Folder) (string, error) {
	// Set Path
	oldPath := f.GetPath()

	// generate path
	path, err := folder.GenPath(oldPath)
	if err != nil {
		return "", err
	}

	// Define Check Path Function
	f.Path = path
	return f.Path, nil
}
