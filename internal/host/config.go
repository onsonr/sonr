package host

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"

	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/crypto"
	disco "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	disc "github.com/libp2p/go-libp2p/p2p/discovery"
	ma "github.com/multiformats/go-multiaddr"
	md "github.com/sonr-io/core/internal/models"
)

type DiscNotifee struct {
	h   host.Host
	ctx context.Context
}

// ^ Contains Host Configuration ^ //
type HostConfig struct {
	Connectivity md.ConnectionRequest_Connectivity
	DHT          *dht.IpfsDHT
	ListenAddrs  []ma.Multiaddr
	Interval     time.Duration
	MDNS         disc.Service
	OLC          string
	Point        string
	PrivateKey   crypto.PrivKey
	Routing      *discovery.RoutingDiscovery
	Bootstrap    struct {
		P2P struct {
			RDVP []struct {
				Maddr string `json:"maddr"`
			} `json:"rdvp"`
			RelayHack []string `json:"relayHack" yaml:"relayHack"`
		} `json:"p2p"`
	}
}

// ^ Creates new host configuration ^ //
func NewHostConfig(req *md.ConnectionRequest) (HostConfig, error) {
	// Initialize
	var config HostConfig
	config.Connectivity = req.Connectivity
	config.OLC = olc.Encode(req.Latitude, req.Longitude, 8)
	config.Point = "/sonr/" + config.OLC
	config.Interval = time.Second * 4

	// Get Private Key
	privKey, err := getPrivateKey(req.Directories)
	if err != nil {
		return config, err
	}

	// Get Addresses
	listenAddrs, err := getAddresses()
	if err != nil {
		return config, err
	}

	// Get Bootstrap Nodes
	nodeDat, err := ioutil.ReadFile("nodes.json")
	if err != nil {
		return config, err
	}

	// Get Bootstrap List
	err = json.Unmarshal(nodeDat, &config.Bootstrap)
	if err != nil {
		return config, err
	}

	// Set and Return Config
	config.PrivateKey = privKey
	config.ListenAddrs = listenAddrs
	return config, nil
}

// ^ Find Peers from Routing Discovery ^ //
func (hc *HostConfig) FindPeers(ctx context.Context, limit int) (<-chan peer.AddrInfo, error) {
	if hc.Routing != nil {
		return hc.Routing.FindPeers(ctx, hc.Point, disco.Limit(limit))
	}
	return nil, nil
}

// ^ Listen Addresses Returns MultiAddr of Listening Addresses ^
func getAddresses() ([]ma.Multiaddr, error) {
	ipv4 := ipv4()
	ipv6 := ipv6()
	addrs := make([]ma.Multiaddr, 4)

	udpv4, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/%s/udp/0/quic", ipv4))
	if err != nil {
		return addrs, err
	}

	udpv6, err := ma.NewMultiaddr(fmt.Sprintf("/ip6/%s/udp/0/quic", ipv6))
	if err != nil {
		return addrs, err
	}

	tcpv4, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/0", ipv4))
	if err != nil {
		return addrs, err
	}

	tcpv6, err := ma.NewMultiaddr(fmt.Sprintf("/ip6/%s/tcp/0", ipv6))
	if err != nil {
		return addrs, err
	}

	addrs = append(addrs, udpv4, tcpv4, udpv6, tcpv6)
	return addrs, nil
}

// ^ Get Keys: Returns Private/Public keys from disk if found ^ //
func getPrivateKey(dirs *md.Directories) (crypto.PrivKey, error) {
	// Set Path
	path := filepath.Join(dirs.Documents, ".sonr-priv-key")

	// @ Path Doesnt Exist Generate Keys
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Generate Keys
		privKey, _, err := crypto.GenerateRSAKeyPair(2048, rand.Reader)
		if err != nil {
			return nil, err
		}

		// Get Key Bytes
		privDat, err := crypto.MarshalPrivateKey(privKey)
		if err != nil {
			return nil, err
		}

		// Write Private/Pub To File
		err = ioutil.WriteFile(path, privDat, 0644)
		if err != nil {
			return nil, err
		}
		return privKey, nil
	}
	// @ Keys Exist Load Keys
	// Load Private Key Bytes from File
	privDat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal PrivKey from Bytes
	privKey, err := crypto.UnmarshalPrivateKey(privDat)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}

// ^ IPv4 returns the non loopback local IP of the host as IPv4 ^
func ipv4() string {
	// @1. Find IPv4 Address
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	var ipv4Ref string

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv4 := addr.To4(); ipv4 != nil {
			ipv4Ref = ipv4.String()
		} else {
			ipv4Ref = "0.0.0.0"
		}
	}
	// No IPv4 Found
	return ipv4Ref
}

// ^ IPv4 returns the non loopback local IP of the host as IPv4 ^
func ipv6() string {
	// @1. Find IPv4 Address
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	var ipv6Ref string

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv6 := addr.To16(); ipv6 != nil {
			ipv6Ref = ipv6.String()
		} else {
			ipv6Ref = "::"
		}
	}
	// No IPv4 Found
	return ipv6Ref
}
