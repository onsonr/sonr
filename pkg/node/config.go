package node

import (
	"github.com/sonr-hq/sonr/core/common"
)

// NodeOption is a function that configures a Node
type NodeOption func(*NodeConfig) error

// NodeConfig is the configuration for the node that automatically configures itself based on if its a Motor
type NodeConfig struct {
	// BootstrapMultiaddrs is the list of multiaddresses to bootstrap to
	BootstrapMultiaddrs []string

	// MotorCallback is the callback for the motor
	MotorCallback common.MotorCallback

	// RendezvousString is the rendezvous string for the motor
	RendezvousString string
}

// defaultNodeConfig returns the default configuration for the node
func defaultNodeConfig() *NodeConfig {
	return &NodeConfig{
		BootstrapMultiaddrs: []string{
			// IPFS Bootstrapper nodes.
			"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
			// "/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
			// "/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
			// "/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",

			// IPFS Cluster Pinning nodes
			// "/ip4/138.201.67.219/tcp/4001/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
			// "/ip4/138.201.67.219/udp/4001/quic/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
			// "/ip4/138.201.67.220/tcp/4001/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
			// "/ip4/138.201.67.220/udp/4001/quic/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
			// "/ip4/138.201.68.74/tcp/4001/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
			// "/ip4/138.201.68.74/udp/4001/quic/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
			// "/ip4/94.130.135.167/tcp/4001/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",
			// "/ip4/94.130.135.167/udp/4001/quic/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",

			// You can add more nodes here, for example, another IPFS node you might have running locally, mine was:
			// "/ip4/127.0.0.1/tcp/4010/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
			// "/ip4/127.0.0.1/udp/4010/quic/p2p/QmZp2fhDLxjYue2RiUvLwT9MWdnbDxam32qYFnGmxZDh5L",
		},
		RendezvousString: "sonr",
	}
}

// AddBootstrappers adds additional nodes to start initial connections with
func AddBootstrappers(bootstrappers []string) NodeOption {
	return func(c *NodeConfig) error {
		c.BootstrapMultiaddrs = append(c.BootstrapMultiaddrs, bootstrappers...)
		return nil
	}
}

// WithMotorCallback sets the callback for the motor
func WithMotorCallback(callback common.MotorCallback) NodeOption {
	return func(c *NodeConfig) error {
		c.MotorCallback = callback
		return nil
	}
}
