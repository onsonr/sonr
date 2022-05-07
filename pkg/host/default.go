package host

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"errors"

	"github.com/libp2p/go-libp2p"
	cmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/sonr/pkg/config"
	device "github.com/sonr-io/sonr/pkg/fs"
	t "go.buf.build/grpc/go/sonr-io/core/types/v1"
	types "go.buf.build/grpc/go/sonr-io/core/types/v1"
	"google.golang.org/protobuf/proto"
)

// defaultHostImpl type - a p2p host implementing one or more p2p protocols
type defaultHostImpl struct {
	// Standard Node Implementation
	host host.Host
	SonrHost
	role device.Role

	// Host and context
	connection   types.Connection
	privKey      crypto.PrivKey
	mdnsPeerChan chan peer.AddrInfo
	dhtPeerChan  <-chan peer.AddrInfo

	// Properties
	ctx context.Context

	*dht.IpfsDHT
	*ps.PubSub

	// State
	flag   uint64
	Chn    chan bool
	status HostStatus
}

// NewMachineHost Creates a Sonr libp2p Host with the given config
func NewMachineHost(ctx context.Context, config *config.Config) (SonrHost, error) {
	var err error
	// Create the host.
	hn := &defaultHostImpl{
		ctx:          ctx,
		status:       Status_IDLE,
		mdnsPeerChan: make(chan peer.AddrInfo),
		role:         config.Role,
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
func (hn *defaultHostImpl) Host() host.Host {
	return hn.host
}

// HostID returns the ID of the Host
func (n *defaultHostImpl) HostID() peer.ID {
	return n.host.ID()
}

// Ping sends a ping to the peer
func (n *defaultHostImpl) Ping(pid string) error {
	return nil
}

// PrivateKey returns the private key of the node
func (n *defaultHostImpl) PrivateKey() (ed25519.PrivateKey, error) {
	// Get Raw Private Key
	buf, err := n.privKey.Raw()
	if err != nil {
		logger.Errorf("%s - Failed to get Raw Private Key", err)
		return nil, err
	}
	return ed25519.PrivateKey(buf), nil
}

// Role returns the role of the node
func (n *defaultHostImpl) Role() device.Role {
	return n.role
}

// AuthenticateMessage Authenticates incoming p2p message
func (n *defaultHostImpl) AuthenticateMessage(msg proto.Message, metadata *t.Metadata) bool {
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
func (hn *defaultHostImpl) Connect(pi peer.AddrInfo) error {
	// Check if host is ready
	if err := hn.HasRouting(); err != nil {
		logger.Warn("Connect: Underlying host is not ready, failed to call Connect()")
		return err
	}

	// Call Underlying Host to Connect
	return hn.host.Connect(hn.ctx, pi)
}

// HandlePeerFound is to be called when new  peer is found
func (hn *defaultHostImpl) HandlePeerFound(pi peer.AddrInfo) {
	hn.mdnsPeerChan <- pi
}

// HasRouting returns no-error if the host is ready for connect
func (h *defaultHostImpl) HasRouting() error {
	if h.IpfsDHT == nil || h.host == nil {
		return errors.New("Host is not ready")
	}
	return nil
}

// Join wraps around PubSub.Join and returns topic. Checks wether the host is ready before joining.
func (hn *defaultHostImpl) Join(topic string, opts ...ps.TopicOpt) (*ps.Topic, error) {
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
func (n *defaultHostImpl) NewStream(ctx context.Context, pid peer.ID, pids ...protocol.ID) (network.Stream, error) {
	return n.host.NewStream(ctx, pid, pids...)
}

// NewTopic creates a new topic
func (n *defaultHostImpl) NewTopic(name string, opts ...ps.TopicOpt) (*ps.Topic, *ps.TopicEventHandler, *ps.Subscription, error) {
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
func (hn *defaultHostImpl) Router(h host.Host) (routing.PeerRouting, error) {
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
func (hn *defaultHostImpl) Pubsub() *ps.PubSub {
	return hn.PubSub
}

// Routing returns the host node Peer Routing Function
func (hn *defaultHostImpl) Routing() routing.Routing {
	return hn.IpfsDHT
}

// SetStreamHandler sets the handler for a given protocol
func (n *defaultHostImpl) SetStreamHandler(protocol protocol.ID, handler network.StreamHandler) {
	n.host.SetStreamHandler(protocol, handler)
}

// SendMessage writes a protobuf go data object to a network stream
func (h *defaultHostImpl) SendMessage(id peer.ID, p protocol.ID, data proto.Message) error {
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
func (hn *defaultHostImpl) Stat() (map[string]string, error) {
	// Return Host Stat
	return map[string]string{
		"ID":        hn.host.ID().String(),
		"Status":    hn.status.String(),
		"MultiAddr": hn.host.Addrs()[0].String(),
	}, nil
}

// Serve handles incoming peer Addr Info
func (hn *defaultHostImpl) Serve() {
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
func (n *defaultHostImpl) VerifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool {
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

// createDHTDiscovery is a Helper Method to initialize the DHT Discovery
func (hn *defaultHostImpl) createDHTDiscovery(c *config.Config) error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := dsc.NewRoutingDiscovery(hn.IpfsDHT)
	dsc.Advertise(hn.ctx, routingDiscovery, c.Libp2pRendezvous, c.Libp2pTTL)

	// Create Pub Sub
	hn.PubSub, err = ps.NewGossipSub(hn.ctx, hn.host, ps.WithDiscovery(routingDiscovery))
	if err != nil {
		hn.SetStatus(Status_FAIL)
		logger.Errorf("%s - Failed to Create new Gossip Sub", err)
		return err
	}

	// Handle DHT Peers
	hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, c.Libp2pRendezvous, c.Libp2pTTL)
	if err != nil {
		hn.SetStatus(Status_FAIL)
		logger.Errorf("%s - Failed to create FindPeers Discovery channel", err)
		return err
	}
	hn.SetStatus(Status_READY)
	return nil
}

// TODO Migrate MDNS Service to latesat libp2p spec
// // createMdnsDiscovery is a Helper Method to initialize the MDNS Discovery
// func (hn *hostImpl) createMdnsDiscovery(c *config.Config) {
// 	if hn.Role() == device.Role_MOTOR {
// 		// Create MDNS Service
// 		ser := mdns.NewMdnsService(hn.host, c.Libp2pRendezvous)

// 		// Handle Events
// 		ser.RegisterNotifee(hn)
// 	}
// }
