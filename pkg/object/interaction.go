package object

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (ao *AppObjectInternalImpl) UploadObject(schemaReference *st.SchemaReference, object map[string]interface{}) (*ObjectUploadResult, error) {
	schemaCid := schemaReference.Cid
	tmpPath := fmt.Sprintf("%s%s%s", os.TempDir(), schemaCid, time.Now().Unix())
	ao.shell.Get(schemaCid, tmpPath)

	err := ao.schemaInternal.VerifyObject(object, schemaDefinition)

	if err != nil {
		return nil, err
	}

	n, err := ao.schemaInternal.BuildNodesFromDefinition(schemaDefinition, object)
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
