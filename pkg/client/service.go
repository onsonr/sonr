package client

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
)

func (c *clientStub) BroadcastTx(tx []byte) (*sdk.TxResponse, error) {
	res, err := c.cctx.BroadcastTx(tx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DecodeTxResponseData(d string, v proto.Unmarshaler) error {
	data, err := hex.DecodeString(d)
	if err != nil {
		return err
	}

	anyWrapper := new(types.Any)
	if err := proto.Unmarshal(data, anyWrapper); err != nil {
		return err
	}

	// TODO: figure out if there's a better 'cosmos' way of doing this
	// you have to unwrap the Any twice, and the first time the bytes get decoded
	// in the 'TypeUrl' field instead of 'Value' field
	any := new(types.Any)
	if err := proto.Unmarshal([]byte(anyWrapper.TypeUrl), any); err != nil {
		return err
	}

	return v.Unmarshal(any.Value)
}
