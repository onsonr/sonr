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
	psub "github.com/libp2p/go-libp2p-pubsub"
	mdns "github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/keychain"
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

// WithConnection sets the connection to the host. Default is WIFI.
func WithConnection(c common.Connection) HostOption {
	return func(o hostOptions) {
		o.Connection = c
	}
}

// WithConnOptions sets the connection manager options. Defaults are (lowWater: 15, highWater: 40, gracePeriod: 5m)
func WithConnOptions(low int, hi int, grace time.Duration) HostOption {
	return func(o hostOptions) {
		o.LowWater = low
		o.HighWater = hi
		o.GracePeriod = grace
	}
}

// WithInterval sets the interval for the host. Default is 5 seconds.
func WithInterval(interval time.Duration) HostOption {
	return func(o hostOptions) {
		o.Interval = interval
	}
}

// WithTTL sets the ttl for the host. Default is 2 minutes.
func WithTTL(ttl time.Duration) HostOption {
	return func(o hostOptions) {
		o.TTL = ttl
	}
}

// hostOptions is a collection of options for the SnrHost.
type hostOptions struct {
	// Properties
	Connection  common.Connection
	LowWater    int
	HighWater   int
	GracePeriod time.Duration
	MultiAddrs  []multiaddr.Multiaddr
	Rendezvous  string
	Interval    time.Duration
	TTL         time.Duration
}

// defaultHostOptions returns the default host options.
func defaultHostOptions() hostOptions {
	return hostOptions{
		Connection:  common.Connection_WIFI,
		LowWater:    100,
		HighWater:   200,
		GracePeriod: time.Second * 20,
		Rendezvous:  "/sonr/rendevouz/0.9.2",
		MultiAddrs:  make([]multiaddr.Multiaddr, 0),
		Interval:    time.Second * 5,
		TTL:         time.Minute * 2,
	}
}

// Apply applies the host options and returns new SNRHost
func (opts hostOptions) Apply(ctx context.Context, em *state.Emitter, options ...HostOption) (*SNRHost, error) {
	// Check if emitter is set
	if em == nil {
		return nil, errors.New("Emitter is not set")
	}

	// Iterate over the options.
	var err error
	for _, opt := range options {
		opt(opts)
	}

	// Create the host.
	hn := &SNRHost{
		ctx:          ctx,
		status:       Status_IDLE,
		emitter:      em,
		mdnsPeerChan: make(chan peer.AddrInfo),
		connection:   opts.Connection,
		rendezvous:   opts.Rendezvous,
		ttl:          opts.TTL,
		interval:     opts.Interval,
	}

	// // Check if the listener is set.
	// if ho.listener != nil {
	// 	logger.Debug("TCP Listener provided, using for MultiAddr")
	// 	// Get MultiAddr from listener
	// 	addr, err := ho.listener.Multiaddr()
	// 	if err != nil {
	// 		logger.Warn("Failed to add MultiAddr, Skipping...", err)
	// 	} else {
	// 		ho.MultiAddrs = append(ho.MultiAddrs, addr)
	// 	}
	// } else {
	// 	logger.Debug("No TCP Listener provided, using default MultiAddr's")
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
	hn.Peerstore().AddAddrs(pi.ID, pi.Addrs, hn.ttl)
	return true
}

// createDHTDiscovery is a Helper Method to initialize the DHT Discovery
func (hn *SNRHost) createDHTDiscovery() error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := dsc.NewRoutingDiscovery(hn.IpfsDHT)
	dsc.Advertise(hn.ctx, routingDiscovery, hn.rendezvous, dscl.TTL(hn.ttl))

	// Create Pub Sub
	hn.PubSub, err = psub.NewGossipSub(hn.ctx, hn.Host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		hn.SetStatus(Status_FAIL)
		logger.Error("Failed to Create new Gossip Sub", err)
		return err
	}

	// Handle DHT Peers
	hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, hn.rendezvous, dscl.TTL(hn.ttl))
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
	if !hn.connection.IsMdnsCompatible() {
		logger.Error("Failed to Start MDNS Discovery ", ErrMDNSInvalidConn)
		return
	}

	// Create MDNS Service
	ser := mdns.NewMdnsService(hn.Host, hn.rendezvous)

	// Handle Events
	ser.RegisterNotifee(hn)
}
