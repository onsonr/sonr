package config

import (
	"os"
	"time"

	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"

	ma "github.com/multiformats/go-multiaddr"
	t "go.buf.build/grpc/go/sonr-io/motor/core/v1"
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

	// WebAuthn
	WebAuthNRPDisplayName string
	WebAuthNRPID          string
	WebAuthNRPOrigin      string
	WebAuthNRPIcon        string
	WebAuthNDebug         bool

	// Cosmos SDK
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
	Location       *t.Location
	HomeDirPath    string
	SupportDirPath string
	TempDirPath    string

	// Matrix Config
	MatrixServerName string
}

// DefaultConfig returns the default configuration for the Highway node
func DefaultConfig(r Role) *Config {
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
		LogLevel:                 string(InfoLevel),
		Role:                     r,
		Libp2pMdnsDisabled:       true,
		HighwayGRPCNetwork:       "tcp",
		Libp2pBootstrapPeers:     ds,
		Libp2pLowWater:           200,
		Libp2pHighWater:          400,
		Libp2pGracePeriod:        time.Second * 20,
		Libp2pRendezvous:         "/sonr/rendevouz/0.9.2",
		Libp2pInterval:           time.Second * 5,
		Libp2pTTL:                dscl.TTL(time.Minute * 2),
		WebAuthNRPDisplayName:    "Sonr",
		WebAuthNRPID:             "localhost",
		WebAuthNRPOrigin:         "http://localhost:8989",
		WebAuthNRPIcon:           "",
		WebAuthNDebug:            true,
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
		HighwayGRPCEndpoint:      "localhost:26225",
		HighwayHTTPEndpoint:      ":8081",
		MatrixServerName:         "sonr-matrix-1",
		DeviceID:                 "",
		HostName:                 "",
		Location:                 nil,
	}

	// Role specific configuration
	if r == Role_MOTOR {
		conf.Location = DefaultLocation()

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

// DefaultLocation returns the Sonr HQ as default location
func DefaultLocation() *t.Location {
	return &t.Location{
		Latitude:  float64(40.673010),
		Longitude: float64(-73.994450),
		Placemark: &t.Location_Placemark{
			Name:                  "Sonr HQ",
			Street:                "94 9th St.",
			IsoCountryCode:        "US",
			Country:               "United States",
			AdministrativeArea:    "New York",
			SubAdministrativeArea: "Brooklyn",
			Locality:              "Brooklyn",
			SubLocality:           "Gowanus",
			PostalCode:            "11215",
		},
	}
}
