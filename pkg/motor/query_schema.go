package motor

import (
	"context"
	"fmt"

	mt "github.com/sonr-io/sonr/thirdparty/types/motor"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (mtr *motorNodeImpl) QueryWhatIs(ctx context.Context, request mt.QueryWhatIsRequest) (mt.QueryWhatIsResponse, error) {
	resp, err := mtr.schemaQueryClient.WhatIs(ctx, &st.QueryWhatIsRequest{
		Creator: request.Creator,
		Did:     request.Did,
	})
	if err != nil {
		return mt.QueryWhatIsResponse{}, err
	}

	// store reference to schema
	_, err = mtr.Resources.StoreWhatIs(resp.WhatIs)
	if err != nil {
		return mt.QueryWhatIsResponse{}, fmt.Errorf("store WhatIs: %s", err)
	}

	return mt.QueryWhatIsResponse{
		WhatIs: resp.WhatIs,
	}, nil
}
