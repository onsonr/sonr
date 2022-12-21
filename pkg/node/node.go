package node

import (
	"context"
	"fmt"
	"os"
	"strings"

	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/ipfs/kubo/core"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/sonr-hq/sonr/pkg/common"
)

// Node represents a Interface to the IPFS node
type Node struct {
	icore.CoreAPI
	node     *core.IpfsNode
	ctx      context.Context
	callback common.MotorCallback
	topics   []string
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
		topics:  make([]string, 0),
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
