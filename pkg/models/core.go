package models

import (
	"path/filepath"
	"sync"
	"sync/atomic"
)

// ** ─── CALLBACK MANAGEMENT ────────────────────────────────────────────────────────
// Define Function Types
type GetStatus func() Status
type SetStatus func(s Status)
type GetContact func() *Contact
type OnProtobuf func([]byte)
type OnInvite func(data []byte)
type OnProgress func(data float32)
type OnReceived func(data *TransferCard)
type OnTransmitted func(data *TransferCard)
type OnError func(err *SonrError)
type NodeCallback struct {
	Contact     GetContact
	Invited     OnInvite
	Refreshed   OnProtobuf
	Event       OnProtobuf
	RemoteStart OnProtobuf
	Responded   OnProtobuf
	Progressed  OnProgress
	Received    OnReceived
	Status      SetStatus
	Transmitted OnTransmitted
	Error       OnError
	GetStatus   GetStatus
}

// @ Binary State Management ** //
type state struct {
	flag uint64
	chn  chan bool
}

var (
	instance *state
	once     sync.Once
)

func GetState() *state {
	once.Do(func() {
		chn := make(chan bool)
		close(chn)

		instance = &state{chn: chn}
	})

	return instance
}

// Checks rather to wait or does not need
func (c *state) NeedsWait() {
	<-c.chn
}

// Says all of goroutines to resume execution
func (c *state) Resume() {
	if atomic.LoadUint64(&c.flag) == 1 {
		close(c.chn)
		atomic.StoreUint64(&c.flag, 0)
	}
}

// Says all of goroutines to pause execution
func (c *state) Pause() {
	if atomic.LoadUint64(&c.flag) == 0 {
		atomic.StoreUint64(&c.flag, 1)
		c.chn = make(chan bool)
	}
}

// Returns Path for Application/User Data
func (d *Directories) DataSavePath(fileName string, IsDesktop bool) string {
	// Check for Desktop
	if IsDesktop {
		return filepath.Join(d.GetLibrary(), fileName)
	} else {
		return filepath.Join(d.GetSupport(), fileName)
	}
}

// Returns Path for Transferred File
func (d *Directories) TransferSavePath(fileName string, mime *MIME, IsDesktop bool) string {
	// @ Check for Media
	if mime.IsMedia() {
		// Check for Desktop
		if IsDesktop {
			return filepath.Join(d.GetDownloads(), fileName)
		} else {
			return filepath.Join(d.GetTemporary(), fileName)
		}
	} else {
		// Check for Desktop
		if IsDesktop {
			return filepath.Join(d.GetDownloads(), fileName)
		} else {
			return filepath.Join(d.GetDocuments(), fileName)
		}
	}
}
