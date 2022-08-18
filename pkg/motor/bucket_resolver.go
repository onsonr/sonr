package motor

import (
	"context"
	"errors"

	"github.com/sonr-io/sonr/internal/bucket"
	"github.com/sonr-io/sonr/pkg/motor/types"
)

func (mtr *motorNodeImpl) NewBucketResolver(context context.Context, creator string, did string) (bucket.Bucket, error) {
	if creator == "" || did == "" {
		return nil, errors.New("creator address and did must be defined within the request")
	}
	addr, err := mtr.Wallet.Address()
	if err != nil {
		return nil, err
	}

	if _, ok := mtr.Resources.whereIsStore[did]; !ok {
		wiReq, err := mtr.QueryWhereIs(context, types.QueryWhereIsRequest{
			Creator: creator,
			Did:     did,
		})

		if err != nil {
			return nil, err
		}

		mtr.Resources.whereIsStore[did] = wiReq.WhereIs
	}
	wi := mtr.Resources.whereIsStore[did]
	s := mtr.Resources.shell
	bq := mtr.Resources.bucketQueryClient

	b := bucket.New(addr, wi, s, bq)

	mtr.Resources.bucketStore[did] = b

	return b, nil
}
