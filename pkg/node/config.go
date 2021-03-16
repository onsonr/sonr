package node

import (
	"encoding/json"
	"net"
	"os"

	sentry "github.com/getsentry/sentry-go"
	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ User Node Info ^ //
// @ ID Returns Peer ID
func (n *Node) ID() *md.Peer_ID {
	return n.fs.GetPeerID(n.hostOpts.ConnRequest, n.profile, n.host.ID().String())
}

// @ Info returns ALL Peer Data as Bytes
func (n *Node) Info() []byte {
	// Convert to bytes to view in plugin
	data, err := proto.Marshal(n.ID())
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}
	return data
}

// @ Peer returns Current Peer Info
func (n *Node) Peer() *md.Peer {
	return n.peer
}

// ^ Bootstrap Nodes ^ //
type AddrConfig struct {
	P2P struct {
		RDVP []struct {
			Maddr string `json:"maddr"`
		} `json:"rdvp"`
	} `json:"p2p"`
}

// @ Returns Current Addr List
func addrList() string {
	return `
{
  "p2p": {
    "rdvp": [
      {
        "maddr": "/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN"
      },
      {
        "maddr": "/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa"
      },
      {
        "maddr": "/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb"
      },
      {
        "maddr": "/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt"
      },
      {
        "maddr": "/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ"
      },
      {
        "maddr": "/ip4/104.131.131.82/udp/4001/quic/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ"
      }
    ],
  }
}`
}

// ^ Host Config ^ //
type HostOptions struct {
	BootStrappers []peer.AddrInfo
	ConnRequest   *md.ConnectionRequest
	IPv4          string
	IPv6          string
	OLC           string
	Namespace     string
	Prefix        protocol.ID
}

// @ Returns new Host Config
func newHostOpts(req *md.ConnectionRequest) (*HostOptions, error) {
	// Get Open Location Code
	olcValue := olc.Encode(float64(req.Latitude), float64(req.Longitude), 8)

	// Get Addr List
	input := addrList()
	config := AddrConfig{}
	err := json.Unmarshal([]byte(input), &config)
	if err != nil {
		panic(err)
	}

	// Create Bootstrapper List
	var bootstrappers []peer.AddrInfo

	// Get Known Addr List
	for _, maddrString := range config.P2P.RDVP {
		// Convert String to MultiAddr
		maddr, err := multiaddr.NewMultiaddr(maddrString.Maddr)
		if err != nil {
			sentry.CaptureException(err)
		}

		// Get Addr Info
		pi, err := peer.AddrInfoFromP2pAddr(maddr)
		if err != nil {
			sentry.CaptureException(err)
			return nil, errors.Wrap(err, "parsing bootstrapper node address info from p2p address")
		}
		bootstrappers = append(bootstrappers, *pi)
	}

	// Find IPv4/IPv6 Addresses
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	ipv4Ref := "0.0.0.0"
	ipv6Ref := "::"

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv4 := addr.To4(); ipv4 != nil {
			ipv4Ref = ipv4.String()
		}
	}

	// @1. Find IPv4 Address
	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv6 := addr.To16(); ipv6 != nil {
			ipv6Ref = ipv6.String()
		}
	}

	// Set Host Options
	return &HostOptions{
		BootStrappers: bootstrappers,
		ConnRequest:   req,
		IPv4:          ipv4Ref,
		IPv6:          ipv6Ref,
		OLC:           olcValue,
		Prefix:        protocol.ID("/sonr/kad/0.9.1"),
		Namespace:     "/sonr/kad/0.9.1",
	}, nil
}
