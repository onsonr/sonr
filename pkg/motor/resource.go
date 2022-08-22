package motor

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kataras/golog"
	"github.com/sonr-io/sonr/pkg/client"
	st "github.com/sonr-io/sonr/x/schema/types"
)

type motorResources struct {
	config            *client.Client
	schemaQueryClient st.QueryClient
	logger            *golog.Logger
	whatIsStore       map[string]*st.WhatIs
	schemaStore       map[string]*st.SchemaDefinition
}

func newMotorResources(config *client.Client, logger *golog.Logger, schemaQueryClient st.QueryClient) *motorResources {
	return &motorResources{
		config:            config,
		schemaQueryClient: schemaQueryClient,
		logger:            logger,
		whatIsStore:       make(map[string]*st.WhatIs),
		schemaStore:       make(map[string]*st.SchemaDefinition),
	}
}

// StoreWhatIs fetches the schema definition from IPFS and caches it
func (r *motorResources) StoreWhatIs(whatIs *st.WhatIs) (*st.SchemaDefinition, error) {
	r.logger.Debugf("Storing WhatIs with did: %s", whatIs.Did)
	r.whatIsStore[whatIs.Did] = whatIs

	if whatIs.Schema == nil {
		r.logger.Errorf("Error while storing what is, schema cannot be nil: %s", whatIs.Did)
		return nil, fmt.Errorf("WhatIs '%s' has no schema", whatIs.Did)
	}
	if schema, ok := r.schemaStore[whatIs.Schema.Cid]; ok {
		return schema, nil
	}

	r.logger.Debug("Querying for schema with cid: %s", whatIs.Schema.Cid)
	resp, err := http.Get(fmt.Sprintf("%s/ipfs/%s", r.config.GetIPFSAddress(), whatIs.Schema.Cid))
	if err != nil {
		return nil, fmt.Errorf("error getting cid '%s': %s", whatIs.Schema.Cid, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.logger.Errorf("Error while querying schema: %s", err)
		return nil, fmt.Errorf("error reading body: %s", err)
	}

	definition := &st.SchemaDefinition{}
	if err = definition.Unmarshal(body); err != nil {
		r.logger.Errorf("Error while unmarshaling body: %s", err)
		return nil, fmt.Errorf("error unmarshalling body: %s", err)
	}

	r.schemaStore[whatIs.Schema.Cid] = definition
	return definition, nil
}

func (r *motorResources) GetSchema(did string) (*st.WhatIs, *st.SchemaDefinition, bool) {
	var whatIs *st.WhatIs
	if w, ok := r.whatIsStore[did]; !ok {
		r.logger.Info("Could not find WhatIs with did: %s within cache", whatIs.Did)
		return nil, nil, false
	} else if w.Schema == nil {
		r.logger.Info("Schema not found within WhatIs: %s", whatIs.Did)
		return nil, nil, false
	} else {
		whatIs = w
	}

	if def, ok := r.schemaStore[whatIs.Schema.Cid]; ok {
		r.logger.Info("Resolved schema for cid: %s", whatIs.Schema.Cid)
		return r.whatIsStore[did], def, true
	}

	return nil, nil, false
}
