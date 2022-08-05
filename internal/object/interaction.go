package object

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sonr-io/sonr/pkg/motor/x/object"
)

func (ao *objectImpl) CreateObject(
	label string,
	obj map[string]interface{}) (*object.ObjectUploadResult, error) {
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
	did := fmt.Sprintf("did:snr:%s", uuid.New().String())
	if err != nil {
		return nil, err
	}

	return &object.ObjectUploadResult{
		Code: 200,
		Reference: &object.ObjectReference{
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
