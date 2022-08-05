package motor

import (
	"context"
	"sync"

	st "github.com/sonr-io/sonr/x/schema/types"
)

type motorResources struct {
	mu sync.Mutex

	schemaQueryClient st.QueryClient

	whatIsStore map[string]*st.WhatIs
	schemaStore map[string]*st.SchemaDefinition
}

func newMotorResources(schemaQueryClient st.QueryClient) *motorResources {
	return &motorResources{
		schemaQueryClient: schemaQueryClient,

		whatIsStore: make(map[string]*st.WhatIs),
		schemaStore: make(map[string]*st.SchemaDefinition),
	}
}

/// StoreWhatIs fetches
func (r *motorResources) StoreWhatIs(whatIs *st.WhatIs) (*st.SchemaDefinition, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.whatIsStore[whatIs.Did] = whatIs

	if schema, ok := r.schemaStore[whatIs.Schema.Cid]; ok {
		return schema, nil
	}

	// send QuerySchemaRequest to get fields
	resp, err := r.schemaQueryClient.Schema(context.Background(), &st.QuerySchemaRequest{
		Creator: whatIs.Creator,
		Did:     whatIs.Did,
	})
	if err != nil {
		return nil, err
	}

	r.schemaStore[whatIs.Schema.Cid] = resp.Definition
	return resp.Definition, nil
}

func (r *motorResources) GetSchema(did string) (*st.WhatIs, *st.SchemaDefinition, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.whatIsStore[did]; !ok {
		return nil, nil, false
	}
	if def, ok := r.schemaStore[r.whatIsStore[did].Schema.Cid]; ok {
		return r.whatIsStore[did], def, true
	}

	return nil, nil, false
}
