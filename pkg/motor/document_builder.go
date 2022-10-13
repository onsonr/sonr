package motor

import (
	"encoding/json"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	id "github.com/sonr-io/sonr/internal/document"
	"github.com/sonr-io/sonr/internal/schemas"
	"github.com/sonr-io/sonr/pkg/motor/x/document"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (mtr *motorNodeImpl) NewDocumentBuilder(did string) (*document.DocumentBuilder, error) {
	whatIs, _, found := mtr.Resources.GetSchema(did)
	if !found {
		return nil, fmt.Errorf("could not find WhatIs with did '%s'", did)
	}

	schemaImpl := schemas.NewWithClient(mtr.GetClient(), whatIs)
	objCli := id.New(schemaImpl, shell.NewShell(mtr.Cosmos.GetIPFSApiAddress()))
	return document.NewBuilder(schemaImpl, objCli), nil
}

func (mtr *motorNodeImpl) GetDocument(req mt.GetDocumentRequest) (*mt.GetDocumentResponse, error) {
	dag, err := mtr.queryDocument(req.GetCid())
	if err != nil {
		return nil, fmt.Errorf("query document: %s", err)
	}

	schemaDid, ok := dag[st.IPLD_SCHEMA_DID].(string)
	if !ok {
		return nil, fmt.Errorf("could not get schema did from DAG")
	}

	schemaRes, err := mtr.QueryWhatIsByDid(schemaDid)
	if err != nil {
		return nil, fmt.Errorf("fetch WhatIs: %s", err)
	}

	schema := schemas.NewWithClient(mtr.GetClient(), schemaRes.WhatIs)

	// convert dag float values to int if necessary
	// json unmarshalling uses float64 by default
	d, ok := dag[st.IPLD_DOCUMENT].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("missing document in dag")
	}
	for _, f := range schema.GetFields() {
		if f.GetKind() == st.Kind_INT {
			d[f.Name] = int(d[f.Name].(float64))
		}
	}

	doc, err := id.NewDocumentFromDag(dag, schema)
	if err != nil {
		return nil, fmt.Errorf("create document from DAG: %s", err)
	}

	return &mt.GetDocumentResponse{
		Status:   200,
		Document: doc,
		Cid:      req.GetCid(),
	}, nil
}

func (mtr *motorNodeImpl) UploadDocument(req mt.UploadDocumentRequest) (*mt.UploadDocumentResponse, error) {
	var doc map[string]interface{}
	if err := json.Unmarshal(req.GetDocument(), &doc); err != nil {
		return nil, fmt.Errorf("error decoding document JSON")
	}

	builder, err := mtr.NewDocumentBuilder(req.GetSchemaDid())
	if err != nil {
		return nil, err
	}

	// Normalize values
	// json.Unmarshal decodes all numbers as float64 by default
	// json.Unmarshal decodes base64 encoded bytes as strings
	for _, f := range builder.GetSchema().GetFields() {
		if f.GetKind() == st.Kind_INT {
			doc[f.Name] = int(doc[f.Name].(float64))
		} else if f.GetKind() == st.Kind_BYTES {
			doc[f.Name] = []byte(doc[f.Name].([]byte))
		}
	}

	builder.SetLabel(req.GetLabel())
	for k, v := range doc {
		if err = builder.Set(k, v); err != nil {
			return nil, fmt.Errorf("error setting document field: %s", err)
		}
	}

	resp, err := builder.Upload()
	if err != nil {
		return nil, err
	}

	return &mt.UploadDocumentResponse{
		Status:   resp.Status,
		Cid:      resp.Cid,
		Document: resp.Document,
	}, nil
}
