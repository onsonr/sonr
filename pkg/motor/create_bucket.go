package motor

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/sonr-io/sonr/internal/bucket"
	"github.com/sonr-io/sonr/pkg/client"
	mt "github.com/sonr-io/sonr/pkg/motor/types"
	"github.com/sonr-io/sonr/pkg/tx"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (mtr *motorNodeImpl) CreateBucket(ctx context.Context, request mt.CreateBucketRequest) (bucket.Bucket, error) {

	if request.Creator == "" {
		return nil, errors.New("Invalid Address")
	}

	if request.Label == "" {
		return nil, errors.New("Label nust be defined")
	}

	createWhereIsRequest := bt.NewMsgCreateWhereIs(request.Creator)
	createWhereIsRequest.Label = request.Label
	createWhereIsRequest.Role = request.Role
	createWhereIsRequest.Visibility = request.Visibility
	createWhereIsRequest.Content = request.Content

	txRaw, err := tx.SignTxWithWallet(mtr.Wallet, "/sonrio.sonr.bucket.MsgCreateWhereIs", createWhereIsRequest)
	if err != nil {
		return nil, fmt.Errorf("sign tx with wallet: %s", err)
	}

	resp, err := mtr.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return nil, fmt.Errorf("broadcast tx: %s", err)
	}

	cbresp := &bt.MsgCreateWhereIsResponse{}
	if err := client.DecodeTxResponseData(resp.TxResponse.Data, cbresp); err != nil {
		return nil, fmt.Errorf("decode MsgCreateWhereIsResponse: %s", err)
	}

	if cbresp.Status != http.StatusAccepted {
		return nil, errors.New("Non success status from Create bucket Request")
	}

	mtr.Resources.whereIsStore[cbresp.WhereIs.Did] = cbresp.WhereIs
	addr, err := mtr.Wallet.Address()
	if err != nil {
		return nil, err
	}

	b := bucket.New(addr,
		mtr.Resources.whereIsStore[cbresp.WhereIs.Did],
		mtr.Resources.shell,
		mtr.Resources.bucketQueryClient)

	mtr.Resources.bucketStore[cbresp.WhereIs.Did] = b

	return b, nil
}
