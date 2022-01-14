package node

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"

	"github.com/pkg/errors"
	"github.com/sonr-io/core/common"
)

// HostStatus is the status of the host
type HostStatus int

// SNRHostStatus Definitions
const (
	Status_IDLE       HostStatus = iota // Host is idle, default state
	Status_STANDBY                      // Host is standby, waiting for connection
	Status_CONNECTING                   // Host is connecting
	Status_READY                        // Host is ready
	Status_FAIL                         // Host failed to connect
	Status_CLOSED                       // Host is closed
)

// Equals returns true if given SNRHostStatus matches this one
func (s HostStatus) Equals(other HostStatus) bool {
	return s == other
}

// IsNotIdle returns true if the SNRHostStatus != Status_IDLE
func (s HostStatus) IsNotIdle() bool {
	return s != Status_IDLE
}

// IsStandby returns true if the SNRHostStatus == Status_STANDBY
func (s HostStatus) IsStandby() bool {
	return s == Status_STANDBY
}

// IsReady returns true if the SNRHostStatus == Status_READY
func (s HostStatus) IsReady() bool {
	return s == Status_READY
}

// IsConnecting returns true if the SNRHostStatus == Status_CONNECTING
func (s HostStatus) IsConnecting() bool {
	return s == Status_CONNECTING
}

// IsFail returns true if the SNRHostStatus == Status_FAIL
func (s HostStatus) IsFail() bool {
	return s == Status_FAIL
}

// IsClosed returns true if the SNRHostStatus == Status_CLOSED
func (s HostStatus) IsClosed() bool {
	return s == Status_CLOSED
}

// String returns the string representation of the SNRHostStatus
func (s HostStatus) String() string {
	switch s {
	case Status_IDLE:
		return "IDLE"
	case Status_STANDBY:
		return "STANDBY"
	case Status_CONNECTING:
		return "CONNECTING"
	case Status_READY:
		return "READY"
	case Status_FAIL:
		return "FAIL"
	case Status_CLOSED:
		return "CLOSED"
	}
	return "UNKNOWN"
}

// SetStatus sets the host status and emits the event
func (h *node) SetStatus(s HostStatus) {
	// Check if status is changed
	if h.status == s {
		return
	}

	// Update Status
	h.status = s
}

// Close closes the node
func (n *node) Close() {
	// Update Status
	n.SetStatus(Status_CLOSED)
	n.IpfsDHT.Close()

	// Close Store
	if err := n.store.Close(); err != nil {
		logger.Errorf("%s - Failed to close store, ", err)
	}

	// Close Host
	if err := n.Host.Close(); err != nil {
		logger.Errorf("%s - Failed to close host, ", err)
	}
}

// NeedsWait checks if state is Resumed or Paused and blocks channel if needed
func (c *node) NeedsWait() {
	<-c.Chn
}

// Resume tells all of goroutines to resume execution
func (c *node) Resume() {
	if atomic.LoadUint64(&c.flag) == 1 {
		close(c.Chn)
		atomic.StoreUint64(&c.flag, 0)
	}
}

// Pause tells all of goroutines to pause execution
func (c *node) Pause() {
	if atomic.LoadUint64(&c.flag) == 0 {
		atomic.StoreUint64(&c.flag, 1)
		c.Chn = make(chan bool)
	}
}

// Filename is base 36 encoded to avoid conflict on case-insensitive filesystems
var maxSockFilename = strconv.FormatUint(math.MaxUint32, 36)
var paddingFormatStr = "%0" + strconv.Itoa(len(maxSockFilename)) + "s"

const UDSDir = "sock"

type SockManager struct {
	sockDirPath string

	counter   uint32
	muCounter sync.Mutex
}

func NewSockManager(path string) (*SockManager, error) {
	abspath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	_, err = os.Stat(abspath)
	if os.IsNotExist(err) {
		return nil, errors.Wrap(err, "sock parent folder doesn't exist")
	}

	sockDirPath := filepath.Join(abspath, UDSDir)

	// max length for a unix domain socket path is around 103 char (108 - '/unix')
	maxSockPath := filepath.Join("/unix", sockDirPath, maxSockFilename)
	if len(maxSockPath) > syscall.SizeofSockaddrAny {
		return nil, errors.New("path length exceeded")
	}

	// remove sock folder from previous app run if exists
	_, err = os.Stat(sockDirPath)
	if !os.IsNotExist(err) {
		err := os.RemoveAll(sockDirPath)
		if err != nil {
			return nil, errors.Wrap(err, "can't cleanup old sock folder")
		}
	}

	if err := os.MkdirAll(sockDirPath, 0700); err != nil {
		return nil, errors.Wrap(err, "can't create sock folder")
	}

	return &SockManager{
		sockDirPath: sockDirPath,
	}, nil
}

func (sm *SockManager) NewSockPath() (string, error) {
	sm.muCounter.Lock()
	if sm.counter == math.MaxUint32 {
		// TODO: do something smarter knowing that a socket may have been
		// removed in the meantime
		sm.muCounter.Unlock()
		return "", errors.New("max number of socket exceeded")
	}
	sockFilename := fmt.Sprintf(paddingFormatStr, strconv.FormatUint(uint64(sm.counter), 36))
	sm.counter++
	sm.muCounter.Unlock()

	sockPath := filepath.Join(sm.sockDirPath, sockFilename)
	_, err := os.Stat(sockPath)
	if os.IsNotExist(err) {
		return sockPath, nil
	} else if err == nil {
		return "", errors.New("sock already exists: " + sockPath)
	}

	return "", errors.Wrap(err, "can't create new sock")
}

// ParseDid converts string into a DID struct
func (n *node) ParseDid(did string) (*common.Did, error) {
	return common.ParseDid(did)
}

// ResolveDid resolves a DID to a Did Document
func (n *node) ResolveDid(did string) (*common.DidDocument, error) {
	doc := &common.DidDocument{}
	return doc, nil
}
