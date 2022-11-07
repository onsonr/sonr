package service

import (
	"context"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func DeleteBucketItem(sh *shell.Shell, whereIs *bt.BucketConfig, address string, item string) error {
	err := sh.FilesRm(context.Background(), whereIs.GetPath(address, item), true)
	if err != nil {
		return fmt.Errorf("Failed to Delete item at %s - %e", whereIs.GetPath(address, item), err)
	}
	return nil
}

func PurgeBucketItems(sh *shell.Shell, whereIs *bt.BucketConfig, address string) error {
	entries, err := LsBucket(sh, whereIs, address)

    fmt.Printf("%+v", entries)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		err := sh.FilesRm(context.Background(), whereIs.GetPath(address, entry.Name), true)
		if err != nil {
			return fmt.Errorf("Failed to Delete item at %s - %e", entry.Name, err)
		}
	}
	return nil
}
