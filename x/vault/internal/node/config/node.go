package config

import (
	"berty.tech/go-orbit-db/iface"
	"github.com/gogo/protobuf/proto"
	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

// `Node` is an interface that defines the methods that a node must implement to be used by the
// Motor library.
// @property PeerID - The peer ID of the node.
// @property {error} Connect - Connect to a peer
// @property {string} MultiAddrs - The multiaddr of the node
// @property {error} Close - Close the node
type Node interface {
	// PeerID returns the peer ID of the node
	PeerID() peer.ID

	// Connect to a peer
	Connect(peers ...string) error

	// MultiAddrs returns the multiaddr of the node
	MultiAddrs() string

	// Close the node
	Close() error
}

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
	Node

	// Get the IPFS Core API
	CoreAPI() icore.CoreAPI

	// Add a file to the network
	Add(data []byte) (string, error)

	// AddPath adds a file to the network
	AddPath(path string) (string, error)

	// Get a file from the network
	Get(hash string) ([]byte, error)

	// GetPath gets a file from the network
	GetPath(hash string) (map[string]files.Node, error)

	// It's returning a DocumentStore for the given username.
	LoadDocsStore(username string, opts *iface.CreateDocumentDBOptions) (iface.DocumentStore, error)

	// It's returning a DocumentStore for the given username.
	LoadEventLogStore(username string) (iface.EventLogStore, error)

	// It's returning a KeyValueStore for the given username.
	LoadKeyValueStore(username string) (iface.KeyValueStore, error)
}

// `PeerNode` is an interface that defines the methods that a node must implement to be used by the
// Motor library.
// @property PeerID - The peer ID of the node.
// @property {error} Connect - Connect to a peer
// @property {string} MultiAddrs - The multiaddresses of the node.
// @property {error} NewStream - This is the function that allows you to create a new stream to a peer.
// @property {error} Publish - Publish a message to a topic.
// @property SetStreamHandler - This is a function that sets the handler for a given protocol.
// @property Subscribe - Subscribe to a topic.
// @property {error} Close - Closes the node.
type PeerNode interface {
	Node

	// NewStream creates a new stream to a peer
	NewStream(to peer.ID, protocol protocol.ID, msg proto.Message) error

	// SetStreamHandler sets the handler for a given protocol
	SetStreamHandler(protocol protocol.ID, handler network.StreamHandler)

	// Publish a message to a topic
	Publish(topic string, message []byte, opts ...ps.TopicOpt) error

	// Subscribe to a topic
	Subscribe(topic string, handlers ...func(msg *ps.Message)) (*ps.Subscription, error)
}
