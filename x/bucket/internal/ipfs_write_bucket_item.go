package internal

import (
	"bytes"
	"context"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func WriteBucketItem(sh *shell.Shell, b *bt.BucketConfig, address string, item bt.BucketItem) (string, error) {
	itemPath := b.GetPath(address, item.Name())
	err := sh.FilesWrite(context.Background(), itemPath, bytes.NewReader(item.Content()), shell.FilesWrite.Create(true))
	if err != nil {
		return "", fmt.Errorf("%e = Failed to write %s with content [%v]", err, itemPath, item.Content())
	}
	return b.GetURI(address, item.Name()), nil
}
