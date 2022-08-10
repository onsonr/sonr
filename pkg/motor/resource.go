package motor

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sonr-io/sonr/pkg/client"
	st "github.com/sonr-io/sonr/x/schema/types"
)

type motorResources struct {
	config            *client.Client
	schemaQueryClient st.QueryClient

	whatIsStore map[string]*st.WhatIs
	schemaStore map[string]*st.SchemaDefinition
}

func newMotorResources(config *client.Client, schemaQueryClient st.QueryClient) *motorResources {
	return &motorResources{
		config:            config,
		schemaQueryClient: schemaQueryClient,

		whatIsStore: make(map[string]*st.WhatIs),
		schemaStore: make(map[string]*st.SchemaDefinition),
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
