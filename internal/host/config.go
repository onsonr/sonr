package host

import (
	"context"
	"crypto/rand"

	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/tools/internet"
	net "github.com/sonr-io/core/tools/internet"
	"github.com/sonr-io/core/tools/state"
)


// Error Definitions
var (
	logger              = golog.Child("internal/host")
	ErrRoutingNotSet    = errors.New("DHT and Host have not been set by Routing Function")
	ErrListenerRequired = errors.New("Listener was not Provided")
	ErrMDNSInvalidConn  = errors.New("Invalid Connection, cannot begin MDNS Service")
)

// HostOption is a function that modifies the node options.
type HostOption func(hostOptions)

// WithBootstrappers sets the bootstrap peers.
func WithBootstrappers(pis []peer.AddrInfo) HostOption {
	return func(o hostOptions) {
		o.BootstrapPeers = pis
	}
}

// WithConnection sets the connection to the host.
func WithConnection(c common.Connection) HostOption {
	return func(o hostOptions) {
		o.Connection = c
	}
}

// WithConnOptions sets the connection manager options.
func WithConnOptions(low int, hi int, grace time.Duration) HostOption {
	return func(o hostOptions) {
		o.LowWater = low
		o.HighWater = hi
		o.GracePeriod = grace
	}
}

// WithInterval sets the interval for the host.
func WithInterval(interval time.Duration) HostOption {
	return func(o hostOptions) {
		o.Interval = interval
	}
}

// WithTTL sets the ttl for the host.
func WithTTL(ttl time.Duration) HostOption {
	return func(o hostOptions) {
		o.TTL = ttl
	}
}

// hostOptions is a collection of options for the SnrHost.
type hostOptions struct {
	// Properties
	Connection     common.Connection
	BootstrapPeers []peer.AddrInfo
	LowWater       int
	HighWater      int
	GracePeriod    time.Duration
	Rendezvous     string
	Interval       time.Duration
	TTL            time.Duration

	// Parameters
	ctx      context.Context
	listener *net.TCPListener
}

// defaultHostOptions returns the default host options.
func defaultHostOptions(ctx context.Context, l *internet.TCPListener) hostOptions {
	return hostOptions{
		ctx:            ctx,
		Connection:     common.Connection_WIFI,
		BootstrapPeers: dht.GetDefaultBootstrapPeerAddrInfos(),
		LowWater:       10,
		HighWater:      15,
		GracePeriod:    time.Second * 5,
		Rendezvous:     "/sonr/rendevouz/0.9.2",
		Interval:       time.Second * 5,
		TTL:            time.Minute * 1,
		listener:       l,
	}
}

// Apply applies the host options and returns new SNRHost
func (ho hostOptions) Apply(options ...HostOption) (*SNRHost, error) {
	// Iterate over the options.
	var err error
	for _, opt := range options {
		opt(ho)
	}

	// Check if the listener is set.
	if ho.listener == nil {
		return nil, errors.Wrap(ErrListenerRequired, "Failed to apply host options: TCPListener")
	}

	// Create the host.
	hn := &SNRHost{
		ctx:          ho.ctx,
		opts:         ho,
		status:       Status_IDLE,
		Emitter:      state.NewEmitter(2048),
	}

	// Get MultiAddr from listener
	hn.multiAddr, err = ho.listener.Multiaddr()
	if err != nil {
		logger.Error("Failed to parse MultiAddr from LocalHost")
		return nil, errors.Wrap(err, "Failed to apply host options: MultiAddr")
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
	hn.privKey, err = findPrivKey()
	if err != nil {
		logger.Error("Failed to create host, invalid hostOptions")
		return nil, errors.Wrap(err, "Failed to apply host options: PrivKey")
	}

	// Set the private key.
	return hn, nil
}
