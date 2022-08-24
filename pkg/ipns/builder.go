package ipns

import (
	"fmt"

	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
)

type IPNSURIBuilder struct {
	prefix string
	delim  string
	cid    string
}

func NewBuilder() *IPNSURIBuilder {
	return &IPNSURIBuilder{
		prefix: "ipfs",
		delim:  "/",
	}
}

func (iub *IPNSURIBuilder) SetCid(cid string) {
	iub.cid = cid
}

func (iub *IPNSURIBuilder) String() string {
	return fmt.Sprintf("%s%s%s%s", iub.delim, iub.prefix, iub.delim, iub.cid)
}

func (iub *IPNSURIBuilder) BuildService() did.Service {
	url := ssi.MustParseURI(iub.String())
	return did.Service{
		ID:   url,
		Type: "ipns",
		ServiceEndpoint: map[string]string{
			"cid": iub.cid,
		},
	}
}
