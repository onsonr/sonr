package types

import (
	fmt "fmt"
	io "io"
	"strings"

	//"github.com/sonr-io/sonr/pkg/did"
    rt "github.com/sonr-io/sonr/x/registry/types"
	"github.com/sonr-io/sonr/third_party/types/common"
)

func (b *BucketConfig) GetPath(address string, segments ...string) string {
	path := fmt.Sprintf("/%s/%s", address, b.Uuid)
	for _, segment := range segments {
		path = fmt.Sprintf("%s/%s", path, segment)
	}
	return path
}

func (b *BucketConfig) GetDidService(address string, cid string) *rt.Service {
	segments := strings.Split(address, "snr")
	service := fmt.Sprintf("did:snr:%s#%s", segments[1], b.Uuid)

	return &rt.Service{
		Id:   service,
		Type: "LinkedResource",
		ServiceEndpoint: b.GetURI(address, cid),
	}
}

func (b *BucketConfig) GetURI(address string, cid string) string {
	params := NewParams()
	path := fmt.Sprintf("%s/ipfs/%s", params.IpfsGateway, cid)
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
