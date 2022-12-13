package node

import (
	"errors"
	"fmt"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ps "github.com/libp2p/go-libp2p-pubsub"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/discovery"

	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	dscrouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-msgio"
)

// Host returns the host of the node
func (hn *hostImpl) Host() host.Host {
	return hn.host
}

// HostID returns the ID of the Host
func (n *hostImpl) ID() (ID, error) {
	return ParseID(n.host.ID())
}

// createDHTDiscovery is a Helper Method to initialize the DHT Discovery
func (hn *hostImpl) createDHTDiscovery() error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := dscrouting.NewRoutingDiscovery(hn.idht)
	routingDiscovery.Advertise(hn.ctx, libp2pRendevouz)

	// Create Pub Sub
	hn.ps, err = pubsub.NewGossipSub(hn.ctx, hn.host, pubsub.WithDiscovery(routingDiscovery))
	if err != nil {
		return err
	}

	// Handle DHT Peers
	hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, libp2pRendevouz, discovery.Limit(10))
	if err != nil {
		return err
	}
	return nil
}

func (hn *hostImpl) Close() error {
	err := hn.host.Close()
	if err != nil {
		return err
	}
	return nil
}

// createMdnsDiscovery is a Helper Method to initialize the MDNS Discovery
func (hn *hostImpl) createMdnsDiscovery() {

	fmt.Println("Starting MDNS Discovery...")
	// Create MDNS Service
	ser := mdns.NewMdnsService(hn.host, libp2pRendevouz, hn)
	if err := ser.Start(); err != nil {
		fmt.Println("Error starting MDNS Service: ", err)
		return
	}

}

// Connect connects with `peer.AddrInfo` if underlying Host is ready
func (hn *hostImpl) Connect(pi peer.AddrInfo) error {
	// Call Underlying Host to Connect
	return hn.host.Connect(hn.ctx, pi)
}

// HandlePeerFound is to be called when new  peer is found
func (hn *hostImpl) HandlePeerFound(pi peer.AddrInfo) {
	hn.mdnsPeerChan <- pi
}

// JoinTopic wraps around PubSub.Join and returns topic. Checks wether the host is ready before joining.
func (hn *hostImpl) JoinTopic(topic string, opts ...ps.TopicOpt) (*ps.Topic, error) {
	// Check if PubSub is Set
	if hn.ps == nil {
		return nil, errors.New("Join: Pubsub has not been set on SNRHost")
	}

	// Check if topic is valid
	if topic == "" {
		return nil, errors.New("Join: Empty topic string provided to Join for host.Pubsub")
	}

	// Call Underlying Pubsub to Connect
	return hn.ps.Join(topic, opts...)
}

// NewStream opens a new stream to the peer with given peer id
func (n *hostImpl) NewStream(pid peer.ID, pids ...protocol.ID) (network.Stream, error) {
	return n.host.NewStream(n.ctx, pid, pids...)
}

// PubSub returns the host node PubSub Function
func (hn *hostImpl) Pubsub() *ps.PubSub {
	return hn.ps
}

// Router returns the host node Peer Routing Function
func (hn *hostImpl) Router(h host.Host) (routing.PeerRouting, error) {
	// Create DHT
	kdht, err := dht.New(hn.ctx, h)
	if err != nil {
		return nil, err
	}

	// Set Properties
	hn.idht = kdht

	// Setup Properties
	return hn.idht, nil
}

// Routing returns the host node Peer Routing Function
func (hn *hostImpl) Routing() routing.Routing {
	return hn.idht
}

// SetStreamHandler sets the handler for a given protocol
func (n *hostImpl) SetStreamHandler(protocol protocol.ID, handler network.StreamHandler) {
	n.host.SetStreamHandler(protocol, handler)
}

// SendMessage writes a protobuf go data object to a network stream
func (h *hostImpl) Send(id peer.ID, p protocol.ID, data []byte) error {
	s, err := h.NewStream(id, p)
	if err != nil {
		return err
	}
	defer s.Close()

	// Create Writer and write data to stream
	w := msgio.NewWriter(s)
	if err := w.WriteMsg(data); err != nil {
		return err
	}
	return nil
}
