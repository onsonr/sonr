package motor

import (
	"fmt"

	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/tx"
	st "github.com/sonr-io/sonr/x/schema/types"
	mt "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

func (mtr *motorNodeImpl) CreateSchema(request mt.CreateSchemaRequest) (mt.CreateSchemaResponse, error) {

	listFields, err := convertFields(request.Fields)
	if err != nil {
		return mt.CreateSchemaResponse{}, fmt.Errorf("process fields: %s", err)
	}
	createSchemaMsg := st.NewMsgCreateSchema(&st.SchemaDefinition{
		Creator: mtr.Address,
		Label:   request.Label,
		Fields:  listFields,
	})

	txRaw, err := tx.SignTxWithWallet(mtr.Wallet, "/sonrio.sonr.schema.MsgCreateSchema", createSchemaMsg)
	if err != nil {
		return mt.CreateSchemaResponse{}, fmt.Errorf("sign tx with wallet: %s", err)
	}

	resp, err := mtr.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return mt.CreateSchemaResponse{}, fmt.Errorf("broadcast tx: %s", err)
	}

	csresp := &st.MsgCreateSchemaResponse{}
	if err := client.DecodeTxResponseData(resp.TxResponse.Data, csresp); err != nil {
		return mt.CreateSchemaResponse{}, fmt.Errorf("decode MsgCreateSchemaResponse: %s", err)
	}

	whatIsBytes, err := csresp.WhatIs.Marshal()
	if err != nil {
		return mt.CreateSchemaResponse{}, fmt.Errorf("marshal WhatIs: %s", err)
	}
	return mt.CreateSchemaResponse{
		WhatIs: whatIsBytes,
	}, nil
}

func convertFields(fields map[string]mt.CreateSchemaRequest_SchemaKind) ([]*st.SchemaKindDefinition, error) {
	result := make([]*st.SchemaKindDefinition, len(fields))
	var i int32
	for k, v := range fields {
		// Note: This will work while mt and st schema types stay in sync.
		// This should be refactored such that there is only one proto for these types
		sk, ok := mt.CreateSchemaRequest_SchemaKind_value[v.String()]
		if !ok {
			return nil, fmt.Errorf("invalid schema kind: %s", v)
		}
		result[i] = &st.SchemaKindDefinition{
			Name:  k,
			Field: st.SchemaKind(sk),
		}
		i += 1
	}

	return result, nil
}
