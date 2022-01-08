package node

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/common"

	"github.com/sonr-io/core/wallet"
	"google.golang.org/protobuf/proto"
)

// AuthenticateId verifies UUID value and signature
func (h *node) AuthenticateId(id *wallet.UUID) (bool, error) {
	// Get local node's public key
	pubKey, err := wallet.DevicePubKey()
	if err != nil {
		logger.Errorf("%s - AuthenticateId: Failed to get local host's public key", err)
		return false, err
	}

	// verify UUID value
	result, err := pubKey.Verify([]byte(id.GetValue()), []byte(id.GetSignature()))
	if err != nil {
		logger.Errorf("%s - AuthenticateId: Failed to verify signature of UUID", err)
		return false, err
	}
	return result, nil
}

// AuthenticateMessage Authenticates incoming p2p message
func (n *node) AuthenticateMessage(msg proto.Message, metadata *common.Metadata) bool {
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
func (hn *node) Connect(pi peer.AddrInfo) error {
	// Check if host is ready
	if err := hn.HasRouting(); err != nil {
		logger.Warn("Connect: Underlying host is not ready, failed to call Connect()")
		return err
	}

	// Call Underlying Host to Connect
	return hn.Host.Connect(hn.ctx, pi)
}

// HandlePeerFound is to be called when new  peer is found
func (hn *node) HandlePeerFound(pi peer.AddrInfo) {
	hn.mdnsPeerChan <- pi
}

// HasRouting returns no-error if the host is ready for connect
func (h *node) HasRouting() error {
	if h.IpfsDHT == nil || h.Host == nil {
		return ErrRoutingNotSet
	}
	return nil
}

// Join wraps around PubSub.Join and returns topic. Checks wether the host is ready before joining.
func (hn *node) Join(topic string, opts ...ps.TopicOpt) (*ps.Topic, error) {
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
func (n *node) NewStream(ctx context.Context, pid peer.ID, pids ...protocol.ID) (network.Stream, error) {
	return n.Host.NewStream(ctx, pid, pids...)
}

// Router returns the host node Peer Routing Function
func (hn *node) Router(h host.Host) (routing.PeerRouting, error) {
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

// SetStreamHandler sets the handler for a given protocol
func (n *node) SetStreamHandler(protocol protocol.ID, handler network.StreamHandler) {
	n.Host.SetStreamHandler(protocol, handler)
}

// SendMessage writes a protobuf go data object to a network stream
func (h *node) SendMessage(id peer.ID, p protocol.ID, data proto.Message) error {
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

// SignData signs an outgoing p2p message payload
func (n *node) SignData(data []byte) ([]byte, error) {
	// Get local node's private key
	res, err := wallet.Sign(data)
	if err != nil {
		logger.Errorf("%s - SignData: Failed to get local host's private key", err)
		return nil, err
	}
	return res, nil
}

// SignMessage signs an outgoing p2p message payload
func (n *node) SignMessage(message proto.Message) ([]byte, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		logger.Errorf("%s - SignMessage: Failed to Sign Message", err)
		return nil, err
	}
	return n.SignData(data)
}

// Stat returns the host stat info
func (hn *node) Stat() (map[string]string, error) {
	// Return Host Stat
	return map[string]string{
		"ID":        hn.ID().String(),
		"Status":    hn.status.String(),
		"MultiAddr": hn.Addrs()[0].String(),
	}, nil
}

// Serve handles incoming peer Addr Info
func (hn *node) Serve() {
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
func (n *node) VerifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool {
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
