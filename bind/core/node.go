// ready to use gomobile package for ipfs

// This package intend to only be use with gomobile bind directly if you
// want to use it in your own gomobile project, you may want to use host/node package directly

package core

// Main API exposed to the ios/android

import (
	"context"
	"log"
	"net"
	"sync"

	mobile_host "github.com/ipfs-shipyard/gomobile-ipfs/go/pkg/host"
	mobile_node "github.com/ipfs-shipyard/gomobile-ipfs/go/pkg/node"

	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"

	ipfs_bs "github.com/ipfs/go-ipfs/core/bootstrap"
	// ipfs_log "github.com/ipfs/go-log"
)

type Node struct {
	listeners   []manet.Listener
	muListeners sync.Mutex

	ipfsMobile *mobile_node.IpfsMobile
}

func (n *Node) Close() error {
	n.muListeners.Lock()
	for _, l := range n.listeners {
		l.Close()
	}
	n.muListeners.Unlock()

	return n.ipfsMobile.Close()
}

func (n *Node) ServeUnixSocketAPI(sockpath string) (err error) {
	_, err = n.ServeMultiaddr("/unix/" + sockpath)
	return
}

// Serve API on the given port and return the current listening maddr
func (n *Node) ServeTCPAPI(port string) (string, error) {
	return n.ServeMultiaddr("/ip4/127.0.0.1/tcp/" + port)
}

func (n *Node) ServeConfigAPI() error {
	cfg, err := n.ipfsMobile.Repo.Config()
	if err != nil {
		return err
	}

	if len(cfg.Addresses.API) > 0 {
		for _, maddr := range cfg.Addresses.API {
			if _, err := n.ServeMultiaddr(maddr); err != nil {
				log.Printf("cannot serve `%s`: %s", maddr, err.Error())
			}
		}
	}

	return nil
}

func (n *Node) ServeMultiaddr(smaddr string) (string, error) {
	maddr, err := ma.NewMultiaddr(smaddr)
	if err != nil {
		return "", err
	}

	ml, err := manet.Listen(maddr)
	if err != nil {
		return "", err
	}

	n.muListeners.Lock()
	n.listeners = append(n.listeners, ml)
	n.muListeners.Unlock()

	go func(l net.Listener) {
		if err := n.ipfsMobile.Serve(l); err != nil {
			log.Printf("serve error: %s", err.Error())
		}
	}(manet.NetListener(ml))

	return ml.Multiaddr().String(), nil
}

func NewNode(r *Repo) (*Node, error) {
	ctx := context.Background()

	if _, err := loadPlugins(r.mr.Path); err != nil {
		return nil, err
	}

	mnode, err := mobile_node.NewNode(ctx, r.mr, &mobile_host.MobileConfig{})
	if err != nil {
		return nil, err
	}

	if err := mnode.IpfsNode.Bootstrap(ipfs_bs.DefaultBootstrapConfig); err != nil {
		log.Printf("failed to bootstrap node: `%s`", err)
	}

	return &Node{
		ipfsMobile: mnode,
	}, nil
}

func init() {
	//      ipfs_log.SetDebugLogging()
}
