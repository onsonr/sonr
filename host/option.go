package host

// import (
// 	"context"
// 	"crypto/rand"
// 	"fmt"
// 	"net"
// 	"time"

// 	"github.com/duo-labs/webauthn/webauthn"
// 	"github.com/kataras/golog"
// 	"github.com/libp2p/go-libp2p-core/crypto"
// 	dscl "github.com/libp2p/go-libp2p-core/discovery"
// 	"github.com/libp2p/go-libp2p-core/peer"
// 	dsc "github.com/libp2p/go-libp2p-discovery"
// 	psub "github.com/libp2p/go-libp2p-pubsub"
// 	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
// 	ma "github.com/multiformats/go-multiaddr"
// 	"github.com/pkg/errors"
// 	"github.com/sonr-io/core/device"
// 	"github.com/sonr-io/core/highway/config"
// 	types "go.buf.build/grpc/go/sonr-io/core/types/v1"
// )

// // LogLevel is the type for the log level
// type LogLevel string

// const (
// 	// DebugLevel is the debug log level
// 	DebugLevel LogLevel = "debug"
// 	// InfoLevel is the info log level
// 	InfoLevel LogLevel = "info"
// 	// WarnLevel is the warn log level
// 	WarnLevel LogLevel = "warn"
// 	// ErrorLevel is the error log level
// 	ErrorLevel LogLevel = "error"
// 	// FatalLevel is the fatal log level
// 	FatalLevel LogLevel = "fatal"
// )

// // Option is a function that modifies the node options.
// type Option func(*options)

// // WithLogLevel sets the log level for Logger
// func WithLogLevel(level LogLevel) Option {
// 	return func(o *options) {
// 		o.logLevel = string(level)
// 	}
// }

// // WithAddress sets the host address for the Node Stub Client Host
// func WithAddress(host string) Option {
// 	return func(o *options) {
// 		o.host = host
// 	}
// }

// // WithCosmosAccountName sets the cosmos account name to use. defaults to "alice"
// func WithCosmosAccountName(host string) Option {
// 	return func(o *options) {
// 		o.host = host
// 	}
// }

// // WithConnOptions sets the connection manager options. Defaults are (lowWater: 15, highWater: 40, gracePeriod: 5m)
// func WithConnOptions(low int, hi int, grace time.Duration) Option {
// 	return func(o *options) {
// 		o.LowWater = low
// 		o.HighWater = hi
// 		o.GracePeriod = grace
// 	}
// }

// // WithInterval sets the interval for the host. Default is 5 seconds.
// func WithInterval(interval time.Duration) Option {
// 	return func(o *options) {
// 		o.Interval = interval
// 	}
// }

// // WithTTL sets the ttl for the host. Default is 2 minutes.
// func WithTTL(ttl time.Duration) Option {
// 	return func(o *options) {
// 		o.TTL = dscl.TTL(ttl)
// 	}
// }

// // WithPort sets the port for the Node Stub Client
// func WithPort(port int) Option {
// 	return func(o *options) {
// 		o.port = port
// 	}
// }

// // WithWebAuthn sets the webauthn server Properties
// func WithWebAuthn(displayName string, rpId string, rpOrigin string, isDebug bool) Option {
// 	return func(o *options) {
// 		o.RPDisplayName = displayName
// 		o.RPID = rpId
// 		o.RPOrigin = rpOrigin
// 		o.Debug = isDebug
// 	}
// }

// // DisableMDNS sets the non-priority of MDNS Discovery
// func (hn *hostImpl) DisableMDNS() Option {
// 	return func(o *options) {
// 		o.mdnsDisabled = true
// 	}
// }

// // options is a collection of options for the node.
// type options struct {
// 	role device.Role

// 	// Host
// 	BootstrapPeers []peer.AddrInfo
// 	Connection     types.Connection
// 	LowWater       int
// 	HighWater      int
// 	GracePeriod    time.Duration
// 	MultiAddrs     []ma.Multiaddr
// 	Rendezvous     string
// 	Interval       time.Duration
// 	TTL            dscl.Option

// 	// Session
// 	host         string
// 	logLevel     string
// 	network      string
// 	port         int
// 	mdnsDisabled bool

// 	// WebAuthn
// 	RPDisplayName string
// 	RPID          string
// 	RPOrigin      string
// 	RPIcon        string
// 	Debug         bool

// 	// Cosmos SDK
// 	accountName string
// }

// // defaultOptions returns the default options
// func defaultOptions(r device.Role) *options {
// 	// Create Bootstrapper List
// 	var bootstrappers []ma.Multiaddr
// 	for _, s := range config.BootstrapAddrStrs {
// 		ma, err := ma.NewMultiaddr(s)
// 		if err != nil {
// 			continue
// 		}
// 		bootstrappers = append(bootstrappers, ma)
// 	}

