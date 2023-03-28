package host

import (
	"context"
	"errors"
	"time"

	ggio "github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/proto"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/pkg/node/config"
)

// A P2PHost is a host.Host with a private key, a channel of mDNS peers, a channel of DHT peers, a
// context, a map of topics, and a DHT and PubSub.
// @property host - The host is the main object of the libp2p library. It represents the local node and
// provides all the functionality to interact with the network.
// @property {string} accAddr - The address of the account that is being used to connect to the
// network.
// @property privKey - The private key of the host.
// @property mdnsPeerChan - This is a channel that will receive peer.AddrInfo objects from the mdns
// service.
// @property dhtPeerChan - A channel that will receive peer.AddrInfo objects when a peer is found via
// the DHT.
// @property ctx - The context of the P2PHost.
// @property topics - A map of topic names to the PubSub topic object.
// @property  - `host`: The libp2p host.
// @property  - `host`: The libp2p host.
type hostImpl struct {
	// Standard Node Implementation
	host     host.Host
	callback config.NodeCallback

	// Host and context
	privKey      crypto.PrivKey
	mdnsPeerChan chan peer.AddrInfo
	dhtPeerChan  <-chan peer.AddrInfo

	// Properties
	ctx    context.Context
	topics map[string]*ps.Topic

	*dht.IpfsDHT
	*ps.PubSub

	config *config.Config
}

// Initialize Creates a Sonr libp2p Host with the given config
func Initialize(config *config.Config) (config.PeerNode, error) {
	// Create Host and apply options
	hn := defaultNode(config)
	// Initialize Host
	if err := initializeHost(hn); err != nil {
		return nil, err
	}

	// Bootstrap DHT
	if err := hn.Bootstrap(hn.ctx); err != nil {
		return nil, err
	}

	// Connect to Bootstrap Nodes
	snrctx := local.NewContext()
	for _, pistr := range snrctx.BsMultiaddrs {
		if err := hn.Connect(pistr); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := setupRoutingDiscovery(hn); err != nil {
		return nil, err
	}
	return hn, nil
}


// PeerID returns the ID of the Host
func (n *hostImpl) PeerID() peer.ID {
	return n.host.ID()
}

// Connect connects with `peer.AddrInfo` if underlying Host is ready
func (hn *hostImpl) Connect(pi ...string) error {
	ctx, cancel := context.WithTimeout(hn.ctx, 10*time.Second)
	defer cancel()
	for _, p := range pi {
		pi, err := peer.AddrInfoFromString(p)
		if err != nil {
			return err
		}
		doneChan := make(chan struct{})
		go func() {
			err = hn.host.Connect(ctx, *pi)
			close(doneChan)
		}()
		select {
		case <-doneChan:
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return errors.New("timeout")
		}
	}
	return nil
}

// HandlePeerFound is to be called when new  peer is found
func (hn *hostImpl) HandlePeerFound(pi peer.AddrInfo) {
	hn.mdnsPeerChan <- pi
}

// MultiAddrs returns the MultiAddresses of the Host as single string
func (hn *hostImpl) MultiAddrs() string {
	maddrs := hn.host.Addrs()
	maddr := ma.Join(maddrs...)
	return maddr.String()
}

// NewStream opens a new stream to the peer with given peer id
func (n *hostImpl) NewStream(to peer.ID, protocol protocol.ID, msg proto.Message) error {
	stream, err := n.host.NewStream(context.Background(), to, protocol)
	if err != nil {
		return err
	}
	defer stream.Close()

	writer := ggio.NewFullWriter(stream)
	if err := writer.WriteMsg(msg); err != nil {
		return err
	}
	return nil
}

// JoinTopic creates a new topic
func (n *hostImpl) Publish(topic string, message []byte, opts ...ps.TopicOpt) error {
	ctx, cancel := context.WithTimeout(n.ctx, 10*time.Second)
	defer cancel()
	// Check if PubSub is Set
	if n.PubSub == nil {
		return nil
	}

	// Check if topic is valid
	t, ok := n.topics[topic]
	if ok {
		return t.Publish(ctx, message)
	}

	// Call Underlying Pubsub to Connect
	t, err := n.PubSub.Join(topic, opts...)
	if err != nil {
		return err
	}

	// Create Subscriber
	n.topics[topic] = t
	return t.Publish(ctx, message)
}

// SetStreamHandler sets the handler for a given protocol
func (n *hostImpl) SetStreamHandler(protocol protocol.ID, handler network.StreamHandler) {
	n.host.SetStreamHandler(protocol, handler)
}

// Join wraps around PubSub.Join and returns topic. Checks wether the host is ready before joining.
func (hn *hostImpl) Subscribe(topic string, handlers ...func(msg *ps.Message)) (*ps.Subscription, error) {
	// Check if PubSub is Set
	if hn.PubSub == nil {
		return nil, errors.New("Join: Pubsub has not been set on SNRHost")
	}

	// Check if topic is already joined
	if t, ok := hn.topics[topic]; ok {
		return t.Subscribe()
	}

	// Call Underlying Pubsub to Connect
	t, err := hn.PubSub.Join(topic)
	if err != nil {
		return nil, err
	}
	hn.topics[topic] = t

	// Subscribe to Topic
	sub, err := t.Subscribe()
	if err != nil {
		return nil, err
	}

	// Handle Subscription
	if len(handlers) > 0 {
		go hn.handleSubscription(sub, handlers[0])
	}
	return sub, nil
}

// handleSubscription handles the subscription to a topic
func (hn *hostImpl) handleSubscription(sub *ps.Subscription, handler func(msg *ps.Message)) {
	for {
		msg, err := sub.Next(hn.ctx)
		if err != nil {
			return
		}
		handler(msg)

		select {
		case <-hn.ctx.Done():
			return
		default:
		}
	}
}

// Closing the host.
func (hn *hostImpl) Close() error {
	err := hn.host.Close()
	if err != nil {
		return err
	}
	return nil
}
