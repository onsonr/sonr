package motor

import (
	"context"
	"fmt"

	mt "github.com/sonr-io/sonr/pkg/motor/types"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (mtr *motorNodeImpl) QueryWhatIs(ctx context.Context, request mt.QueryWhatIsRequest) (mt.QueryWhatIsResponse, error) {
	mtr.Logger.Infof("querying for schema with did: %s and creator: %s", request.Did, request.Creator)
	resp, err := mtr.schemaQueryClient.WhatIs(ctx, &st.QueryWhatIsRequest{
		Creator: request.Creator,
		Did:     request.Did,
	})
	if err != nil {
		mtr.Logger.Errorf("error while querying WhatIs: %s", err)
		return mt.QueryWhatIsResponse{}, err
	}

	// store reference to schema
	_, err = mtr.Resources.StoreWhatIs(resp.WhatIs)
	if err != nil {
		mtr.Logger.Errorf("Error while querying WhatIs: %s", err)
		return mt.QueryWhatIsResponse{}, fmt.Errorf("store WhatIs: %s", err)
	}

	return mt.QueryWhatIsResponse{
		WhatIs: resp.WhatIs,
	}, nil
}
