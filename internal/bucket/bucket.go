package bucket

import (
	"context"
	"errors"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

var (
	errContentNotFound = func(id string) error {
		if id != "" {
			return fmt.Errorf("could not find content with id: %s", id)
		}

		return errors.New("could not find content")
	}
)

type bucketImpl struct {
	address      string
	whereIs      *bt.Bucket
	contentCache map[string]*bt.BucketContent
	shell        *shell.Shell
	rpcClient    *client.Client
}

func New(address string, shell *shell.Shell, rpcClient *client.Client) *bucketImpl {
	return &bucketImpl{
		address:      address,
		shell:        shell,
		contentCache: make(map[string]*bt.BucketContent),
		rpcClient:    rpcClient,
	}
}

func GenerateBucket(sh *shell.Shell, whereIs *bt.Bucket, address string) (*did.Service, error) {
	allocPath := whereIs.GetPath(address)
	allocDid := whereIs.GetDid(address)
	serviceEndpoint := whereIs.GetServiceEndpoint(address)

	err := sh.FilesMkdir(context.Background(), allocPath, shell.FilesWrite.Create(true))
	if err != nil {
		return nil, err
	}

	res, err := sh.FilesStat(context.Background(), allocPath)
	if err != nil {
		return nil, err
	}
	cid := res.Hash
	err = sh.Publish(cid, allocDid)
	if err != nil {
		return nil, err
	}

	return &did.Service{
		ID:              ssi.MustParseURI(allocDid),
		Type:            "LinkedResource",
		ServiceEndpoint: serviceEndpoint,
	}, nil
}
