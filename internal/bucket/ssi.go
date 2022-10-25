package bucket

import (
	"fmt"

	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
)

func (b *bucketImpl) CreateBucketServiceEndpoint(urlPath string, id string) did.Service {
	return did.Service{
		ID:              ssi.MustParseURI(fmt.Sprintf("did:snr:%s#%s", b.address, b.whereIs.Name)),
		Type:            "LinkedResource",
		ServiceEndpoint: fmt.Sprintf("%s/sonr-io/sonr/bucket/where_is/%s/%s", b.rpcClient.GetAPIAddress(), b.address, b.whereIs.Uuid),
	}
}
