package motor

import (
	"fmt"

	ct "github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	_ "golang.org/x/mobile/bind"
)

func GenerateBucket(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.GenerateBucketRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.GenerateBucket(request)
	if err != nil {
		return nil, err
	}

	return resp.Marshal()
}

func AddBucketItems(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.AddBucketItemsRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.AddBucketItems(request)
	if err != nil {
		return nil, err
	}

	return resp.Marshal()
}

func GetBucketItems(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.GetBucketItemsRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.GetBucketItems(request)
	if err != nil {
		return nil, err
	}

	return resp.Marshal()
}
func QueryBuckets(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.FindBucketConfigRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.QueryBuckets(request)
	if err != nil {
		return nil, err
	}

	return resp.Marshal()
}

func BurnBucket(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.BurnBucketRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.BurnBucket(request)
	if err != nil {
		return nil, err
	}

	return resp.Marshal()
}
