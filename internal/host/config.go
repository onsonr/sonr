package host

import (
	"context"
	"crypto/rand"

	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/crypto"
	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	mdns "github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/keychain"
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
func defaultHostOptions(ctx context.Context) hostOptions {
	return hostOptions{
		ctx:            ctx,
		Connection:     common.Connection_WIFI,
		BootstrapPeers: dht.GetDefaultBootstrapPeerAddrInfos(),
		LowWater:       15,
		HighWater:      40,
		GracePeriod:    time.Minute * 5,
		Rendezvous:     "/sonr/rendevouz/0.9.2",
		Interval:       time.Second * 5,
		TTL:            time.Minute * 2,
		// listener:       l,
	}
}

// Apply applies the host options and returns new SNRHost
func (ho hostOptions) Apply(em *state.Emitter, options ...HostOption) (*SNRHost, error) {
	// Check if emitter is set
	if em == nil {
		return nil, errors.New("Emitter is not set")
	}

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
		emitter:      em,
		mdnsPeerChan: make(chan peer.AddrInfo),
	}

	// // Get MultiAddr from listener
	// hn.multiAddr, err = ho.listener.Multiaddr()
	// if err != nil {
	// 	return nil, errors.Wrap(err, "Failed to apply host options: MultiAddr")
	// }

	// findPrivKey returns the private key for the host.
	findPrivKey := func() (crypto.PrivKey, error) {
		privKey, err := device.KeyChain.GetPrivKey(keychain.Account)
		if err == nil {
			return privKey, nil
		}
		privKey, _, err = crypto.GenerateEd25519Key(rand.Reader)
		if err == nil {
			return privKey, nil
		}
		return nil, err
	}

	// Fetch the private key.
	hn.privKey, err = findPrivKey()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to apply host options: PrivKey")
	}

	// Set the private key.
	return hn, nil
}

// checkUnknown is a Helper Method checks if Peer AddrInfo is Unknown
func (hn *SNRHost) checkUnknown(pi peer.AddrInfo) bool {
	// Iterate and Check
	if len(hn.Peerstore().Addrs(pi.ID)) > 0 {
		return false
	}

	// Add to PeerStore
	hn.Peerstore().AddAddrs(pi.ID, pi.Addrs, hn.opts.TTL)
	return true
}

// createDHTDiscovery is a Helper Method to initialize the DHT Discovery
func (hn *SNRHost) createDHTDiscovery() error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := dsc.NewRoutingDiscovery(hn.IpfsDHT)
	dsc.Advertise(hn.ctx, routingDiscovery, hn.opts.Rendezvous, dscl.TTL(hn.opts.TTL))

	// Create Pub Sub
	hn.PubSub, err = psub.NewGossipSub(hn.ctx, hn.Host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		hn.SetStatus(Status_FAIL)
		logger.Error("Failed to Create new Gossip Sub", err)
		return err
	}

	// Handle DHT Peers
	hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, hn.opts.Rendezvous, dscl.TTL(hn.opts.TTL))
	if err != nil {
		hn.SetStatus(Status_FAIL)
		logger.Error("Failed to create FindPeers Discovery channel", err)
		return err
	}
	hn.SetStatus(Status_READY)
	return nil
}

// createMdnsDiscovery is a Helper Method to initialize the MDNS Discovery
func (hn *SNRHost) createMdnsDiscovery() {
	// Verify if MDNS is Enabled
	if !hn.opts.Connection.IsMdnsCompatible() {
		logger.Error("Failed to Start MDNS Discovery ", ErrMDNSInvalidConn)
		return
	}

	// Create MDNS Service
	ser := mdns.NewMdnsService(hn.Host, hn.opts.Rendezvous)

	// Handle Events
	ser.RegisterNotifee(hn)
}
