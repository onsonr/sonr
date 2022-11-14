package motor

import (
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/pkg/client"
	st "github.com/sonr-io/sonr/x/schema/types"
)

type motorResources struct {
	config      *client.Client
	shell       *shell.Shell
	whatIsStore map[string]*st.WhatIs
	schemaStore map[string]*st.Schema
}

func newMotorResources(
	config *client.Client,
	shell *shell.Shell) *motorResources {
	return &motorResources{
		config:      config,
		shell:       shell,
		whatIsStore: make(map[string]*st.WhatIs),
		schemaStore: make(map[string]*st.Schema),
	}
}

// StoreWhatIs fetches the schema definition from IPFS and caches it
func (r *motorResources) StoreWhatIs(whatIs *st.WhatIs) (*st.Schema, error) {

	r.whatIsStore[whatIs.Did] = whatIs

	if whatIs.Schema == nil {
		return nil, fmt.Errorf("WhatIs '%s' has no schema", whatIs.Did)
	}
	if schema, ok := r.schemaStore[whatIs.Schema.Did]; ok {
		return schema, nil
	}

	var definition *st.Schema = whatIs.Schema
	definition.Did = whatIs.Schema.Did

	r.schemaStore[whatIs.Schema.Did] = definition
	return definition, nil
}

func (r *motorResources) GetSchema(did string) (*st.WhatIs, *st.Schema, bool) {
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
