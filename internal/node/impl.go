package node

import (
	"context"
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
func (n *hostImpl) HostID() peer.ID {
	return n.host.ID()
}

// Address returns the address of the underlying wallet
func (h *hostImpl) Address() string {
	return h.accAddr
}

// createDHTDiscovery is a Helper Method to initialize the DHT Discovery
func (hn *hostImpl) createDHTDiscovery() error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := dscrouting.NewRoutingDiscovery(hn.IpfsDHT)
	routingDiscovery.Advertise(hn.ctx, libp2pRendevouz)

	// Create Pub Sub
	hn.PubSub, err = pubsub.NewGossipSub(hn.ctx, hn.host, pubsub.WithDiscovery(routingDiscovery))
	if err != nil {
		hn.fsm.SetState(Status_FAIL)
		return err
	}

	// Handle DHT Peers
	hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, libp2pRendevouz, discovery.Limit(10))
	if err != nil {
		hn.fsm.SetState(Status_FAIL)
		return err
	}

	hn.fsm.SetState(Status_READY)
	return nil
}

func (hn *hostImpl) Close() error {
	err := hn.host.Close()
	if err != nil {
		hn.fsm.SetState(Status_FAIL)
		return err
	}

	hn.fsm.SetState(Status_STANDBY)

	return nil
}

// NeedsWait checks if state is Resumed or Paused and blocks channel if needed
func (hn *hostImpl) NeedsWait() {
	<-hn.fsm.Chn
}

/*
Stops the libp2p host, dhcp, and sets the host status to IDLE
*/
func (hn *hostImpl) Stop() error {
	err := hn.host.Close()
	if err != nil {
		hn.fsm.SetState(Status_FAIL)
		return err
	}
	hn.Pause()

	return nil
}

/*
Stops the libp2p host, dhcp, and sets the host status to ready
*/
func (hn *hostImpl) Pause() error {
	defer hn.fsm.PauseOperation()
	hn.fsm.SetState(Status_STANDBY)
	return nil
}

func (hn *hostImpl) Resume() error {
	defer hn.fsm.ResumeOperation()
	hn.fsm.SetState(Status_STANDBY)

	return nil
}

func (hn *hostImpl) Status() HostStatus {
	return hn.fsm.CurrentStatus
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
	// Check if host is ready
	if !hn.HasRouting() {
		return fmt.Errorf("Host does not have routing")
	}

	// Call Underlying Host to Connect
	return hn.host.Connect(hn.ctx, pi)
}

// HandlePeerFound is to be called when new  peer is found
func (hn *hostImpl) HandlePeerFound(pi peer.AddrInfo) {
	hn.mdnsPeerChan <- pi
}

// HasRouting returns no-error if the host is ready for connect
func (h *hostImpl) HasRouting() bool {
	return h.IpfsDHT != nil && h.host != nil
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



// Router returns the host node Peer Routing Function
func (hn *hostImpl) Router(h host.Host) (routing.PeerRouting, error) {
	// Create DHT
	kdht, err := dht.New(hn.ctx, h)
	if err != nil {
		hn.fsm.SetState(Status_FAIL)
		return nil, err
	}

	// Set Properties
	hn.IpfsDHT = kdht

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
func (h *hostImpl) Send(id peer.ID, p protocol.ID, data []byte) error {
	if !h.HasRouting() {
		return fmt.Errorf("Host does not have routing")
	}

	s, err := h.NewStream(h.ctx, id, p)
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

type HostStat struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	MultiAddr string `json:"multi_addr"`
}

// Stat returns the host stat info
func (hn *hostImpl) Stat() HostStat {
	// Return Host Stat
	return HostStat{
		ID:        hn.host.ID().String(),
		Status:    string(hn.fsm.CurrentStatus),
		MultiAddr: hn.host.Addrs()[0].String(),
	}
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
