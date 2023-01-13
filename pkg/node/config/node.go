package config

import (
	"github.com/gogo/protobuf/proto"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	ps "github.com/libp2p/go-libp2p-pubsub"
)

// `IPFSNode` is an interface that defines the methods that a Highway node must implement.
// @property Add - This is the function that adds a file to the IPFS network.
// @property {error} Connect - Connects to a peer
// @property CoreAPI - This is the IPFS Core API.
// @property Get - Get a file from the network
// @property {string} MultiAddr - The multiaddr of the node
// @property PeerID - The peer ID of the node
// @property GetDecrypted - This is a method that takes a cid and a public key and returns the
// decrypted file.
// @property AddEncrypted - Add a file to the network, encrypted with the given public key.
type IPFSNode interface {
	// Add a file to the network
	Add(data []byte) (string, error)

	// Connect to a peer
	Connect(peers ...string) error

	// Get the IPFS Core API
	CoreAPI() icore.CoreAPI

	// Get a file from the network
	Get(hash string) ([]byte, error)

	// MultiAddr returns the multiaddr of the node
	MultiAddrs() string

	// PeerID returns the peer ID of the node
	PeerID() peer.ID

	// GetDecrypted takes a cid and a public key and returns the decrypted file.
	GetDecrypted(cidStr string, pubKey []byte) ([]byte, error)

	// AddEncrypted adds a file to the network, encrypted with the given public key.
	AddEncrypted(file []byte, pubKey []byte) (string, error)
}

// `P2PNode` is an interface that defines the methods that a node must implement to be used by the
// Motor library.
// @property PeerID - The peer ID of the node.
// @property {error} Connect - Connect to a peer
// @property {string} MultiAddrs - The multiaddresses of the node.
// @property {error} NewStream - This is the function that allows you to create a new stream to a peer.
// @property {error} Publish - Publish a message to a topic.
// @property SetStreamHandler - This is a function that sets the handler for a given protocol.
// @property Subscribe - Subscribe to a topic.
// @property {error} Close - Closes the node.
type P2PNode interface {
	// PeerID returns the peer ID of the node
	PeerID() peer.ID

	// Connect to a peer
	Connect(pi interface{}) error

	// MultiAddrs returns the multiaddresses of the node
	MultiAddrs() string
	NewStream(to peer.ID, protocol protocol.ID, msg proto.Message) error
	Publish(topic string, message []byte, opts ...ps.TopicOpt) error
	SetStreamHandler(protocol protocol.ID, handler network.StreamHandler)
	Subscribe(topic string, handlers ...func(msg *ps.Message)) (*ps.Subscription, error)
	Close() error
}
