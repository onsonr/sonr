package host

import (
	"context"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	ma "github.com/multiformats/go-multiaddr"
	md "github.com/sonr-io/core/internal/models"
)

type DiscNotifee struct {
	h   host.Host
	ctx context.Context
}

// ^ Contains Host Configuration ^ //
type HostConfig struct {
	UDPv4      ma.Multiaddr
	TCPv4      ma.Multiaddr
	UDPv6      ma.Multiaddr
	TCPv6      ma.Multiaddr
	OLC        string
	Point      string
	PrivateKey crypto.PrivKey
}

// ^ Creates new host configuration ^ //
func NewHostConfig(req *md.ConnectionRequest) (HostConfig, error) {
	// Initialize
	olc := olc.Encode(req.Latitude, req.Longitude, 8)
	config := HostConfig{
		OLC:   olc,
		Point: "/sonr/" + olc,
	}

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
	return config, nil
}

// ^ Listen Addresses Returns MultiAddr of Listening Addresses ^
func (hc *HostConfig) setAddresses() error {
	ipv4 := ipv4()
	ipv6 := ipv6()

	udpv4, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/%s/udp/8081/quic", ipv4))
	if err != nil {
		return err
	} else {
		hc.UDPv4 = udpv4
	}

	udpv6, err := ma.NewMultiaddr(fmt.Sprintf("/ip6/%s/udp/8081/quic", ipv6))
	if err != nil {
		return err
	} else {
		hc.UDPv6 = udpv6
	}

	tcpv4, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/8081", ipv4))
	if err != nil {
		return err
	} else {
		hc.TCPv4 = tcpv4
	}

	tcpv6, err := ma.NewMultiaddr(fmt.Sprintf("/ip6/%s/tcp/8081", ipv6))
	if err != nil {
		return err
	} else {
		hc.TCPv6 = tcpv6
	}
	return nil
}

// ^ Get Keys: Returns Private/Public keys from disk if found ^ //
func (hc *HostConfig) setPrivateKey(dirs *md.Directories) error {
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
