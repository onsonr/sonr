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

// ** ─── State MANAGEMENT ────────────────────────────────────────────────────────
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

// ** ─── Directories MANAGEMENT ────────────────────────────────────────────────────────
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

// ** ─── PROGRESS MANAGEMENT ────────────────────────────────────────────────────────
// Returned Progress Struct
type Progress struct {
	HasMetInterval bool
	ItemComplete   bool
	ItemProgress   float32
}

// Individual Item Progress Struct
type ItemProgress struct {
	CurrentChunk int // Current Chunk Number
	Interval     int // Interval for Callback Progress
	ItemSize     int // Current Item Size
	ItemTotal    int // Current Size of Item
}

// @ NewProgress: Creates new ItemProgress Instance
func NewProgress(total int32) *ItemProgress {
	return &ItemProgress{
		Interval:     0,
		CurrentChunk: 0,
		ItemTotal:    int(total),
	}
}

// @ Add: Updates current chunk and written bytes
func (p *ItemProgress) Add(curr int, n int) *Progress {
	p.CurrentChunk = curr
	p.ItemSize = p.ItemSize + n

	return &Progress{
		HasMetInterval: p.checkInterval(),
		ItemProgress:   p.getProgress(),
		ItemComplete:   p.checkItemComplete(),
	}
}

// @ Set: Sets the first chunk from stream to initialize tracking
func (p *ItemProgress) Set(c *Chunk) {
	// Calculate Tracking
	itemChunks := p.ItemTotal / K_CHUNK_SIZE
	interval := itemChunks / 100

	// Update Properties
	p.CurrentChunk = 0
	p.Interval = interval
	p.ItemSize = int(c.Size)
}

// # Helper: Checks wether callback Interval has been met
func (p *ItemProgress) checkInterval() bool {
	if p.Interval != 0 {
		return p.CurrentChunk%p.Interval == 0
	}
	return false
}

// # Helper: Checks wether item has completed transfer
func (p *ItemProgress) checkItemComplete() bool {
	if p.ItemTotal < p.ItemSize {
		return false
	}
	return true
}

func (p *ItemProgress) getProgress() float32 {
	return float32(p.ItemSize) / float32(p.ItemTotal)
}
