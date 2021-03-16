package node

import (
	"fmt"
	"net"
	"os"

	sentry "github.com/getsentry/sentry-go"
	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/protocol"
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

// ^ Host Config ^ //
type HostOptions struct {
	BootStrappers []string
	ConnRequest   *md.ConnectionRequest
	OLC           string
	Point         string
	Prefix        protocol.ID
}

// @ Returns new Host Config
func NewHostOpts(req *md.ConnectionRequest) (*HostOptions, error) {
	// Get Open Location Code
	olcValue := olc.Encode(float64(req.Latitude), float64(req.Longitude), 8)

	// Set Host Options
	return &HostOptions{
		//BootStrappers: bootstrappers,
		ConnRequest: req,
		OLC:         olcValue,
		Prefix:      protocol.ID("/sonr"),
		Point:       fmt.Sprintf("/sonr/%s", olcValue),
	}, nil
}

// @ Returns new Bootstrapped Host Config
func NewBootstrappedHostOpts(req *md.ConnectionRequest) (*HostOptions, error) {
	// Create Bootstrapper List
	olcValue := olc.Encode(float64(req.Latitude), float64(req.Longitude), 8)
	bootstrappers := []string{
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		"/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
		"/ip4/104.131.131.82/udp/4001/quic/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	}

	// Set Host Options
	return &HostOptions{
		BootStrappers: bootstrappers,
		ConnRequest:   req,
		OLC:           olcValue,
		Prefix:        protocol.ID("/sonr"),
		Point:         fmt.Sprintf("/sonr/%s", olcValue),
	}, nil
}

// @ Returns Node Public IPv4 Address
func IPv4() string {
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	ipv4Ref := "0.0.0.0"
	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv4 := addr.To4(); ipv4 != nil {
			ipv4Ref = ipv4.String()
		}
	}
	return ipv4Ref
}

// @ Returns Node Public IPv6 Address
func IPv6() string {
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	ipv6Ref := "::"

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv6 := addr.To16(); ipv6 != nil {
			ipv6Ref = ipv6.String()
		}
	}
	return ipv6Ref
}
