package object

import (
	"bytes"
	"fmt"

	"github.com/google/uuid"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (ao *AppObjectInternalImpl) UploadObject(fields []*st.SchemaKindDefinition, object map[string]interface{}) (*ObjectUploadResult, error) {
	err := ao.schemaInternal.VerifyObject(object, fields)

	if err != nil {
		return nil, err
	}

	n, err := ao.schemaInternal.BuildNodesFromDefinition(fields, object)
	if err != nil {
		return nil, err
	}

	enc, err := ao.schemaInternal.EncodeDagCbor(n)

	if err != nil {
		return nil, err
	}

	cid, err := ao.shell.Add(bytes.NewReader(enc))
	did := fmt.Sprintf("did:snr:%s", uuid.New().String())
	if err != nil {
		return nil, err
	}

	return &ObjectUploadResult{
		Code: 200,
		Definition: &ObjectDefinition{
			Cid: cid,
			Did: did,
		},
		Message: "Object uploaded",
	}, nil
}

func (ao *AppObjectInternalImpl) GetObject(cid string) ([]byte, error) {
	return nil, nil
}
