package host

import (
	"context"
	"errors"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-msgio"
	"github.com/multiformats/go-multiaddr"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/tools/internet"
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
	host.Host

	// Properties
	ctx          context.Context
	bootstrapped bool
	opts         hostOptions
	multiAddr    multiaddr.Multiaddr
	privKey      crypto.PrivKey
	readyChan    chan bool

	// Discovery
	*dht.IpfsDHT
	*ps.PubSub
}

// NewHost creates a new host
func NewHost(ctx context.Context, listener *internet.TCPListener, options ...HostOption) (*SNRHost, error) {
	// Initialize DHT
	opts := defaultHostOptions(ctx, listener)
	hn, err := opts.Apply(options...)
	if err != nil {
		return nil, err
	}

	// Start Host
	hn.Host, err = libp2p.New(ctx,
		libp2p.Identity(hn.privKey),
		libp2p.ConnectionManager(connmgr.NewConnManager(
			opts.LowWater,    // Lowwater
			opts.HighWater,   // HighWater,
			opts.GracePeriod, // GracePeriod
		)),
		libp2p.ListenAddrs(hn.multiAddr),
		libp2p.DefaultListenAddrs,
		libp2p.DefaultStaticRelays(),
		libp2p.Routing(hn.Router),
		libp2p.NATPortMap(),
		libp2p.EnableAutoRelay())
	if err != nil {
		logger.Error("Failed to create libp2p host", err)
		return nil, err
	}
	return hn.Setup()
}

// Connect connects with `peer.AddrInfo` if underlying Host is ready
func (hn *SNRHost) Connect(ctx context.Context, pi peer.AddrInfo) error {
	// Check if host is ready
	if err := hn.IsReady(); err != nil {
		logger.Warn("Underlying host is not ready, failed to call Connect()")
		return err
	}

	// Connect to peer concurrently
	errChan := make(chan error, 1)
	ctxTO, cancel := context.WithTimeout(ctx, HOST_TIMEOUT)
	defer cancel()
	go func(context context.Context, errorChannel chan error) {
		err := hn.Host.Connect(ctxTO, pi)
		errorChannel <- err
	}(ctxTO, errChan)

	// Await for result
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

// Join wraps around PubSub.Join and returns the joined topic. Checks wether the
// host is ready before joining.
func (hn *SNRHost) Join(topic string, opts ...ps.TopicOpt) (*ps.Topic, error) {
	if hn.PubSub == nil {
		return nil, errors.New("Pubsub has not been set on SNRHost")
	}
	if topic == "" {
		return nil, errors.New("Empty topic string provided to Join for host.Pubsub")
	}
	return hn.PubSub.Join(topic, opts...)
}

// Close closes the underlying host
func (hn *SNRHost) Close() error {
	hn.IpfsDHT.Close()
	return hn.Host.Close()
}

// Router returns the host node Peer Routing Function
func (hn *SNRHost) Router(h host.Host) (routing.PeerRouting, error) {
	// Create DHT
	kdht, err := dht.New(hn.ctx, h)
	if err != nil {
		hn.readyChan <- false
		return nil, err
	}

	// Set Properties
	hn.IpfsDHT = kdht
	hn.Host = h
	logger.Info("Host and DHT have been set for SNRNode")

	// Setup Properties
	return hn.IpfsDHT, nil
}

// ** ─── Host Info ────────────────────────────────────────────────────────
// IsReady returns no-error if the host is ready for connect
func (hn *SNRHost) IsReady() error {
	if hn.IpfsDHT == nil || hn.Host == nil {
		return ErrRoutingNotSet
	}
	return nil
}

// SendMessage writes a protobuf go data object to a network stream
func (hn *SNRHost) SendMessage(id peer.ID, p protocol.ID, data proto.Message) error {
	err := hn.IsReady()
	if err != nil {
		return err
	}

	s, err := hn.NewStream(hn.ctx, id, p)
	if err != nil {
		logger.Error("Failed to start stream", err)
		return err
	}
	defer s.Close()

	// marshall data to protobufs3 binary format
	bin, err := proto.Marshal(data)
	if err != nil {
		logger.Error("Failed to marshal pb", err)
		return err
	}

	// Create Writer and write data to stream
	w := msgio.NewWriter(s)
	if err := w.WriteMsg(bin); err != nil {
		logger.Error("Failed to write message to stream.", err)
		return err
	}
	return nil
}

// Stat returns the host stat info
func (hn *SNRHost) Stat() (*SNRHostStat, error) {
	// Get Public Key
	pubKey, err := device.KeyChain.GetPubKey(keychain.Account)
	if err != nil {
		logger.Error("Failed to get public key", err)
		return nil, err
	}

	// Marshal Public Key
	buf, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		logger.Error("Failed to marshal public key", err)
		return nil, err
	}

	// Return Host Stat
	return &SNRHostStat{
		ID:        hn.ID(),
		PublicKey: string(buf),
		PeerID:    hn.ID().Pretty(),
		MultAddr:  hn.Addrs()[0].String(),
	}, nil
}

// WaitForReady blocks until the host is ready
func (hn *SNRHost) WaitForReady() error {
	timeout := time.After(HOST_TIMEOUT)
	for {
		select {
		case <-timeout:
			return errors.New("Timeout occurred waiting for host to be ready")
		case ready := <-hn.readyChan:
			if !ready {
				return errors.New("Host failed to be ready")
			}
			return nil
		case <-hn.ctx.Done():
			return errors.New("Context ended before host became ready")
		}
	}
}
