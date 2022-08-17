package bucket

func (b *bucketImpl) GetContentById(id string) (*BucketContent, error) {
	if b.contentCache[id] == nil {
		return nil, errContentNotFound(id)
	}

	cnt := b.contentCache[id]

	return cnt, nil
}

func (b *bucketImpl) GetContent() []*BucketContent {
	content := make([]*BucketContent, len(b.contentCache))

	for _, v := range b.contentCache {
		content = append(content, v)
	}

	return content
}

func (b *bucketImpl) ContentExists(id string) bool {
	return b.contentCache[id] == nil
}
