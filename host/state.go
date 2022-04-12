package host

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"

	"github.com/kataras/golog"
	dsc "github.com/libp2p/go-libp2p-discovery"
	psub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/device"
	"github.com/sonr-io/core/highway/config"
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
func (h *hostImpl) SetStatus(s HostStatus) {
	// Check if status is changed
	if h.status == s {
		return
	}

	// Update Status
	h.status = s
}

// Close closes the node
func (n *hostImpl) Close() {
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
func (c *hostImpl) NeedsWait() {
	<-c.Chn
}

// Resume tells all of goroutines to resume execution
func (c *hostImpl) Resume() {
	if atomic.LoadUint64(&c.flag) == 1 {
		close(c.Chn)
		atomic.StoreUint64(&c.flag, 0)
	}
}

// Pause tells all of goroutines to pause execution
func (c *hostImpl) Pause() {
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

// persist contains the main loop for the Node
func (n *hostImpl) Persist() {
	// Check if node is highway
	if n.Role() != device.Role_HIGHWAY {
		golog.Default.Child("(app)").Errorf("%s - Persist: Node is not a highway node", n.HostID())
		return
	}

	golog.Default.Child("(app)").Infof("Starting GRPC Server on %s", n.listener.Addr().String())

	// Wait for Exit Signal
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		exit(0, n.ctx)
	}()

	// Hold until Exit Signal
	for {
		select {
		case <-n.ctx.Done():
			golog.Default.Child("(app)").Info("Context Done")
			n.listener.Close()
			return
		}
	}
}

// Exit handles cleanup on Sonr Node
func exit(code int, ctx context.Context) {
	golog.Default.Child("(app)").Debug("Cleaning up Node on Exit...")
	defer ctx.Done()

}

// createDHTDiscovery is a Helper Method to initialize the DHT Discovery
func (hn *hostImpl) createDHTDiscovery(c *config.Config) error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := dsc.NewRoutingDiscovery(hn.IpfsDHT)
	dsc.Advertise(hn.ctx, routingDiscovery, c.Libp2pRendezvous, c.Libp2pTTL)

	// Create Pub Sub
	hn.PubSub, err = psub.NewGossipSub(hn.ctx, hn.Host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		hn.SetStatus(Status_FAIL)
		logger.Errorf("%s - Failed to Create new Gossip Sub", err)
		return err
	}

	// Handle DHT Peers
	hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, c.Libp2pRendezvous, c.Libp2pTTL)
	if err != nil {
		hn.SetStatus(Status_FAIL)
		logger.Errorf("%s - Failed to create FindPeers Discovery channel", err)
		return err
	}
	hn.SetStatus(Status_READY)
	return nil
}

// createMdnsDiscovery is a Helper Method to initialize the MDNS Discovery
func (hn *hostImpl) createMdnsDiscovery(c *config.Config) {
	if hn.Role() == device.Role_MOTOR {
		// Create MDNS Service
		ser := mdns.NewMdnsService(hn.Host, c.CosmosKeyringServiceName, hn)

		ser.Start()
		// Handle Events
		// ser.RegisterNotifee(hn)
	}
}
