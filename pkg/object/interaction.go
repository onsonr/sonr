package object

import (
	"fmt"

	"github.com/google/uuid"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (ao *AppObjectInternalImpl) CreateObject(
	label string,
	fields []*st.SchemaKindDefinition,
	object map[string]interface{}) (*ObjectUploadResult, error) {
	err := ao.schemaInternal.VerifyObject(object, fields)

	if err != nil {
		return nil, err
	}

	n, err := ao.schemaInternal.BuildNodesFromDefinition(fields, object)
	if err != nil {
		return nil, err
	}

	enc, err := ao.schemaInternal.EncodeDagJson(n)

	if err != nil {
		return nil, err
	}

	cid, err := ao.shell.DagPut(enc, "dag-json", "dag-cbor")
	did := fmt.Sprintf("did:snr:%s", uuid.New().String())
	if err != nil {
		return nil, err
	}

	return &ObjectUploadResult{
		Code: 200,
		Definition: &ObjectDefinition{
			Did:   did,
			Cid:   cid,
			Label: label,
		},
		Message: "Object uploaded",
	}, nil
}

func (ao *AppObjectInternalImpl) GetObject(cid string) (map[string]interface{}, error) {
	var dag map[string]interface{}
	err := ao.shell.DagGet(cid, &dag)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return nil, nil
}