// 	// Create Address Info List
// 	ds := make([]peer.AddrInfo, 0, len(bootstrappers))
// 	for i := range bootstrappers {
// 		info, err := peer.AddrInfoFromP2pAddr(bootstrappers[i])
// 		if err != nil {
// 			continue
// 		}
// 		ds = append(ds, *info)
// 	}

// 	return &options{
// 		host:           ":",
// 		port:           26225,
// 		role:           r,
// 		network:        "tcp",
// 		mdnsDisabled:   false,
// 		logLevel:       string(InfoLevel),
// 		LowWater:       200,
// 		HighWater:      400,
// 		GracePeriod:    time.Second * 20,
// 		Rendezvous:     "/sonr/rendevouz/0.9.2",
// 		MultiAddrs:     make([]ma.Multiaddr, 0),
// 		Interval:       time.Second * 5,
// 		BootstrapPeers: ds,
// 		TTL:            dscl.TTL(time.Minute * 2),
// 		RPDisplayName:  "Sonr",
// 		RPID:           "localhost",
// 		RPOrigin:       "localhost:8080",
// 		RPIcon:         "",
// 		Debug:          true,
// 		accountName:    "alice",
// 	}
// }

// // Address returns the address of the node.
// func (opts *options) Address() string {
// 	return fmt.Sprintf("%s%d", opts.host, opts.port)
// }

// // Apply applies the host options and returns new SNRHost
// func (opts *options) Apply(ctx context.Context, options ...Option) (*hostImpl, error) {
// 	// Iterate over the options.
// 	var err error
// 	for _, opt := range options {
// 		opt(opts)
// 	}

// 	// Create the host.
// 	hn := &hostImpl{
// 		ctx:          ctx,
// 		status:       Status_IDLE,
// 		mdnsPeerChan: make(chan peer.AddrInfo),
// 		connection:   opts.Connection,
// 		role:         opts.role,
// 		webauthnConfig: &webauthn.Config{
// 			RPDisplayName: opts.RPDisplayName,
// 			RPID:          opts.RPID,
// 			RPOrigin:      opts.RPOrigin,
// 			RPIcon:        opts.RPIcon,
// 			Debug:         opts.Debug,
// 		},
// 		cosmosAccountName: opts.accountName,
// 	}

// 	// Open Listener on Port
// 	if opts.role == device.Role_HIGHWAY {
// 		hn.listener, err = net.Listen(opts.network, opts.Address())
// 		if err != nil {
// 			golog.Default.Child("(app)").Fatalf("%s - Failed to Create New Listener", err)
// 			return nil, err
// 		}
// 		logger.Infof("(app) - Listening on %s", hn.listener.Addr().String())
// 	}

// 	// findPrivKey returns the private key for the host.
// 	findPrivKey := func() (crypto.PrivKey, error) {
// 		privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
// 		if err == nil {
// 			logger.Warn("Generated new Libp2p Private Key")
// 			return privKey, nil
// 		}
// 		return nil, err
// 	}

// 	// Fetch the private key.
// 	hn.privKey, err = findPrivKey()
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Failed to apply host options: PrivKey")
// 	}

// 	// Set the private key.
// 	return hn, nil
// }

// // createDHTDiscovery is a Helper Method to initialize the DHT Discovery
// func (hn *node) createDHTDiscovery(opts *options) error {
// 	// Set Routing Discovery, Find Peers
// 	var err error
// 	routingDiscovery := dsc.NewRoutingDiscovery(hn.IpfsDHT)
// 	dsc.Advertise(hn.ctx, routingDiscovery, opts.Rendezvous, opts.TTL)

// 	// Create Pub Sub
// 	hn.PubSub, err = psub.NewGossipSub(hn.ctx, hn.Host, psub.WithDiscovery(routingDiscovery))
// 	if err != nil {
// 		hn.SetStatus(Status_FAIL)
// 		logger.Errorf("%s - Failed to Create new Gossip Sub", err)
// 		return err
// 	}

// 	// Handle DHT Peers
// 	hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, opts.Rendezvous, opts.TTL)
// 	if err != nil {
// 		hn.SetStatus(Status_FAIL)
// 		logger.Errorf("%s - Failed to create FindPeers Discovery channel", err)
// 		return err
// 	}
// 	hn.SetStatus(Status_READY)
// 	return nil
// }

// // createMdnsDiscovery is a Helper Method to initialize the MDNS Discovery
// func (hn *node) createMdnsDiscovery(opts *options) {
// 	if hn.Role() == device.Role_MOTOR {
// 		// Create MDNS Service
// 		ser := mdns.NewMdnsService(hn.Host, opts.Rendezvous, hn)

// 		ser.Start()
// 		// Handle Events
// 		// ser.RegisterNotifee(hn)
// 	}
// }
