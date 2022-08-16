package bucket

import (
	"fmt"

	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
)

func (b *bucketImpl) CreateBucketService() did.Service {
	return did.Service{
		ID:   ssi.MustParseURI(fmt.Sprintf("v1.sonr.ws:1337/sonr-io/sonr/bucket/where_is/%s/%s", b.adress, b.whereIs.Did)),
		Type: "Bucket",
		ServiceEndpoint: map[string]string{
			"did":     b.whereIs.Did,
			"creator": b.whereIs.Creator,
		},
	}
}

func (b *bucketImpl) CreateObjectServices() did.Services {
	return make([]did.Service, 0)
}
