package host

import (
	"context"
	"crypto/ed25519"
	"sync/atomic"

	"github.com/kataras/go-events"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/sonr/pkg/config"

	types "go.buf.build/grpc/go/sonr-io/core/types/v1"
	"google.golang.org/protobuf/proto"
)

// SonrHost returns the SonrHost for the Main Node
type SonrHost interface {
	// AuthenticateMessage authenticates a message
	AuthenticateMessage(msg proto.Message, metadata *types.Metadata) bool

	// Close closes the node
	Close()

	// Config returns the configuration of the node
	Config() *config.Config

	// Connect to a peer
	Connect(pi peer.AddrInfo) error

	// Events returns the events manager of the node
	Events() events.EventEmmiter

	// HasRouting returns true if the node has routing
	HasRouting() error

	// Host returns the Host
	Host() host.Host

	// HostID returns the ID of the Host
	HostID() peer.ID

	// Join subsrcibes to a topic
	Join(topic string, opts ...ps.TopicOpt) (*ps.Topic, error)

	// NewStream opens a new stream to a peer
	NewStream(ctx context.Context, pid peer.ID, pids ...protocol.ID) (network.Stream, error)

	// NewTopic creates a new pubsub topic with event handler and subscription
	NewTopic(topic string, opts ...ps.TopicOpt) (*ps.Topic, *ps.TopicEventHandler, *ps.Subscription, error)

	// NeedsWait checks if state is Resumed or Paused and blocks channel if needed
	NeedsWait()

	// Pause tells all of goroutines to pause execution
	Pause()

	// Ping sends a ping to a peer to check if it is alive
	Ping(id string) error

	// Peer returns the peer of the node
	Peer() (*types.Peer, error)

	// Pubsub returns the pubsub of the node
	Pubsub() *ps.PubSub

	// Resume tells all of goroutines to resume execution
	Resume()

	// Role returns the role of the node
	Role() config.Role

	// Router returns the routing.Router
	Router(h host.Host) (routing.PeerRouting, error)

	// Routing returns the routing.Routing
	Routing() routing.Routing

	// PrivateKey returns the ed25519 private key instance of the libp2p host
	PrivateKey() (ed25519.PrivateKey, error)

	// SendMessage sends a message to a peer
	SendMessage(id peer.ID, p protocol.ID, data proto.Message) error

	// SetStreamHandler sets the handler for a protocol
	SetStreamHandler(protocol protocol.ID, handler network.StreamHandler)

	// SignData signs the data with the private key
	SignData(data []byte) ([]byte, error)

	// SignMessage signs a message with the node's private key
	SignMessage(message proto.Message) ([]byte, error)

	// VerifyData verifies the data signature
	VerifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool
}

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
	n.host.Close()
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
