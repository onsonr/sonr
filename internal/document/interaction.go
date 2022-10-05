package document

import (
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (ao *documentImpl) CreateDocument(
	label string,
	schemaDid string,
	obj map[string]interface{}) (*mt.UploadDocumentResponse, error) {
	err := ao.schema.VerifyDocument(obj)

	if err != nil {
		return nil, err
	}

	err = ao.schema.BuildNodesFromDefinition(label, schemaDid, obj)
	if err != nil {
		return nil, err
	}

	enc, err := ao.schema.EncodeDagJson()
	if err != nil {
		return nil, err
	}

	cid, err := ao.shell.DagPut(enc, "dag-json", "dag-cbor")
	if err != nil {
		return nil, err
	}

	doc, err := NewDocumentFromDag(map[string]interface{}{
		st.IPLD_LABEL:      label,
		st.IPLD_SCHEMA_DID: schemaDid,
		st.IPLD_DOCUMENT:   obj,
	}, ao.schema)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return NewDocumentFromDag(dag, ao.schema)
}
