package host

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/tools/logger"
	"google.golang.org/protobuf/proto"
)

// SNRHostStat is the host stat info
type SNRHostStat struct {
	ID        peer.ID
	PublicKey string
	PeerID    string
	MultAddr  string
	Address   string
}

// SNRHost is the host wrapper for the Sonr Network
type SNRHost struct {
	// Properties
	ctx  context.Context
	opts hostOptions

	host.Host
	*dht.IpfsDHT
	*psub.PubSub

	// Libp2p
	disc *dsc.RoutingDiscovery
}

// NewHost creates a new host
func NewHost(ctx context.Context, options ...HostOption) (*SNRHost, error) {
	// Initialize DHT
	var err error
	opts := defaultHostOptions()
	for _, opt := range options {
		opt(opts)
	}

	// Create Host
	hn := &SNRHost{
		ctx:  ctx,
		opts: opts,
	}

	// Start Host
	hn.Host, err = libp2p.New(ctx,
		libp2p.Identity(opts.privateKey),
		libp2p.ConnectionManager(connmgr.NewConnManager(
			opts.lowWater,    // Lowwater
			opts.highWater,   // HighWater,
			opts.gracePeriod, // GracePeriod
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
		libp2p.EnableAutoRelay())
	if err != nil {
		return nil, logger.Error("Failed to initialize libp2p host", err)
	}

	// Bootstrap Host
	err = hn.Bootstrap()
	if err != nil {
		return nil, logger.Error("Failed to bootstrap libp2p Host", err)
	}
	return hn, nil
}

// Close closes the host
func (hn *SNRHost) Close() error {
	return hn.Host.Close()
}

// ** ─── Host Info ────────────────────────────────────────────────────────
// SendMessage writes a protobuf go data object to a network stream
func (hn *SNRHost) SendMessage(id peer.ID, p protocol.ID, data proto.Message) error {
	ctxCancel, cancel := context.WithDeadline(hn.ctx, <-time.After(10*time.Second))
	defer cancel()

	s, err := hn.NewStream(ctxCancel, id, p)
	if err != nil {
		return logger.Error("Failed to start stream", err)
	}
	defer s.Close()

	// marshall data to protobufs3 binary format
	bin, err := proto.Marshal(data)
	if err != nil {
		return logger.Error("Failed to marshal pb", err)
	}

	// Create Writer and write data to stream
	w := msgio.NewWriter(s)
	if err := w.WriteMsg(bin); err != nil {
		return logger.Error("Failed to write message to stream.", err)
	}
	return nil
}

// Stat returns the host stat info
func (hn *SNRHost) Stat() (*SNRHostStat, error) {
	// Get Public Key
	pubKey, err := device.KeyChain.GetPubKey(keychain.Account)
	if err != nil {
		return nil, logger.Error("Failed to get public key", err)
	}

	// Marshal Public Key
	buf, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		return nil, logger.Error("Failed to marshal public key", err)
	}

	// Return Host Stat
	return &SNRHostStat{
		ID:        hn.ID(),
		PublicKey: string(buf),
		PeerID:    hn.ID().Pretty(),
		MultAddr:  hn.Addrs()[0].String(),
	}, nil
}
