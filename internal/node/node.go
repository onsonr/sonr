package node

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/ipfs/kubo/core"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-hq/sonr/core/common"
)

// Node represents a Interface to the IPFS node
type Node struct {
	api icore.CoreAPI
	*core.IpfsNode
	ctx      context.Context
	callback common.MotorCallback
	topics   []string
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
		api:      ipfsA,
		IpfsNode: nodeA,
		ctx:      ctx,
		topics:   make([]string, 0),
	}
	return n, nil
}

// Get returns a file from the network given its CID
func (n *Node) Get(cidString string) ([]byte, error) {
	ctx, cancel := context.WithCancel(n.ctx)
	defer cancel()
	cid := icorepath.New(cidString)

	// Get the file from the network
	fileNode, err := n.api.Unixfs().Get(ctx, cid)
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
	cid, err := n.api.Unixfs().Add(ctx, fileNode)
	if err != nil {
		return "", err
	}
	return cid.String(), nil
}

// Connect connects to a peer with a given multiaddress
func (n *Node) Connect(addrInfo peer.AddrInfo) error {

	err := n.api.Swarm().Connect(n.ctx, addrInfo)
	if err != nil {
		log.Printf("failed to connect to %s: %s", addrInfo.ID, err)
	}
	return nil
}

// ID returns the node's ID
func (n *Node) ID() peer.ID {
	return n.IpfsNode.Identity
}

// MultiAddr returns the node's multiaddress as a string
func (n *Node) AddrInfo() peer.AddrInfo {
	// Connection Addrs
	addrs, err := n.api.Swarm().LocalAddrs(n.ctx)
	if err != nil {
		return peer.AddrInfo{}
	}

	return peer.AddrInfo{
		ID:    n.IpfsNode.Identity,
		Addrs: addrs,
	}

}

// // Send sends a message to a peer. This is by using the p2p node to create a new stream
func (n *Node) Send(peerID string, message []byte, pid protocol.ID) error {
	// Create a new stream
	stream, err := n.IpfsNode.PeerHost.NewStream(n.ctx, peer.ID(peerID), pid)
	if err != nil {
		return err
	}

	wr := msgio.NewWriter(stream)
	err = wr.WriteMsg(message)
	if err != nil {
		return err
	}
	return nil
}

// HandleProtocol Sets a Stream Handler for the underlying PeerHost
func (n *Node) SetStreamHandler(pid protocol.ID, handler network.StreamHandler) {
	n.IpfsNode.PeerHost.SetStreamHandler(pid, handler)
}
