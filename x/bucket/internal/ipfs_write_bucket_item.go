package internal

import (
	"bytes"
	"context"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/third_party/types/common"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func WriteBucketItems(sh *shell.Shell, b *bt.BucketConfig, address string, items ...*common.BucketItem) (map[string]string, error) {
	endpoints := make(map[string]string, 0)
	for _, buckItem := range items {
		item := bt.NewItemWrapperFromCommon(buckItem)
		itemPath := b.GetPath(address, item.Name())
		err := sh.FilesWrite(context.Background(), itemPath, bytes.NewReader(item.Content()), shell.FilesWrite.Create(true))
		if err != nil {
			return endpoints, fmt.Errorf("%e = Failed to write %s with content [%v]", err, itemPath, item.Content())
		}
		uri := b.GetURI(address, item.Name())
		endpoints[item.Name()] = uri
	}
	return endpoints, nil
}
