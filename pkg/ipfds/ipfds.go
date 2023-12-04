package ipfds

import (
	"context"
	"fmt"

	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
)

// New creates a new IPFS node and pins a given file by its CID
func New() {
    // "Connect" to local node
    node, err := rpc.NewLocalApi()
    if err != nil {
        fmt.Printf("failed to connect to local node: %s", err)
        return
    }

    // Create a context and a path
    ctx := context.Background()
    cid := "bafkreidtuosuw37f5xmn65b3ksdiikajy7pwjjslzj2lxxz2vc4wdy3zku"
    p, err := path.NewPath(cid)
    if err != nil {
        fmt.Printf("failed to parse %s: %s", cid, err)
        return
    }

    // Pin a given file by its CID
    err = node.Pin().Add(ctx, p)
    if err != nil {
        fmt.Printf("failed to pin %s: %s", cid, err)
        return
    }
    fmt.Printf("pinned %s", cid)
}
