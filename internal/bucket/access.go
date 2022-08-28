package bucket

import (
	mt "github.com/sonr-io/sonr/third_party/types/motor"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (b *bucketImpl) GetContentById(id string) (*mt.BucketContent, error) {
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

func (b *bucketImpl) GetContent() []*mt.BucketContent {
	var content []*mt.BucketContent = make([]*mt.BucketContent, 0)
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

func (b *bucketImpl) ResolveContentBySchema(did string) ([]*mt.BucketContent, error) {
	var matchedContent []*mt.BucketContent = make([]*mt.BucketContent, 0)
	for _, c := range b.whereIs.Content {
		if c.Type == bt.ResourceIdentifier_CID {
			if c.SchemaDefinition.Did == did {
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
