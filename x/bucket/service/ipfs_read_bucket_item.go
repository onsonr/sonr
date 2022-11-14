package service

import (
	"context"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func ReadBucketItem(sh *shell.Shell, whereIs *bt.BucketConfig, address string, item string) (bt.ItemWrapper, error) {
	reader, err := sh.FilesRead(context.Background(), whereIs.GetPath(address, item))
	if err != nil {
		return nil, fmt.Errorf("Failed to read item at %s - %e", whereIs.GetPath(address, item), err)
	}
	defer reader.Close()
	return bt.NewItemWrapperFromReader(item, reader)
}
