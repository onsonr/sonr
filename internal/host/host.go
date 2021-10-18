package host

import (
	"context"
	"errors"

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
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/keychain"
	"github.com/sonr-io/core/internal/wallet"
	"google.golang.org/protobuf/proto"
)

// Transfer Emission Events
const (
	Event_STATUS = "host-status"
)

// SNRHost is the host wrapper for the Sonr Network
type SNRHost struct {
	host.Host
	mdnsPeerChan chan peer.AddrInfo
	dhtPeerChan  <-chan peer.AddrInfo

	// Properties
	ctx        context.Context
	privKey    crypto.PrivKey
	connection common.Connection

	// State
	status SNRHostStatus

	// Discovery
	*dht.IpfsDHT
	*ps.PubSub
}

// NewHost creates a new host
func NewHost(ctx context.Context, options ...HostOption) (*SNRHost, error) {
	// Initialize DHT
	opts := defaultHostOptions()
	hn, err := opts.Apply(ctx, options...)
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
		libp2p.DefaultListenAddrs,
		libp2p.Routing(hn.Router),
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		logger.Error("NewHost: Failed to create libp2p host", err)
		return nil, err
	}
	hn.SetStatus(Status_CONNECTING)

	// Bootstrap DHT
	if err := hn.Bootstrap(context.Background()); err != nil {
		logger.Error("Failed to Bootstrap KDHT to Host", err)
		hn.SetStatus(Status_FAIL)
		return nil, err
	}

	// Connect to Bootstrap Nodes
	for _, pi := range opts.BootstrapPeers {
		if err := hn.Connect(pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := hn.createDHTDiscovery(opts); err != nil {
		logger.Fatal("Could not start DHT Discovery", err)
		hn.SetStatus(Status_FAIL)
		return nil, err
	}

	// Initialize Discovery for MDNS
	hn.createMdnsDiscovery(opts)
	hn.SetStatus(Status_READY)
	go hn.Serve()
	return hn, nil
}

// AuthenticateId verifies UUID value and signature
func (h *SNRHost) AuthenticateId(id *wallet.UUID) (bool, error) {
	// Get local node's public key
	pubKey, err := keychain.Primary.GetPubKey(keychain.Account)
	if err != nil {
		logger.Error("AuthenticateId: Failed to get local host's public key", err)
		return false, err
	}

	// verify UUID value
	result, err := pubKey.Verify([]byte(id.GetValue()), []byte(id.GetSignature()))
	if err != nil {
		logger.Error("AuthenticateId: Failed to verify signature of UUID", err)
		return false, err
	}
	return result, nil
}

// AuthenticateMessage Authenticates incoming p2p message
func (n *SNRHost) AuthenticateMessage(msg proto.Message, metadata *common.Metadata) bool {
	// store a temp ref to signature and remove it from message data
	// sign is a string to allow easy reset to zero-value (empty string)
	sign := metadata.Signature
	metadata.Signature = nil

	// marshall data without the signature to protobufs3 binary format
	buf, err := proto.Marshal(msg)
	if err != nil {
		logger.Error("AuthenticateMessage: Failed to marshal Protobuf Message.", err)
		return false
	}

	// restore sig in message data (for possible future use)
	metadata.Signature = sign

	// restore peer id binary format from base58 encoded node id data
	peerId, err := peer.Decode(metadata.NodeId)
	if err != nil {
		logger.Error("AuthenticateMessage: Failed to decode node id from base58.", err)
		return false
	}

	// verify the data was authored by the signing peer identified by the public key
	// and signature included in the message
	return n.VerifyData(buf, []byte(sign), peerId, metadata.PublicKey)
}

// Close closes the underlying host
func (hn *SNRHost) Close() error {
	hn.SetStatus(Status_CLOSED)
	hn.IpfsDHT.Close()
	return hn.Host.Close()
}

// Connect connects with `peer.AddrInfo` if underlying Host is ready
func (hn *SNRHost) Connect(pi peer.AddrInfo) error {
	// Check if host is ready
	if err := hn.HasRouting(); err != nil {
		logger.Warn("Connect: Underlying host is not ready, failed to call Connect()")
		return err
	}

	// Call Underlying Host to Connect
	return hn.Host.Connect(hn.ctx, pi)
}

// HandlePeerFound is to be called when new  peer is found
func (hn *SNRHost) HandlePeerFound(pi peer.AddrInfo) {
	hn.mdnsPeerChan <- pi
}

// HasRouting returns no-error if the host is ready for connect
func (h *SNRHost) HasRouting() error {
	if h.IpfsDHT == nil || h.Host == nil {
		return ErrRoutingNotSet
	}
	return nil
}

// Join wraps around PubSub.Join and returns topic. Checks wether the host is ready before joining.
func (hn *SNRHost) Join(topic string, opts ...ps.TopicOpt) (*ps.Topic, error) {
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
	logger.Debug("Router: Host and DHT have been set for SNRNode")

	// Setup Properties
	return hn.IpfsDHT, nil
}

// SendMessage writes a protobuf go data object to a network stream
func (h *SNRHost) SendMessage(id peer.ID, p protocol.ID, data proto.Message) error {
	err := h.HasRouting()
	if err != nil {
		return err
	}

	s, err := h.NewStream(h.ctx, id, p)
	if err != nil {
		logger.Error("SendMessage: Failed to start stream", err)
		return err
	}
	defer s.Close()

	// marshall data to protobufs3 binary format
	bin, err := proto.Marshal(data)
	if err != nil {
		logger.Error("SendMessage: Failed to marshal pb", err)
		return err
	}

	// Create Writer and write data to stream
	w := msgio.NewWriter(s)
	if err := w.WriteMsg(bin); err != nil {
		logger.Error("SendMessage: Failed to write message to stream.", err)
		return err
	}
	return nil
}

// SignData signs an outgoing p2p message payload
func (n *SNRHost) SignData(data []byte) ([]byte, error) {
	// Get local node's private key
	res, err := keychain.Primary.SignWith(keychain.Account, data)
	if err != nil {
		logger.Error("SignData: Failed to get local host's private key", err)
		return nil, err
	}
	return res, nil
}

// SignMessage signs an outgoing p2p message payload
func (n *SNRHost) SignMessage(message proto.Message) ([]byte, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		logger.Error("SignMessage: Failed to Sign Message", err)
		return nil, err
	}
	return n.SignData(data)
}

