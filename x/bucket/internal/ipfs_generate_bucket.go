package internal

import (
	"context"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/pkg/did"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func GenerateBucket(sh *shell.Shell, b *bt.Bucket, address string) (*did.Service, error) {
	path := b.GetPath(address)
	service := b.GetDidService(address)
	err := sh.FilesMkdir(context.Background(), path, shell.FilesWrite.Create(true))
	if err != nil {
		return service, fmt.Errorf("Failed to Make Directory %s in IPFS Node, %e", path, err)
	}
	res, err := sh.FilesStat(context.Background(), path)
	if err != nil {
		return service, fmt.Errorf("Failed to Get Directory %s CID on IPFS Node, %e", path, err)
	}
	err = sh.Publish(res.Hash, path)
	if err != nil {
		return service, fmt.Errorf("Failed to Publish Directory %s CID %s to IPNS in IPFS Node, %e", path, res.Hash, err)
	}
	return service, nil
}
