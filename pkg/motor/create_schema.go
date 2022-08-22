package motor

import (
	"fmt"

	"github.com/sonr-io/sonr/pkg/client"
	mt "github.com/sonr-io/sonr/pkg/motor/types"
	"github.com/sonr-io/sonr/pkg/tx"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (mtr *motorNodeImpl) CreateSchema(request mt.CreateSchemaRequest) (mt.CreateSchemaResponse, error) {
	mtr.Logger.Debug("Creating schema definition from array: %#v\n", request.Fields)
	listFields, err := convertFields(request.Fields)
	if err != nil {
		mtr.Logger.Error("proccessing fields failed with error: %s", err)
		return mt.CreateSchemaResponse{}, fmt.Errorf("process fields: %s", err)
	}
	createSchemaMsg := st.NewMsgCreateSchema(convertMetadata(request.Metadata), &st.SchemaDefinition{
		Creator: mtr.Address,
		Label:   request.Label,
		Fields:  listFields,
	})

	txRaw, err := tx.SignTxWithWallet(mtr.Wallet, "/sonrio.sonr.schema.MsgCreateSchema", createSchemaMsg)
	if err != nil {
		return mt.CreateSchemaResponse{}, fmt.Errorf("sign tx with wallet: %s", err)
	}
	mtr.Logger.Info("Broadcasting create schema request")
	resp, err := mtr.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return mt.CreateSchemaResponse{}, fmt.Errorf("broadcast tx: %s", err)
	}
	mtr.Logger.Info("Broadcasting create schema request returned wtih: %d", resp.TxResponse.Code)
	csresp := &st.MsgCreateSchemaResponse{}
	if err := client.DecodeTxResponseData(resp.TxResponse.Data, csresp); err != nil {
		mtr.Logger.Error("Error while unmarshalling Create Schema response: %s", err)
		return mt.CreateSchemaResponse{}, fmt.Errorf("decode MsgCreateSchemaResponse: %s", err)
	}

	// store reference to newly created WhatIs
	_, err = mtr.Resources.StoreWhatIs(csresp.WhatIs)
	if err != nil {
		mtr.Logger.Error("Error while storing WhatIs %s", err)
		return mt.CreateSchemaResponse{}, fmt.Errorf("store WhatIs: %s", err)
	}

	return mt.CreateSchemaResponse{
		WhatIs: csresp.WhatIs,
	}, nil
}

func convertFields(fields map[string]st.SchemaKind) ([]*st.SchemaKindDefinition, error) {
	result := make([]*st.SchemaKindDefinition, len(fields))
	var i int32
	for k, v := range fields {
		result[i] = &st.SchemaKindDefinition{
			Name:  k,
			Field: v,
		}
		i += 1
	}

	return result, nil
}

func convertMetadata(m map[string]string) []*st.MetadataDefintion {
	result := make([]*st.MetadataDefintion, 0)
	for k, v := range m {
		result = append(result, &st.MetadataDefintion{
			Key:   k,
			Value: v,
		})
	}
	return result
}
