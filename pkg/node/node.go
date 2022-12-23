package node

import (
	"context"
	"fmt"
	"os"
	"strings"

	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/ipfs/kubo/core"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/sonr-hq/sonr/pkg/common"
	cv1 "github.com/sonr-hq/sonr/pkg/common"
)

// Node represents a Interface to the IPFS node
type Node struct {
	icore.CoreAPI
	node     *core.IpfsNode
	ctx      context.Context
	callback common.NodeCallback
	config   *NodeConfig
}

// New creates a new node with the given options
func New(ctx context.Context, options ...NodeOption) (*Node, error) {
	// Apply the options
	c := defaultNodeConfig()
	for _, option := range options {
		option(c)
	}
	// Spawn a local peer using a temporary path, for testing purposes
	ipfsA, nodeA, err := c.spawnEphemeral(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to spawn ephemeral node: %s", err)
	}

	// Connect to the bootstrap nodes
	err = c.connectToPeers(ctx, ipfsA, c.BootstrapMultiaddrs)
	if err != nil {
		return nil, err
	}

	// Create the node
	n := &Node{
		CoreAPI: ipfsA,
		node:    nodeA,
		ctx:     ctx,
		config:  c,
	}
	return n, nil
}

// Get returns a file from the network given its CID
func (n *Node) Get(cidString string) ([]byte, error) {
	ctx, cancel := context.WithCancel(n.ctx)
	defer cancel()
	cid := icorepath.New(cidString)

	// Get the file from the network
	fileNode, err := n.CoreAPI.Unixfs().Get(ctx, cid)
	if err != nil {
		return nil, err
	}

	// Create a temporary directory to store the file
	outputBasePath, err := os.MkdirTemp("", "example")
	if err != nil {
		return nil, err
	}

	// Set the output path
	outputPath := outputBasePath + strings.Split(cidString, "/")[2]

	// Write the file to the output path
	err = files.WriteTo(fileNode, outputPath)
	if err != nil {
		return nil, err
	}

	// Read the file
	file, err := os.ReadFile(outputPath)
	if err != nil {
		return nil, err
	}

	// Delete the temporary directory
	err = os.RemoveAll(outputBasePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Add adds a file to the network
func (n *Node) Add(file []byte) (string, error) {
	ctx, cancel := context.WithCancel(n.ctx)
	defer cancel()

	// Generate a temporary directory
	inputBasePath, err := os.MkdirTemp("", "example")
	if err != nil {
		return "", err
	}

	// Write contents to a temporary file
	inputPath := fmt.Sprintf("%s/%s", inputBasePath, "file")
	err = os.WriteFile(inputPath, file, 0644)
	if err != nil {
		return "", err
	}

	// Get File Node
	fileNode, err := getUnixfsNode(inputPath)
	if err != nil {
		return "", err
	}

	// Add the file to the network
	cid, err := n.CoreAPI.Unixfs().Add(ctx, fileNode)
	if err != nil {
		return "", err
	}
	return cid.String(), nil
}

// Connect connects to a peer with a given multiaddress
func (n *Node) Connect(addrStrs ...string) error {
	return n.config.connectToPeers(n.ctx, n.CoreAPI, addrStrs)
}

// ID returns the node's ID
func (n *Node) ID() peer.ID {
	return n.node.Identity
}

// MultiAddr returns the node's multiaddress as a string
func (n *Node) MultiAddr() string {
	return fmt.Sprintf("/ip4/127.0.0.1/udp/4010/p2p/%s", n.node.Identity.String())
}

// Peer returns the node's peer info
func (n *Node) Peer() *cv1.Peer {
	return &cv1.Peer{
		PeerId:    n.ID().String(),
		Multiaddr: n.MultiAddr(),
		Type:      n.config.PeerType,
	}
}

// Publish publishes a message to a topic
func (n *Node) Publish(topic string, message []byte) error {
	errChan := make(chan error)
	go func() {
		err := n.PubSub().Publish(n.ctx, topic, message)
		errChan <- err
	}()
	select {
	case err := <-errChan:
		return err
	case <-n.ctx.Done():
		return n.ctx.Err()
	}
}

// Subscribe subscribes to a topic
func (n *Node) Subscribe(topic string, initialPeers ...string) (icore.PubSubSubscription, error) {
	err := n.Connect(initialPeers...)
	if err != nil {
		return nil, err
	}
	sub, err := n.PubSub().Subscribe(n.ctx, topic, options.PubSub.Discover(true))
	if err != nil {
		return nil, err
	}
	return sub, nil
}

// ListTopics lists the topics the node is subscribed to
func (n *Node) ListTopics() ([]string, error) {
	return n.PubSub().Ls(n.ctx)
}

// ListPeers lists the peers the node is connected to on a given topic
func (n *Node) ListPeers(topic string) ([]peer.ID, error) {
	return n.PubSub().Peers(n.ctx, options.PubSub.Topic(topic))
}
