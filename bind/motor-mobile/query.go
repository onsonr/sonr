package motor

import (
	"fmt"

	ct "github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	_ "golang.org/x/mobile/bind"
)

func QuerySchema(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.QueryWhatIsRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	res, err := instance.QueryWhatIs(request)
	if err != nil {
		return nil, err
	}
	return res.Marshal()
}

func QuerySchemaByCreator(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.QueryWhatIsByCreatorRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	res, err := instance.QueryWhatIsByCreator(request)
	if err != nil {
		return nil, err
	}
	return res.Marshal()
}

func QuerySchemaByDid(did string) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	res, err := instance.QueryWhatIsByDid(did)
	if err != nil {
		return nil, err
	}
	return res.Marshal()
}
