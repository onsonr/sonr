package bucket

import (
	"context"
	"errors"

	"github.com/sonr-io/sonr/x/bucket/types"
)

func (b *bucketImpl) ResolveBuckets(address string) error {
	if b.whereIs == nil {
		return errors.New("top level bucket not provided")
	}
	var buckets []string = make([]string, 0)
	for _, b := range b.whereIs.Content {
		if b.Type == types.ResourceIdentifier_DID {
			buckets = append(buckets, b.Uri)
		}
	}

	for len(buckets) > 0 {
		key := buckets[len(buckets)-1]
		buckets = buckets[:len(buckets)-1]
		resp, err := b.queryClient.WhereIs(context.Background(), &types.QueryGetWhereIsRequest{
			Creator: address,
			Did:     key,
		})

		if err != nil {
			return err
		}
		b.contentCache[key] = &BucketContent{
			Item:        New(b.adress, &resp.WhereIs, b.shell, b.queryClient),
			Id:          key,
			ContentType: types.ResourceIdentifier_DID,
		}
	}

	return nil
}

func (b *bucketImpl) ResolveContent() error {
	if b.whereIs == nil {
		return errors.New("top level bucket not provided")
	}
	var cids []string = make([]string, 0)
	for _, b := range b.whereIs.Content {
		if b.Type == types.ResourceIdentifier_CID {
			cids = append(cids, b.Uri)
		}
	}

	for len(cids) > 0 {
		cid := cids[len(cids)-1]
		cids = cids[:len(cids)-1]

		var dag map[string]interface{}
		err := b.shell.DagGet(cid, &dag)

		if err != nil {
			return err
		}
		b.contentCache[cid] = &BucketContent{
			Item:        dag,
			Id:          cid,
			ContentType: types.ResourceIdentifier_DID,
		}
	}

	return nil
}
