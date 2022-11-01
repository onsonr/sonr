package host

import (
	"context"
	"crypto/ed25519"
	"strings"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/sonr/pkg/config"
	ct "github.com/sonr-io/sonr/third_party/types/common"
)

// SonrHost returns the SonrHost for the Main Node
type SonrHost interface {
	// AuthenticateMessage authenticates a message
	// AuthenticateMessage(msg proto.Message, metadata *types.Metadata) error
	// Address returns the account address of the underlying wallet for the host
	Address() string

	// Config returns the configuration of the node
	Config() *config.Config

	// Connect to a peer
	Connect(pi peer.AddrInfo) error

	// HasRouting returns true if the node has routing
	HasRouting() bool

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

	// Ping sends a ping to a peer to check if it is alive
	Ping(id string) error

	// TODO: implement
	// Peer returns the peer of the node
	Peer() (*ct.Peer, error)

	// SignData signs the data with the private key
	SignData(data []byte) ([]byte, error)
	// SignMessage signs a message with the node's private key
	// SignMessage(message proto.Message) ([]byte, error)

	// Pubsub returns the pubsub of the node
	Pubsub() *ps.PubSub

	// Role returns the role of the node
	Role() config.Role

	// Router returns the routing.Router
	Router(h host.Host) (routing.PeerRouting, error)

	// Routing returns the routing.Routing
	Routing() routing.Routing

	// PrivateKey returns the ed25519 private key instance of the libp2p host
	PrivateKey() (ed25519.PrivateKey, error)

	// Send(ctx context.Context, target string, data interface{}, protocol protocol.ID) error

	// SendMessage sends a message to a peer
	Send(id peer.ID, p protocol.ID, data []byte) error

	// SetStreamHandler sets the handler for a protocol
	SetStreamHandler(protocol protocol.ID, handler network.StreamHandler)

	// VerifyData verifies the data signature
	VerifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) error

	// Close closes the node
	Close() error

	Start() error

	Stop() error

	// Pauses tells all of goroutines to pause execution
	Pause() error

	// Resume tells all of goroutines to resume execution
	Resume() error

	Status() HostStatus
}

func addrToDidUrl(addr string) string {
	if strings.Contains(addr, "snr") {
		rawAddr := strings.TrimLeft(addr, "snr")
		return "did:snr:" + rawAddr
	}
	return addr
}
