package bucket

import (
	"fmt"

	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
)

func (b *bucketImpl) CreateBucketServiceEndpoint() did.Service {
	return did.Service{
		ID:   ssi.MustParseURI(fmt.Sprintf("%s/sonr-io/sonr/bucket/where_is/%s/%s", b.rpcClient.GetAPIAddress(), b.address, b.whereIs.Did)),
		Type: "Bucket",
		ServiceEndpoint: map[string]string{
			"did":     b.whereIs.Did,
			"creator": b.whereIs.Creator,
		},
	}
}
