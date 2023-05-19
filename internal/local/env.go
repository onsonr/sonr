package local

import (
	"os"
)

const (
	// Standard ports for the sonr grpc and rpc api endpoints.
	SonrGrpcPort = "0.0.0.0:9090"
	SonrRpcPort  = "0.0.0.0:26657"

	// CurrentChainID is the current chain ID.
	CurrentChainID = "sonrdevnet-1"
)

// Default configuration
var (
	// defaultBootstrapMultiaddrs is the default list of bootstrap nodes
	defaultBootstrapMultiaddrs = []string{
		// IPFS Bootstrapper nodes.
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		// "/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
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
	}

	// defaultRendezvousString is the default rendezvous string for the motor
	defaultRendezvousString = "sonr"
)

func currGrpcEndpoint() string {
	if env := os.Getenv("ENVIRONMENT"); env != "prod" {
		return SonrGrpcPort
	}
	return SonrGrpcPort
}

func currRpcEndpoint() string {
	if env := os.Getenv("ENVIRONMENT"); env != "prod" {
		return SonrRpcPort
	}
	return SonrRpcPort
}

func getServerPort() string {
	if port := os.Getenv("PUBLC_HIGHWAY_PORT"); port != "" {
		return port
	}
	return "8080"
}

func getTLSCert() string {
	if cert := os.Getenv("TLS_CERT_FILE"); cert != "" {
		return cert
	}
	return ""
}

func getTLSKey() string {
	if key := os.Getenv("TLS_KEY_FILE"); key != "" {
		return key
	}
	return ""
}

func getHomeDir() string {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		homeDir = os.Getenv("USERPROFILE") // windows
	}
	return homeDir
}

func currPublicHostIP() string {
	if ip := os.Getenv("PUBLC_HOST_IP"); ip != "" {
		return ip
	}
	return "localhost"
}

func ValidatorAddress() (string, bool) {
	if address := os.Getenv("SONR_VALIDATOR_ADDRESS"); address != "" {
		return address, true
	}
	return "", false
}
