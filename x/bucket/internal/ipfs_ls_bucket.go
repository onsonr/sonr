package internal

import (
	"context"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func LsBucket(sh *shell.Shell, whereIs *bt.Bucket, address string) ([]*shell.MfsLsEntry, error) {
	files, err := sh.FilesLs(context.Background(), whereIs.GetPath(address))
	if err != nil {
		return nil, fmt.Errorf("Failed to read item at %s - %e", whereIs.GetPath(address), err)
	}
	return files, nil
}
