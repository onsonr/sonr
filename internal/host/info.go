package host

import (
	"fmt"

	ps "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

// ^ Returns HostNode Peer Addr Info ^ //
func (hn *HostNode) Info() ps.AddrInfo {
	peerInfo := ps.AddrInfo{
		ID:    hn.Host.ID(),
		Addrs: hn.Host.Addrs(),
	}
	return peerInfo
}

// ^ Returns Host Node MultiAddr ^ //
func (hn *HostNode) MultiAddr() (multiaddr.Multiaddr, error) {
	pi := hn.Info()
	addrs, err := ps.AddrInfoToP2pAddrs(&pi)
	if err != nil {
		return nil, err
	}
	fmt.Println("libp2p node address:", addrs[0])
	return addrs[0], nil
}
