package object

import (
	"fmt"

	"github.com/google/uuid"
)

func (ao *objectImpl) CreateObject(
	label string,
	object map[string]interface{}) (*ObjectUploadResult, error) {
	err := ao.schema.VerifyObject(object)

	if err != nil {
		return nil, err
	}

	err = ao.schema.BuildNodesFromDefinition(object)
	if err != nil {
		return nil, err
	}

	enc, err := ao.schema.EncodeDagJson()

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
		Reference: &ObjectReference{
			Did:   did,
			Cid:   cid,
			Label: label,
		},
		Message: "Object uploaded",
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
