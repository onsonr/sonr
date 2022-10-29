package motor

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/tx"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

func (mtr *motorNodeImpl) CreateBucket(request mt.CreateBucketRequest) (*mt.CreateBucketResponse, error) {

	if request.Creator == "" {
		return nil,  errors.New("invalid Address")
	}

	if request.Label == "" {
		return nil,  errors.New("label nust be defined")
	}

	createWhereIsRequest := bt.NewMsgDefineBucket(request.Creator, request.Label)

	txRaw, err := tx.SignTxWithWallet(mtr.Wallet, "/sonrio.sonr.bucket.MsgDefineBucket", createWhereIsRequest)
	if err != nil {
		return nil,  fmt.Errorf("sign tx with wallet: %s", err)
	}

	resp, err := mtr.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return nil,  fmt.Errorf("broadcast tx: %s", err)
	}

	cbresp := &bt.MsgDefineBucketResponse{}
	if err := client.DecodeTxResponseData(resp.TxResponse.Data, cbresp); err != nil {
		return nil,  fmt.Errorf("decode MsgDefineBucketResponse: %s", err)
	}

	if cbresp.Status != http.StatusAccepted {
		return nil,  fmt.Errorf("non success status from Create bucket Reques: %d", cbresp.Status)
	}
	return nil, nil
}
