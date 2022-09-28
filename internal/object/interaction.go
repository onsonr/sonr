package object

import (
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (ao *objectImpl) CreateObject(
	label string,
	obj map[string]interface{}) (*mt.UploadDocumentResponse, error) {
	err := ao.schema.VerifyObject(obj)

	if err != nil {
		return nil, err
	}

	err = ao.schema.BuildNodesFromDefinition(obj)
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

	doc := st.NewDocumentFromMap(cid, obj)

	return &mt.UploadDocumentResponse{
		Status:   200,
		Cid:      cid,
		Document: doc,
	}, nil
}

func (ao *objectImpl) GetObject(cid string) (map[string]interface{}, error) {
	var dag map[string]interface{}
	err := ao.shell.DagGet(cid, &dag)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return dag, err
}