// Stat returns the host stat info
func (hn *SNRHost) Stat() (map[string]string, error) {
	// Return Host Stat
	return map[string]string{
		"ID":        hn.ID().String(),
		"Status":    hn.status.String(),
		"MultiAddr": hn.Addrs()[0].String(),
	}, nil
}

// Serve handles incoming peer Addr Info
func (hn *SNRHost) Serve() {
	for {
		select {
		case mdnsPI := <-hn.mdnsPeerChan:
			if err := hn.Connect(mdnsPI); err != nil {
				hn.Peerstore().ClearAddrs(mdnsPI.ID)
				continue
			}

		case dhtPI := <-hn.dhtPeerChan:
			if err := hn.Connect(dhtPI); err != nil {
				hn.Peerstore().ClearAddrs(dhtPI.ID)
				continue
			}
		case <-hn.ctx.Done():
			return
		}
	}
}

// VerifyData verifies incoming p2p message data integrity
func (n *SNRHost) VerifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool {
	key, err := crypto.UnmarshalPublicKey(pubKeyData)
	if err != nil {
		logger.Error("Failed to extract key from message key data", err)
		return false
	}

	// extract node id from the provided public key
	idFromKey, err := peer.IDFromPublicKey(key)
	if err != nil {
		logger.Error("VerifyData: Failed to extract peer id from public key", err)
		return false
	}

	// verify that message author node id matches the provided node public key
	if idFromKey != peerId {
		logger.Error("VerifyData: Node id and provided public key mismatch", err)
		return false
	}

	res, err := key.Verify(data, signature)
	if err != nil {
		logger.Error("VerifyData: Error authenticating data", err)
		return false
	}
	return res
}
