package local

import (
	"os"
	"path/filepath"
)

// LocalContext is a struct that holds the current context of the application.
type LocalContext struct {
	grpcApiEndpoint string
	rpcApiEndpoint  string

	TlsCertPath  string
	TlsKeyPath   string
	HomeDir      string
	NodeHome     string
	IPFSRepoPath string
	OrbitDBPath  string
	Rendevouz    string
	BsMultiaddrs []string
	isProd       bool
}

// Option is a function that configures the local context
type Option func(LocalContext)

// SetProd sets the current context to production
func SetProd() Option {
	return func(c LocalContext) {
		c.isProd = true
	}
}

// Context returns the current context of the Sonr blockchain application.
func Context(opts ...Option) LocalContext {
	c := LocalContext{
		grpcApiEndpoint: currGrpcEndpoint(),
		rpcApiEndpoint:  currRpcEndpoint(),
		TlsCertPath:     getTLSCert(),
		TlsKeyPath:      getTLSKey(),
		HomeDir:         filepath.Join(getHomeDir()),
		NodeHome:        filepath.Join(getHomeDir(), ".sonr"),
		IPFSRepoPath:    filepath.Join(getHomeDir(), ".sonr", "adapters", "ipfs"),
		OrbitDBPath:     filepath.Join(getHomeDir(), ".sonr", "adapters", "orbitdb"),
		Rendevouz:       defaultRendezvousString,
		BsMultiaddrs:    defaultBootstrapMultiaddrs,
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

// ChainID returns the chain id of the current context
func (c LocalContext) ChainID() string {
	val := os.Getenv("SONR_CHAIN_ID")
	if val == "" {
		return "sonr"
	}
	return val
}

// FaucetEndpoint returns the faucet endpoint of the current context
func (c LocalContext) FaucetEndpoint() string {
	if c.IsDev() {
		return "http://localhost:4500"
	}
	return "https://faucet.sonr.ws"
}

// IsDev returns true if the current context is a development context
func (c LocalContext) IsDev() bool {
	if c.isProd {
		return false
	}
	return os.Getenv("ENVIRONMENT") == "dev"
}

// IsProd returns true if the current context is a production context
func (c LocalContext) IsProd() bool {
	if c.isProd {
		return true
	}
	return !c.IsDev()
}

// HasTlsCert returns true if the current context has a TLS certificate
func (c LocalContext) HasTlsCert() bool {
	return c.TlsCertPath != "" && c.TlsKeyPath != "" && c.IsProd()
}

// GrpcEndpoint returns the grpc endpoint of the current context
func (c LocalContext) GrpcEndpoint() string {
	return c.grpcApiEndpoint
}

// RpcEndpoint returns the rpc endpoint of the current context
func (c LocalContext) RpcEndpoint() string {
	return c.rpcApiEndpoint
}
