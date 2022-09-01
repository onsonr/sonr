package bucket

import (
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (b *bucketImpl) GetContentById(id string) (*bt.BucketContent, error) {
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

func (b *bucketImpl) GetContent() []*bt.BucketContent {
	var content []*bt.BucketContent = make([]*bt.BucketContent, 0)
	for _, v := range b.contentCache {
		if v.ContentType == bt.ResourceIdentifier_CID {
			content = append(content, v)
		}
	}

	return content
}

func (b *bucketImpl) GetBuckets() []Bucket {
	var content []Bucket = make([]Bucket, 0)
	for _, v := range b.bucketCache {
		content = append(content, v)
	}

	return content
}

func (b *bucketImpl) ResolveContentBySchema(did string) ([]*bt.BucketContent, error) {
	var matchedContent []*bt.BucketContent = make([]*bt.BucketContent, 0)
	for _, c := range b.whereIs.Content {
		if c.Type == bt.ResourceIdentifier_CID {
			if c.SchemaDid == did {
				matchedContent = append(matchedContent, b.contentCache[c.Uri])
			}
		}
	}

	return matchedContent, nil
}

func (b *bucketImpl) GetDID() string {
	return b.whereIs.Did
}

func (b *bucketImpl) GetVisibility() bt.BucketVisibility {
	return b.whereIs.Visibility
}

func (b *bucketImpl) GetRole() bt.BucketRole {
	return b.whereIs.Role
}

func (b *bucketImpl) GetCreator() string {
	return b.whereIs.Creator
}
