package motor

import (
	"fmt"

	mt "github.com/sonr-io/sonr/thirdparty/types/motor"
)

// DEPRECATED - use Query instead
func (mtr *motorNodeImpl) QueryWhatIs(request mt.QueryWhatIsRequest) (mt.QueryWhatIsResponse, error) {
	resp, err := mtr.GetClient().QueryWhatIs(
		request.Creator,
		request.Did,
	)
	if err != nil {
		return mt.QueryWhatIsResponse{}, err
	}

	// store reference to schema
	_, err = mtr.Resources.StoreWhatIs(resp)
	if err != nil {
		return mt.QueryWhatIsResponse{}, fmt.Errorf("store WhatIs: %s", err)
	}

	return mt.QueryWhatIsResponse{
		WhatIs: resp,
	}, nil
}
