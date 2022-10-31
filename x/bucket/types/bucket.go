package types

import (
	fmt "fmt"
	io "io"
	"strings"

	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/third_party/types/common"
)

func (b *BucketConfig) GetPath(address string, segments ...string) string {
	path := fmt.Sprintf("/%s/%s", address, b.Uuid)
	for _, segment := range segments {
		path = fmt.Sprintf("%s/%s", path, segment)
	}
	return path
}

func (b *BucketConfig) GetDidService(address string) *did.Service {
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

func (b *BucketConfig) GetURI(address string, items ...string) string {
	params := NewParams()
	bucketPath := b.GetPath(address)
	path := fmt.Sprintf("%s/ipns/%s", params.IpfsGateway, bucketPath)
	for _, item := range items {
		path = fmt.Sprintf("%s/%s", path, item)
	}
	return path
}

type bucketWrapperImpl struct {
	name    string
	content []byte
}

func (b *bucketWrapperImpl) Name() string {
	return b.name
}

func (b *bucketWrapperImpl) Content() []byte {
	return b.content
}

type ItemWrapper interface {
	Name() string
	Content() []byte
}

func NewItemWrapperFromReader(name string, reader io.Reader) (ItemWrapper, error) {
	buf, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return &bucketWrapperImpl{
		name:    name,
		content: buf,
	}, nil
}

func NewItemWrapperFromBytes(name string, content []byte) (ItemWrapper, error) {
	if content == nil {
		return nil, fmt.Errorf("content cannot be nil")
	}
	return &bucketWrapperImpl{
		name:    name,
		content: content,
	}, nil
}

func NewItemWrapperFromCommon(item *common.BucketItem) ItemWrapper {
	return &bucketWrapperImpl{
		name:    item.Name,
		content: item.Value,
	}
}
