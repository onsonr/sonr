package node

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	sentry "github.com/getsentry/sentry-go"
	md "github.com/sonr-io/core/pkg/models"
)

// ^ Host Config ^ //
type HostOptions struct {
	BootstrapAddrs []multiaddr.Multiaddr
	ConnRequest    *md.ConnectionRequest
}

// @ Returns new Host Config
func NewHostOpts(req *md.ConnectionRequest) (*HostOptions, error) {
	// Create Bootstrapper List
	var bootstrappers []multiaddr.Multiaddr
	for _, s := range []string{
		// Libp2p Default
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		"/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
		"/ip4/104.131.131.82/udp/4001/quic/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	} {
		ma, err := multiaddr.NewMultiaddr(s)
		if err != nil {
			panic(err)
		}
		bootstrappers = append(bootstrappers, ma)
	}

	// Set Host Options
	return &HostOptions{
		BootstrapAddrs: bootstrappers,
		ConnRequest:    req,
	}, nil
}

// ^ Return Bootstrap List Address Info ^ //
func (ho *HostOptions) GetBootstrapAddrInfo() []peer.AddrInfo {
	ds := make([]peer.AddrInfo, 0, len(ho.BootstrapAddrs))
	for i := range ho.BootstrapAddrs {
		info, err := peer.AddrInfoFromP2pAddr(ho.BootstrapAddrs[i])
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, fmt.Sprintf("failed to convert bootstrapper address to peer addr info addr: %s",
				ho.BootstrapAddrs[i].String())))
			continue
		}
		ds = append(ds, *info)
	}
	return ds
}

// ^ User Node Info ^ //
// @ ID Returns Peer ID
func (n *Node) ID() *md.Peer_ID {
	return n.fs.GetPeerID(n.device, n.profile, n.host.ID().String())
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
