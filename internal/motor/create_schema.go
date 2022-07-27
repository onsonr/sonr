package motor

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sonr-io/sonr/pkg/tx"
	st "github.com/sonr-io/sonr/x/schema/types"
	mt "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

func (mtr *MotorNode) CreateSchema(requestBytes []byte) (mt.CreateSchemaResponse, error) {
	var request mt.CreateSchemaRequest
	if err := json.Unmarshal(requestBytes, &request); err != nil {
		return mt.CreateSchemaResponse{}, fmt.Errorf("error unmarshalling request: %s", err)
	}

	listFields, err := convertFields(request.Fields)
	if err != nil {
		return mt.CreateSchemaResponse{}, fmt.Errorf("error processing fields: %s", err)
	}
	createSchemaMsg := st.NewMsgCreateSchema(&st.SchemaDefinition{
		Creator: mtr.Address,
		Label:   request.Label,
		Fields:  listFields,
	})

	txRaw, err := tx.SignTxWithWallet(mtr.Wallet, "/sonrio.sonr.schema.MsgCreateSchema", createSchemaMsg)
	if err != nil {
		return mt.CreateSchemaResponse{}, err
	}

	resp, err := mtr.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return mt.CreateSchemaResponse{}, err
	}

	if resp.TxResponse.RawLog != "[]" {
		return mt.CreateSchemaResponse{}, errors.New(resp.TxResponse.RawLog)
	}

	fmt.Print("done: ")
	fmt.Println(resp)

	// TODO
	return mt.CreateSchemaResponse{}, nil
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
