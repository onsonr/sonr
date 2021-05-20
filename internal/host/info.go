package host

import (
	"fmt"

	ps "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
	md "github.com/sonr-io/core/pkg/models"
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
func (hn *HostNode) MultiAddr() (multiaddr.Multiaddr, *md.SonrError) {
	pi := hn.Info()
	addrs, err := ps.AddrInfoToP2pAddrs(&pi)
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_HOST_INFO)
	}
	fmt.Println("libp2p node address:", addrs[0])
	return addrs[0], nil
}

// ^ Adds Username for PeerID ^ //
func (hn *HostNode) AddUsername(u string) error {
	// Marshal User
	value, err := hn.Host.ID().MarshalBinary()
	if err != nil {
		return err
	}

	// Place Value
	hn.KDHT.PutValue(hn.ctx, u, value)
	return nil
}

// ^ Finds PeerID from Username
func (hn *HostNode) FindAddress(u string) (string, error) {
	// Find Value
	value, err := hn.KDHT.GetValue(hn.ctx, u)
	if err != nil {
		return "", err
	}
	return string(value), nil
}
