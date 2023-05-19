package node

import (
	"context"
	"fmt"

	"berty.tech/go-orbit-db/iface"
	"github.com/sonrhq/core/x/vault/internal/node/config"
	"github.com/sonrhq/core/x/vault/internal/node/ipfs"
)

// IPFSKVStore is an alias for a iface.KeyValueStore.
type IPFSKVStore = iface.KeyValueStore

// IPFSEventLogStore is an alias for a iface.EventLogStore.
type IPFSEventLogStore = iface.EventLogStore

// IPFSDocsStore is an alias for a iface.DocumentStore.
type IPFSDocsStore = iface.DocumentStore

// Callback is an alias for a common.NodeCallback
type Callback = config.NodeCallback

// IPFS is an alias for a common.IPFSNode.
type IPFS = config.IPFSNode

// P2P is an alias for a common.P2PNode.
type P2P = config.PeerNode

var (
	local IPFS
)

// StartLocalIPFS initializes a local IPFS node.
func StartLocalIPFS() error {

	config := config.DefaultConfig()
	err := config.Apply()
	if err != nil {
		return err
	}
	i, err := ipfs.Initialize(config)
	if err != nil {
		return err
	}
	local = i
	return nil
}

// OpenKeyValueStore creates a new IPFSKVStore. This requires a valid Sonr Account Public Key.
func OpenKeyValueStore(ctx context.Context, controllerAddr string) (IPFSKVStore, error) {
	if local == nil {
		err := StartLocalIPFS()
		if err != nil {
			return nil, fmt.Errorf("local IPFS node not initialized: %w", err)
		}
	}

	kv, err := local.LoadKeyValueStore(controllerAddr)
	if err != nil {
		return nil, err
	}
	return kv, nil
}

// OpenEventLogStore creates a new IPFSEventLogStore. This requires a valid Sonr Account Public Key.
func OpenEventLogStore(ctx context.Context, controllerAddr string) (IPFSEventLogStore, error) {
	if local == nil {
		err := StartLocalIPFS()
		if err != nil {
			return nil, fmt.Errorf("local IPFS node not initialized: %w", err)
		}
	}

	el, err := local.LoadEventLogStore(controllerAddr)
	if err != nil {
		return nil, err
	}
	return el, nil
}

// OpenDocumentStore creates a new IPFSDocsStore. This requires a valid Sonr Account Public Key.
func OpenDocumentStore(ctx context.Context, controllerAddr string, opts *iface.CreateDocumentDBOptions) (IPFSDocsStore, error) {
	if local == nil {
		err := StartLocalIPFS()
		if err != nil {
			return nil, fmt.Errorf("local IPFS node not initialized: %w", err)
		}
	}

	ds, err := local.LoadDocsStore(controllerAddr, opts)
	if err != nil {
		return nil, err
	}
	return ds, nil
}
