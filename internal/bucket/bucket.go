package bucket

import (
	"errors"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/pkg/client"
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
	whereIs      *bt.WhereIs
	contentCache map[string]*bt.BucketContent
	bucketCache  map[string]Bucket
	shell        *shell.Shell
	rpcClient    *client.Client
}

func New(
	address string,
	whereIs *bt.WhereIs,
	shell *shell.Shell,
	rpcClient *client.Client) *bucketImpl {

	return &bucketImpl{
		address:      address,
		whereIs:      whereIs,
		shell:        shell,
		contentCache: make(map[string]*bt.BucketContent),
		rpcClient:    rpcClient,
	}
}
