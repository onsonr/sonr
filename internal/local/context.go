package local

import (
	"os"
	"path/filepath"
)


const (
	// Standard ports for the sonr grpc and rpc api endpoints.
	SonrGrpcPort = "0.0.0.0:9090"
	SonrRpcPort  = "0.0.0.0:26657"

	// CurrentChainID is the current chain ID.
	CurrentChainID = "sonrdevnet-1"
)



// LocalContext is a struct that holds the current context of the application.
type LocalContext struct {
	HomeDir      string
	NodeHome     string
	IPFSRepoPath string
	OrbitDBPath  string
	ConfigDirPath string
	ConfigTomlPath string
	Rendevouz    string
	BsMultiaddrs []string
	isProd       bool
}

// Option is a function that configures the local context
type Option func(LocalContext)


// Context returns the current context of the Sonr blockchain application.
func Context(opts ...Option) LocalContext {
	c := LocalContext{
		HomeDir:         filepath.Join(HomeDir()),
		NodeHome:        filepath.Join(HomeDir(), ".sonr"),
		IPFSRepoPath:    filepath.Join(HomeDir(), ".sonr", "adapters", "ipfs"),
		OrbitDBPath:     filepath.Join(HomeDir(), ".sonr", "adapters", "orbitdb"),
		ConfigDirPath:   filepath.Join(HomeDir(), ".sonr", "config"),
		ConfigTomlPath:  filepath.Join(HomeDir(), ".sonr", "config", "config.toml"),
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}


func GrpcEndpoint() string {
	if env := os.Getenv("ENVIRONMENT"); env != "prod" {
		return SonrGrpcPort
	}
	return SonrGrpcPort
}

func RpcEndpoint() string {
	if env := os.Getenv("ENVIRONMENT"); env != "prod" {
		return SonrRpcPort
	}
	return SonrRpcPort
}

func HomeDir() string {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		homeDir = os.Getenv("USERPROFILE") // windows
	}
	return homeDir
}

func ValidatorAddress() (string, bool) {
	if address := os.Getenv("SONR_VALIDATOR_ADDRESS"); address != "" {
		return address, true
	}
	return "", false
}
