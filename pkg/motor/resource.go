package motor

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/internal/bucket"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/motor/types"
	bt "github.com/sonr-io/sonr/x/bucket/types"
	st "github.com/sonr-io/sonr/x/schema/types"
)

type motorResources struct {
	config            *client.Client
	schemaQueryClient st.QueryClient
	bucketQueryClient bt.QueryClient
	shell             *shell.Shell
	whatIsStore       map[string]*st.WhatIs
	whereIsStore      map[string]*bt.WhereIs
	schemaStore       map[string]*st.SchemaDefinition
	bucketStore       map[string]bucket.Bucket
}

func newMotorResources(
	config *client.Client,
	bucketQueryClient bt.QueryClient,
	schemaQueryClient st.QueryClient,
	shell *shell.Shell) *motorResources {
	return &motorResources{
		config:            config,
		schemaQueryClient: schemaQueryClient,
		bucketQueryClient: bucketQueryClient,
		shell:             shell,
		bucketStore:       make(map[string]bucket.Bucket),
		whatIsStore:       make(map[string]*st.WhatIs),
		whereIsStore:      make(map[string]*bt.WhereIs),
		schemaStore:       make(map[string]*st.SchemaDefinition),
	}
}

// StoreWhatIs fetches the schema definition from IPFS and caches it
func (r *motorResources) StoreWhatIs(whatIs *st.WhatIs) (*st.SchemaDefinition, error) {

	r.whatIsStore[whatIs.Did] = whatIs

	if whatIs.Schema == nil {
		return nil, fmt.Errorf("WhatIs '%s' has no schema", whatIs.Did)
	}
	if schema, ok := r.schemaStore[whatIs.Schema.Cid]; ok {
		return schema, nil
	}

	resp, err := http.Get(fmt.Sprintf("%s/ipfs/%s", r.config.GetIPFSAddress(), whatIs.Schema.Cid))
	if err != nil {
		return nil, fmt.Errorf("error getting cid '%s': %s", whatIs.Schema.Cid, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %s", err)
	}

	definition := &st.SchemaDefinition{}
	if err = definition.Unmarshal(body); err != nil {
		return nil, fmt.Errorf("error unmarshalling body: %s", err)
	}

	r.schemaStore[whatIs.Schema.Cid] = definition
	return definition, nil
}

func (r *motorResources) StoreWhereIs(whereis *bt.WhereIs) bool {
	_, ok := r.whatIsStore[whereis.Did]
	r.whereIsStore[whereis.Did] = whereis

	return ok
}

func (r *motorResources) StoreBucket(did string, b bucket.Bucket) bool {
	_, ok := r.whatIsStore[did]
	r.bucketStore[did] = b

	return ok
}

func (r *motorResources) GetSchema(did string) (*st.WhatIs, *st.SchemaDefinition, bool) {
	var whatIs *st.WhatIs
	if w, ok := r.whatIsStore[did]; !ok {
		return nil, nil, false
	} else if w.Schema == nil {
		return nil, nil, false
	} else {
		whatIs = w
	}

	if def, ok := r.schemaStore[whatIs.Schema.Cid]; ok {
		return r.whatIsStore[did], def, true
	}

	return nil, nil, false
}

func (r *motorResources) GetWhereIs(ctx context.Context, did string, address string) (*bt.WhereIs, error) {
	if did == "" {
		return nil, errors.New("did invalid for Get WhereIs by Creator request")
	}

	resp, err := r.bucketQueryClient.WhereIs(ctx, &bt.QueryGetWhereIsRequest{
		Creator: address,
		Did:     did,
	})

	if err != nil {
		return nil, err
	}

	res := types.QueryWhereIsResponse{
		WhereIs: &resp.WhereIs,
	}

	r.whereIsStore[res.WhereIs.Did] = res.WhereIs

	return res.WhereIs, nil
}

func (r *motorResources) GetWhereIsByCreator(ctx context.Context, address string) ([]*bt.WhereIs, error) {
	res, err := r.bucketQueryClient.WhereIsByCreator(ctx, &bt.QueryGetWhereIsByCreatorRequest{
		Creator:    address,
		Pagination: nil,
	})

	if err != nil {
		return nil, err
	}

	var ptrArr []*bt.WhereIs = make([]*bt.WhereIs, len(res.WhereIs))
	for _, wi := range res.WhereIs {
		r.whereIsStore[wi.Did] = &wi
		ptrArr = append(ptrArr, &wi)
	}

	return ptrArr, nil
}
