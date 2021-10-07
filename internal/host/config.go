package host

import (
	"crypto/rand"
	"errors"
	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/multiformats/go-multiaddr"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/keychain"
	net "github.com/sonr-io/core/tools/internet"
)

var (
	logger = golog.Child("host")
)

// HostOption is a function that modifies the node options.
type HostOption func(hostOptions)

// WithBootstrappers sets the bootstrap peers.
func WithBootstrappers(pis []peer.AddrInfo) HostOption {
	return func(o hostOptions) {
		o.bootstrapPeers = pis
	}
}

// WithConnection sets the connection to the host.
func WithConnection(c common.Connection) HostOption {
	return func(o hostOptions) {
		o.connection = c
	}
}

// WithConnOptions sets the connection manager options.
func WithConnOptions(low int, hi int, grace time.Duration) HostOption {
	return func(o hostOptions) {
		o.lowWater = low
		o.highWater = hi
		o.gracePeriod = grace
	}
}

// WithInterval sets the interval for the host.
func WithInterval(interval time.Duration) HostOption {
	return func(o hostOptions) {
		o.interval = interval
	}
}

// TCPListener sets the listener for the host. (This is Required for the
// host to start)
func TCPListener(l *net.TCPListener) (HostOption, error) {
	// Get the address.
	ma, err := l.Multiaddr()
	if err != nil {
		return nil, err
	}

	// Return the option.
	return func(o hostOptions) {
		o.setAddr = true
		o.multiAddrs = []multiaddr.Multiaddr{ma}
	}, nil
}

// WithPrivKey sets the private key for the host.
func WithPrivKey(pk crypto.PrivKey) HostOption {
	return func(o hostOptions) {
		o.privateKey = pk
	}
}

// WithTTL sets the ttl for the host.
func WithTTL(ttl time.Duration) HostOption {
	return func(o hostOptions) {
		o.ttl = ttl
	}
}

// hostOptions is a collection of options for the SnrHost.
type hostOptions struct {
	setAddr        bool
	connection     common.Connection
	bootstrapPeers []peer.AddrInfo
	multiAddrs     []multiaddr.Multiaddr
	lowWater       int
	highWater      int
	gracePeriod    time.Duration
	privateKey     crypto.PrivKey
	rendezvous     string
	interval       time.Duration
	ttl            time.Duration
}

// defaultHostOptions returns the default host options.
func defaultHostOptions() hostOptions {
	return hostOptions{
		setAddr:        false,
		connection:     common.Connection_WIFI,
		bootstrapPeers: dht.GetDefaultBootstrapPeerAddrInfos(),
		lowWater:       10,
		highWater:      15,
		gracePeriod:    time.Second * 5,
		rendezvous:     "/sonr/rendevouz/0.9.2",
		interval:       time.Second * 5,
		ttl:            time.Minute * 1,
	}
}

// Apply applies the host options
func (ho hostOptions) Apply(options ...HostOption) error {
	// Iterate over the options.
	for _, opt := range options {
		opt(ho)
	}

	// Check if Listener is set.
	if !ho.setAddr {
		logger.Error("Failed to create host, invalid hostOptions")
		return errors.New("Listener Address was not set")
	}

	// findPrivKey returns the private key for the host.
	findPrivKey := func() (crypto.PrivKey, error) {
		privKey, err := device.KeyChain.GetPrivKey(keychain.Account)
		if err == nil {
			return privKey, nil
		}
		logger.Warn("Failed to get Account Private Key for Host", err)
		privKey, _, err = crypto.GenerateEd25519Key(rand.Reader)
		if err == nil {
			return privKey, nil
		}
		return nil, err
	}

	// Fetch the private key.
	privKey, err := findPrivKey()
	if err != nil {
		logger.Error("Failed to create host, invalid hostOptions")
		return err
	}

	// Set the private key.
	ho.privateKey = privKey
	return nil
}
