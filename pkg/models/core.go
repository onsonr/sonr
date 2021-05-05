package models

import (
	"path/filepath"
	"sync"
	"sync/atomic"
)

const K_BUF_CHUNK = 32000
const K_B64_CHUNK = 31998 // Adjusted for Base64 -- has to be divisible by 3

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
type Progress struct {
	HasMetInterval bool
	ItemComplete   bool
	ItemProgress   float32
	TotalComplete  bool
	TotalProgress  float32
}

type TransferProgress struct {
	CurrentChunk int // Current Chunk Number
	Interval     int // Interval for Callback Progress
	ItemSize     int // Current Item Size
	ItemTotal    int // Current Size of Item
	TransferSize int // Current Size of Transfer
	TotalSize    int // Total Size of Transfer
	TotalCount   int // Total Items in Transfer
}

func NewProgress(totalCount int, totalSize int) *TransferProgress {
	return &TransferProgress{
		Interval:     0,
		CurrentChunk: 0,
		TransferSize: 0,
		ItemTotal:    0,
		TotalCount:   totalCount,
		TotalSize:    totalSize,
	}
}

func (p *TransferProgress) Add(n int) *Progress {
	p.CurrentChunk = p.CurrentChunk + 1
	p.ItemSize = p.ItemSize + n
	p.TransferSize = p.TransferSize + n

	return &Progress{
		HasMetInterval: p.checkInterval(),
		ItemProgress:   float32(p.ItemSize) / float32(p.ItemTotal),
		ItemComplete:   p.ItemTotal >= p.ItemSize,
		TotalProgress:  float32(p.TransferSize) / float32(p.TotalSize),
		TotalComplete:  p.TransferSize >= p.TotalSize,
	}
}

func (p *TransferProgress) Next(c *Chunk) {
	// Calculate Tracking
	itemTotal := int(c.GetTotal())
	itemChunks := itemTotal / K_B64_CHUNK
	interval := itemChunks / 100

	// Update Properties
	p.CurrentChunk = 1
	p.Interval = interval
	p.ItemSize = int(c.Size)
	p.ItemTotal = itemTotal
}

func (p *TransferProgress) checkInterval() bool {
	if p.Interval != 0 {
		return p.CurrentChunk%p.Interval == 0
	}
	return false
}
