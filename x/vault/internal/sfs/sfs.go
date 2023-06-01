package sfs

import (
	"context"

	"github.com/sonrhq/core/x/vault/internal/node"
	"github.com/sonrhq/core/x/vault/types"
)

var (
	ksTable   node.IPFSKVStore
	mailTable node.IPFSDocsStore

	ctx = context.Background()
)

// The function "Init" is declared with the parameter "error" and its purpose is not specified as the code is incomplete.
func Init() error {
	err := node.StartLocalIPFS()
	if err != nil {
		return err
	}
	params := types.NewParams()

	kv, err := node.OpenKeyValueStore(ctx, params.KeyshareSeedFragment)
	if err != nil {
		return err
	}

	docs, err := node.OpenDocumentStore(ctx, params.InboxSeedFragment, nil)
	if err != nil {
		return err
	}
	ksTable = kv
	mailTable = docs
	return nil
}
