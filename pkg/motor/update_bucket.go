package motor

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/sonr-io/sonr/internal/bucket"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/tx"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (mtr *motorNodeImpl) UpdateBucket(req mt.UpdateBucketRequest) (bucket.Bucket, error) {
	if mtr.Address == "" {
		return nil, errors.New("invalid Address")
	}

	updateWhereIsRequest := bt.NewMsgUpdateWhereIs(mtr.Address, req.Did)
	updateWhereIsRequest.Label = req.Label
	updateWhereIsRequest.Role = req.Role
	updateWhereIsRequest.Visibility = req.Visibility
	updateWhereIsRequest.Content = req.Content

	txRaw, err := tx.SignTxWithWallet(mtr.Wallet, "/sonrio.sonr.bucket.MsgUpdateWhereIs", updateWhereIsRequest)
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
		return nil, errors.New("non success status from Update bucket Request")
	}

	mtr.Resources.whereIsStore[cbresp.WhereIs.Did] = cbresp.WhereIs
	addr, err := mtr.Wallet.Address()
	if err != nil {
		return nil, err
	}

	if cbresp.Status != http.StatusAccepted {
		return nil, errors.New("non success status from Update bucket Request")
	}
	mtr.Resources.bucketStore[req.Did] = nil

	b := bucket.New(addr,
		mtr.Resources.whereIsStore[cbresp.WhereIs.Did],
		mtr.Resources.shell,
		mtr.GetClient())

	mtr.Resources.bucketStore[req.Did] = b

	err = b.ResolveBuckets()

	if err != nil {
		return nil, err
	}

	err = b.ResolveContent()

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (mtr *motorNodeImpl) UpdateBucketItems(context context.Context, did string, items []*bt.BucketItem) (bucket.Bucket, error) {
	if _, ok := mtr.Resources.bucketStore[did]; !ok {
		return nil, errors.New("cannot resolve content for bucket, not found")
	}

	wi := mtr.Resources.whereIsStore[did]
	updateReq := mt.UpdateBucketRequest{
		Creator:    mtr.Address,
		Did:        did,
		Label:      wi.Label,
		Role:       wi.Role,
		Visibility: wi.Visibility,
		Content:    items,
	}
	b, err := mtr.UpdateBucket(updateReq)

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
	updateReq := mt.UpdateBucketRequest{
		Creator:    mtr.Address,
		Did:        did,
		Label:      label,
		Role:       wi.Role,
		Visibility: wi.Visibility,
		Content:    wi.Content,
	}

	b, err := mtr.UpdateBucket(updateReq)

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
	updateReq := mt.UpdateBucketRequest{
		Creator:    mtr.Address,
		Did:        did,
		Label:      wi.Label,
		Role:       wi.Role,
		Visibility: visibility,
		Content:    wi.Content,
	}

	b, err := mtr.UpdateBucket(updateReq)

	if err != nil {
		return nil, err
	}

	return b, nil
}
