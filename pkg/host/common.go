package host

import (
	"context"
	"crypto/ed25519"
	"errors"
	"fmt"

	"github.com/kataras/go-events"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/sonr/pkg/config"
	"google.golang.org/protobuf/proto"
)

// Config returns the configuration of the node
func (n *hostImpl) Config() *config.Config {
	return n.config
}

// Host returns the host of the node
func (hn *hostImpl) Host() host.Host {
	return hn.host
}

// HostID returns the ID of the Host
func (n *hostImpl) HostID() peer.ID {
	return n.host.ID()
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
		return nil, err
	}
	return ed25519.PrivateKey(buf), nil
}

// Role returns the role of the node
func (n *hostImpl) Role() config.Role {
	return n.config.Role
}

// // AuthenticateMessage Authenticates incoming p2p message
// func (n *hostImpl) AuthenticateMessage(msg proto.Message, metadata *t.Metadata) error {
// 	// store a temp ref to signature and remove it from message data
// 	// sign is a string to allow easy reset to zero-value (empty string)
// 	sign := metadata.Signature
// 	metadata.Signature = nil

// 	// marshall data without the signature to protobufs3 binary format
// 	buf, err := proto.Marshal(msg)
// 	if err != nil {
// 		return err
// 	}

// 	// restore sig in message data (for possible future use)
// 	metadata.Signature = sign

// 	// restore peer id binary format from base58 encoded node id data
// 	peerId, err := peer.Decode(metadata.NodeId)
// 	if err != nil {
// 		return err
// 	}

// 	// verify the data was authored by the signing peer identified by the public key
// 	// and signature included in the message
// 	return n.VerifyData(buf, []byte(sign), peerId, metadata.PublicKey)
// }

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

// NewTopic creates a new topic
func (n *hostImpl) NewTopic(name string, opts ...ps.TopicOpt) (*ps.Topic, *ps.TopicEventHandler, *ps.Subscription, error) {
	// Check if PubSub is Set
	if n.PubSub == nil {
		return nil, nil, nil, errors.New("NewTopic: Pubsub has not been set on SNRHost")
	}

	// Call Underlying Pubsub to Connect
	t, err := n.Join(name, opts...)
	if err != nil {
		return nil, nil, nil, err
	}

	// Create Event Handler
	h, err := t.EventHandler()
	if err != nil {
		return nil, nil, nil, err
	}

	// Create Subscriber
	s, err := t.Subscribe()
	if err != nil {
		return nil, nil, nil, err
	}
	return t, h, s, nil
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
func (h *hostImpl) Send(id peer.ID, p protocol.ID, data proto.Message) error {
	if !h.HasRouting() {
		return fmt.Errorf("Host does not have routing")
	}

	s, err := h.NewStream(h.ctx, id, p)
	if err != nil {
		return err
	}
	defer s.Close()

	// marshall data to protobufs3 binary format
	bin, err := proto.Marshal(data)
	if err != nil {
		return err
	}

	// Create Writer and write data to stream
	w := msgio.NewWriter(s)
	if err := w.WriteMsg(bin); err != nil {
		return err
	}
	return nil
}

// TODO
func (hn *hostImpl) Events() events.EventEmmiter {
	return events.New()
}

// // TODO
// func (hn *hostImpl) Peer() (*types.Peer, error) {
// 	return nil, nil
// }

// TODO
func (hn *hostImpl) SignData(data []byte) ([]byte, error) {
	return nil, nil
}

// // TODO
// func (hn *hostImpl) SignMessage(message proto.Message) ([]byte, error) {
// 	return nil, nil
// }

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

// VerifyData verifies incoming p2p message data integrity
func (n *hostImpl) VerifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) error {
	key, err := crypto.UnmarshalPublicKey(pubKeyData)
	if err != nil {
		return err
	}

	// extract node id from the provided public key
	idFromKey, err := peer.IDFromPublicKey(key)
	if err != nil {
		return err
	}

	// verify that message author node id matches the provided node public key
	if idFromKey != peerId {
		return err
	}

	_, err = key.Verify(data, signature)
	if err != nil {
		return err
	}
	return nil
}
