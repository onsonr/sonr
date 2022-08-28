package bucket

import (
	"fmt"

	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
)

func (b *bucketImpl) CreateBucketServiceEndpoint(baseURI string) did.Service {
	return did.Service{
		ID:   ssi.MustParseURI(fmt.Sprintf("%s/sonr-io/sonr/bucket/where_is/%s/%s", baseURI, b.address, b.whereIs.Did)),
		Type: "Bucket",
		ServiceEndpoint: map[string]string{
			"did":     b.whereIs.Did,
			"creator": b.whereIs.Creator,
		},
	}
}
