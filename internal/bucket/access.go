package bucket

import (
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (b *bucketImpl) GetContentById(id string) (*BucketContent, error) {
	if b.contentCache[id] == nil {
		return nil, errContentNotFound(id)
	}

	cnt := b.contentCache[id]

	return cnt, nil
}

func (b *bucketImpl) GetBucketItems() []*bt.BucketItem {
	if b.whereIs == nil {
		return make([]*bt.BucketItem, 0)
	}

	return b.whereIs.Content
}

func (b *bucketImpl) ContentExists(id string) bool {
	return b.contentCache[id] == nil
}
