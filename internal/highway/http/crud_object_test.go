package core_test

import (
	"context"
	"errors"
	"testing"

	hw "github.com/sonr-io/sonr/internal/highway"
	ot "github.com/sonr-io/sonr/x/object/types"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateObject_ValidateBasic(t *testing.T) {
	ctx := context.Background()
	_, err := hw.NewHighway(ctx)
	require.NoError(t, err)

	tests := []struct {
		name     string
		msg      ot.MsgCreateObject
		expected ot.MsgCreateObjectResponse
		err      error
	}{
		{
			name: "invalid object",
			msg:  ot.MsgCreateObject{},
			err:  errors.New("object to register must have fields"),
		}, {
			name: "valid object",
			msg: ot.MsgCreateObject{
				Creator:     "snr1ulu0a0eew3w3sj8nk8lx5z3cmaw46u383kek9t",
				Label:       "a label",
				Description: "a description",
				InitialFields: []*ot.TypeField{
					{
						Name: "another label",
						Kind: 0,
					},
				},
			},
			expected: ot.MsgCreateObjectResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// resp, err := highway.CreateObject(ct)
			// if tt.err != nil {
			// 	require.Error(t, err, tt.err)
			// 	return
			// }

			//require.ElementsMatch(t, resp, tt.expected)

			//	require.NoError(t, err)
		})
	}
}
