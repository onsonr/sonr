package bucket

import (
	"encoding/json"
	"errors"

	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (b *bucketImpl) ResolveBuckets() error {
	if b.whereIs == nil {
		return errors.New("top level bucket not provided")
	}

	for _, bi := range b.whereIs.Content {
		if bi.Type == bt.ResourceIdentifier_DID {
			resp, err := b.rpcClient.QueryWhereIs(bi.Uri, b.address)

			if err != nil {
				return err
			}
			b.bucketCache[bi.Uri] = New(b.address, resp, b.shell, b.rpcClient)
		}
	}

	return nil
}

func (b *bucketImpl) ResolveContent() error {
	if b.whereIs == nil {
		return errors.New("top level bucket not provided")
	}

	for _, bi := range b.whereIs.Content {
		if bi.Type == bt.ResourceIdentifier_CID {
			var dag map[string]interface{}
			err := b.shell.DagGet(bi.Uri, &dag)

			if err != nil {
				return err
			}
			dag_bytes, err := json.Marshal(dag)
			if err != nil {
				return err
			}
			b.contentCache[bi.Uri] = &bt.BucketContent{
				Item:        dag_bytes,
				Id:          bi.Uri,
				ContentType: bt.ResourceIdentifier_CID,
			}
		}
	}

	return nil
}
