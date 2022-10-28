package resolver

import (
	"fmt"
	"time"

	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
)

type ResolverImpl struct {
	client *client.Client
	did    did.Document
	wallet *mpc.Wallet
}

func New(c *client.Client, w *mpc.Wallet) did.Resolver {
	return &ResolverImpl{
		client: c,
		wallet: w,
	}
}

func (r *ResolverImpl) Resolve(inputDID string) (*did.Document, *did.DocumentMetadata, error) {
	if r.client == nil {
		return nil, nil, fmt.Errorf("client is not established")
	}

	res, err := r.client.QueryWhoIs(inputDID)

	if err != nil {
		return nil, nil, err
	}

	doc, err := res.DidDocument.ToPkgDoc()

	if err != nil {
		return nil, nil, err
	}
	t := time.Unix(res.Timestamp, 0)
	md := make(map[string]interface{})

	md["updated"] = res.Timestamp
	md["deactivated"] = !res.IsActive

	metaData := did.DocumentMetadata{
		Created:    nil,
		Updated:    &t,
		Properties: md,
	}
	r.did = doc

	return &doc, &metaData, nil
}
