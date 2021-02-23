package host

import (
	"context"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	dht "github.com/libp2p/go-libp2p-kad-dht"
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
	UDPv4        ma.Multiaddr
	TCPv4        ma.Multiaddr
	UDPv6        ma.Multiaddr
	TCPv6        ma.Multiaddr
	Interval     time.Duration
	OLC          string
	Point        string
	PrivateKey   crypto.PrivKey
	Bootstrap    BootstrapConfig
}

// ^ Creates new host configuration ^ //
func NewHostConfig(req *md.ConnectionRequest) (HostConfig, error) {
	// Initialize
	var config HostConfig
	config.setInfo(req)

	// Get Private Key
	err := config.setPrivateKey(req.Directories)
	if err != nil {
		return config, err
	}

	// Get Addresses
	err = config.setAddresses()
	if err != nil {
		return config, err
	}

	// Get Bootstrap Nodes
	config.Bootstrap, err = getBootstrap()
	if err != nil {
		return config, err
	}
	return config, nil
}

// ^ Listen Addresses Returns MultiAddr of Listening Addresses ^
func (hc HostConfig) setAddresses() error {
	ipv4 := ipv4()
	ipv6 := ipv6()

	udpv4, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/%s/udp/0/quic", ipv4))
	if err != nil {
		return err
	} else {
		hc.UDPv4 = udpv4
	}

	udpv6, err := ma.NewMultiaddr(fmt.Sprintf("/ip6/%s/udp/0/quic", ipv6))
	if err != nil {
		return err
	} else {
		hc.UDPv6 = udpv6
	}

	tcpv4, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/0", ipv4))
	if err != nil {
		return err
	} else {
		hc.TCPv4 = tcpv4
	}

	tcpv6, err := ma.NewMultiaddr(fmt.Sprintf("/ip6/%s/tcp/0", ipv6))
	if err != nil {
		return err
	} else {
		hc.TCPv6 = tcpv6
	}
	return nil
}

// ^ Set Config Info ^ //
func (hc HostConfig) setInfo(req *md.ConnectionRequest) {
	olc := olc.Encode(req.Latitude, req.Longitude, 8)
	hc.Connectivity = req.Connectivity
	hc.OLC = olc
	hc.Point = "/sonr/" + olc
	hc.Interval = time.Second * 4
}

// ^ Get Keys: Returns Private/Public keys from disk if found ^ //
func (hc HostConfig) setPrivateKey(dirs *md.Directories) error {
	// Set Path
	path := filepath.Join(dirs.Documents, ".sonr-priv-key")

	// @ Path Doesnt Exist Generate Keys
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Generate Keys
		privKey, _, err := crypto.GenerateRSAKeyPair(2048, rand.Reader)
		if err != nil {
			return err
		}

		// Get Key Bytes
		privDat, err := crypto.MarshalPrivateKey(privKey)
		if err != nil {
			return err
		}

		// Write Private/Pub To File
		err = ioutil.WriteFile(path, privDat, 0644)
		if err != nil {
			return err
		}
		hc.PrivateKey = privKey
		return nil
	}
	// @ Keys Exist Load Keys
	// Load Private Key Bytes from File
	privDat, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Unmarshal PrivKey from Bytes
	privKey, err := crypto.UnmarshalPrivateKey(privDat)
	if err != nil {
		return err
	}
	hc.PrivateKey = privKey
	return nil
}
