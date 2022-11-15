package service

import (
	"context"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	//"github.com/sonr-io/sonr/pkg/did"
	bt "github.com/sonr-io/sonr/x/bucket/types"
	rt "github.com/sonr-io/sonr/x/registry/types"
)

func GenerateBucket(sh *shell.Shell, b *bt.BucketConfig, address string) (*rt.Service, error) {
	path := b.GetPath(address)
	err := sh.FilesMkdir(context.Background(), path, shell.FilesWrite.Create(true), shell.FilesWrite.Parents(true))
	if err != nil {
		return nil, fmt.Errorf("failed to Make Directory %s in IPFS Node, %e", path, err)
	}
	res, err := sh.FilesStat(context.Background(), path)
	if err != nil {
		return nil, fmt.Errorf("failed to Get Directory %s CID on IPFS Node, %e", path, err)
	}
	service := b.GetDidService(address, res.Hash)
	/*
		err = sh.Publish(res.Hash, path)
		if err != nil {
			return service, fmt.Errorf("failed to Publish Directory %s CID %s to IPNS in IPFS Node, %e", path, res.Hash, err)
		}
	*/
	return service, nil
}
