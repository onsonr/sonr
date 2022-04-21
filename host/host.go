package host

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"net"

	"git.mills.io/prologic/bitcask"
	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p"
	cmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/device"
	"github.com/sonr-io/core/highway/config"
	t "go.buf.build/sonr-io/grpc-gateway/sonr-io/core/types/v1"
	types "go.buf.build/sonr-io/grpc-gateway/sonr-io/core/types/v1"
	"google.golang.org/protobuf/proto"
)

var (
	logger = golog.Default.Child("core/node")
)

// HostImpl returns the HostImpl for the Main Node
type HostImpl interface {
	// AuthenticateMessage authenticates a message
	AuthenticateMessage(msg proto.Message, metadata *types.Metadata) bool

	// Close closes the node
	Close()

	// Connect to a peer
	Connect(pi peer.AddrInfo) error

	// HasRouting returns true if the node has routing
	HasRouting() error

	// Host returns the Host
	Host() host.Host

	// HostID returns the ID of the Host
	HostID() peer.ID

	// Join subsrcibes to a topic
	Join(topic string, opts ...ps.TopicOpt) (*ps.Topic, error)

	// Listener returns the listener of the node
	Listener() (net.Listener, error)

	// NewStream opens a new stream to a peer
	NewStream(ctx context.Context, pid peer.ID, pids ...protocol.ID) (network.Stream, error)

	// NewTopic creates a new pubsub topic with event handler and subscription
	NewTopic(topic string, opts ...ps.TopicOpt) (*ps.Topic, *ps.TopicEventHandler, *ps.Subscription, error)

	// NeedsWait checks if state is Resumed or Paused and blocks channel if needed
	NeedsWait()

	// Pause tells all of goroutines to pause execution
	Pause()

	// Persist persists the node to the port and address
	Persist()

	// Ping sends a ping to a peer to check if it is alive
	Ping(id string) error

	// Peer returns the peer of the node
	Peer() (*types.Peer, error)

	// Profile returns the profile of the node from Local Store
	Profile() (*types.Profile, error)

	// Publish publishes a message to a topic
	Publish(topic string, msg proto.Message, metadata *types.Metadata) error

	// Pubsub returns the pubsub of the node
	Pubsub() *ps.PubSub

	// Resume tells all of goroutines to resume execution
	Resume()

	// Role returns the role of the node
	Role() device.Role

	// Router returns the routing.Router
	Router(h host.Host) (routing.PeerRouting, error)

	// Routing returns the routing.Routing
	Routing() routing.Routing

	// PrivateKey returns the ed25519 private key instance of the libp2p host
	PrivateKey() (ed25519.PrivateKey, error)

	// SendMessage sends a message to a peer
	SendMessage(id peer.ID, p protocol.ID, data proto.Message) error

	// SetStreamHandler sets the handler for a protocol
	SetStreamHandler(protocol protocol.ID, handler network.StreamHandler)

	// SignData signs the data with the private key
	SignData(data []byte) ([]byte, error)

	// SignMessage signs a message with the node's private key
	SignMessage(message proto.Message) ([]byte, error)

	// VerifyData verifies the data signature
	VerifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool
}

// hostImpl type - a p2p host implementing one or more p2p protocols
type hostImpl struct {
	// Standard Node Implementation
	host host.Host
	HostImpl
	role device.Role

	// Host and context
	connection   types.Connection
	listener     net.Listener
	privKey      crypto.PrivKey
	mdnsPeerChan chan peer.AddrInfo
	dhtPeerChan  <-chan peer.AddrInfo

	// Properties
	ctx   context.Context
	store *bitcask.Bitcask

	*dht.IpfsDHT
	*ps.PubSub

	// State
	flag   uint64
	Chn    chan bool
	status HostStatus
}

