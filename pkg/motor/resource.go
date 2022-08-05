package motor

import (
	"context"
	"fmt"
	"sync"

	st "github.com/sonr-io/sonr/x/schema/types"
)

// SchemaCallback is a callback used when a schema CID is fetched
type SchemaCallback func(*st.SchemaDefinition, error)
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
func (r *motorResources) StoreWhatIs(whatIs *st.WhatIs, cb SchemaCallback) {
	callback := func(s *st.SchemaDefinition, e error) {
		r.mu.Unlock()
		cb(s, e)
	}

	go func() {
		r.mu.Lock()

		r.whatIsStore[whatIs.Did] = whatIs

		if whatIs.Schema == nil {
			callback(nil, fmt.Errorf("WhatIs '%s' has no schema", whatIs.Did))
			return
		}
		if schema, ok := r.schemaStore[whatIs.Schema.Cid]; ok {
			callback(schema, nil)
			return
		}

		// send QuerySchemaRequest to get fields
		resp, err := r.schemaQueryClient.Schema(context.Background(), &st.QuerySchemaRequest{
			Creator: whatIs.Creator,
			Did:     whatIs.Did,
		})
		if err != nil {
			callback(nil, err)
			return
		}

		r.schemaStore[whatIs.Schema.Cid] = resp.Definition
		callback(resp.Definition, nil)
	}()
}

func (r *motorResources) GetSchema(did string) (*st.WhatIs, *st.SchemaDefinition, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

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

func noSchemaCallback(_ *st.SchemaDefinition, _ error) {}
