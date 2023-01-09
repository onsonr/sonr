package ipfs

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/ipfs/kubo/core"
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/sonr-hq/sonr/pkg/common"
	cv1 "github.com/sonr-hq/sonr/pkg/common"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
)

// `IPFS` is a struct that contains a `CoreAPI` and a `IpfsNode` and a `WalletShare` and a
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
type IPFS struct {
	icore.CoreAPI
	node       *core.IpfsNode
	repoPath   string
	rendezvous string

	callback    common.NodeCallback
	peerType    cv1.PeerType
	walletShare common.WalletShare

	ctx                context.Context
	bootstrappers      []string
	topicEventHandlers map[string]TopicMessageHandler

	mpcPeerIds []peer.ID
	partyId    party.ID
}

// New creates a new node with the given options
func New(ctx context.Context, options ...NodeOption) (*IPFS, error) {
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

// Connect connects to a peer with a given multiaddress
func (n *IPFS) Connect(peers ...string) error {
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

// Add adds a file to the network
func (n *IPFS) Add(file []byte) (string, error) {
	filename := uuid.New().String()
	// Generate a temporary directory
	inputBasePath, err := os.MkdirTemp("", filename)
	if err != nil {
		return "", err
	}

	// Write contents to a temporary file
	inputPath := filepath.Join(inputBasePath, filename)
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
	cid, err := n.CoreAPI.Unixfs().Add(n.ctx, fileNode, options.Unixfs.Pin(true))
	if err != nil {
		return "", err
	}
	if err := os.RemoveAll(inputBasePath); err != nil {
		fmt.Printf("Failed to cleanup Temporary IPFS directory: %s", err)
	}
	return cid.String(), nil
}

// AddEncrypted utilizes the NACL Secret box to encrypt data on behalf of a user
func (n *IPFS) AddEncrypted(file []byte, pubKey []byte) (string, error) {
return "", errors.New("Unimplemented method")
}

// Get returns a file from the network given its CID
func (n *IPFS) Get(cidStr string) ([]byte, error) {
	filename := uuid.New().String()
	cid, err := cid.Parse(cidStr)
	if err != nil {
		return nil, err
	}

	// Get the file from the network
	fileNode, err := n.CoreAPI.Unixfs().Get(n.ctx, icorepath.IpfsPath(cid))
	if err != nil {
		return nil, err
	}

	// Create a temporary directory to store the file
	outputBasePath, err := os.MkdirTemp("", filename)
	if err != nil {
		return nil, err
	}

	// Set the output path
	outputPath := filepath.Join(outputBasePath, filename)
	err = files.WriteTo(fileNode, outputPath)
	if err != nil {
		return nil, err
	}

	// Read the file
	file, err := os.ReadFile(outputPath)
	if err != nil {
		return nil, err
	}
	if err := os.RemoveAll(outputBasePath); err != nil {
		fmt.Printf("Failed to cleanup Temporary IPFS directory: %s", err)
	}
	return file, nil
}

// GetDecrypted decrypts a file from a cid hash using the pubKey
func (n *IPFS) GetDecrypted(cidStr string, pubKey []byte) ([]byte, error) {
	return nil, errors.New("Unimplemented method")
}

// PeerID returns the node's PeerID
func (n *IPFS) PeerID() peer.ID {
	return n.node.Identity
}

// ListTopics lists the topics the node is subscribed to
func (n *IPFS) ListTopics() ([]string, error) {
	return n.PubSub().Ls(n.ctx)
}

// MultiAddr returns the node's multiaddress as a string
func (n *IPFS) MultiAddr() string {
	return fmt.Sprintf("/ip4/127.0.0.1/udp/4010/p2p/%s", n.node.Identity.String())
}

// PartyID returns the node's party ID
func (n *IPFS) PartyID() party.ID {
	return n.partyId
}

// GroupPartyIDs returns the node's party IDs
func (n *IPFS) GroupPartyIDs() []party.ID {
	pids := make([]party.ID, len(n.mpcPeerIds))
	for i, pid := range n.mpcPeerIds {
		pids[i] = party.ID(pid)
	}
	return pids
}

// Peer returns the node's peer info
func (n *IPFS) Peer() *cv1.PeerInfo {
	return &cv1.PeerInfo{
		Name:      string(n.PartyID()),
		PeerId:    n.PeerID().String(),
		Multiaddr: n.MultiAddr(),
		Type:      n.peerType,
	}
}

// Publish publishes a message to a topic
func (n *IPFS) Publish(topic string, message []byte) error {
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
func (n *IPFS) Subscribe(ctx context.Context, topic string, handler ...TopicMessageHandler) error {
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
func (n *IPFS) handleSubscription(ctx context.Context, topic string, sub icore.PubSubSubscription) {
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
