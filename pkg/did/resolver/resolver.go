package did

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jws"
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
		did:    nil,
	}
}

func (r *ResolverImpl) Resolve(inputDID string) (*did.Document, *did.DocumentMetadata, error) {
	if r.client == nil {
		return nil, nil, fmt.Errorf("client is not established")
	}

	_, err := did.ParseDID(inputDID)

	if err != nil {
		return nil, nil, fmt.Errorf("invalid DID: \n%s", err)
	}

	res, err := r.client.QueryWhoIs(inputDID)

	if err != nil {
		return nil, nil, fmt.Errorf("unable to find did %s\n%s", inputDID, err)
	}

	doc, err := res.DidDocument.ToPkgDoc()

	if err != nil {
		return nil, nil, err
	}

	t := time.Unix(res.Timestamp, 0)

	// TODO: this should not be manually mapped and instead be stored within the did registry.
	// Will need to update registry module to also hold Document metadata
	// need updated, deactivated, versionId, nextUpdate, nextVersionId
	// see: https://www.w3.org/TR/did-spec-registries/#did-document-metadata
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

func (r *ResolverImpl) Sign(content []byte, id string) ([]byte, error) {
	if r.did == nil {
		return nil, fmt.Errorf("did has not been resolved, call Resolve before calling this method")
	}
	parsedId, err := did.ParseDID(id)

	if err != nil {
		return nil, fmt.Errorf("could not parse into did id: %s\n%s", id, err)
	}

	vm := r.did.FindAssertionMethod(*parsedId)

	keyMap := vm.PublicKeyJwk
	b, err := json.Marshal(keyMap)
	if err != nil {
		return nil, err
	}

	key, err := jwk.ParseKey(b)
	if err != nil {
		return nil, err
	}

	return jws.Sign(content, jws.WithKey(jwa.ES256K, key))
}

func (r *ResolverImpl) Verify(content []byte, id string) ([]byte, error) {
	if r.did == nil {
		return nil, fmt.Errorf("did has not been resolved, call Resolve before calling this method")
	}
	parsedId, err := did.ParseDID(id)

	if err != nil {
		return nil, fmt.Errorf("could not parse into did id: %s\n%s", id, err)
	}

	vm := r.did.FindAssertionMethod(*parsedId)

	keyMap := vm.PublicKeyJwk
	b, err := json.Marshal(keyMap)
	if err != nil {
		return nil, err
	}

	key, err := jwk.ParseKey(b)
	if err != nil {
		return nil, err
	}

	return jws.Verify(content, jws.WithKey(jwa.ES256K, key))
}
