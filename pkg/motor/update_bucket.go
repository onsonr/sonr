package motor

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/sonr-io/sonr/internal/bucket"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/tx"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (mtr *motorNodeImpl) UpdateBucket(label string, did string, role bt.BucketRole, visibility bt.BucketVisibility, content []*bt.BucketItem) (bucket.Bucket, error) {
	if mtr.Address == "" {
		return nil, errors.New("invalid Address")
	}

	if label == "" {
		return nil, errors.New("label nust be defined")
	}

	createWhereIsRequest := bt.NewMsgUpdateWhereIs(mtr.Address, did)
	createWhereIsRequest.Label = label
	createWhereIsRequest.Role = role
	createWhereIsRequest.Visibility = visibility
	createWhereIsRequest.Content = content

	txRaw, err := tx.SignTxWithWallet(mtr.Wallet, "/sonrio.sonr.bucket.UpdateWhereIs", createWhereIsRequest)
	if err != nil {
		return nil, fmt.Errorf("sign tx with wallet: %s", err)
	}

	resp, err := mtr.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return nil, fmt.Errorf("broadcast tx: %s", err)
	}

	cbresp := &bt.MsgUpdateWhereIsResponse{}
	if err := client.DecodeTxResponseData(resp.TxResponse.Data, cbresp); err != nil {
		return nil, fmt.Errorf("decode MsgUpdateWhereIsResponse: %s", err)
	}

	if cbresp.Status != http.StatusAccepted {
		return nil, errors.New("Non success status from Update bucket Request")
	}

	mtr.Resources.whereIsStore[cbresp.WhereIs.Did] = cbresp.WhereIs
	addr, err := mtr.Wallet.Address()
	if err != nil {
		return nil, err
	}

	if cbresp.Status != http.StatusAccepted {
		return nil, errors.New("Non success status from Update bucket Request")
	}
	mtr.Resources.bucketStore[did] = nil

	b := bucket.New(addr,
		mtr.Resources.whereIsStore[cbresp.WhereIs.Did],
		mtr.Resources.shell,
		mtr.Resources.bucketQueryClient)

	mtr.Resources.bucketStore[did] = b

	return b, nil
}

func (mtr *motorNodeImpl) UpdateBucketItems(context context.Context, did string, items []*bt.BucketItem) (bucket.Bucket, error) {
	if _, ok := mtr.Resources.bucketStore[did]; !ok {
		return nil, errors.New("cannot resolve content for bucket, not found")
	}

	wi := mtr.Resources.whereIsStore[did]

	b, err := mtr.UpdateBucket(wi.Label, did, wi.Role, wi.Visibility, items)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (mtr *motorNodeImpl) UpdateBucketLabel(context context.Context, did string, label string) (bucket.Bucket, error) {
	if _, ok := mtr.Resources.bucketStore[did]; !ok {
		return nil, errors.New("cannot resolve content for bucket, not found")
	}

	wi := mtr.Resources.whereIsStore[did]

	b, err := mtr.UpdateBucket(label, did, wi.Role, wi.Visibility, wi.Content)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (mtr *motorNodeImpl) UpdateBucketVisibility(context context.Context, did string, visibility bt.BucketVisibility) (bucket.Bucket, error) {
	if _, ok := mtr.Resources.bucketStore[did]; !ok {
		return nil, errors.New("cannot resolve content for bucket, not found")
	}

	wi := mtr.Resources.whereIsStore[did]

	b, err := mtr.UpdateBucket(wi.Label, did, wi.Role, visibility, wi.Content)

	if err != nil {
		return nil, err
	}

	return b, nil
}
