package node

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/ipfs/kubo/core"
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/sonr-hq/sonr/pkg/common"
	cv1 "github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/wallet"
)

// `Node` is a struct that contains a `CoreAPI` and a `IpfsNode` and a `WalletShare` and a
// `NodeCallback` and a `Context` and a `[]string` and a `Peer_Type` and a `string`.
// @property  - `icore.CoreAPI` is the interface that the node will use to communicate with the IPFS
// daemon.
// @property node - The IPFS node
// @property {string} repoPath - The path to the IPFS repository.
// @property walletShare - This is the wallet share object that is used to share the wallet with other
// nodes.
// @property callback - This is a callback function that will be called when the node is ready.
// @property ctx - The context of the node.
// @property {[]string} bootstrappers - The list of bootstrap nodes to connect to.
// @property peerType - The type of peer, which can be either a bootstrap node or a normal node.
// @property {string} rendezvous - The rendezvous string is a unique identifier for the swarm. It is
// used to find other peers in the swarm.
type Node struct {
	icore.CoreAPI
	node       *core.IpfsNode
	repoPath   string
	rendezvous string

	callback    common.NodeCallback
	peerType    cv1.Peer_Type
	walletShare wallet.WalletShare

	ctx                context.Context
	bootstrappers      []string
	topicEventHandlers map[string]TopicMessageHandler

	network    *Network
	mpcPeerIds []peer.ID
}

// New creates a new node with the given options
func New(ctx context.Context, options ...NodeOption) (*Node, error) {
	// Apply the options
	n := defaultNode(ctx)
	err := n.Apply(options...)
	if err != nil {
		return nil, err
	}

	// Connect to the bootstrap nodes
	err = n.Connect(n.bootstrappers...)
	if err != nil {
		return nil, err
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
func (n *Node) Connect(peers ...string) error {
	var wg sync.WaitGroup
	peerInfos := make(map[peer.ID]*peer.AddrInfo, len(peers))
	for _, addrStr := range peers {
		addr, err := ma.NewMultiaddr(addrStr)
		if err != nil {
			return err
		}
		pii, err := peer.AddrInfoFromP2pAddr(addr)
		if err != nil {
			return err
		}
		pi, ok := peerInfos[pii.ID]
		if !ok {
			pi = &peer.AddrInfo{ID: pii.ID}
			peerInfos[pi.ID] = pi
		}
		pi.Addrs = append(pi.Addrs, pii.Addrs...)
	}

	wg.Add(len(peerInfos))
	for _, peerInfo := range peerInfos {
		go func(peerInfo *peer.AddrInfo) {
			defer wg.Done()
			err := n.CoreAPI.Swarm().Connect(n.ctx, *peerInfo)
			if err != nil {
				log.Printf("failed to connect to %s: %s", peerInfo.ID, err)
			}
		}(peerInfo)
	}
	wg.Wait()
	return nil
}

// ID returns the node's ID
func (n *Node) ID() peer.ID {
	return n.node.Identity
}

// ListPeers lists the peers the node is connected to on a given topic
func (n *Node) ListPeers(topic string) ([]peer.ID, error) {
	return n.PubSub().Peers(n.ctx, options.PubSub.Topic(topic))
}

// ListTopics lists the topics the node is subscribed to
func (n *Node) ListTopics() ([]string, error) {
	return n.PubSub().Ls(n.ctx)
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
		Type:      n.peerType,
	}
}

// Publish publishes a message to a topic
func (n *Node) Publish(topic string, message []byte) error {
	ctx, cancel := context.WithCancel(n.ctx)
	defer cancel()

	errChan := make(chan error)
	go func() {
		err := n.PubSub().Publish(ctx, topic, message)
		errChan <- err
	}()
	select {
	case err := <-errChan:
		return err
	case <-n.ctx.Done():
		return n.ctx.Err()
	}
}

// Subscribing to a topic and then calling the `handleSubscription` function.
func (n *Node) Subscribe(ctx context.Context, topic string, handler ...TopicMessageHandler) error {
	sub, err := n.PubSub().Subscribe(ctx, topic, options.PubSub.Discover(true))
	if err != nil {
		return err
	}
	if len(handler) > 0 {
		n.topicEventHandlers[topic] = handler[0]
	}
	go n.handleSubscription(ctx, topic, sub)
	return nil
}

//
// Private methods
//

// handleTopics handles the topics the node is subscribed to
func (n *Node) handleSubscription(ctx context.Context, topic string, sub icore.PubSubSubscription) {
	for {
		msg, err := sub.Next(n.ctx)
		if err != nil {
			log.Printf("failed to get next message: %s", err)
			return
		}
		if handler, ok := n.topicEventHandlers[topic]; ok {
			handler(topic, msg)
		}
		select {
		case <-n.ctx.Done():
			return
		default:
		}
	}
}
