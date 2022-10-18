package document

import (
	"fmt"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (ao *documentImpl) CreateDocument(
	label string,
	schemaDid string,
	docMap map[string]interface{}) (*mt.UploadDocumentResponse, error) {
	if err := ao.schema.VerifyDocument(docMap); err != nil {
		return nil, fmt.Errorf("verify document: %s", err)
	}

	if err := ao.schema.BuildNodesFromDefinition(label, schemaDid, docMap); err != nil {
		return nil, fmt.Errorf("build DAG: %s", err)
	}

	enc, err := ao.schema.EncodeDagJson()
	if err != nil {
		return nil, fmt.Errorf("encode DAG: %s", err)
	}

	cid, err := ao.shell.DagPut(enc, "dag-json", "dag-cbor")
	if err != nil {
		return nil, fmt.Errorf("put DAG: %s", err)
	}

	doc, err := NewDocumentFromDag(map[string]interface{}{
		st.IPLD_LABEL:      label,
		st.IPLD_SCHEMA_DID: schemaDid,
		st.IPLD_DOCUMENT:   docMap,
	}, ao.schema)
	if err != nil {
		return nil, fmt.Errorf("create Document: %s", err)
	}

	return &mt.UploadDocumentResponse{
		Status:   200,
		Cid:      cid,
		Document: doc,
	}, nil
}

func (ao *documentImpl) GetDocument(cid string) (*st.SchemaDocument, error) {
	var dag map[string]interface{}
	err := ao.shell.DagGet(cid, &dag)
	if err != nil {
		return nil, fmt.Errorf("get DAG: %s", err)
	}

	return NewDocumentFromDag(dag, ao.schema)
}
