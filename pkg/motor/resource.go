package motor

import (
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/internal/bucket"
	"github.com/sonr-io/sonr/pkg/client"
	bt "github.com/sonr-io/sonr/x/bucket/types"
	st "github.com/sonr-io/sonr/x/schema/types"
)

type motorResources struct {
	config       *client.Client
	shell        *shell.Shell
	whatIsStore  map[string]*st.WhatIs
	whereIsStore map[string]*bt.WhereIs
	schemaStore  map[string]*st.SchemaDefinition
	bucketStore  map[string]bucket.Bucket
}

func newMotorResources(
	config *client.Client,
	shell *shell.Shell) *motorResources {
	return &motorResources{
		config:       config,
		shell:        shell,
		bucketStore:  make(map[string]bucket.Bucket),
		whatIsStore:  make(map[string]*st.WhatIs),
		whereIsStore: make(map[string]*bt.WhereIs),
		schemaStore:  make(map[string]*st.SchemaDefinition),
	}
}

// StoreWhatIs fetches the schema definition from IPFS and caches it
func (r *motorResources) StoreWhatIs(whatIs *st.WhatIs) (*st.SchemaDefinition, error) {

	r.whatIsStore[whatIs.Did] = whatIs

	if whatIs.Schema == nil {
		return nil, fmt.Errorf("WhatIs '%s' has no schema", whatIs.Did)
	}
	if schema, ok := r.schemaStore[whatIs.Schema.Did]; ok {
		return schema, nil
	}

	var definition *st.SchemaDefinition = whatIs.Schema
	definition.Did = whatIs.Schema.Did

	r.schemaStore[whatIs.Schema.Did] = definition
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

	if def, ok := r.schemaStore[whatIs.Schema.Did]; ok {
		return r.whatIsStore[did], def, true
	}

	return nil, nil, false
}
