package resolver

import (
	"fmt"
	"time"

	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
)

type ResolverImpl struct {
	wallet *mpc.Wallet
	client *client.Client
	did.Resolver
}

func New() did.Resolver {
	return &ResolverImpl{}
}

// this could be abstracted into a provider for connections.
// right now binded to the client pkg
func (res *ResolverImpl) WithClient(client *client.Client) {
	res.client = client
}

func (res *ResolverImpl) WithWallet(wallet *mpc.Wallet) {
	res.wallet = wallet
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

	return &doc, &metaData, nil
}
