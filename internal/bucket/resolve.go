package bucket

func (b *bucketImpl) GetContent(id string) (*BucketContent, error) {
	if b.contentCache[id] == nil {
		return nil, errContentNotFound(id)
	}

	cnt := b.contentCache[id]

	return cnt, nil
}

func (b *bucketImpl) ContentExists(id string) bool {
	return b.contentCache[id] == nil
}
