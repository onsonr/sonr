package bucket

import (
	"context"
	"errors"

	"github.com/sonr-io/sonr/x/bucket/types"
)

func (b *bucketImpl) TraverseTopLevelBucket(address string) error {
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
			Item:        resp.WhereIs,
			Id:          key,
			ContentType: types.ResourceIdentifier_DID,
		}

		for _, b := range resp.WhereIs.Content {
			if b.Type == types.ResourceIdentifier_DID {
				buckets = append(buckets, b.Uri)
			}
		}
	}

	return nil
}
