package node

import (
	"net"

	ipfs_corehttp "github.com/ipfs/go-ipfs/core/corehttp"
)

// Serve api on the given multiaddr
func (im *IpfsMobile) Serve(l net.Listener) error {
	gatewayOpt := ipfs_corehttp.GatewayOption(false, ipfs_corehttp.WebUIPaths...)
	opts := []ipfs_corehttp.ServeOption{
		ipfs_corehttp.WebUIOption,
		gatewayOpt,
		ipfs_corehttp.CommandsOption(im.commandCtx),
	}

	return ipfs_corehttp.Serve(im.IpfsNode, l, opts...)
}
