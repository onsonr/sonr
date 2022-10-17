package motor

import (
	"encoding/base64"
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
	whatIsResp, err := mtr.QueryWhatIsByDid(did)
	if err != nil {
		return nil, fmt.Errorf("could not find WhatIs with did '%s': %s", did, err)
	}

	schemaImpl := schemas.NewWithClient(mtr.GetClient(), whatIsResp.WhatIs)
	objCli := id.New(schemaImpl, shell.NewShell(mtr.Cosmos.GetIPFSApiAddress()))
	return document.NewBuilder(schemaImpl, objCli), nil
}

func (mtr *motorNodeImpl) NewDocumentBuilderFromWhatIs(whatIs *st.WhatIs) (*document.DocumentBuilder, error) {
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

	dag[st.IPLD_DOCUMENT], err = normalizeDocument(schema, d)
	if err != nil {
		return nil, fmt.Errorf("error normalizing document: %s", err)
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

	var err error
	var builder *document.DocumentBuilder
	if req.WhatIsReference == nil {
		builder, err = mtr.NewDocumentBuilder(req.SchemaDid)
		if err != nil {
			return nil, fmt.Errorf("error creating document builder: %s", err)
		}
	} else {
		builder, err = mtr.NewDocumentBuilderFromWhatIs(req.WhatIsReference)
		if err != nil {
			return nil, fmt.Errorf("error creating document builder: %s", err)
		}
	}

	doc, err = normalizeDocument(builder.GetSchema(), doc)
	if err != nil {
		return nil, fmt.Errorf("error normalizing document: %s", err)
	}

	builder.SetLabel(req.GetLabel())
	for k, v := range doc {
		if err = builder.Set(k, v); err != nil {
			return nil, fmt.Errorf("error setting document field '%s': %s", k, err)
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

func normalizeDocument(schema id.Schema, doc map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for _, f := range schema.GetFields() {
		k, v := f.Name, doc[f.Name]
		switch f.GetKind() {
		// json.Unmarshal decodes all numbers as float64 by default
		case st.Kind_INT:
			if fl, ok := v.(float64); ok {
				result[k] = int(fl)
			} else {
				result[k] = v
			}
		// json.Unmarshal encodes byte arrays as base64 strings
		case st.Kind_BYTES:
			if by, ok := v.(string); ok {
				res := make([]byte, base64.StdEncoding.EncodedLen(len(by)))
				if _, err := base64.StdEncoding.Decode(res, []byte(by)); err != nil {
					return nil, err
				}
				result[k] = res
			} else {
				result[k] = v
			}
		default:
			result[k] = v
		}
	}

	return result, nil
}
