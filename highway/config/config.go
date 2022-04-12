package config

import (
	"fmt"
	"time"

	"github.com/duo-labs/webauthn/webauthn"
	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/device"
	"github.com/tendermint/starport/starport/pkg/cosmosaccount"

	ma "github.com/multiformats/go-multiaddr"
)

var (
	// Default P2P Properties
	BootstrapAddrStrs = []string{
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		"/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
		"/ip4/104.131.131.82/udp/4001/quic/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	}
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
	LogLevel             string
	Role                 device.Role
	Libp2pBootstrapPeers []peer.AddrInfo

	Libp2pLowWater    int
	Libp2pHighWater   int
	Libp2pGracePeriod time.Duration
	Libp2pRendezvous  string
	Libp2pInterval    time.Duration
	Libp2pTTL         dscl.Option

	// Session
	Libp2pHost         string
	Libp2pNetwork      string
	Libp2pPort         int
	Libp2pMdnsDisabled bool

	// WebAuthn
	WebAuthNRPDisplayName string
	WebAuthNRPID          string
	WebAuthNRPOrigin      string
	WebAuthNRPIcon        string
	WebAuthNDebug         bool

	// Cosmos SDK
	CosmosAccountName     string
	CosmosAddressPrefix   string
	CosmosNodeAddress     string
	CosmosUseFaucet       bool
	CosmosFaucetAddress   string
	CosmosFaucetDenom     string
	CosmosFaucetMinAmount uint64

	CosmosHomePath           string
	CosmosKeyringBackend     cosmosaccount.KeyringBackend
	CosmosKeyringServiceName string
}

func DefaultConfig() *Config {
	// Create Bootstrapper List
	var bootstrappers []ma.Multiaddr
	for _, s := range BootstrapAddrStrs {
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
	return &Config{
		LogLevel:                 string(InfoLevel),
		Role:                     device.Role_HIGHWAY,
		Libp2pHost:               ":",
		Libp2pPort:               26225,
		Libp2pMdnsDisabled:       true,
		Libp2pNetwork:            "tcp",
		Libp2pBootstrapPeers:     ds,
		Libp2pLowWater:           200,
		Libp2pHighWater:          400,
		Libp2pGracePeriod:        time.Second * 20,
		Libp2pRendezvous:         "/sonr/rendevouz/0.9.2",
		Libp2pInterval:           time.Second * 5,
		Libp2pTTL:                dscl.TTL(time.Minute * 2),
		WebAuthNRPDisplayName:    "Sonr",
		WebAuthNRPID:             "localhost",
		WebAuthNRPOrigin:         "localhost:8080",
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
		CosmosKeyringBackend:     cosmosaccount.KeyringTest,
		CosmosKeyringServiceName: "sonr",
	}
}

func (c *Config) Libp2pAddress() string {
	return fmt.Sprintf("%s%d", c.Libp2pHost, c.Libp2pPort)
}

func (c *Config) WebauthnConfig() *webauthn.Config {
	return &webauthn.Config{
		RPDisplayName: c.WebAuthNRPDisplayName,
		RPID:          c.WebAuthNRPID,
		RPOrigin:      c.WebAuthNRPOrigin,
		RPIcon:        c.WebAuthNRPIcon,
		Debug:         c.WebAuthNDebug,
	}

}

// Option configures your client.
type Option func(*Config)
