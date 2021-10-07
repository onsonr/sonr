package host

import (
	"context"
	"crypto/rand"
	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/tools/net"
)

var (
	logger = golog.Child("host")
)

// HostOption is a function that modifies the node options.
type HostOption func(hostOptions)

// FindListenAddrs will set Host to determine the local addresses
func FindListenAddrs() HostOption {
	return func(o hostOptions) {
		o.findAddrs = true
	}
}

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

// WithPrivKey sets the private key for the host.
func WithPrivKey(pk crypto.PrivKey) HostOption {
	return func(o hostOptions) {
		o.privateKey = pk
	}
}

// WithRendevouz sets the rendevouz address.
func WithRendevouz(addr string) HostOption {
	return func(o hostOptions) {
		o.rendezvous = addr
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
	findAddrs      bool
	connection     common.Connection
	bootstrapPeers []peer.AddrInfo
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
	var privKey crypto.PrivKey
	needsGen := false
	privKey, err := device.KeyChain.GetPrivKey(keychain.Account)
	if err != nil {
		logger.Warn("Failed to get Account Private Key for Host", err)
		needsGen = true
	}

	if needsGen {
		privKey, _, err = crypto.GenerateEd25519Key(rand.Reader)
		if err != nil {
			logger.Fatal("Failed to generate Host Private Key", err)
		}
	}

	return hostOptions{
		findAddrs:      false,
		connection:     common.Connection_WIFI,
		bootstrapPeers: dht.GetDefaultBootstrapPeerAddrInfos(),
		lowWater:       10,
		highWater:      15,
		gracePeriod:    time.Second * 5,
		privateKey:     privKey,
		rendezvous:     "/sonr/rendevouz/0.9.2",
		interval:       time.Second * 5,
		ttl:            time.Minute * 1,
	}
}

// Apply creates slice of libp2p.Option from the host options.
func (no hostOptions) Apply(ctx context.Context, hn *SNRHost) []libp2p.Option {
	opts := []libp2p.Option{
		libp2p.Identity(no.privateKey),
		libp2p.ConnectionManager(connmgr.NewConnManager(
			no.lowWater,    // Lowwater
			no.highWater,   // HighWater,
			no.gracePeriod, // GracePeriod
		)),
		libp2p.DefaultStaticRelays(),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			// Create DHT
			kdht, err := dht.New(ctx, h)
			if err != nil {
				return nil, err
			}

			// Set DHT
			hn.IpfsDHT = kdht
			return kdht, nil
		}),
		libp2p.NATPortMap(),
		libp2p.EnableAutoRelay(),
	}

	// Check if we should find ListenAddrStrings
	if no.findAddrs {
		// Get Listening Addresses
		listenAddrs, err := net.PublicAddrStrs()
		if err != nil {
			logger.Error("Failed to get Public Listening Addresses", err)
			return opts
		}

		// Return options
		opts = append(opts, libp2p.ListenAddrStrings(listenAddrs...))
	}
	return opts
}
