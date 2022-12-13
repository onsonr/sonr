package node

import (
	"crypto/rand"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/net/connmgr"
	"github.com/multiformats/go-multiaddr"
	"github.com/sonr-hq/sonr/core/common"
)

// NodeOption is a function that configures a Node
type NodeOption func(*NodeConfig) error

// NodeConfig is the configuration for the node that automatically configures itself based on if its a Motor
type NodeConfig struct {
	// privateKey for Identity
	privateKey crypto.PrivKey

	BootstrapPeers []peer.AddrInfo

	// EnableRelay for the node to enable relay
	EnableAutoRelay bool

	// EnableMDNS for the node to enable mdns
	EnableMDNS bool

	// IPFS API URL for Shell Access
	IPFSAPIURL string

	// IPFS Gateway URL for Shell Access
	IPFSGatewayURL string

	// ConnManager for the node
	ConnManager *connmgr.BasicConnMgr

	// Sonr Rendevouz Point
	Rendezvous string

	// Default Stream Handlers for the node
	DefaultStreamHandlers map[protocol.ID]network.StreamHandler

	// MotorCallback is the callback for the motor
	MotorCallback common.MotorCallback
}

// defaultNodeConfig returns the default configuration for the node
func defaultNodeConfig() *NodeConfig {
	// Create Connection Manager
	connmgr, _ := connmgr.NewConnManager(
		10, // Lowwater
		20, // HighWater,
		connmgr.WithGracePeriod(time.Second*4),
	)

	// Define the default bootstrappers
	bootstrapAddrStrs := []string{
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		"/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
		"/ip4/104.131.131.82/udp/4001/quic/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	}

	// Create Bootstrapper List
	var bootstrappers []multiaddr.Multiaddr
	for _, s := range bootstrapAddrStrs {
		ma, err := multiaddr.NewMultiaddr(s)
		if err != nil {
			continue
		}
		bootstrappers = append(bootstrappers, ma)
	}

	// Create Address Info List
	ds := make([]peer.AddrInfo, 0, len(bootstrappers))
	for i := range bootstrappers {
		info, err := peer.AddrInfoFromP2pAddr(bootstrappers[i])
		if err != nil {
			continue
		}
		ds = append(ds, *info)
	}

	return &NodeConfig{
		privateKey:            nil,
		BootstrapPeers:        ds,
		EnableAutoRelay:       true,
		EnableMDNS:            false,
		IPFSAPIURL:            "https://api.ipfs.sonr.ws",
		IPFSGatewayURL:        "https://ipfs.sonr.ws",
		ConnManager:           connmgr,
		Rendezvous:            "sonr",
		DefaultStreamHandlers: map[protocol.ID]network.StreamHandler{
			// "/sonr/1.0.0/message": handleMessageStream,
			// "/sonr/1.0.0/identity": handleIdentityStream,
			// "/sonr/1.0.0/did":      handleDIDStream,
		},
	}
}

// WithPrivateKey sets the PrivateKey for the node
func WithPrivateKey(key crypto.PrivKey) NodeOption {
	return func(c *NodeConfig) error {
		c.privateKey = key
		return nil
	}
}

// WithPathToPrivateKey sets the PathToPrivateKey for the node
func WithPathToPrivateKey(path string) NodeOption {
	return func(c *NodeConfig) error {
		privKey, err := common.LoadPrivKeyFromJsonPath(path)
		if err != nil {
			return err
		}

		c.privateKey = privKey
		return nil
	}
}

// WithEnableRelay sets the EnableRelay for the node
func WithEnableRelay(enable bool) NodeOption {
	return func(c *NodeConfig) error {
		c.EnableAutoRelay = enable
		return nil
	}
}

// WithEnableMDNS sets the EnableMDNS for the node
func WithEnableMDNS(enable bool) NodeOption {
	return func(c *NodeConfig) error {
		c.EnableMDNS = enable
		return nil
	}
}

// WithIPFSAPIURL sets the IPFSAPIURL for the node
func WithIPFSAPIURL(url string) NodeOption {
	return func(c *NodeConfig) error {
		c.IPFSAPIURL = url
		return nil
	}
}

// WithIPFSGatewayURL sets the IPFSGatewayURL for the node
func WithIPFSGatewayURL(url string) NodeOption {
	return func(c *NodeConfig) error {
		c.IPFSGatewayURL = url
		return nil
	}
}

// WithConnMgrLowWater sets the ConnMgrLowWater for the node
func WithConnMgrOptions(low int, high int, ttl time.Duration) NodeOption {
	return func(c *NodeConfig) error {
		// Create Connection Manager
		connmgr, err := connmgr.NewConnManager(
			100, // Lowwater
			400, // HighWater,
			connmgr.WithGracePeriod(time.Minute),
		)
		if err != nil {
			return err
		}
		c.ConnManager = connmgr
		return nil
	}
}

// WithRendezvous sets the Rendezvous for the node
func WithRendezvous(rendezvous string) NodeOption {
	return func(c *NodeConfig) error {
		c.Rendezvous = rendezvous
		return nil
	}
}

// WithDefaultStreamHandlers sets the DefaultStreamHandlers for the node
func WithDefaultStreamHandlers(handlers map[protocol.ID]network.StreamHandler) NodeOption {
	return func(c *NodeConfig) error {
		c.DefaultStreamHandlers = handlers
		return nil
	}
}

// WithMotorCallback sets the MotorCallback for the node
func WithMotorCallback(callback common.MotorCallback) NodeOption {
	return func(c *NodeConfig) error {
		c.MotorCallback = callback
		return nil
	}
}

// GetPrivateKey returns the PrivateKey for the node
func (c *NodeConfig) GetPrivateKey() crypto.PrivKey {
	if c.privateKey != nil {
		return c.privateKey
	}
	privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err == nil {
		return privKey
	}
	return nil
}

// ToLibp2pOptions converts the NodeConfig to libp2p options
func (c *NodeConfig) ToLibp2pOptions(options ...libp2p.Option) []libp2p.Option {
	opts := []libp2p.Option{
		libp2p.Identity(c.GetPrivateKey()),
		libp2p.ConnectionManager(c.ConnManager),
		libp2p.DefaultListenAddrs,
	}
	if c.EnableAutoRelay {
		opts = append(opts, libp2p.EnableAutoRelay())
	}
	if c.EnableMDNS {
		opts = append(opts, libp2p.EnableNATService())
	}
	return append(opts, options...)
}
