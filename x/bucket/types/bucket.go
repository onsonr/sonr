package types

import (
	fmt "fmt"
	io "io"
	"strings"

	//"github.com/sonr-io/sonr/pkg/did"
	//"github.com/sonr-io/sonr/pkg/did/ssi"
	rt "github.com/sonr-io/sonr/x/registry/types"
)

func (b *Bucket) GetPath(address string, segments ...string) string {
	path := fmt.Sprintf("/%s/%s", address, b.Uuid)
	for _, segment := range segments {
		path = fmt.Sprintf("%s/%s", path, segment)
	}
	return path
}

func (b *Bucket) GetDidService(address string) *rt.Service {
	segments := strings.Split(address, "snr")
	service := fmt.Sprintf("did:snr:%s#%s", segments[1], b.Uuid)

	return &rt.Service{
		Id:   service,
		Type: "LinkedResource",
		ServiceEndpoint: &rt.ServiceEndpoint{
			Key:   "uri",
			Value: []string{b.GetURI(address)},
		},
	}
}

func (b *Bucket) GetURI(address string, items ...string) string {
	params := NewParams()
	bucketPath := b.GetPath(address)
	path := fmt.Sprintf("%s/ipns/%s", params.IpfsGateway, bucketPath)
	for _, item := range items {
		path = fmt.Sprintf("%s/%s", path, item)
	}
	return path
}

type bucketItemImpl struct {
	name    string
	content []byte
}

func (b *bucketItemImpl) Name() string {
	return b.name
}

func (b *bucketItemImpl) Content() []byte {
	return b.content
}

type BucketItem interface {
	Name() string
	Content() []byte
}

func NewBucketItemFromReader(name string, reader io.Reader) (BucketItem, error) {
	buf, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return &bucketItemImpl{
		name:    name,
		content: buf,
	}, nil
}

func NewBucketItemFromBytes(name string, content []byte) (BucketItem, error) {
	if content == nil {
		return nil, fmt.Errorf("content cannot be nil")
	}
	return &bucketItemImpl{
		name:    name,
		content: content,
	}, nil
}
