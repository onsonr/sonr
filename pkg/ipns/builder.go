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

func (iub *IPNSURIBuilder) WithCid(cid string) {
	iub.cid = cid
}

func (iub *IPNSURIBuilder) BuildString() string {
	return fmt.Sprintf("%s%s%s%s", iub.delim, iub.prefix, iub.delim, iub.cid)
}

func (iub *IPNSURIBuilder) BuildService() did.Service {
	url := ssi.MustParseURI(fmt.Sprintf("%s%s%s%s", iub.delim, iub.prefix, iub.delim, iub.cid))
	return did.Service{
		ID:   url,
		Type: "ipns",
	}
}
