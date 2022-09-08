package object

import (
	"fmt"

	"github.com/google/uuid"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
)

func (ao *objectImpl) CreateObject(
	label string,
	obj map[string]interface{}) (*mt.UploadObjectResponse, error) {
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

	return &mt.UploadObjectResponse{
		Code: 200,
		Reference: &mt.ObjectReference{
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
