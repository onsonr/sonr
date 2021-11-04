package ipfs

import (
	"context"

	orbitdb "berty.tech/go-orbit-db"
	"berty.tech/go-orbit-db/iface"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	"github.com/libp2p/go-libp2p-core/host"
)

type SonrIPFS struct {
	// IPFS node
	*core.IpfsNode

	host  host.Host
	ctx   context.Context
	orbit iface.OrbitDB
}

func New(ctx context.Context, h host.Host) (*SonrIPFS, error) {
	// Basic ipfsnode setup
	r, err := fsrepo.Open("~/.ipfs")
	if err != nil {
		return nil, err
	}

	cfg := &core.BuildCfg{
		Repo:   r,
		Online: true,
	}
	nd, err := core.NewNode(ctx, cfg)
	if err != nil {
		return nil, err
	}

	api, err := coreapi.NewCoreAPI(nd)
	if err != nil {
		return nil, err
	}

	// OrbitDB setup
	orbitDb, err := orbitdb.NewOrbitDB(ctx, api, nil)
	if err != nil {
		return nil, err
	}

	// Create SonrIPFS
	si := &SonrIPFS{
		IpfsNode: nd,
		host:     h,
		ctx:      ctx,
		orbit:    orbitDb,
	}
	return si, nil
}
