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
	ma "github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/common"
	"github.com/sonr-io/core/wallet"
)

// Error Definitions
var (
	logger              = golog.Default.Child("internal/host")
	ErrRoutingNotSet    = errors.New("DHT and Host have not been set by Routing Function")
	ErrListenerRequired = errors.New("Listener was not Provided")
	ErrMDNSInvalidConn  = errors.New("Invalid Connection, cannot begin MDNS Service")
)

var (
	bootstrapAddrStrs = []string{
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		"/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
		"/ip4/104.131.131.82/udp/4001/quic/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	}
	addrStoreTTL = time.Minute * 5
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
		o.TTL = dscl.TTL(ttl)
	}
}

// hostOptions is a collection of options for the SnrHost.
type hostOptions struct {
	// Properties
	BootstrapPeers []peer.AddrInfo
	Connection     common.Connection
	LowWater       int
	HighWater      int
	GracePeriod    time.Duration
	MultiAddrs     []ma.Multiaddr
	Rendezvous     string
	Interval       time.Duration
	TTL            dscl.Option
}

// defaultHostOptions returns the default host options.
func defaultHostOptions() hostOptions {
	// Create Bootstrapper List
	var bootstrappers []ma.Multiaddr
	for _, s := range bootstrapAddrStrs {
		ma, err := ma.NewMultiaddr(s)
		if err != nil {
			continue
		}
		bootstrappers = append(bootstrappers, ma)
	}

	// Create Address Info List
	ds := make([]peer.AddrInfo, 0, len(bootstrappers))
	for i := range bootstrappers {
		info, err := peer.AddrInfoFromP2pAddr(bootstrappers[i])
		if err != nil {
			continue
		}
		ds = append(ds, *info)
	}

	// Return Host Options
	return hostOptions{
		Connection:     common.Connection_WIFI,
		LowWater:       200,
		HighWater:      400,
		GracePeriod:    time.Second * 20,
		Rendezvous:     "/sonr/rendevouz/0.9.2",
		MultiAddrs:     make([]ma.Multiaddr, 0),
		Interval:       time.Second * 5,
		BootstrapPeers: ds,
		TTL:            dscl.TTL(time.Minute * 2),
	}
}

// Apply applies the host options and returns new SNRHost
func (opts hostOptions) Apply(ctx context.Context, options ...HostOption) (*SNRHost, error) {

	// Iterate over the options.
	var err error
	for _, opt := range options {
		opt(opts)
	}

	// Create the host.
	hn := &SNRHost{
		ctx:          ctx,
		status:       Status_IDLE,
		mdnsPeerChan: make(chan peer.AddrInfo),
		connection:   opts.Connection,
	}

	// findPrivKey returns the private key for the host.
	findPrivKey := func() (crypto.PrivKey, error) {
		privKey, err := wallet.DevicePrivKey()
		if err == nil {
			return privKey, nil
		}
		privKey, _, err = crypto.GenerateEd25519Key(rand.Reader)
		if err == nil {
			logger.Warn("Generated new Account Private Key")
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

// createDHTDiscovery is a Helper Method to initialize the DHT Discovery
func (hn *SNRHost) createDHTDiscovery(opts hostOptions) error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := dsc.NewRoutingDiscovery(hn.IpfsDHT)
	dsc.Advertise(hn.ctx, routingDiscovery, opts.Rendezvous, opts.TTL)

	// Create Pub Sub
	hn.PubSub, err = psub.NewGossipSub(hn.ctx, hn.Host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		hn.SetStatus(Status_FAIL)
		logger.Errorf("%s - Failed to Create new Gossip Sub", err)
		return err
	}

	// Handle DHT Peers
	hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, opts.Rendezvous, opts.TTL)
	if err != nil {
		hn.SetStatus(Status_FAIL)
		logger.Errorf("%s - Failed to create FindPeers Discovery channel", err)
		return err
	}
	hn.SetStatus(Status_READY)
	return nil
}

// createMdnsDiscovery is a Helper Method to initialize the MDNS Discovery
func (hn *SNRHost) createMdnsDiscovery(opts hostOptions) {
	// Verify if MDNS is Enabled
	if !hn.connection.IsMdnsCompatible() {
		logger.Errorf("%s - Failed to Start MDNS Discovery ", ErrMDNSInvalidConn)
		return
	}

	// Create MDNS Service
	ser := mdns.NewMdnsService(hn.Host, opts.Rendezvous)

	// Handle Events
	ser.RegisterNotifee(hn)
}
