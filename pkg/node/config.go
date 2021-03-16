package node

import (
	"encoding/json"

	sentry "github.com/getsentry/sentry-go"
	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	md "github.com/sonr-io/core/pkg/models"
)

// ^ Bootstrap Nodes ^ //
type AddrConfig struct {
	P2P struct {
		RDVP []struct {
			Maddr string `json:"maddr"`
		} `json:"rdvp"`
		RelayHack []string `json:"relayHack" yaml:"relayHack"`
	} `json:"p2p"`
}

// ^ Returns Current Addr List ^ //
func addrList() string {
	return `
{
  "p2p": {
    "rdvp": [
      {
        "maddr": "/ip4/51.159.21.214/tcp/4040/p2p/QmdT7AmhhnbuwvCpa5PH1ySK9HJVB82jr3fo1bxMxBPW6p"
      },
      {
        "maddr": "/ip4/51.159.21.214/udp/4040/quic/p2p/QmdT7AmhhnbuwvCpa5PH1ySK9HJVB82jr3fo1bxMxBPW6p"
      },
      {
        "maddr": "/ip4/51.15.25.224/tcp/4040/p2p/12D3KooWHhDBv6DJJ4XDWjzEXq6sVNEs6VuxsV1WyBBEhPENHzcZ"
      },
      {
        "maddr": "/ip4/51.15.25.224/udp/4040/quic/p2p/12D3KooWHhDBv6DJJ4XDWjzEXq6sVNEs6VuxsV1WyBBEhPENHzcZ"
      },
      {
        "maddr": "/ip4/51.75.127.200/tcp/4141/p2p/12D3KooWPwRwwKatdy5yzRVCYPHib3fntYgbFB4nqrJPHWAqXD7z"
      },
      {
        "maddr": "/ip4/51.75.127.200/udp/4141/quic/p2p/12D3KooWPwRwwKatdy5yzRVCYPHib3fntYgbFB4nqrJPHWAqXD7z"
      }
    ],
    "relayHack": [
      "/ip4/51.159.21.214/udp/4040/quic/p2p/QmdT7AmhhnbuwvCpa5PH1ySK9HJVB82jr3fo1bxMxBPW6p",
      "/ip4/51.15.25.224/udp/4040/quic/p2p/12D3KooWHhDBv6DJJ4XDWjzEXq6sVNEs6VuxsV1WyBBEhPENHzcZ",
      "/ip4/51.75.127.200/udp/4141/quic/p2p/12D3KooWPwRwwKatdy5yzRVCYPHib3fntYgbFB4nqrJPHWAqXD7z"
    ]
  }
}`
}

// ^ Host Config ^ //
type HostOptions struct {
	BootStrappers []peer.AddrInfo
	ConnRequest   *md.ConnectionRequest
	OLC           string
	Point         string
}

// ^ Returns new Host Config ^ //
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

	// Set Host Options
	return &HostOptions{
		OLC:           olcValue,
		Point:         "/sonr/" + olcValue,
		BootStrappers: bootstrappers,
		ConnRequest:   req,
	}, nil
}
