package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"

	ma "github.com/multiformats/go-multiaddr"
)

// LogLevel is the type for the log level
type LogLevel string

const (
	// DebugLevel is the debug log level
	DebugLevel LogLevel = "debug"
	// InfoLevel is the info log level
	InfoLevel LogLevel = "info"
	// WarnLevel is the warn log level
	WarnLevel LogLevel = "warn"
	// ErrorLevel is the error log level
	ErrorLevel LogLevel = "error"
	// FatalLevel is the fatal log level
	FatalLevel LogLevel = "fatal"
)

// Config is the configuration for the entire Highway node
type Config struct {
	// Host
	Role                 Role
	LogLevel             string
	Libp2pMdnsDisabled   bool
	Libp2pBootstrapPeers []peer.AddrInfo
	Libp2pLowWater       int
	Libp2pHighWater      int
	Libp2pGracePeriod    time.Duration
	Libp2pRendezvous     string
	Libp2pInterval       time.Duration
	Libp2pTTL            dscl.Option

	// Highway Config
	HighwayGRPCNetwork  string
	HighwayGRPCEndpoint string
	HighwayHTTPEndpoint string

	// JWT
	JWTSecret        string
	JWTSigningMethod jwt.SigningMethod
	JWTExpiration    int64

	// Cosmos SDK
	AccountAddress           string
	CosmosAccountName        string
	CosmosAddressPrefix      string
	CosmosNodeAddress        string
	CosmosUseFaucet          bool
	CosmosFaucetAddress      string
	CosmosFaucetDenom        string
	CosmosFaucetMinAmount    uint64
	CosmosHomePath           string
	CosmosKeyringBackend     string
	CosmosKeyringServiceName string

	// Device Config
	DeviceID       string
	HostName       string
	HomeDirPath    string
	SupportDirPath string
	TempDirPath    string

	// Matrix Config
	MatrixServerName string
}

// DefaultConfig returns the default configuration for the Highway node
func DefaultConfig(r Role, addr string) *Config {
	var conf *Config
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
	var bootstrappers []ma.Multiaddr
	for _, s := range bootstrapAddrStrs {
		ma, err := ma.NewMultiaddr(s)
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

	// Create the default configuration
	conf = &Config{
		LogLevel:       string(InfoLevel),
		Role:           r,
		AccountAddress: addr,

		Libp2pMdnsDisabled:       true,
		HighwayGRPCNetwork:       "tcp",
		Libp2pBootstrapPeers:     ds,
		Libp2pLowWater:           10,
		Libp2pHighWater:          20,
		Libp2pGracePeriod:        time.Second * 4,
		Libp2pRendezvous:         "/sonr/rendevouz/0.9.2",
		Libp2pInterval:           time.Second * 5,
		Libp2pTTL:                dscl.TTL(time.Minute * 2),
		HighwayGRPCEndpoint:      "localhost:26225",
		HighwayHTTPEndpoint:      ":8081",
		MatrixServerName:         "sonr-matrix-1",
		DeviceID:                 "",
		HostName:                 "",
		CosmosAccountName:        "alice",
		CosmosAddressPrefix:      "snr",
		CosmosNodeAddress:        "http://localhost:26657",
		CosmosUseFaucet:          false,
		CosmosFaucetAddress:      "",
		CosmosFaucetDenom:        "uatom",
		CosmosFaucetMinAmount:    100,
		CosmosHomePath:           "~/.sonr",
		CosmosKeyringBackend:     "test",
		CosmosKeyringServiceName: "sonr",
	}

	// Role specific configuration
	if r == Role_MOTOR {
		// Check for non-mobile device
		if !IsMobile() {
			// Set Home Directory
			if hdir, err := os.UserHomeDir(); err == nil {
				conf.HomeDirPath = hdir
			}

			// Set Support Directory
			if sdir, err := os.UserConfigDir(); err == nil {
				conf.SupportDirPath = sdir
			}

			// Set Temp Directory
			if tdir, err := os.UserCacheDir(); err == nil {
				conf.TempDirPath = tdir
			}
		}
	}
	return conf
}
