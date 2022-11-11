package motor

import (
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/tx"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	bi "github.com/sonr-io/sonr/x/bucket/service"
	rt "github.com/sonr-io/sonr/x/registry/types"
)

func (mtr *motorNodeImpl) GenerateBucket(request mt.GenerateBucketRequest) (*mt.GenerateBucketResponse, error) {
	// Validate the request
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	config, err := mtr.fetchBucketConfig(request.Bucket, request.Uuid, request.Creator, request.Name)
	if err != nil {
		return nil, err
	}

	service, cid, err := bi.GenerateBucket(mtr.sh, config, mtr.GetAddress())
	if err != nil {
		return nil, err
	}

	services, err := rt.ConvertServices([]*rt.Service{service})
	if err != nil {
		return nil, err
	}
	
	mtr.DIDDocument.AddService(services[0])
	docBz, err := mtr.DIDDocument.MarshalJSON()
	if err != nil {
		return nil, err
	}

	doc, err := rt.NewDIDDocumentFromBytes(docBz)
	if err != nil {
		return nil, err
	}

	msg1 := rt.NewMsgUpdateWhoIs(mtr.GetAddress(), docBz)
	txRaw, err := tx.SignTxWithWallet(mtr.GetWallet(), "/sonrio.sonr.registry.MsgUpdateWhoIs", msg1)
	if err != nil {
		return nil, err
	}

	resp, err := mtr.GetClient().BroadcastTx(txRaw)
	if err != nil {
		return nil, err
	}

	cwir := &rt.MsgUpdateWhoIsResponse{}
	if err := client.DecodeTxResponseData(resp.TxResponse.Data, cwir); err != nil {
		return nil, err
	}
	return &mt.GenerateBucketResponse{
		DidDocument: doc,
		Uri:         service.ServiceEndpoint,
		Bucket:      config,
		Cid:         cid,			
	}, nil
}