// NewHost Creates a Sonr libp2p Host with the given config
func NewHost(ctx context.Context, r device.Role, config *config.Config) (HostImpl, error) {
	var err error
	// Create the host.
	hn := &hostImpl{
		ctx:          ctx,
		status:       Status_IDLE,
		mdnsPeerChan: make(chan peer.AddrInfo),
		role:         config.Role,
	}

	// Open Listener on Port
	if config.Role == device.Role_HIGHWAY {
		hn.listener, err = net.Listen(config.Libp2pNetwork, config.Libp2pAddress())
		if err != nil {
			golog.Default.Child("(app)").Fatalf("%s - Failed to Create New Listener", err)
			return nil, err
		}
		logger.Infof("(app) - Listening on %s", hn.listener.Addr().String())
	}

	// findPrivKey returns the private key for the host.
	findPrivKey := func() (crypto.PrivKey, error) {
		privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
		if err == nil {
			logger.Warn("Generated new Libp2p Private Key")
			return privKey, nil
		}
		return nil, err
	}

	// Fetch the private key.
	hn.privKey, err = findPrivKey()
	if err != nil {
		return nil, err
	}

	// Create Connection Manager
	cnnmgr, err := cmgr.NewConnManager(config.Libp2pLowWater, config.Libp2pHighWater)
	if err != nil {
		return nil, err
	}

	// Start Host
	hn.host, err = libp2p.New(
		libp2p.Identity(hn.privKey),
		libp2p.ConnectionManager(cnnmgr),
		libp2p.DefaultListenAddrs,
		libp2p.Routing(hn.Router),
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		logger.Errorf("%s - NewHost: Failed to create libp2p host", err)
		return nil, err
	}
	hn.SetStatus(Status_CONNECTING)

	// Bootstrap DHT
	if err := hn.Bootstrap(context.Background()); err != nil {
		logger.Errorf("%s - Failed to Bootstrap KDHT to Host", err)
		hn.SetStatus(Status_FAIL)
		return nil, err
	}

	// Connect to Bootstrap Nodes
	for _, pi := range config.Libp2pBootstrapPeers {
		if err := hn.Connect(pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := hn.createDHTDiscovery(config); err != nil {
		// Check if we need to close the listener
		hn.SetStatus(Status_FAIL)
		logger.Fatal("Could not start DHT Discovery", err)
		return nil, err
	}

	// Initialize Discovery for MDNS
	if !config.Libp2pMdnsDisabled && hn.role != device.Role_HIGHWAY {
		// hn.createMdnsDiscovery(config)
	}

	hn.SetStatus(Status_READY)
	go hn.Serve()
	return hn, nil
}

// Host returns the host of the node
func (hn *hostImpl) Host() host.Host {
	return hn.host
}

// HostID returns the ID of the Host
func (n *hostImpl) HostID() peer.ID {
	return n.host.ID()
}

// Listener returns the listener of the node
func (n *hostImpl) Listener() (net.Listener, error) {
	if n.listener == nil {
		return nil, errors.New("Host is not listening")
	}
	return n.listener, nil
}

// Ping sends a ping to the peer
func (n *hostImpl) Ping(pid string) error {
	return nil
}

// PrivateKey returns the private key of the node
func (n *hostImpl) PrivateKey() (ed25519.PrivateKey, error) {
	// Get Raw Private Key
	buf, err := n.privKey.Raw()
	if err != nil {
		logger.Errorf("%s - Failed to get Raw Private Key", err)
		return nil, err
	}
	return ed25519.PrivateKey(buf), nil
}

// Publish publishes a message to the network
func (n *hostImpl) Publish(t string, message proto.Message, metadata *types.Metadata) error {
	return nil
}

// Role returns the role of the node
func (n *hostImpl) Role() device.Role {
	return n.role
}

// AuthenticateMessage Authenticates incoming p2p message
func (n *hostImpl) AuthenticateMessage(msg proto.Message, metadata *t.Metadata) bool {
	// store a temp ref to signature and remove it from message data
	// sign is a string to allow easy reset to zero-value (empty string)
	sign := metadata.Signature
	metadata.Signature = nil

	// marshall data without the signature to protobufs3 binary format
	buf, err := proto.Marshal(msg)
	if err != nil {
		logger.Errorf("%s - AuthenticateMessage: Failed to marshal Protobuf Message.", err)
		return false
	}

	// restore sig in message data (for possible future use)
	metadata.Signature = sign

	// restore peer id binary format from base58 encoded node id data
	peerId, err := peer.Decode(metadata.NodeId)
	if err != nil {
		logger.Errorf("%s - AuthenticateMessage: Failed to decode node id from base58.", err)
		return false
	}

	// verify the data was authored by the signing peer identified by the public key
	// and signature included in the message
	return n.VerifyData(buf, []byte(sign), peerId, metadata.PublicKey)
}

// Connect connects with `peer.AddrInfo` if underlying Host is ready
func (hn *hostImpl) Connect(pi peer.AddrInfo) error {
	// Check if host is ready
	if err := hn.HasRouting(); err != nil {
		logger.Warn("Connect: Underlying host is not ready, failed to call Connect()")
		return err
	}

	// Call Underlying Host to Connect
	return hn.host.Connect(hn.ctx, pi)
}

// HandlePeerFound is to be called when new  peer is found
func (hn *hostImpl) HandlePeerFound(pi peer.AddrInfo) {
	hn.mdnsPeerChan <- pi
}

// HasRouting returns no-error if the host is ready for connect
func (h *hostImpl) HasRouting() error {
	if h.IpfsDHT == nil || h.host == nil {
		return errors.New("Host is not ready")
	}
	return nil
}

// Join wraps around PubSub.Join and returns topic. Checks wether the host is ready before joining.
func (hn *hostImpl) Join(topic string, opts ...ps.TopicOpt) (*ps.Topic, error) {
	// Check if PubSub is Set
	if hn.PubSub == nil {
		return nil, errors.New("Join: Pubsub has not been set on SNRHost")
	}

	// Check if topic is valid
	if topic == "" {
		return nil, errors.New("Join: Empty topic string provided to Join for host.Pubsub")
	}

	// Call Underlying Pubsub to Connect
	return hn.PubSub.Join(topic, opts...)
}

// NewStream opens a new stream to the peer with given peer id
func (n *hostImpl) NewStream(ctx context.Context, pid peer.ID, pids ...protocol.ID) (network.Stream, error) {
	return n.host.NewStream(ctx, pid, pids...)
}

// NewTopic creates a new topic
func (n *hostImpl) NewTopic(name string, opts ...ps.TopicOpt) (*ps.Topic, *ps.TopicEventHandler, *ps.Subscription, error) {
	// Check if PubSub is Set
	if n.PubSub == nil {
		return nil, nil, nil, errors.New("NewTopic: Pubsub has not been set on SNRHost")
	}

	// Call Underlying Pubsub to Connect
	t, err := n.Join(name, opts...)
	if err != nil {
		logger.Errorf("%s - NewTopic: Failed to create new topic", err)
		return nil, nil, nil, err
	}

	// Create Event Handler
	h, err := t.EventHandler()
	if err != nil {
		logger.Errorf("%s - NewTopic: Failed to create new topic event handler", err)
		return nil, nil, nil, err
	}

	// Create Subscriber
	s, err := t.Subscribe()
	if err != nil {
		logger.Errorf("%s - NewTopic: Failed to create new topic subscriber", err)
		return nil, nil, nil, err
	}
	return t, h, s, nil
}

// Router returns the host node Peer Routing Function
func (hn *hostImpl) Router(h host.Host) (routing.PeerRouting, error) {
	// Create DHT
	kdht, err := dht.New(hn.ctx, h)
	if err != nil {
		hn.SetStatus(Status_FAIL)
		return nil, err
	}

	// Set Properties
	hn.IpfsDHT = kdht
	logger.Debug("Router: Host and DHT have been set for SNRNode")

	// Setup Properties
	return hn.IpfsDHT, nil
}

// PubSub returns the host node PubSub Function
func (hn *hostImpl) Pubsub() *ps.PubSub {
	return hn.PubSub
}

// Routing returns the host node Peer Routing Function
func (hn *hostImpl) Routing() routing.Routing {
	return hn.IpfsDHT
}

// SetStreamHandler sets the handler for a given protocol
func (n *hostImpl) SetStreamHandler(protocol protocol.ID, handler network.StreamHandler) {
	n.host.SetStreamHandler(protocol, handler)
}

// SendMessage writes a protobuf go data object to a network stream
func (h *hostImpl) SendMessage(id peer.ID, p protocol.ID, data proto.Message) error {
	err := h.HasRouting()
	if err != nil {
		return err
	}

	s, err := h.NewStream(h.ctx, id, p)
	if err != nil {
		logger.Errorf("%s - SendMessage: Failed to start stream", err)
		return err
	}
	defer s.Close()

	// marshall data to protobufs3 binary format
	bin, err := proto.Marshal(data)
	if err != nil {
		logger.Errorf("%s - SendMessage: Failed to marshal pb", err)
		return err
	}

	// Create Writer and write data to stream
	w := msgio.NewWriter(s)
	if err := w.WriteMsg(bin); err != nil {
		logger.Errorf("%s - SendMessage: Failed to write message to stream.", err)
		return err
	}
	return nil
}

// Stat returns the host stat info
func (hn *hostImpl) Stat() (map[string]string, error) {
	// Return Host Stat
	return map[string]string{
		"ID":        hn.host.ID().String(),
		"Status":    hn.status.String(),
		"MultiAddr": hn.host.Addrs()[0].String(),
	}, nil
}

// Serve handles incoming peer Addr Info
func (hn *hostImpl) Serve() {
	for {
		select {
		case mdnsPI := <-hn.mdnsPeerChan:
			if err := hn.Connect(mdnsPI); err != nil {
				hn.host.Peerstore().ClearAddrs(mdnsPI.ID)
				continue
			}

		case dhtPI := <-hn.dhtPeerChan:
			if err := hn.Connect(dhtPI); err != nil {
				hn.host.Peerstore().ClearAddrs(dhtPI.ID)
				continue
			}
		case <-hn.ctx.Done():
			return
		}
	}
}

// VerifyData verifies incoming p2p message data integrity
func (n *hostImpl) VerifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool {
	key, err := crypto.UnmarshalPublicKey(pubKeyData)
	if err != nil {
		logger.Errorf("%s - Failed to extract key from message key data", err)
		return false
	}

	// extract node id from the provided public key
	idFromKey, err := peer.IDFromPublicKey(key)
	if err != nil {
		logger.Errorf("%s - VerifyData: Failed to extract peer id from public key", err)
		return false
	}

	// verify that message author node id matches the provided node public key
	if idFromKey != peerId {
		logger.Errorf("%s - VerifyData: Node id and provided public key mismatch", err)
		return false
	}

	res, err := key.Verify(data, signature)
	if err != nil {
		logger.Errorf("%s - VerifyData: Error authenticating data", err)
		return false
	}
	return res
}
