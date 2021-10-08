package host

import (
	"context"
	"errors"
	"sync"

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
	"github.com/sonr-io/core/tools/internet"
	"github.com/sonr-io/core/tools/state"
	"google.golang.org/protobuf/proto"
)

// SNRHostStat is the host stat info
type SNRHostStat struct {
	ID       peer.ID
	PeerID   string
	MultAddr string
	Address  string
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

	// State
	mu sync.Mutex
	*state.Emitter
	status SNRHostStatus

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
	hn.SetStatus(Status_CONNECTING)
	return hn.Setup()
}

// Connect connects with `peer.AddrInfo` if underlying Host is ready
func (hn *SNRHost) Connect(ctx context.Context, pi peer.AddrInfo) error {
	// Check if host is ready
	if err := hn.HasRouting(); err != nil {
		logger.Warn("Underlying host is not ready, failed to call Connect()")
		return err
	}

	// Call Underlying Host to Connect
	return hn.Host.Connect(ctx, pi)
}

// Join wraps around PubSub.Join and returns topic. Checks wether the host is ready before joining.
func (hn *SNRHost) Join(topic string, opts ...ps.TopicOpt) (*ps.Topic, error) {
	// Check if PubSub is Set
	if hn.PubSub == nil {
		return nil, errors.New("Pubsub has not been set on SNRHost")
	}

	// Check if topic is valid
	if topic == "" {
		return nil, errors.New("Empty topic string provided to Join for host.Pubsub")
	}

	// Call Underlying Pubsub to Connect
	return hn.PubSub.Join(topic, opts...)
}

// Close closes the underlying host
func (hn *SNRHost) Close() error {
	hn.SetStatus(Status_CLOSED)
	hn.IpfsDHT.Close()
	return hn.Host.Close()
}

// Router returns the host node Peer Routing Function
func (hn *SNRHost) Router(h host.Host) (routing.PeerRouting, error) {
	// Create DHT
	kdht, err := dht.New(hn.ctx, h)
	if err != nil {
		hn.SetStatus(Status_FAIL)
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
// HasRouting returns no-error if the host is ready for connect
func (h *SNRHost) HasRouting() error {
	if h.IpfsDHT == nil || h.Host == nil {
		return ErrRoutingNotSet
	}
	return nil
}

// IsStatus returns true if the host is in the provided status
func (hn *SNRHost) IsStatus(s SNRHostStatus) bool {
	hn.mu.Lock()
	defer hn.mu.Unlock()
	return hn.status == s
}

// SendMessage writes a protobuf go data object to a network stream
func (h *SNRHost) SendMessage(id peer.ID, p protocol.ID, data proto.Message) error {
	err := h.HasRouting()
	if err != nil {
		return err
	}

	s, err := h.NewStream(h.ctx, id, p)
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

// SetStatus sets the host status and emits the event
func (h *SNRHost) SetStatus(s SNRHostStatus) {
	h.mu.Lock()
	h.status = s
	h.Emit(Event_STATUS, s)
	h.mu.Unlock()
}

// Stat returns the host stat info
func (hn *SNRHost) Stat() (*SNRHostStat, error) {
	// Return Host Stat
	return &SNRHostStat{
		ID:       hn.ID(),
		PeerID:   hn.ID().Pretty(),
		MultAddr: hn.Addrs()[0].String(),
	}, nil
}

// OnReady registers a function to be called when the host is ready
func (hn *SNRHost) OnReady(f StatusFunc) {
	finished := make(chan bool)
	logger.Info("OnReady: Created Status Worker for - Status_READY")
	go createEventLoop(hn, WithDoneChannel(finished), WithMiddlewareFunc(f), WithTargetEvent(Status_READY))
	logger.Info("OnReady: Waiting for status worker to finish...")
	<-finished
}

// OnFail registers a function to be called when the host failed to connect
func (hn *SNRHost) OnFail(f StatusFunc) {
	finished := make(chan bool)
	logger.Info("OnFail: Created Status Worker for - Status_FAIL")
	go createEventLoop(hn, WithDoneChannel(finished), WithMiddlewareFunc(f), WithTargetEvent(Status_FAIL))
	logger.Info("OnFail: Waiting for status worker to finish...")
	<-finished
}

// WaitForReady waits for the host to be ready to accept connections
func (hn *SNRHost) WaitForReady() {
	finished := make(chan bool)
	logger.Info("WaitForReady: Created Status Worker for - Status_READY")
	go createEventLoop(hn, WithDoneChannel(finished))
	logger.Info("WaitForReady: Waiting for status worker to finish...")
	<-finished
	logger.Info("WaitForReady: Completed")
	return
}
