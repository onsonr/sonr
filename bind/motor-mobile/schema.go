package motor

import (
	"fmt"

	ct "github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	_ "golang.org/x/mobile/bind"
)

func CreateSchema(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.CreateSchemaRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	if res, err := instance.CreateSchema(request); err == nil {
		return res.Marshal()
	} else {
		return nil, err
	}
}

func GetDocument(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.GetDocumentRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	// resp, err := instance.GetDocument(request)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil // resp.Marshal()
}

func UploadDocument(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.UploadDocumentRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.UploadDocument(request)
	if err != nil {
		return nil, err
	}
	return resp.Marshal()
}
