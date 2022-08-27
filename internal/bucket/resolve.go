package bucket

import (
	"errors"

	"github.com/sonr-io/sonr/x/bucket/types"
)

func (b *bucketImpl) ResolveBuckets() error {
	if b.whereIs == nil {
		return errors.New("top level bucket not provided")
	}

	for _, bi := range b.whereIs.Content {
		if bi.Type == types.ResourceIdentifier_DID {
			resp, err := b.rpcClient.QueryWhereIs(bi.Uri, b.address)

			if err != nil {
				return err
			}
			b.contentCache[bi.Uri] = &BucketContent{
				Item:        New(b.address, resp, b.shell, b.rpcClient),
				Id:          bi.Uri,
				ContentType: types.ResourceIdentifier_DID,
			}
		}
	}

	return nil
}

func (b *bucketImpl) ResolveContent() error {
	if b.whereIs == nil {
		return errors.New("top level bucket not provided")
	}

	for _, bi := range b.whereIs.Content {
		if bi.Type == types.ResourceIdentifier_CID {
			var dag map[string]interface{}
			err := b.shell.DagGet(bi.Uri, &dag)

			if err != nil {
				return err
			}
			b.contentCache[bi.Uri] = &BucketContent{
				Item:        dag,
				Id:          bi.Uri,
				ContentType: types.ResourceIdentifier_CID,
			}
		}
	}

	return nil
}
