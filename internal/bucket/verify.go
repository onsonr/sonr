package bucket

import "github.com/sonr-io/sonr/x/bucket/types"

func (b *bucketImpl) IsBucket(id string) bool {
	c, ok := b.contentCache[id]

	if !ok {
		return false
	}

	if c.ContentType == types.ResourceIdentifier_DID {
		return true
	} else {
		return false
	}
}

func (b *bucketImpl) IsContent(id string) bool {
	c, ok := b.contentCache[id]

	if !ok {
		return false
	}

	if c.ContentType == types.ResourceIdentifier_CID {
		return true
	} else {
		return false
	}
}
